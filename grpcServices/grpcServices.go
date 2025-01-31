package sharedServices

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

// NewGRPCService - builds a reusable gRPC Service that creates an instance name and builds a connection. The GRPCPort must be at or above 50051.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGRPCService(
	extensionName string,
	config GRPCConfiguration,
) (gRPCServicePtr *GRPCService, errorInfo errs.ErrorInfo) {

	var (
		tCertificate          tls.Certificate
		tCACertPool           *x509.CertPool
		tGRPCListener         net.Listener
		tTLSConfig            *tls.Config
		tTransportCredentials credentials.TransportCredentials
	)

	if errorInfo = hlps.CheckValueNotEmpty(extensionName, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCHost, errs.ErrRequiredParameterMissing, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
		return
	}
	if config.GRPCPort < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCPortInvalid, errs.BuildLabelValue(ctv.LBL_GRPC_PORT, strconv.Itoa(config.GRPCPort)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCTLSInfo.TLSCertFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCTLSInfo.TLSPrivateKeyFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
		return
	}

	gRPCServicePtr = &GRPCService{
		host: config.GRPCHost,
	}
	if tGRPCListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", config.GRPCHost, config.GRPCPort)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	gRPCServicePtr.GRPCListenerPtr = &tGRPCListener

	tTLSConfig = &tls.Config{}
	switch {
	case config.GRPCSecure.ServerSide:
		if tCertificate, errorInfo = LoadTLSCredentialsWithKey(config.GRPCTLSInfo); errorInfo.Error != nil {
			return
		}

		tTLSConfig.Certificates = []tls.Certificate{tCertificate}
		tTLSConfig.ClientAuth = tls.NoClientCert
		tTransportCredentials = credentials.NewTLS(tTLSConfig)

		gRPCServicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tTransportCredentials))
		if hlps.CheckPointerNotNil(gRPCServicePtr.GRPCServerPtr, errs.ErrPointerMissing, ctv.FN_GRPC_SERVER_POINTER); errorInfo.Error != nil {
			return
		}
	case config.GRPCSecure.Mutual:
		if tCertificate, errorInfo = LoadTLSCredentialsWithKey(config.GRPCTLSInfo); errorInfo.Error != nil {
			return
		}

		if tCACertPool, errorInfo = LoadTLSCACertificate(config.GRPCTLSInfo); errorInfo.Error != nil {
			return
		}

		tTLSConfig.Certificates = []tls.Certificate{tCertificate}
		tTLSConfig.RootCAs = tCACertPool
		tTLSConfig.ClientCAs = tCACertPool
		tTLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tTransportCredentials = credentials.NewTLS(tTLSConfig)
		gRPCServicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tTransportCredentials))
	default:
		// This is the default security if server side and mutual are both set to false.
		tTLSConfig.ClientAuth = tls.NoClientCert
		gRPCServicePtr.GRPCServerPtr = grpc.NewServer()
	}

	return
}

// LoadTLSCACertificate - loads the CA Bundle certificate into a x509 certificate pool.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCACertificate(tlsConfig jwts.TLSInfo) (caCertPoolPtr *x509.CertPool, errorInfo errs.ErrorInfo) {

	var (
		tCACertificateFile []byte
	)

	if tCACertificateFile, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	caCertPoolPtr = x509.NewCertPool()
	if ok := caCertPoolPtr.AppendCertsFromPEM(tCACertificateFile); !ok {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
	}

	return
}

// LoadTLSCredentialsWithKey - load the x509 certificate with private key
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCredentialsWithKey(tlsConfig jwts.TLSInfo) (certificate tls.Certificate, errorInfo errs.ErrorInfo) {

	if certificate, errorInfo.Error = tls.LoadX509KeyPair(tlsConfig.TLSCertFQN, tlsConfig.TLSPrivateKeyFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error, fmt.Sprintf(
				"%s, %s", errs.BuildLabelValue(ctv.LBL_TLS_CERTIFICATE_FILENAME, tlsConfig.TLSCertFQN), errs.BuildLabelValue(
					ctv.LBL_TLS_PRIVATE_KEY_FILENAME,
					tlsConfig.TLSPrivateKey,
				),
			),
		)
	}

	return
}

//  Private Functions
