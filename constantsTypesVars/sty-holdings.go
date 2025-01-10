package sharedServices

//goland:noinspection All
const (
	// Text strings
	TXT_AI2C_KEY = "AI2C KEY: "
)

//goland:noinspection All
const (
	// Fully qualified filenames
	FQN_CERTIFICATE      = "CertificateFQN"
	FQN_GCP_CREDENTIALS  = "GCPCredentialsFQN"
	FQN_NATS_CREDENTIALS = "NATSCredsFQN"
	FQN_TLS_CABUNDLE     = "TLS:TLSCABundle"
	FQN_TLS_CERTIFICATE  = "TLS:TLSCert"
	FQN_TLS_PRIVATE_KEY  = "TLS:TLSKey"
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
)

//goland:noinspection All
const (
	// Special values used to trigger requests
	PAYMENT_METHOD_LIST = "list"
)

//goland:noinspection All
const (
	INSTR_CATEGORY_PROMPT_COMPARISON_NAME        = "CategoryPromptComparison"
	INSTR_TIME_PERIOD_VALUES_NAME                = "TimePeriodValues"
	INSTR_TIME_PERIOD_WORDS_PRESENT_NAME         = "TimePeriodWordsPresent"
	INSTR_TIME_PERIOD_SPECIAL_WORDS_PRESENT_NAME = "TimePeriodSpecialWordsPresent"
)

//goland:noinspection All
const (
	ANS_DATE_TIME_FORMAT = "The time for %s is %s"
)

//goland:noinspection All
const (
	// Define flag positions
	FLAG_YEARS = iota
	FLAG_QUARTERS
	FLAG_MONTHS
	FLAG_WEEKS
	FLAG_DAYS
	YEAR_QUARTER                = 1<<FLAG_YEARS | 1<<FLAG_QUARTERS
	YEAR_QUARTER_MONTH          = 1<<FLAG_YEARS | 1<<FLAG_QUARTERS | 1<<FLAG_MONTHS
	YEAR_QUARTER_MONTH_WEEK     = 1<<FLAG_YEARS | 1<<FLAG_QUARTERS | 1<<FLAG_MONTHS | 1<<FLAG_WEEKS
	YEAR_QUARTER_MONTH_WEEK_DAY = 1<<FLAG_YEARS | 1<<FLAG_QUARTERS | 1<<FLAG_MONTHS | 1<<FLAG_WEEKS | 1<<FLAG_DAYS
)

type TimePeriodSpecialWordsPresent struct {
	Current      bool `json:"current"`
	Today        bool `json:"today"`
	LastPrevious bool `json:"last_previous"`
	Next         bool `json:"next"`
}

type TimePeriodWordsPresent struct {
	Year    bool `json:"year"`
	Quarter bool `json:"quarter"`
	Month   bool `json:"month"`
	Week    bool `json:"week"`
	Day     bool `json:"day"`
}
