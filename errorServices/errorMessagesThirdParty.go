package sharedServices

//goland:noinspection ALL
const (
	//
	// String that are used to determine third party error messages
	//
	// FIREBASE
	FIREBASE_AUTH_3P_BAD_CREDENTIALS = "cannot read credentials"
	FIREBASE_AUTH_3P_USER_NOT_FOUND  = "cannot find user from UID:"
	// GRPC
	GRPC_SHUTDOWN = "use of closed network connection"
	// JWT
	JWT_3P_TOKEN_SIGNATURE_INVALID = "token signature is invalid: crypto/rsa: verification error"
	// NATS
	NATS_3P_NOT_CONNECTED      = "nats: message is not bound to subscription/connection"
	NATS_3P_INVALID_CONNECTION = "nats: invalid connection"
	// STRIPE
	STRIPE_3P_INVALID_API_KEY = "Invalid API Key provided"
	STRIPE_NO_SUCH_RESOURCE   = "no such resource"
)

// POSTGRESQL
// Error Codes
//
//goland:noinspection all
const (
	PSQL_SUCCESS                        = "00000" // successful_completion
	PSQL_NO_DATA                        = "02000"
	PSQL_CONN_EXCEPTION                 = "08000" // connection_exception - General connection error.
	PSQL_CLIENT_CONN_EXCEPTION          = "08001" // sqlclient_unable_to_establish_sqlconnection - The client could not establish a connection to the server.
	PSQL_CONN_NOT_EXIST                 = "08003" // connection_does_not_exist - The connection has been lost.
	PSQL_CONN_FAILURE                   = "08006" // connection_failure - General connection failure.
	PSQL_PROTOCOL_VIOLATION             = "08P01" // protocol_violation - A protocol violation occurred.
	PSQL_INVALID_DATETIME               = "22007" // invalid_datetime_format
	PSQL_INTEGRITY_CONSTRAINT_VIOLATION = "23000" // integrity_constraint_violation
	PSQL_NOT_NULL_VIOLATIOIN            = "23502" // not_null_violation
	PSQL_FOREIGN_KEY_VIOLATIOIN         = "23503" // foreign_key_violation
	PSQL_UNIQUE_VIOLATION               = "23505" // unique_violation
	PSQL_DUPLICATE_KEY                  = "23505" // Duplicate key constriant violation
	PSQL_SYNTAX_ERROR                   = "42601" // syntax_error
	PSQL_UNDEFINED_COLUMN               = "42703" // undefined_column
	PSQL_UNDEFINED_TABLE                = "42P01" // undefined_table
	PSQL_INSUFFICIENT_RESOURCES         = "53000" // insufficient_resources
	PSQL_DISK_FULL                      = "53100" //  disk_full
	PSQL_OUT_OF_MEMORY                  = "53200" //  out_of_memory
	PSQL_TOO_MANY_CONNECTIONS           = "53300" //  too_many_connections
)
