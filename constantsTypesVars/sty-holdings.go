package sharedServices

import (
	"time"

	"cloud.google.com/go/vertexai/genai"
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
	IDP_FIREBASE            = "firebase"
	IDP_COGNITO             = "cognito"
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
	COMBINATION_YEAR_MONTH                  = "00101"
	COMBINATION_YEAR_MONTH_DAY              = "10101"
)

type AnalyzedQuestion struct {
	AnalysisId       string           `json:"AnalysisId"`
	UserQuestion     string           `json:"user_question"`
	CategorySentence CategorySentence `json:"category_sentence"`
	SpecialWords     SpecialWords     `json:"special_words"`
	TimePeriodValues TimePeriodValues `json:"q_time_period_values"`
}

type CategorySentence struct {
	AIPrompt       string   `json:"ai_prompt"`
	Categories     []string `json:"categories"`
	CountBySubject []struct {
		Subject string `json:"subject"`
		Count   string `json:"count"`
	} `json:"count_by_subject"`
	Prompt                 string   `json:"prompt"`
	SentenceSubjects       []string `json:"sentence_subjects"`
	SentenceSubjectAdverbs []struct {
		Subject string `json:"subject"`
		Adverb  string `json:"adverb"`
	} `json:"sentence_subjects_adverb"`
	TokenCount genai.UsageMetadata `json:"-"`
}

type ExtractDateParts struct {
	DateString             string
	Day                    int
	StartDateTime          time.Time
	EndDateTime            time.Time
	GreaterThanDateTimeInt int64 // Unix epoch using the process date minus one day with timezone
	LesserThanDateTimeInt  int64 // Unix epoch using the process date plus one day with timezone
	Month                  int
	Quarter                int
	WeekStart              string
	WeekEnd                string
	Year                   int
}

type SpecialWords struct {
	AverageFlag     bool                `json:"average"`
	CompoundFlag    bool                `json:"compound"`
	ComparisonFlag  bool                `json:"comparison"`
	CountFlag       bool                `json:"count"`
	DetailFlag      bool                `json:"detail"`
	ForecastFlag    bool                `json:"forecast"`
	MaximumFlag     bool                `json:"maximum"`
	MinimumFlag     bool                `json:"minimum"`
	PercentageFlag  bool                `json:"percentage"`
	PredictionFlag  bool                `json:"predict"`
	RecommendFlag   bool                `json:"recommend"`
	ReportFlag      bool                `json:"report"`
	SubTotalFlag    bool                `json:"subtotal"`
	SummaryFlag     bool                `json:"summary"`
	TransactionFlag bool                `json:"transaction"`
	TotalFlag       bool                `json:"total"`
	TrendFlag       bool                `json:"trend"`
	WeekFlag        bool                `json:"week"`
	TokenCount      genai.UsageMetadata `json:"-"`
}

type TimePeriodValues struct {
	Years        []int               `json:"year_values"`
	Quarters     []int               `json:"quarter_values"`
	Months       []int               `json:"month_values"`
	Weeks        []int               `json:"week_values"`
	Days         []int               `json:"day_values"`
	ToDate       bool                `json:"to_date"`
	RelativeTime string              `json:"relative_time"`
	SundayDate   []string            `json:"sunday_date"`
	TokenCount   genai.UsageMetadata `json:"-"`
}

type UserInfo struct {
	KeyB64           string
	internalClientID string
	internalUserID   string
}
