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
		tCACertificateFile    []byte
		tCertificate          tls.Certificate
		tCACertPool           *x509.CertPool
		tGRPCListener         net.Listener
		tTLSConfig            *tls.Config
		tTransportCredentials credentials.TransportCredentials
	)

	if errorInfo = hlps.CheckValueNotEmpty(extensionName, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCHost, errs.ErrRequiredParameterMissing, ctv.LBL_HOST); errorInfo.Error != nil {
		return
	}
	if config.GRPCPort < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCPortInvalud, errs.BuildLabelValue(ctv.LBL_GRPC_PORT, strconv.Itoa(config.GRPCPort)))
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
	if errorInfo = hlps.CheckValueNotEmpty(config.GRPCTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
		return
	}

	gRPCServicePtr = &GRPCService{
		host:     config.GRPCHost,
		secure:   config.GRPCSecure,
		userInfo: ctv.UserInfo{},
	}

	if tGRPCListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", config.GRPCHost, config.GRPCPort)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	gRPCServicePtr.GRPCListenerPtr = &tGRPCListener

	tTLSConfig = &tls.Config{}
	if config.GRPCSecure {
		if tCertificate, errorInfo.Error = tls.LoadX509KeyPair(config.GRPCTLSInfo.TLSCertFQN, config.GRPCTLSInfo.TLSPrivateKeyFQN); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(
				errorInfo.Error, fmt.Sprintf(
					"%s, %s", errs.BuildLabelValue(ctv.LBL_TLS_CERTIFICATE_FILENAME, config.GRPCTLSInfo.TLSCertFQN), errs.BuildLabelValue(
						ctv.LBL_TLS_PRIVATE_KEY_FILENAME,
						config.GRPCTLSInfo.TLSPrivateKeyFQN,
					),
				),
			)
			return
		}
		if tCACertificateFile, errorInfo.Error = os.ReadFile(config.GRPCTLSInfo.TLSCABundleFQN); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, config.GRPCTLSInfo.TLSCABundleFQN))
			return
		}
		tCACertPool = x509.NewCertPool()
		if ok := tCACertPool.AppendCertsFromPEM(tCACertificateFile); !ok {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
			return
		}
		tTransportCredentials = credentials.NewTLS(tTLSConfig)
		tTLSConfig.Certificates = []tls.Certificate{tCertificate}
		tTLSConfig.ClientCAs = tCACertPool
		tTLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		gRPCServicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tTransportCredentials))
		return
	}

	tTLSConfig.ClientAuth = tls.NoClientCert
	gRPCServicePtr.GRPCServerPtr = grpc.NewServer()

	return
}

//  Private Functions
