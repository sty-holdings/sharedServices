package sharedServices

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
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
func NewPSQLServer(configFilename string, environment string) (servicePtr *PSQLService, errorInfo errs.ErrorInfo) {

	var (
		tConfig PSQLConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadPSQLConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateConfig(tConfig, environment); errorInfo.Error != nil {
		return
	}

	servicePtr = &PSQLService{
		debugModeOn: tConfig.Debug,
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

// BatchInsert - will insert up to 100 records. The role is optional.
//
//	Customer Messages: None
//	Errors: ErrEmptyRequiredParameter, returned by BeginTx, returned by Exec
//	Verifications: None
func (psqlServicePtr *PSQLService) BatchInsert(
	database string,
	role string,
	batchName string,
	insertStatement string,
	values [][]any,
) (errorInfo errs.ErrorInfo) {

	var (
		pCommandTag        pgconn.CommandTag
		pTransaction       pgx.Tx
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tQueueString       string
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, database, ctv.LBL_DATABASE); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, batchName, ctv.LBL_PSQL_BATCH); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, insertStatement, ctv.LBL_PSQL_INSERT); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.LBL_SERVICE_PSQL, values, ctv.LBL_VALUE); errorInfo.Error != nil {
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

	if psqlServicePtr.debugModeOn {
		fmt.Printf("%s function %s inserted %d records.\r", strings.ToUpper(ctv.VAL_SERVICE_PSQL), tFunctionName, len(values))
	}

	return
}

func (psqlServicePtr *PSQLService) CheckDuplicateKeyError(pResultsPtr *gorm.DB) bool {

	if psqlServicePtr.ConvertErrorCode(pResultsPtr) == errs.PSQL_DUPLICATE_KEY {
		return true
	}

	return false
}

// Close - shuts down all active connections in the ConnectionPoolPtrs, releasing resources.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) Close() {

	for _, connectionPtr := range psqlServicePtr.ConnectionPoolPtrs {
		connectionPtr.Close()
	}
}

// ConvertErrorCode - extracts and returns the PostgreSQL error code from a GORM error.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: vals.ctv.VAL_EMPTY
func (psqlServicePtr *PSQLService) ConvertErrorCode(pResultsPtr *gorm.DB) string {

	var (
		pgError *pgconn.PgError
	)

	if errors.As(pResultsPtr.Error, &pgError) {
		return pgError.Code
	}

	return ctv.VAL_EMPTY
}

// CommitRollbackTransaction - handles the commit or rollback of a GORM transaction based on the presence of errors in the transaction. Returns detailed error info on failure.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
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

func (psqlServicePtr *PSQLService) ExecuteSelectExist(database string, sqlStatement string) (exists bool, errorInfo errs.ErrorInfo) {

	errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].QueryRow(context.Background(), sqlStatement).Scan(&exists)

	return
}

func (psqlServicePtr *PSQLService) ExecuteSelectSingleRowInt(database string, sqlStatement string) (result int, errorInfo errs.ErrorInfo) {

	errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].QueryRow(context.Background(), sqlStatement).Scan(&result)

	return
}

func (psqlServicePtr *PSQLService) ExecuteSelectSingleRowText(database string, sqlStatement string) (result string, errorInfo errs.ErrorInfo) {

	errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].QueryRow(context.Background(), sqlStatement).Scan(&result)

	return
}

// ExecuteStaticSQL - will execute a static SQL statement in a transaction.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) ExecuteStaticSQL(database string, sqlStatement string, sqlType string) (rowsAffected uint64, errorInfo errs.ErrorInfo, resultPtr *gorm.DB) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTransactionPtr    *gorm.DB
	)

	if tTransactionPtr, errorInfo = psqlServicePtr.StartTransaction(ctv.VAL_EMPTY, database); errorInfo.Error != nil {
		return
	}
	if resultPtr = tTransactionPtr.Exec(sqlStatement); resultPtr.Error != nil {
		errorInfo = errs.NewErrorInfo(resultPtr.Error, errs.BuildLabelSubLabelValueMessage(strings.ToUpper(ctv.VAL_SERVICE_PSQL), sqlType, ctv.LBL_FUNCTION_NAME, tFunctionName, ctv.TXT_FAILED))
		if resultPtr = tTransactionPtr.Rollback(); resultPtr.Error != nil {
			errorInfo = errs.NewErrorInfo(
				errorInfo.Error,
				errs.BuildLabelSubLabelValueMessage(strings.ToUpper(ctv.VAL_SERVICE_PSQL), ctv.LBL_PSQL_ROLLBACK, ctv.LBL_FUNCTION_NAME, tFunctionName, ctv.TXT_FAILED),
			)
			errs.PrintErrorInfo(errorInfo)
		}
		return
	}
	rowsAffected = uint64(resultPtr.RowsAffected)
	if resultPtr = tTransactionPtr.Commit(); resultPtr.Error != nil {
		errorInfo = errs.NewErrorInfo(
			errorInfo.Error,
			errs.BuildLabelSubLabelValueMessage(strings.ToUpper(ctv.VAL_SERVICE_PSQL), ctv.LBL_PSQL_COMMIT, ctv.LBL_FUNCTION_NAME, tFunctionName, ctv.TXT_FAILED),
		)
		return
	}

	if psqlServicePtr.debugModeOn {
		fmt.Printf("%s function %s %s %d records.\r", strings.ToUpper(ctv.VAL_SERVICE_PSQL), tFunctionName, sqlType, rowsAffected)
	}

	return

}

// InsertUpdateUsingStaticSQL - executes an insert SQL and, on duplicate key error, executes an update SQL for a given database. Returns detailed error info on failure.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) InsertUpdateUsingStaticSQL(
	databaseName string,
	insertSQL string,
	updateSQL string,
) (rowsAffected uint64, sqlType string, errorInfo errs.ErrorInfo, resultPtr *gorm.DB) {

	sqlType = ctv.LBL_PSQL_INSERT
	if rowsAffected, errorInfo, resultPtr = psqlServicePtr.ExecuteStaticSQL(databaseName, insertSQL, sqlType); errorInfo.Error != nil {
		if psqlServicePtr.CheckDuplicateKeyError(resultPtr) {
			sqlType = ctv.LBL_PSQL_UPDATE
			if rowsAffected, errorInfo, resultPtr = psqlServicePtr.ExecuteStaticSQL(databaseName, updateSQL, sqlType); errorInfo.Error != nil {
				errorInfo = errs.NewErrorInfo(
					errorInfo.Error,
					errs.BuildLabelValueMessage(strings.ToUpper(ctv.VAL_SERVICE_PSQL), sqlType, ctv.VAL_EMPTY, ctv.TXT_FAILED),
				)
				return
			}
		} else {
			errorInfo = errs.NewErrorInfo(
				errorInfo.Error,
				errs.BuildLabelValueMessage(strings.ToUpper(ctv.VAL_SERVICE_PSQL), sqlType, ctv.VAL_EMPTY, ctv.TXT_FAILED),
			)
		}
	}

	return
}

// StartTransaction - initiates a GORM transaction for the specified database and batchName. Returns a transaction pointer and error information if the transaction fails to begin.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) StartTransaction(batchName string, database string) (transactionPtr *gorm.DB, errorInfo errs.ErrorInfo) {

	if transactionPtr = psqlServicePtr.GORMPoolPtrs[database].Begin(); transactionPtr.Error != nil {
		errorInfo = errs.NewErrorInfo(transactionPtr.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRANSACTION, ctv.LBL_PSQL_BATCH, batchName, ctv.TXT_FAILED))
	}

	return
}

// TruncateTable - clears all data from a specified table in the defined database and schema, while maintaining its structure. Returns an ErrorInfo object if the operation fails.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (psqlServicePtr *PSQLService) TruncateTable(database string, schema string, tableName string) (errorInfo errs.ErrorInfo) {

	var (
		pStatement         string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	pStatement = fmt.Sprintf(TRUNCATE_TABLE, pgx.Identifier{schema}.Sanitize(), pgx.Identifier{tableName}.Sanitize())
	if _, errorInfo.Error = psqlServicePtr.ConnectionPoolPtrs[database].Exec(CTXBackground, pStatement); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TRUNCATE, schema, tableName, ctv.TXT_FAILED))
	}

	if psqlServicePtr.debugModeOn {
		fmt.Printf("%s function %s truncated database %s table %s.%s.\r", strings.ToUpper(ctv.VAL_SERVICE_PSQL), tFunctionName, database, schema, tableName)
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
// DBName:   		Name of the Postgres database
// Host:     		Internet DNS or IP address of the server running the instance of Postgres
// Max Connections:	Must be greater than 0 and less than 100 across all instances.
// Password:        Encrypted password for authentication
// Port:            Interface the connection communicates with Postgres
// SSL Mode:        Which mode will be used to connect to the PSQL server.
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
		tDialector = postgres.New(
			postgres.Config{
				DSN:                  tConnectionString,
				PreferSimpleProtocol: false, // This is the default and provided here for documentation. Only change this to true is there are issues.
			},
		)

		if gormPoolPtr, errorInfo.Error = gorm.Open(
			tDialector, &gorm.Config{
				CreateBatchSize: 100,
				Logger:          logger.Default.LogMode(logger.Silent),
				PrepareStmt:     true,
			},
		); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_GORM_CONNECTION, ctv.TXT_FAILED))
			return
		}

		if config.GORM.LoggerOn {
			newLogger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // Slow SQL threshold
					LogLevel:      logger.Info, // Log level
					Colorful:      true,        // Enable color
				},
			)
			gormPoolPtr.Logger = newLogger
		}

		errorInfo = isGORMConnectionActive(gormPoolPtr)
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

// loadPSQLConfig - reads and returns a psql service pointer
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile
//	Verifications: validateConfiguration
func loadPSQLConfig(configFilename string) (config PSQLConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
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

// validateConfig - validates the given PSQLConfig object for completeness and correctness.
// Returns an ErrorInfo object containing error details if validation fails.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func validateConfig(config PSQLConfig, environment string) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckArrayLengthGTZero(ctv.LBL_SERVICE_PSQL, config.DBNames, ctv.LBL_PSQL_DBNAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.Host, ctv.LBL_PSQL_HOST); errorInfo.Error != nil {
		return
	}
	if config.MaxConnections <= ctv.VAL_ZERO && config.MaxConnections >= ctv.VAL_ONE_HUNDRED {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidPSQLMaxConnections, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_MAX_CONNECTIONS, strconv.Itoa(config.Port)))
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.Host, ctv.LBL_PSQL_PASSWORD); errorInfo.Error != nil {
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
		if environment == ctv.VAL_ENVIRONMENT_LOCAL {
			break
		}
		fallthrough
	case PSQL_SSL_MODE_PREFER:
		fallthrough
	case PSQL_SSL_MODE_REQUIRED:
		fallthrough
	case PSQL_SSL_MODE_VERIFY:
		errorInfo = errs.NewErrorInfo(errs.ErrPSQLSSLModeNotAllowed, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
		return
	case PSQL_SSL_MODE_VERIFY_FULL:
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSCABundleFQN, ctv.LBL_TLS_CA_BUNDLE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSCertFQN, ctv.LBL_TLS_CERTIFICATE_FILENAME); errorInfo.Error != nil {
			return
		}
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.PSQLTLSInfo.TLSPrivateKeyFQN, ctv.LBL_TLS_PRIVATE_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidPSQLTimeout, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_SSL_MODE, config.SSLMode))
	}
	if config.Timeout < ctv.VAL_ONE && config.Timeout >= ctv.VAL_FIVE {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidGRPCTimeout, errs.BuildLabelValue(ctv.LBL_SERVICE_PSQL, ctv.LBL_PSQL_TIMEOUT, strconv.Itoa(config.Timeout)))
	}
	errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_PSQL, config.UserName, ctv.LBL_PSQL_USER_NAME)

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
