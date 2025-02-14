package sharedServices

//goland:noinspection ALL
const (
	//
	// String that are used to determine third party error messages
	JWT_3P_TOKEN_SIGNATURE_INVALID   = "token signature is invalid: crypto/rsa: verification error"
	FIREBASE_AUTH_3P_BAD_CREDENTIALS = "cannot read credentials"
	FIREBASE_AUTH_3P_USER_NOT_FOUND  = "cannot find user from uid:"
	NATS_3P_NOT_CONNECTED            = "nats: message is not bound to subscription/connection"
	NATS_3P_INVALID_CONNECTION       = "nats: invalid connection"
	STRIPE_3P_INVALID_API_KEY        = "Invalid API Key provided"
)
