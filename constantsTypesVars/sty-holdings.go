package sharedServices

//goland:noinspection All
const (
	// Text strings
	TXT_AI2C_KEY = "AI2C KEY: "
)

//goland:noinspection All
const (
	// Subjects
	NCI_PING           = "ping"
	NCI_TURN_DEBUG_OFF = "turn-debug-off"
	NCI_TURN_DEBUG_ON  = "turn-debug-on"
	//
	SUB_GEMINI_ANALYZE_QUESTION = "gemini.analyze-question"
	SUB_GEMINI_SUMMARIZE_ANSWER = "gemini.analyze-answer"
	//
	SUB_STRIPE_CANCEL_PAYMENT_INTENT  = "stripe.cancel-payment-intent"
	SUB_STRIPE_CONFIRM_PAYMENT_INTENT = "stripe.confirm-payment-intent"
	SUB_STRIPE_CREATE_PAYMENT_INTENT  = "stripe.create-payment-intent"
	SUB_STRIPE_LIST_PAYMENT_INTENTS   = "stripe.list-payment-intents"
	SUB_STRIPE_LIST_PAYMENT_METHODS   = "stripe.list-payment-methods"
	SUB_STRIPE_BALANCE                = "stripe.balance"
	SUB_STRIPE_CUSTOMERS              = "stripe.customers"
	//
	SUB_SENDGRID_SEND_EMAIL = "sendgrid.send-email"
	//
	SUB_SYNADIA_GET_PERSONAL_ACCESS_TOKEN   = "synadia.get-personal-access-token"
	SUB_SYNADIA_GET_SYSTEM                  = "synadia.get-system"
	SUB_SYNADIA_GET_SYSTEM_LIMITS           = "synadia.get-system-limits"
	SUB_SYNADIA_GET_TEAM                    = "synadia.get-team"
	SUB_SYNADIA_GET_TEAM_LIMITS             = "synadia.get-team-limits"
	SUB_SYNADIA_GET_VERSION                 = "synadia.get-version"
	SUB_SYNADIA_LIST_ACCOUNT                = "synadia.list-account"
	SUB_SYNADIA_LIST_INFO_APP_USERS_TEAM    = "synadia.list-info-app-users-team"
	SUB_SYNADIA_LIST_NATS_USERS             = "synadia.list-nats-users"
	SUB_SYNADIA_LIST_PERSONAL_ACCESS_TOKENS = "synadia.list-personal-access-tokens"
	SUB_SYNADIA_LIST_SYSTEMS                = "synadia.list-systems"
	SUB_SYNADIA_LIST_SYSTEM_ACCOUN_TINFO    = "synadia.list-system-account-info"
	SUB_SYNADIA_LIST_SYSTEM_SERVER_INFO     = "synadia.list-system-server-info"
	SUB_SYNADIA_LIST_TEAMS                  = "synadia.list-teams"
	SUB_SYNADIA_LIST_TEAM_SERVER_ACCOUNTS   = "synadia.list-team-server-accounts"
)

//goland:noinspection All
const (
	// Extensions
	EXTENSION_GEMINI    = "gemini"
	EXTENSION_HTTP_NATS = "http-nats"
	NC_INTERNAL         = "nc-internal"
	EXTENSION_STRIPE    = "stripe"
	EXTENSION_SENDGRID  = "sendgrid"
	EXTENSION_SYNADIA   = "synadia"
)

//goland:noinspection All
const (
	// Fully qualified filenames
	FQN_CERTIFICATE          = "CertificateFQN"
	FQN_FIREBASE_CREDENTIALS = "FirebaseCredentialsFQN"
	FQN_GCP_CREDENTIALS      = "GCPCredentialsFQN"
	FQN_NATS_CREDENTIALS     = "NATSCredsFQN"
	FQN_TLS_CABUNDLE         = "TLS:TLSCABundle"
	FQN_TLS_CERTIFICATE      = "TLS:TLSCert"
	FQN_TLS_PRIVATE_KEY      = "TLS:TLSKey"
)

//goland:noinspection All
const (
	// Testing
	TEST_POSITIVE_SUCCESS = "Positive Case: Successful: "
	TEST_NEGATIVE_SUCCESS = "Negative Case: Successful: "
)

//goland:noinspection All
const (
	// Software Values
	ATTEMPT_LIMIT_EXCEEDED  = "Attempt limit exceeded"
	AUTHENTICATOR_SERVICE   = "AuthenticatorService"
	DEFAULT_VERSION         = "9999.9999.9999"
	ENVIRONMENT_LOCAL       = "local"
	ENVIRONMENT_DEVELOPMENT = "development"
	ENVIRONMENT_DEMO        = "demo"
	ENVIRONMENT_PRODUCTION  = "production"
	IDP_FIREBASE            = "firebase"
	IDP_COGNITO             = "cognito"
	LOCAL_HOST              = "localhost"
	NATS_NON_TLS_CONNECTION = "NON-TLS"
	NATS_TLS_CONNECTION     = "TLS"
	OPER_DOUBE_EQUAL_SIGN   = "=="
	OPER_EQUAL_SIGN         = "="
	PID_FILENAME            = "server.pid"
	VAL_EMPTY               = ""
	VAL_ZERO                = 0
	VALID                   = "VALID"
)

//goland:noinspection All
const (
	// Special values used to trigger requests
	PAYMENT_METHOD_LIST = "list"
)
