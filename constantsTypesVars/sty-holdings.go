package sharedServices

import (
	"cloud.google.com/go/vertexai/genai"

	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

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
	ANS_DATE_TIME_FORMAT = "The date/time in your location (%s) is %s"
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
	COMBINATION_DAY                         = "10000"
	COMBINATION_MONTH                       = "00100"
	COMBINATION_NONE                        = "00000"
	COMBINATION_QUARTER                     = "00010"
	COMBINATION_WEEK                        = "01000"
	COMBINATION_YEAR                        = "00001"
	COMBINATION_YEAR_QUARTER                = "00011"
	COMBINATION_YEAR_QUARTER_MONTH          = "00111"
	COMBINATION_YEAR_QUARTER_MONTH_WEEK     = "01111"
	COMBINATION_YEAR_QUARTER_MONTH_WEEK_DAY = "11111"
)

type CategoryPromptComparison struct {
	AdverbBySubject    map[string]int      `json:"adverb_by_subject"`
	Category           []string            `json:"category"`
	CompoundQuestion   bool                `json:"compound_question"`
	ComparisonQuestion bool                `json:"comparison_question"`
	CountBySubject     map[string]int      `json:"count_by_subject"`
	Prompt             string              `json:"prompt"`
	QuestionSubject    []string            `json:"question_subject"`
	SaasProvider       []string            `json:"saas_provider"`
	Total              bool                `json:"total"`
	TokenCount         genai.UsageMetadata `json:"-"`
}

type CategoryPromptComparisonWithResponse struct {
	Response struct {
		AdverbBySubject    map[string]int      `json:"adverb_by_subject"`
		Category           []string            `json:"category"`
		CompoundQuestion   bool                `json:"compound_question"`
		ComparisonQuestion bool                `json:"comparison_question"`
		CountBySubject     map[string]int      `json:"count_by_subject"`
		Prompt             string              `json:"prompt"`
		QuestionSubject    []string            `json:"question_subject"`
		SaasProvider       []string            `json:"saas_provider"`
		Total              bool                `json:"total"`
		TokenCount         genai.UsageMetadata `json:"-"`
	} `json:"response"`
}

type SaaSResponse struct {
	SaaSData  string
	ErrorInfo errs.ErrorInfo
}

type TimePeriodSpecialWordsPresent struct {
	Current      bool                `json:"current"`
	Details      bool                `json:"details"`
	Last         bool                `json:"last"`
	Next         bool                `json:"next"`
	Previous     bool                `json:"previous"`
	SubTotal     bool                `json:"subtotal"`
	Today        bool                `json:"today"`
	Transactions bool                `json:"transactions"`
	TokenCount   genai.UsageMetadata `json:"-"`
}

type TimePeriodWordsPresent struct {
	Year       bool                `json:"year"`
	Quarter    bool                `json:"quarter"`
	Month      bool                `json:"month"`
	Week       bool                `json:"week"`
	Day        bool                `json:"day"`
	TokenCount genai.UsageMetadata `json:"-"`
}

type TimePeriodValues struct {
	Years      []int               `json:"years"`
	Quarters   []int               `json:"quarters"`
	Months     []int               `json:"months"`
	Weeks      []int               `json:"weeks"`
	Days       []int               `json:"days"`
	TokenCount genai.UsageMetadata `json:"-"`
}

type UserInfo struct {
	KeyB64       string
	STYHClientId string
	UId          string
}
