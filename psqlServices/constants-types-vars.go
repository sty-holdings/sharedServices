package sharedServices

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

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
	TRUNCATE_TABLE           = "truncate table %s.%s"                                                                                                                                                                                                                                                                     // Disregard the IDE error warning. If a bug.
	INSERT_DAILY_PERFORMANCE = `INSERT INTO dkga.daily_performance (campaign_id, campaign_type, campaign_name, date, clicks, impressions, ctr, cpc, spend, cpm, cost_per_conversion, conversion_rate, conversion_value
								) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);` // Disregard the IDE error warning. If a bug.
)

type PSQLConfig struct {
	DBName         string       `json:"psql_db_name" yaml:"psql_db_name"`
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

type PSQLService struct {
	DebugOn           bool
	ConnectionPoolPtr *pgxpool.Pool
}

// Row and Rows are so pgx package doesn't need to be imported everywhere there are queries to the database.
type Transaction pgx.Tx
type Rows pgx.Rows
type Row pgx.Row
