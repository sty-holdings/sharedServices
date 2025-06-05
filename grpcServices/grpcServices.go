package sharedServices

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
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
		tCreds             credentials.TransportCredentials
		tListener          net.Listener
		tConfig            GRPCConfig
		tGRPCServerOptions []grpc.ServerOption
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_GRPC_SERVER, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadConfig(ctv.LBL_SERVICE_GRPC_SERVER, configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateServerConfig(ctv.LBL_SERVICE_GRPC_SERVER, tConfig); errorInfo.Error != nil {
		return
	}

	servicePtr = &GRPCService{
		debugModeOn:   tConfig.DebugModeOn,
		GRPCClientPtr: nil,
		Host:          tConfig.Host,
		Port:          uint(tConfig.Port),
		Secure: SecureSettings{
			ServerSide: tConfig.Secure.ServerSide,
			Mutual:     tConfig.Secure.Mutual,
		},
		ServerMinPingTime: 0,
		Timeout:           0,
	}

	if tConfig.DebugModeOn {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2WithVerbosity(io.Discard, os.Stdout, os.Stdout, ctv.VAL_ONE))
	}

	// Configure keepalive enforcement policy
	tGRPCServerOptions = append(
		tGRPCServerOptions, grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Duration(tConfig.ServerKeepAlive.ServerEnforcementPolicy.MinTimeClientPingsSec) * time.Second,
				PermitWithoutStream: tConfig.ServerKeepAlive.ServerEnforcementPolicy.PermitWithoutStream,
			},
		),
	)

	// Configure keepalive parameters
	tGRPCServerOptions = append(
		tGRPCServerOptions, grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(tConfig.ServerKeepAlive.ServerParameters.PingIntervalSec) * time.Second,
				Timeout: time.Duration(tConfig.ServerKeepAlive.ServerParameters.PingTimeoutSec) * time.Second,
			},
		),
	)

	if tListener, errorInfo.Error = net.Listen(ctv.VAL_TCP, fmt.Sprintf("%s:%d", servicePtr.Host, servicePtr.Port)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_GRPC_SERVER, ctv.LBL_GRPC_LISTENER, ctv.TXT_FAILED))
		return
	}
	servicePtr.GRPCListenerPtr = &tListener

	if servicePtr.Secure.ServerSide {
		if tCreds, errorInfo = LoadTLSCredentialsCACertKey(ctv.LBL_SERVICE_GRPC_SERVER, tConfig.TLSInfo); errorInfo.Error != nil {
			return
		}
		tGRPCServerOptions = append(tGRPCServerOptions, grpc.Creds(tCreds))
	}

	servicePtr.GRPCServerPtr = grpc.NewServer(tGRPCServerOptions...)
	if errorInfo = hlps.CheckPointerNotNil(ctv.LBL_SERVICE_GRPC_SERVER, servicePtr.GRPCServerPtr, ctv.LBL_POINTER); errorInfo.Error != nil {
		return
	}

	return
}

// NewGRPCClient - initializes and returns a Client GRPCService pointer along with error details.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo
//	Verifications: validateConfig, hlps.CheckValueNotEmpty
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

	if errorInfo = validateClientConfig(ctv.LBL_SERVICE_GRPC_CLIENT, tConfig); errorInfo.Error != nil {
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

	if tConfig.ClientKeepAlive != (ClientKeepAlive{}) {
		tDailOption = grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                time.Duration(tConfig.ClientKeepAlive.PingIntervalSec) * time.Second,
				Timeout:             time.Duration(tConfig.ClientKeepAlive.PingTimeoutSec) * time.Second,
				PermitWithoutStream: tConfig.ClientKeepAlive.PermitWithoutStream,
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

// validateServerConfig - validates the server configuration settings.
//
//	Customer Messages: None
//	Errors: errs.Err if validation fails
//	Verifications: ctv. for label constants, hlps.CheckValueNotEmpty, hlps.CheckValueGreatZero to validate configurations
func validateServerConfig(creator string, config GRPCConfig) (errorInfo errs.ErrorInfo) {

	// The config.DebugModeOn is either true or false. No need to check the value.
	if errorInfo = hlps.CheckValueNotEmpty(creator, config.Host, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueGreatZero(creator, config.ServerKeepAlive.ServerEnforcementPolicy.MinTimeClientPingsSec, ctv.LBL_MIN_TIME_CLIENT_PINGS_SEC); errorInfo.Error != nil {
		return
	}
	// The config.ServerKeepAlive.ServerEnforcementPolicy.PermitWithoutStream is either true or false. No need to check the value.
	// The config.ServerKeepAlive.ServerParameters.MaxConnectionIdleSec uses the default setting and is optional.
	// The config.ServerKeepAlive.ServerParameters.MaxConnectionAgeSec uses the default setting and is optional.
	// Then config.ServerKeepAlive.ServerParameters.MaxConnectionAgeGraceSec uses the default setting and is optional.
	if errorInfo = hlps.CheckValueGreatZero(creator, config.ServerKeepAlive.ServerParameters.PingIntervalSec, ctv.LBL_PING_INTERVAL_SEC); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueGreatZero(creator, config.ServerKeepAlive.ServerParameters.PingTimeoutSec, ctv.LBL_PING_TIMEOUT_SEC); errorInfo.Error != nil {
		return
	}
	if config.Port < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCPort, errs.BuildLabelValue(creator, ctv.LBL_GRPC_PORT, strconv.Itoa(config.Port)))
		return
	}
	// The config.Secure.ServerSide is either true or false. No need to check the value.
	// The config.Secure.Mutual is either true or false and is optional. No need to check the value.
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

// validateClientConfig - validates the configuration for a gRPC client.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrGreaterThanZero, errs.ErrInvalidGRPCPort, errs.ErrInvalidGRPCTimeout
//	Verifications: ctv.
func validateClientConfig(creator string, config GRPCConfig) (errorInfo errs.ErrorInfo) {

	// The config.DebugModeOn is either true or false. No need to check the value.
	if errorInfo = hlps.CheckValueNotEmpty(creator, config.Host, ctv.LBL_GRPC_HOST); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueGreatZero(creator, config.ClientKeepAlive.PingIntervalSec, ctv.LBL_PING_INTERVAL_SEC); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueGreatZero(creator, config.ClientKeepAlive.PingTimeoutSec, ctv.LBL_PING_TIMEOUT_SEC); errorInfo.Error != nil {
		return
	}
	if config.Port < ctv.VAL_GRPC_MIN_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCPort, errs.BuildLabelValue(creator, ctv.LBL_GRPC_PORT, strconv.Itoa(config.Port)))
		return
	}
	// The config.Secure.ServerSide is either true or false. No need to check the value.
	// The config.Secure.Mutual is either true or false and is optional. No need to check the value.
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
