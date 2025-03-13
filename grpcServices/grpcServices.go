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
func NewGRPCServer(configFilename string) (servicePtr *GRPCService, errorInfo errs.ErrorInfo) {

	var (
		tCreds    credentials.TransportCredentials
		tListener net.Listener
		tConfig   GRPCConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(configFilename, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(tConfig); errorInfo.Error != nil {
		return
	}

	servicePtr = &GRPCService{
		DebugOn: tConfig.GRPCDebug,
		Secure: SecureSettings{
			ServerSide: tConfig.GRPCSecure.ServerSide,
			Mutual:     tConfig.GRPCSecure.Mutual,
		},
		Host: tConfig.GRPCHost,
		Port: tConfig.GRPCPort,
	}

	if tConfig.GRPCDebug {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stdout, 2))
	}

	if tListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", servicePtr.Host, servicePtr.Port)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	servicePtr.GRPCListenerPtr = &tListener

	if servicePtr.Secure.ServerSide {
		if tCreds, errorInfo = LoadTLSCredentialsCACertKey(tConfig.GRPCTLSInfo); errorInfo.Error != nil {
			return
		}
		servicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tCreds))
		errorInfo = hlps.CheckPointerNotNil(servicePtr.GRPCServerPtr, errs.ErrPointerMissing, ctv.FN_GRPC_SERVER_POINTER)
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

// loadConfig - reads, and returns a grpc service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile
//	Verifications: validateConfiguration
func loadConfig(configFilename string) (grpcConfig GRPCConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(configFilename, errs.ErrRequiredParameterMissing, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &grpcConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, configFilename))
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

func validateConfig(config GRPCConfig) (errorInfo errs.ErrorInfo) {

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
	if config.GRPCTimeout < ctv.VAL_ONE {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCTimeoutInvalid, errs.BuildLabelValue(ctv.LBL_GRPC_TIMEOUT, strconv.Itoa(config.GRPCTimeout)))
	}

	return
}
