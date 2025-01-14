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
)

//goland:noinspection All
const (
	COMBINATION_YEAR                        = "00001"
	COMBINATION_YEAR_QUARTER                = "00011"
	COMBINATION_YEAR_QUARTER_MONTH          = "00111"
	COMBINATION_YEAR_QUARTER_MONTH_WEEK     = "01111"
	COMBINATION_YEAR_QUARTER_MONTH_WEEK_DAY = "11111"
	COMBINATION_QUARTER                     = "00010"
	COMBINATION_MONTH                       = "00100"
	COMBINATION_WEEK                        = "01000"
	COMBINATION_DAY                         = "10000"
)

type TimePeriodSpecialWordsPresent struct {
	Current      bool `json:"current"`
	Details      bool `json:"details"`
	Last         bool `json:"last"`
	Next         bool `json:"next"`
	Previous     bool `json:"previous"`
	SubTotal     bool `json:"subtotal"`
	Today        bool `json:"today"`
	Transactions bool `json:"transactions"`
}

type TimePeriodWordsPresent struct {
	Year    bool `json:"year"`
	Quarter bool `json:"quarter"`
	Month   bool `json:"month"`
	Week    bool `json:"week"`
	Day     bool `json:"day"`
}
