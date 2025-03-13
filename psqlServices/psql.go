package sharedServices

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v3"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

var (
	CTXBackground = context.Background()
)

// NewPSQLServer - builds a reusable PostgreSQL Service to access Postgres databases.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewPSQLServer(configFilename string) (servicePtr *PSQLService, errorInfo errs.ErrorInfo) {

	var (
		tConfig            PSQLConfig
		tConnectionPoolPtr *pgxpool.Pool
	)

	if errorInfo = hlps.CheckValueNotEmpty(configFilename, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadPSQLConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(tConfig); errorInfo.Error != nil {
		return
	}

	if tConnectionPoolPtr, errorInfo = getConnection(tConfig); errorInfo.Error != nil {
		return
	}
	servicePtr = &PSQLService{
		DebugOn:           tConfig.Debug,
		ConnectionPoolPtr: tConnectionPoolPtr,
	}

	return
}

//  Private Functions

// buildConnectionString - returns the connection string based on the PSQL configuration.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func buildConnectionString(config PSQLConfig) (dbConnString string) {

	return fmt.Sprintf(PSQL_CONN_STRING, config.DBName, config.Host, config.MaxConnections, config.Password, config.Port, config.SSLMode, config.Timeout, config.UserName)

}

// getConnection - will create a connection pool and connect to a database.
// DBName   		Name of the Postgres database
// Host     		Internet DNS or IP address of the server running the instance of Postgres
// Max Connections 	Must be greater than 0 and less than 100 across all instances.
// Password 		Encrypted password for authentication
// Port     		Interface the connection communicates with Postgres
// SSL Mode 		Which mode will be used to connect to the PSQL server.
//
//						Blocked: See PSQL_SSL_MODE_DISABLE, PSQL_SSL_MODE_ALLOW, PSQL_SSL_MODE_PREFER, PSQL_SSL_MODE_REQUIRED
//						Supported: PSQL_SSL_MODE_VERIFY, PSQL_SSL_MODE_VERIFY_FULL
//	         			https://www.postgresql.org/docs/current/libpq-ssl.html
//
// Timeout  		Number of seconds a request must complete (1 is min, 3 seconds is normal, and 5 is max)
// UserName 		Encrypted UserName used to authenticate
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getConnection(config PSQLConfig) (connectionPoolPtr *pgxpool.Pool, errorInfo errs.ErrorInfo) {

	var (
		tCACertPoolPtr *x509.CertPool
		tCert          tls.Certificate
		tConfigPtr     *pgxpool.Config
	)

	if tConfigPtr, errorInfo.Error = pgxpool.ParseConfig(buildConnectionString(config)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_PSQL_CONNECTION, ctv.TXT_FAILED))
		return
	}

	if tCACertPoolPtr, errorInfo = loadTLSCABundle(config.PSQLTLSInfo); errorInfo.Error != nil {
		return
	}
	if tCert, errorInfo = loadTLSCertKeyPair(config.PSQLTLSInfo); errorInfo.Error != nil {
		return
	}

	tConfigPtr.ConnConfig.TLSConfig = &tls.Config{
		RootCAs:            tCACertPoolPtr,
		Certificates:       []tls.Certificate{tCert},
		InsecureSkipVerify: false,
		ServerName:         tConfigPtr.ConnConfig.Host,
	}

	if connectionPoolPtr, errorInfo.Error = pgxpool.ConnectConfig(CTXBackground, tConfigPtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_PSQL_CONNECTION, ctv.TXT_FAILED))
		return
	}

	errorInfo = isConnectionActive(connectionPoolPtr, config.DBName)

	return
}

// loadTLSCABundle - loads the CA Bundle certificate into a x509 certificate pool.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func loadTLSCABundle(tlsConfig jwts.TLSInfo) (caCertPoolPtr *x509.CertPool, errorInfo errs.ErrorInfo) {

	var (
		tCABundleFile []byte
	)

	if tCABundleFile, errorInfo.Error = os.ReadFile(tlsConfig.TLSCABundleFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	caCertPoolPtr = x509.NewCertPool()
	if ok := caCertPoolPtr.AppendCertsFromPEM(tCABundleFile); !ok {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
	}

	return
}

// loadTLSCertKeyPair - load the x509 certificate, and private key
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func loadTLSCertKeyPair(tlsConfig jwts.TLSInfo) (cert tls.Certificate, errorInfo errs.ErrorInfo) {

	if cert, errorInfo.Error = tls.LoadX509KeyPair(tlsConfig.TLSCertFQN, tlsConfig.TLSPrivateKeyFQN); errorInfo.Error != nil {
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

// loadPSQLConfig - reads, and returns a psql service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile
//	Verifications: validateConfiguration
func loadPSQLConfig(configFilename string) (config PSQLConfig, errorInfo errs.ErrorInfo) {

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

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_EXTENSION_CONFIG_FILENAME, configFilename))
		return
	}

	return
}

func validateConfig(config PSQLConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(config.DBName, errs.ErrPSQLDBNameEmpty, ctv.LBL_PSQL_DBNAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.Host, errs.ErrPSQLHostEmpty, ctv.LBL_PSQL_HOST); errorInfo.Error != nil {
		return
	}
	if config.MaxConnections <= ctv.VAL_ZERO && config.MaxConnections >= ctv.VAL_ONE_HUNDRED {
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLMaxConnectionsInvalid, errs.BuildLabelValue(ctv.LBL_PSQL_MAX_CONNECTIONS, strconv.Itoa(config.Port)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.Host, errs.ErrPSQLPasswordEmpty, ctv.LBL_PSQL_PASSWORD); errorInfo.Error != nil {
		return
	}
	if config.Port != ctv.VAL_PSQL_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLPortInvalid, errs.BuildLabelValue(ctv.LBL_PSQL_PORT, strconv.Itoa(config.Port)))
		return
	}
	switch config.SSLMode {
	case PSQL_SSL_MODE_ALLOW:
		fallthrough
	case PSQL_SSL_MODE_DISABLE:
		fallthrough
	case PSQL_SSL_MODE_PREFER:
		fallthrough
	case PSQL_SSL_MODE_REQUIRED:
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLSSLModeNotAllowed, errs.BuildLabelValue(ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
		return
	case PSQL_SSL_MODE_VERIFY:
		if errorInfo = hlps.CheckValueNotEmpty(config.PSQLTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
			return
		}
	case PSQL_SSL_MODE_VERIFY_FULL:
		if errorInfo = hlps.CheckValueNotEmpty(config.PSQLTLSInfo.TLSCABundleFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(config.PSQLTLSInfo.TLSCertFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(config.PSQLTLSInfo.TLSPrivateKeyFQN, errs.ErrRequiredParameterMissing, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLTimeoutInvalid, errs.BuildLabelValue(ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
	}
	if config.Timeout < ctv.VAL_ONE && config.Timeout >= ctv.VAL_FIVE {
		errorInfo = errs.NewErrorInfo(errs.ErrGRPCTimeoutInvalid, errs.BuildLabelValue(ctv.LBL_PSQL_TIMEOUT, strconv.Itoa(config.Timeout)))
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.UserName, errs.ErrPSQLUserEmpty, ctv.LBL_PSQL_USER_NAME); errorInfo.Error != nil {
		return
	}

	return
}

// isConnectionActive - checks the connection by querying the database.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func isConnectionActive(connectionPtr *pgxpool.Pool, dbName string) (errorInfo errs.ErrorInfo) {

	//goland:noinspection ALL
	const checkStatActivity = "SELECT * FROM pg_stat_activity WHERE datname = $1 and state = 'active';"

	var (
		tRows Rows
	)

	if connectionPtr == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLConnFalied, errs.BuildLabelValue(ctv.LBL_PSQL_DBNAME, dbName))
		return
	}

	tRows, errorInfo.Error = connectionPtr.Query(CTXBackground, checkStatActivity, dbName)
	if errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLConnFalied, errs.BuildLabelValue(ctv.LBL_PSQL_DBNAME, dbName))
	}
	defer tRows.Close()

	return
}
