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
	// POSTGRESQL
	// Error Codes
	PSQL_ERROR_DUPLICATE_KEY = "SQLSTATE 23505"
	// STRIPE
	STRIPE_3P_INVALID_API_KEY = "Invalid API Key provided"
)
