package sharedServices

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strconv"

	"github.com/goccy/go-yaml"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

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
		tConfig           PSQLConfig
		tConnectionConfig PSQLConnectionConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, configFilename, errs.ErrEmptyRequiredParameter, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadPSQLConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(tConfig); errorInfo.Error != nil {
		return
	}

	servicePtr = &PSQLService{
		DebugOn:            tConfig.Debug,
		ConnectionPoolPtrs: make(map[string]*pgxpool.Pool),
	}

	tConnectionConfig = PSQLConnectionConfig{
		Debug:          tConfig.Debug,
		Host:           tConfig.Host,
		MaxConnections: tConfig.MaxConnections,
		Password:       tConfig.Password,
		Port:           tConfig.Port,
		SSLMode:        tConfig.SSLMode,
		PSQLTLSInfo:    tConfig.PSQLTLSInfo,
		Timeout:        tConfig.Timeout,
		UserName:       tConfig.UserName,
	}

	for _, database := range tConfig.DBName {
		tConnectionConfig.DBName = database
		if servicePtr.ConnectionPoolPtrs[database], errorInfo = getConnection(tConnectionConfig); errorInfo.Error != nil {
			return
		}
	}

	return
}

func (psqlService *PSQLService) BatchInsert(database string, role string, batchName string, insertStatement string, values [][]any) (errorInfo errs.ErrorInfo) {

	var (
		pCommandTag  pgconn.CommandTag
		pTransaction pgx.Tx
		tQueueString string
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, database, errs.ErrEmptyRequiredParameter, ctv.LBL_DATABASE); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, batchName, errs.ErrEmptyRequiredParameter, ctv.LBL_PSQL_BATCH); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, insertStatement, errs.ErrEmptyRequiredParameter, ctv.LBL_PSQL_INSERT); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.LBL_SERVICE_PSQL, values, errs.ErrEmptyRequiredParameter, ctv.LBL_VALUE); errorInfo.Error != nil {
		return
	}
	if len(values) > ctv.VAL_ONE_HUNDRED {
		errorInfo = errs.NewErrorInfo(errs.ErrArraySizeExceeded, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_BATCH, "values [][]any", ctv.TXT_IS_INVALID))
		return
	}
	
	if pTransaction, errorInfo.Error = psqlService.ConnectionPoolPtrs[database].BeginTx(CTXBackground, pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadWrite}); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRANSACTION, ctv.LBL_SERVICE_PSQL, batchName, ctv.TXT_FAILED))
		return
	}

	if role != ctv.VAL_EMPTY {
		tQueueString = fmt.Sprintf(SET_ROLE, role)
	}
	tQueueString += insertStatement

	for idx, row := range values {
		if pCommandTag, errorInfo.Error = pTransaction.Exec(CTXBackground, insertStatement, row...); errorInfo.Error != nil {
			return
		}
		if pCommandTag.RowsAffected() != ctv.VAL_ONE {
			errorInfo = errs.NewErrorInfo(
				errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_BATCH_INSERT, ctv.LBL_RECORD_NUMBER, strconv.Itoa(idx), ctv.TXT_FAILED),
			)
			errs.PrintErrorInfo(errorInfo)
			if errorInfo.Error = pTransaction.Rollback(CTXBackground); errorInfo.Error != nil {
				errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_COMMIT, ctv.LBL_PSQL_BATCH, batchName, ctv.TXT_FAILED))
				errs.PrintErrorInfo(errorInfo)
			}
			return
		}
	}

	if errorInfo.Error = pTransaction.Commit(CTXBackground); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_COMMIT, ctv.LBL_PSQL_BATCH, batchName, ctv.TXT_FAILED))
	}

	return
}

func (psqlService *PSQLService) Close() {

	for _, connectionPtr := range psqlService.ConnectionPoolPtrs {
		connectionPtr.Close()
	}
}

func (psqlService *PSQLService) TruncateTable(database string, schema string, tableName string) (errorInfo errs.ErrorInfo) {

	var (
		pStatement string
	)

	pStatement = fmt.Sprintf(TRUNCATE_TABLE, pgx.Identifier{schema}.Sanitize(), pgx.Identifier{tableName}.Sanitize())
	if _, errorInfo.Error = psqlService.ConnectionPoolPtrs[database].Exec(CTXBackground, pStatement); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRUNCATE, schema, tableName, ctv.TXT_FAILED))
	}

	return

}

//  Private Functions

// buildConnectionString - returns the connection string based on the PSQL configuration.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func buildConnectionString(config PSQLConnectionConfig) (dbConnString string) {

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
func getConnection(config PSQLConnectionConfig) (connectionPoolPtr *pgxpool.Pool, errorInfo errs.ErrorInfo) {

	var (
		tCACertPoolPtr *x509.CertPool
		tCert          tls.Certificate
		tConfigPtr     *pgxpool.Config
	)

	if tConfigPtr, errorInfo.Error = pgxpool.ParseConfig(buildConnectionString(config)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_PARSE_CONFIG, ctv.TXT_FAILED))
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

	if connectionPoolPtr, errorInfo.Error = pgxpool.NewWithConfig(CTXBackground, tConfigPtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_CONNECTION, ctv.TXT_FAILED))
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
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_TLS_CA_BUNDLE_FILENAME, tlsConfig.TLSCABundleFQN))
		return
	}

	caCertPoolPtr = x509.NewCertPool()
	if ok := caCertPoolPtr.AppendCertsFromPEM(tCABundleFile); !ok {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_TLS_CA_CERT_POOL, ctv.TXT_FAILED))
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
			errorInfo.Error, errs.BuildLabelSubLabelValueMessage(
				ctv.LBL_SERVICE_PSQL, ctv.LBL_TLS_CERTIFICATE_FILENAME, ctv.LBL_TLS_PRIVATE_KEY_FILENAME, ctv.VAL_EMPTY, tlsConfig.TLSPrivateKey,
			), // The tlsConfig.TLSCertFQN, tlsConfig.TLSPrivateKeyFQN values are not output for security.
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

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, configFilename, errs.ErrEmptyRequiredParameter, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	return
}

func validateConfig(config PSQLConfig) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.LBL_SERVICE_PSQL, config.DBName, errs.ErrEmptyPsqlDatabaseName, ctv.LBL_PSQL_DBNAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.Host, errs.ErrEmptyServerHostName, ctv.LBL_PSQL_HOST); errorInfo.Error != nil {
		return
	}
	if config.MaxConnections <= ctv.VAL_ZERO && config.MaxConnections >= ctv.VAL_ONE_HUNDRED {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidPSQLMaxConnections, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_MAX_CONNECTIONS, strconv.Itoa(config.Port)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.Host, errs.ErrEmptyPassword, ctv.LBL_PSQL_PASSWORD); errorInfo.Error != nil {
		return
	}
	if config.Port != ctv.VAL_PSQL_PORT {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidPSQLPort, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_PORT, strconv.Itoa(config.Port)))
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
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLSSLModeNotAllowed, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
		return
	case PSQL_SSL_MODE_VERIFY:
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSCABundleFQN, errs.ErrEmptyRequiredParameter, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
			return
		}
	case PSQL_SSL_MODE_VERIFY_FULL:
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSCABundleFQN, errs.ErrEmptyRequiredParameter, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSCertFQN, errs.ErrEmptyRequiredParameter, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSPrivateKeyFQN, errs.ErrEmptyRequiredParameter, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidPSQLTimeout, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
	}
	if config.Timeout < ctv.VAL_ONE && config.Timeout >= ctv.VAL_FIVE {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCTimeout, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TIMEOUT, strconv.Itoa(config.Timeout)))
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.UserName, errs.ErrEmptyUserName, ctv.LBL_PSQL_USER_NAME); errorInfo.Error != nil {
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

	var (
		tRows pgx.Rows
	)

	if connectionPtr == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrEmptyPointer, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_DBNAME, dbName))
		return
	}

	tRows, errorInfo.Error = connectionPtr.Query(CTXBackground, CHECK_STAT_ACTIVITY, dbName)
	if errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFailedPsqlConn, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_DBNAME, dbName))
	}
	defer tRows.Close()

	return
}
