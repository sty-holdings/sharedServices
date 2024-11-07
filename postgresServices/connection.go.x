package sty_shared

import (
	"context"
	"fmt"
	"strings"


	pi "github.com/sty-holdings/sharedServices/v2024/programInfo"
)

type ConnValues struct {
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Timeout  int    `json:"timeout"`
	SSLMode  string `json:"sslMode"`
}

var (
	CTXBackground = context.Background()
)

// Row and Rows are so pgx package doesn't need to be imported everywhere there are queries to the database.
type Transaction pgx.Tx
type Rows pgx.Rows
type Row pgx.Row

// GetPostgresConnection will create a connection to a database.
// dbName   Name of the Postgres database
// user     Account that GetPostgresConnection will use to authenticate
// password Users password for authentication
// host     Internet DNS or IP address of the server running the instance of Postgres
// sslMode  Type of encryption used for the connection (https://www.postgresql.org/docs/12/libpq-ssl.html for version 12)
// port     Interface the connection communicates with Postgres
// timeout  Number of seconds a request must complete (3 seconds is normal setting)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetPostgresConnection(
	dbName, user, password, host, sslMode string,
	port, timeout, poolMaxConns int,
) (
	DBPoolPtr *pgxpool.Pool,
	errorInfo pi.ErrorInfo,
) {

	if dbName == ctv.EMPTY || user == ctv.EMPTY || password == ctv.EMPTY || host == ctv.EMPTY || coreValidators.IsPostgresSSLModeValid(sslMode) {
		errorInfo.Error = pi.ErrRequiredArgumentMissing
		pi.PrintError(errorInfo)
		return
	}

	if port == 0 || timeout == 0 || poolMaxConns == 0 {
		errorInfo.Error = pi.ErrRequiredArgumentMissing
		pi.PrintError(errorInfo)
		return
	}

	if coreValidators.IsPostgresSSLModeValid(sslMode) == false {
		errorInfo.Error = pi.ErrPostgresSSLMode
		pi.PrintError(errorInfo)
		return
	}

	if DBPoolPtr, errorInfo.Error = pgxpool.New(
		context.Background(),
		setConnectionValues(dbName, user, password, host, sslMode, port, timeout, poolMaxConns),
	); errorInfo.Error != nil {
		if strings.Contains(errorInfo.Error.Error(), "dial") {
			errorInfo.Error = pi.ErrPostgresConnFailed
			pi.PrintError(errorInfo)
		} else {
			errorInfo.Error = pi.ErrServiceFailedPOSTGRES
			pi.PrintError(errorInfo)
		}
	}

	return
}

func setConnectionValues(
	dbName, user, password, host, sslMode string,
	port, timeout, poolMaxConns int,
) (dbConnString string) {

	return fmt.Sprintf(ctv.POSTGRES__CONN_STRING, dbName, user, password, host, port, timeout, sslMode, poolMaxConns)

}

// Verify that the pointer to the database connection is active.
func VerifyConnection(
	dbPoolPtr *pgxpool.Pool,
	dbName string,
) (errorInfo pi.ErrorInfo) {

	var (
		tRows Rows
	)

	if dbPoolPtr == nil {
		errorInfo.Error = pi.ErrPostgresConnEmpty
		pi.PrintError(errorInfo)
	} else {
		qStmt := "SELECT * FROM pg_stat_activity WHERE datname = $1 and state = 'active';"
		tRows, errorInfo.Error = dbPoolPtr.Query(CTXBackground, qStmt, dbName)
		if errorInfo.Error != nil {
			errorInfo.Error = pi.ErrPostgresConnFailed
			pi.PrintError(errorInfo)
		}
		defer tRows.Close()
	}

	return
}
