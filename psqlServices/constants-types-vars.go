package sharedServices

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

//goland:noinspection ALL
const (
	PSQL_SSL_MODE_DISABLE     = "disable"
	PSQL_SSL_MODE_ALLOW       = "allow"
	PSQL_SSL_MODE_PREFER      = "prefer"
	PSQL_SSL_MODE_REQUIRED    = "require"
	PSQL_SSL_MODE_VERIFY      = "verify-ca"
	PSQL_SSL_MODE_VERIFY_FULL = "verify-full"
	//
	PSQL_CONN_STRING = "dbname=%s host=%s pool_max_conns=%d password=%s port=%d sslmode=%s connect_timeout=%d user=%s"
	//
	SET_ROLE       = "SET ROLE %s;\n"
	TRUNCATE_TABLE = "TRUNCATE TABLE %s.%s;\n"
)

type PSQLConfig struct {
	DBName         []string     `json:"psql_db_names" yaml:"psql_db_names"`
	Debug          bool         `json:"psql_debug" yaml:"psql_debug"`
	Host           string       `json:"psql_host" yaml:"psql_host"`
	MaxConnections int          `json:"pgsql_max_connections" yaml:"pgsql_max_connections"`
	Password       string       `json:"psql_password" yaml:"psql_password"`
	Port           int          `json:"psql_port" yaml:"psql_port"`
	SSLMode        string       `json:"psql_ssl_mode" yaml:"psql_ssl_mode"`
	PSQLTLSInfo    jwts.TLSInfo `json:"psql_tls_info" yaml:"psql_tls_info"`
	Timeout        int          `json:"psql_timeout" yaml:"psql_timeout"`
	UserName       string       `json:"psql_user_name" yaml:"psql_user_name"`
}

type PSQLConnectionConfig struct {
	DBName         string
	Debug          bool
	Host           string
	MaxConnections int
	Password       string
	Port           int
	SSLMode        string
	PSQLTLSInfo    jwts.TLSInfo
	Timeout        int
	UserName       string
}

type PSQLService struct {
	DebugOn            bool
	ConnectionPoolPtrs map[string]*pgxpool.Pool
}

// Row and Rows are so pgx package doesn't need to be imported everywhere there are queries to the database.
type Transaction pgx.Tx
type Rows pgx.Rows
type Row pgx.Row
