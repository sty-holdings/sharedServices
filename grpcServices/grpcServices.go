package sharedServices

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

// NewGRPCServer - builds a reusable gRPC Server that creates an instance name and builds a connection.
// The GRPCPort must be at or above 50051. This will not register the server or execute the serve command.
// Here are examples for reference:
//
//		protos_def.RegisterHalServicesServer(gRPCServer.GRPCServerPtr, &Server{})
//		errorInfo.Error = gRPCServer.GRPCServerPtr.Serve(*gRPCServer.GRPCListenerPtr)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGRPCServer(gRPCConfigFilename string) (gRPCServicePtr *GRPCService, errorInfo errs.ErrorInfo) {

	var (
		tCreds        credentials.TransportCredentials
		tGRPCListener net.Listener
		tGRPCConfig   GRPCConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(gRPCConfigFilename, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tGRPCConfig, errorInfo = LoadGRPCConfig(gRPCConfigFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateGRPCConfig(tGRPCConfig); errorInfo.Error != nil {
		return
	}

	gRPCServicePtr = &GRPCService{
		DebugOn: tGRPCConfig.GRPCDebug,
		Secure: SecureSettings{
			ServerSide: tGRPCConfig.GRPCSecure.ServerSide,
			Mutual:     tGRPCConfig.GRPCSecure.Mutual,
		},
		Host: tGRPCConfig.GRPCHost,
		Port: tGRPCConfig.GRPCPort,
	}

	if tGRPCConfig.GRPCDebug {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stdout, 2))
	}

	if tGRPCListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", gRPCServicePtr.Host, gRPCServicePtr.Port)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	gRPCServicePtr.GRPCListenerPtr = &tGRPCListener

	if gRPCServicePtr.Secure.ServerSide {
		if tCreds, errorInfo = LoadTLSCredentialsCACertKey(tGRPCConfig.GRPCTLSInfo); errorInfo.Error != nil {
			return
		}
		gRPCServicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tCreds))
		errorInfo = hlps.CheckPointerNotNil(gRPCServicePtr.GRPCServerPtr, errs.ErrPointerMissing, ctv.FN_GRPC_SERVER_POINTER)
	}

	return
}

// NewGRPCClient - builds a reusable gRPC Client. The GRPCHost and GRPCPort must match the server. This will not create the message client or
// execute any services. Here are examples for reference:
//
//		d := protos_def.NewHalServicesClient(gRPCServicePtr.GRPCClientPtr)
//		myPongReplyPtr, err = d.PingPong(gRPCServicePtr.timeoutContext, &protos_def.PingRequest{Email: "scott-DK@yackofamily.com"},
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGRPCClient(
	gRPCConfigFilename string,
) (gRPCServicePtr *GRPCService, errorInfo errs.ErrorInfo) {

	//var (
	//	tDailOption grpc.DialOption
	//)

	//if errorInfo = hlps.CheckValueNotEmpty(ctv.EXT_SERVICE_GRPC_SERVER, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_NAME); errorInfo.Error != nil {
	//	return
	//}
	//if errorInfo = hlps.CheckValueNotEmpty(config.GRPCHost, errs.ErrRequiredParameterMissing, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
	//	return
	//}
	//if config.GRPCPort < ctv.VAL_GRPC_MIN_PORT {
	//	errorInfo = errs.NewErrorInfo(errs.ErrGRPCPortInvalid, errs.BuildLabelValue(ctv.LBL_GRPC_PORT, strconv.Itoa(config.GRPCPort)))
	//	return
	//}
	//if errorInfo = hlps.CheckValueNotEmpty(config.GRPCTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
	//	return
	//}
	//if config.GRPCTimeout <= ctv.VAL_ZERO {
	//	errorInfo = errs.NewErrorInfo(errs.ErrGRPCTimeoutInvalid, errs.BuildLabelValue(ctv.LBL_GRPC_TIMEOUT, strconv.Itoa(config.GRPCTimeout)))
	//	return
	//}
	//
	//if config.GRPCDebug {
	//	grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stdout, 2))
	//}
	//
	//gRPCServicePtr = &GRPCService{
	//	Host: config.GRPCHost,
	//	Port: config.GRPCPort,
	//}
	//
	//if tDailOption, errorInfo = LoadTLSCABundle(config.GRPCTLSInfo); errorInfo.Error != nil {
	//	return
	//}
	//
	//if gRPCServicePtr.GRPCClientPtr, errorInfo.Error = grpc.NewClient(fmt.Sprintf("%s:%s", config.GRPCHost, strconv.Itoa(config.GRPCPort)), tDailOption); errorInfo.Error != nil {
	//	errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GRPC_CLIENT, ctv.TXT_FAILED))
	//}

	return
}

//  Private Functions

// LoadGRPCConfig - reads, and returns a grpc service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func LoadGRPCConfig(geminiConfigFilename string) (grpcConfig GRPCConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(geminiConfigFilename, errs.ErrRequiredParameterMissing, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(geminiConfigFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, geminiConfigFilename))
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &grpcConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, geminiConfigFilename))
		return
	}

	return
}

// LoadTLSCABundle - loads the CA Bundle certificate into a x509 certificate pool and returns the dailOption.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCABundle(tlsConfig jwts.TLSInfo) (dailOption grpc.DialOption, errorInfo errs.ErrorInfo) {

	var (
		tCABundleFile  []byte
		tCACertPoolPtr *x509.CertPool
	)

	if tCABundleFile, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	tCACertPoolPtr = x509.NewCertPool()
	if ok := tCACertPoolPtr.AppendCertsFromPEM(tCABundleFile); !ok {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
	}

	config := &tls.Config{
		RootCAs: tCACertPoolPtr,
	}

	if dailOption = grpc.WithTransportCredentials(credentials.NewTLS(config)); dailOption == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrDialOptionFailed, errs.BuildLabelValue(ctv.LBL_DIAL_OPTION, ctv.TXT_FAILED))
	}

	return
}

// LoadTLSCredentialsCACertKey - load the CA bundle, x509 certificate, and private key
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCredentialsCACertKey(tlsConfig jwts.TLSInfo) (creds credentials.TransportCredentials, errorInfo errs.ErrorInfo) {

	var (
		tCertPool   *x509.CertPool
		tServerCA   []byte
		tServerCert tls.Certificate
	)

	if tServerCA, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFileUnreadable, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}
	tCertPool = x509.NewCertPool()
	if tCertPool.AppendCertsFromPEM(tServerCA) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrCABundleLoadingFailed, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	if tServerCert, errorInfo.Error = tls.LoadX509KeyPair(tlsConfig.TLSCertFQN, tlsConfig.TLSPrivateKeyFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error, fmt.Sprintf(
				"%s, %s", errs.BuildLabelValue(ctv.LBL_TLS_CERTIFICATE_FILENAME, tlsConfig.TLSCertFQN), errs.BuildLabelValue(
					ctv.LBL_TLS_PRIVATE_KEY_FILENAME,
					tlsConfig.TLSPrivateKey,
				),
			),
		)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{tServerCert},
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    tCertPool,
	}

	creds = credentials.NewTLS(config)

	return
}

func validateGRPCConfig(grpcCConfig GRPCConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(grpcCConfig.GRPCHost, errs.ErrRequiredParameterMissing, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
		return
	}
	if grpcCConfig.GRPCPort < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCPortInvalid, errs.BuildLabelValue(ctv.LBL_GRPC_PORT, strconv.Itoa(grpcCConfig.GRPCPort)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(grpcCConfig.GRPCTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(grpcCConfig.GRPCTLSInfo.TLSCertFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(grpcCConfig.GRPCTLSInfo.TLSPrivateKeyFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
		return
	}
	if grpcCConfig.GRPCTimeout < ctv.VAL_ONE {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCTimeoutInvalid, errs.BuildLabelValue(ctv.FN_GRPC_TIMEOUT, strconv.Itoa(grpcCConfig.GRPCTimeout)))
	}

	return
}
