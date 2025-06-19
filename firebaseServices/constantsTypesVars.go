package sharedServices

// Do not move these to the JWT service. Could introduce circular reference.
//
//goland:noinspection ALL
const (
	JWT_PAYLOAD_SUBJECT_FN      = "SUBJECT"
	JWT_PAYLOAD_CLAIMS_FN       = "CLAIMS"
	JWT_PAYLOAD_AUDIENCE_FN     = "AUDIENCE"
	JWT_PAYLOAD_REQUESTOR_ID_FN = "REQUESTOR_ID"
	JWT_PAYLOAD_EXPIRES_FN      = "EXPIRES"
	JWT_PAYLOAD_ISSUER_FN       = "ISSUER"
	JWT_PAYLOAD_ISSUED_AT_FN    = "ISSUED_AT"
)

//goland:noinspection All
const (
	DATASTORE_ACCESS_CODE = "access"
	DATASTORE_CLIENTS     = "clients"
	DATASTORE_USERS       = "users"
)
