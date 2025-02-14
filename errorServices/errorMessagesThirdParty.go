package sharedServices

//goland:noinspection ALL
const (
	//
	// String that are used to determine third party error messages
	USER_DOES_NOT_EXIST           = "User does not exist."
	NOT_FOUND                     = "not found"
	UNKNOWN                       = "UNKNOWN"
	JWT_TOKEN_SIGNATURE_INVALID   = "token signature is invalid: crypto/rsa: verification error"
	FIREBASE_AUTH_BAD_CREDENTIALS = "cannot read credentials"
	FIREBASE_AUTH_USER_NOT_FOUND  = "cannot find user from uid:"
	NATS_NOT_CONNECTED            = "nats: message is not bound to subscription/connection"
	NATS_INVALID_CONNECTION       = "nats: invalid connection"
	STRIPE_INVALID_API_KEY        = "Invalid API Key provided"
)
