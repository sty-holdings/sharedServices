package sharedServices

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
	pis "github.com/sty-holdings/sharedServices/v2025/programInfo"
)

// NewGRPCServer - builds a reusable gRPC Server that creates an instance name and builds a connection.
// The Port must be at or above 50051. This will not register the server or execute the serve command.
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

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_GRPC_SERVER, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadConfig(ctv.LBL_SERVICE_GRPC_SERVER, configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(ctv.LBL_SERVICE_GRPC_SERVER, tConfig); errorInfo.Error != nil {
		return
	}

	servicePtr = &GRPCService{
		debugModeOn: tConfig.DebugModeOn,
		Secure: SecureSettings{
			ServerSide: tConfig.Secure.ServerSide,
			Mutual:     tConfig.Secure.Mutual,
		},
		Host: tConfig.Host,
		Port: uint(tConfig.Port),
	}

	if tConfig.DebugModeOn {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stdout, 2))
	}

	if tListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", servicePtr.Host, servicePtr.Port)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_GRPC_SERVER, ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	servicePtr.GRPCListenerPtr = &tListener

	if servicePtr.Secure.ServerSide {
		if tCreds, errorInfo = LoadTLSCredentialsCACertKey(ctv.LBL_SERVICE_GRPC_SERVER, tConfig.TLSInfo); errorInfo.Error != nil {
			return
		}
		servicePtr.GRPCServerPtr = grpc.NewServer(grpc.Creds(tCreds))
		errorInfo = hlps.CheckPointerNotNil(ctv.LBL_SERVICE_GRPC_SERVER, servicePtr.GRPCServerPtr, ctv.LBL_POINTER)
		return
	}

	servicePtr.GRPCServerPtr = grpc.NewServer()
	errorInfo = hlps.CheckPointerNotNil(ctv.LBL_SERVICE_GRPC_SERVER, servicePtr.GRPCServerPtr, ctv.LBL_POINTER)

	return
}

// NewGRPCClient - builds a reusable gRPC Client. The Host and Port must match the server. This will not create the message client or
// execute any services. Here are examples for reference:
//
//		d := protos_def.NewHalServicesClient(gRPCServicePtr.GRPCClientPtr)
//		myPongReplyPtr, err = d.PingPong(gRPCServicePtr.timeoutContext, &protos_def.PingRequest{Email: "scott-DK@yackofamily.com"},
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewGRPCClient(configFilename string) (gRPCServicePtr *GRPCService, errorInfo errs.ErrorInfo) {

	var (
		tConfig      GRPCConfig
		tGRPCAddress string
		tDailOptions []grpc.DialOption
		tDailOption  grpc.DialOption
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_GRPC_CLIENT, configFilename, ctv.LBL_CONFIG_GRPC_CLIENT_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadConfig(ctv.LBL_SERVICE_GRPC_CLIENT, configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(ctv.LBL_SERVICE_GRPC_CLIENT, tConfig); errorInfo.Error != nil {
		return
	}

	gRPCServicePtr = &GRPCService{
		debugModeOn: tConfig.DebugModeOn,
		Host:        tConfig.Host,
		Port:        uint(tConfig.Port),
		Secure: SecureSettings{
			ServerSide: tConfig.Secure.ServerSide,
			Mutual:     tConfig.Secure.Mutual,
		},
		Timeout: time.Duration(tConfig.Timeout) * time.Second,
	}

	if tConfig.DebugModeOn {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(os.Stdout, os.Stdout, os.Stdout, 2))
	}

	tGRPCAddress = fmt.Sprintf("%s:%s", tConfig.Host, strconv.Itoa(tConfig.Port)) // Localhost in the host file must point to 127.0.0.1 only.

	if gRPCServicePtr.Secure.ServerSide {
		if tDailOption, errorInfo = LoadTLSCABundle(ctv.LBL_SERVICE_GRPC_CLIENT, tConfig.TLSInfo); errorInfo.Error != nil {
			return
		}
	} else {
		tDailOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	}
	tDailOptions = append(tDailOptions, tDailOption)

	if tConfig.KeepAlive != (KeepAlive{}) {
		tDailOption = grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                time.Duration(tConfig.KeepAlive.PingInternalSec) * time.Second, // Ping interval.
				Timeout:             time.Duration(tConfig.KeepAlive.PingTimeoutSec) * time.Second,
				PermitWithoutStream: tConfig.KeepAlive.PermitWithoutStream,
			},
		)
		tDailOptions = append(tDailOptions, tDailOption)
	}

	if gRPCServicePtr.GRPCClientPtr, errorInfo.Error = grpc.NewClient(tGRPCAddress, tDailOptions...); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_GRPC_CLIENT, ctv.LBL_GRPC_CLIENT, ctv.TXT_FAILED))
	}

	return
}

//  Private Functions

// loadConfig - reads, and returns a grpc service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile
//	Verifications: validateConfiguration
func loadConfig(creator string, configFilename string) (grpcConfig GRPCConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(creator, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(creator, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &grpcConfig); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(creator, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	return
}

// LoadTLSCABundle - loads the CA Bundle certificate into a x509 certificate pool and returns the dailOption.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCABundle(creator string, tlsConfig jwts.TLSInfo) (dailOption grpc.DialOption, errorInfo errs.ErrorInfo) {

	var (
		tCABundleFile  []byte
		tCACertPoolPtr *x509.CertPool
	)

	if tCABundleFile, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(creator, ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	tCACertPoolPtr = x509.NewCertPool()
	if ok := tCACertPoolPtr.AppendCertsFromPEM(tCABundleFile); !ok {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(creator, ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
	}

	config := &tls.Config{
		RootCAs: tCACertPoolPtr,
	}

	if dailOption = grpc.WithTransportCredentials(credentials.NewTLS(config)); dailOption == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFailedGRPCDialOption, errs.BuildLabelValue(creator, ctv.LBL_GRPC_DIAL_OPTION, ctv.TXT_FAILED))
	}

	return
}

// LoadTLSCredentialsCACertKey - load the CA bundle, x509 certificate, and private key
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func LoadTLSCredentialsCACertKey(creator string, tlsConfig jwts.TLSInfo) (creds credentials.TransportCredentials, errorInfo errs.ErrorInfo) {

	var (
		tCertPool   *x509.CertPool
		tServerCA   []byte
		tServerCert tls.Certificate
	)

	if tServerCA, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFailedReadFile, errs.BuildLabelValue(creator, ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}
	tCertPool = x509.NewCertPool()
	if tCertPool.AppendCertsFromPEM(tServerCA) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFailedTlsRootCaLoading, errs.BuildLabelValue(creator, ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	if tServerCert, errorInfo.Error = tls.LoadX509KeyPair(tlsConfig.TLSCertFQN, tlsConfig.TLSPrivateKeyFQN); errorInfo.Error != nil {
		if strings.Contains(errorInfo.Error.Error(), "tls: private key") {
			errorInfo = errs.NewErrorInfo(
				errorInfo.Error, errs.BuildLabelSubLabelValueMessage(creator, pis.GetMyFunctionInfo(true).Name, ctv.LBL_TLS_PRIVATE_KEY_FILENAME, tlsConfig.TLSPrivateKeyFQN, ctv.TXT_FAILED),
			)
			return
		}
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error, errs.BuildLabelValueMessage(creator, pis.GetMyFunctionInfo(true).Name, ctv.VAL_EMPTY, ctv.TXT_FAILED),
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

func validateConfig(creator string, config GRPCConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(creator, config.Host, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
		return
	}
	if config.Port < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCPort, errs.BuildLabelValue(creator, ctv.LBL_GRPC_PORT, strconv.Itoa(config.Port)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(creator, config.TLSInfo.TLSCABundleFQN, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(creator, config.TLSInfo.TLSCertFQN, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(creator, config.TLSInfo.TLSPrivateKeyFQN, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
		return
	}
	if config.Timeout < ctv.VAL_ONE {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCTimeout, errs.BuildLabelValue(creator, ctv.LBL_GRPC_TIMEOUT, strconv.Itoa(config.Timeout)))
	}

	return
}
