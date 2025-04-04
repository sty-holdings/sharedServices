package sharedServices

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

var (
	CTXBackground = context.Background()
)

// NewPSQLServer - builds a reusable PGX or GORM PostgreSQL Service to access Postgres databases.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewPSQLServer(configFilename string) (servicePtr *PSQLService, errorInfo errs.ErrorInfo) {

	var (
		tConfig PSQLConfig
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
		DebugOn: tConfig.Debug,
	}
	servicePtr.GORMPoolPtrs = make(map[string]*gorm.DB)
	servicePtr.ConnectionPoolPtrs = make(map[string]*pgxpool.Pool)

	for _, databaseName := range tConfig.DBNames {
		if servicePtr.ConnectionPoolPtrs[databaseName], servicePtr.GORMPoolPtrs[databaseName], errorInfo = getConnection(tConfig, databaseName); errorInfo.Error != nil {
			return
		}
	}

	return
}

// BatchInsert - will insert upto 100 records. The role is optional.
//
//	Customer Messages: None
//	Errors: ErrEmptyRequiredParameter, returned by BeginTx, returned by Exec
//	Verifications: None
func (psqlServicePtr *PSQLService) BatchInsert(database string, role string, batchName string, insertStatement string, values [][]any) (errorInfo errs.ErrorInfo) {

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

	if pTransaction, errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].BeginTx(
		CTXBackground,
		pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadWrite},
	); errorInfo.Error != nil {
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

func (psqlServicePtr *PSQLService) Close() {

	for _, connectionPtr := range psqlServicePtr.ConnectionPoolPtrs {
		connectionPtr.Close()
	}
}

func (psqlServicePtr *PSQLService) CommitRollbackTransaction(batchName string, transactionPtr *gorm.DB) (errorInfo errs.ErrorInfo) {

	var (
		tResultsPtr *gorm.DB
	)

	if transactionPtr.Error != nil {
		if tResultsPtr = transactionPtr.Rollback(); tResultsPtr.Error != nil {
			errorInfo = errs.NewErrorInfo(
				transactionPtr.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRANSACTION, ctv.LBL_PSQL_ROLLBACK, batchName, ctv.TXT_FAILED),
			)
		}
		return
	}
	if errorInfo.Error = transactionPtr.Commit().Error; errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRANSACTION, ctv.LBL_PSQL_COMMIT, batchName, ctv.TXT_FAILED))
	}

	return
}

func (psqlServicePtr *PSQLService) StartTransaction(batchName string, database string) (transactionPtr *gorm.DB, errorInfo errs.ErrorInfo) {

	if transactionPtr = psqlServicePtr.GORMPoolPtrs[database].Begin(); transactionPtr.Error != nil {
		errorInfo = errs.NewErrorInfo(transactionPtr.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRANSACTION, ctv.LBL_PSQL_BATCH, batchName, ctv.TXT_FAILED))
	}

	return
}

func (psqlServicePtr *PSQLService) TruncateTable(database string, schema string, tableName string) (errorInfo errs.ErrorInfo) {

	var (
		pStatement string
	)

	pStatement = fmt.Sprintf(TRUNCATE_TABLE, pgx.Identifier{schema}.Sanitize(), pgx.Identifier{tableName}.Sanitize())
	if _, errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].Exec(CTXBackground, pStatement); errorInfo.Error != nil {
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
func buildConnectionString(config PSQLConfig, databaseName string) string {

	return fmt.Sprintf(
		PSQL_CONN_STRING,
		databaseName,
		config.Host,
		config.Password,
		config.Port,
		config.SSLMode,
		config.PSQLTLSInfo.TLSCABundleFQN,
		config.PSQLTLSInfo.TLSCertFQN,
		config.PSQLTLSInfo.TLSPrivateKeyFQN,
		config.Timeout,
		config.UserName,
	)
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
func getConnection(config PSQLConfig, databaseName string) (connectionPoolPtr *pgxpool.Pool, gormPoolPtr *gorm.DB, errorInfo errs.ErrorInfo) {

	var (
		newLogger         logger.Interface
		tConfigPtr        *pgxpool.Config
		tConnectionString string
		tDialector        gorm.Dialector
	)

	tConnectionString = buildConnectionString(config, databaseName)

	if config.GORM.UseGorm {
		if config.GORM.LoggerOn {
			newLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // Slow SQL threshold
					LogLevel:      logger.Info, // Log level
					Colorful:      true,        // Enable color
				},
			)
		}

		tDialector = postgres.New(
			postgres.Config{
				DSN:                  tConnectionString,
				PreferSimpleProtocol: false, // This is the default and provided here for documentation. Only change to true is there are issues.
			},
		)

		if gormPoolPtr, errorInfo.Error = gorm.Open(
			tDialector, &gorm.Config{
				CreateBatchSize: 100,
				Logger:          newLogger,
				PrepareStmt:     true,
			},
		); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_GORM_CONNECTION, ctv.TXT_FAILED))
			return
		}

		errorInfo = isGORMConnectionActive(gormPoolPtr)

		return
	}

	if tConfigPtr, errorInfo.Error = pgxpool.ParseConfig(tConnectionString); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_PARSE_CONFIG, ctv.TXT_FAILED))
		return
	}

	if connectionPoolPtr, errorInfo.Error = pgxpool.NewWithConfig(CTXBackground, tConfigPtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_CONNECTION, ctv.TXT_FAILED))
		return
	}

	errorInfo = isConnectionActive(connectionPoolPtr, databaseName)

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

	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.LBL_SERVICE_PSQL, config.DBNames, errs.ErrEmptyPsqlDatabaseName, ctv.LBL_PSQL_DBNAME); errorInfo.Error != nil {
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
		fallthrough
	case PSQL_SSL_MODE_VERIFY:
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLSSLModeNotAllowed, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
		return
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

// isGORMConnectionActive - checks the GORM connection by querying the database.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func isGORMConnectionActive(connectionPtr *gorm.DB) (errorInfo errs.ErrorInfo) {

	var (
		sqlDB *sql.DB
	)

	if sqlDB, errorInfo.Error = connectionPtr.DB(); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_GORM_CONNECTION, ctv.TXT_FAILED))
		return
	}

	if errorInfo.Error = sqlDB.Ping(); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_GORM_CONNECTION, ctv.TXT_FAILED))
	}

	return
}
