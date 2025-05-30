package sharedServices

// Entries are for comparison or setting arguments, parameters, and variables. If the file gets too large, split
// out larger sections to topic-specific files.

//goland:noinspection All
const (
	VAL_ALL                      = "all"
	VAL_BUSINESS                 = "business"
	VAL_CHARGE                   = "charge"
	VAL_CREATE                   = "create"
	VAL_DETERMINE_SUBJECTS       = "determine-subjects"
	VAL_EMPTY                    = ""
	VAL_END_DAY                  = "23:59:59"
	VAL_ENVIRONMENT_SHORT_DEV    = "dev"
	VAL_ENVIRONMENT_SHORT_PROD   = "prod"
	VAL_INDIVIDUAL               = "individual"
	VAL_INVALID_FLAG_SETTING     = "00000"
	VAL_LOGIN                    = "login"
	VAL_MID_NIGHT                = "23:59:59"
	VAL_NA                       = "N/A"
	VAL_NIL                      = "nil"
	VAL_NONE                     = "none"
	VAL_NOT_FOUND                = "not found"
	VAL_NULL                     = "null"
	VAL_PAYOUT                   = "payout"
	VAL_START_DAY                = "00:00:00"
	VAL_START_MONTH              = "-01 00:00:00"
	VAL_TEST                     = "TEST"
	VAL_UNKNOWN                  = "UNKNOWN"
	VAL_USER_DOES_NOT_EXIST      = "User does not exist."
	VAL_VALID                    = "VALID"
	VAL_QUARTER_ONE_START_DATE   = "01-01"
	VAL_QUARTER_ONE_END_DATE     = "03-31"
	VAL_QUARTER_TWO_START_DATE   = "04-01"
	VAL_QUARTER_TWO_END_DATE     = "06-30"
	VAL_QUARTER_THREE_START_DATE = "07-01"
	VAL_QUARTER_THREE_END_DATE   = "09-30"
	VAL_QUARTER_FOUR_START_DATE  = "10-01"
	VAL_QUARTER_FOUR_END_DATE    = "12-31"
)

//goland:noinspection All
const (
	VAL_DAVEKNOWS_NET                = "daveknows.net"
	VAL_ENVIRONMENT_LOCAL            = "local"
	VAL_ENVIRONMENT_DEVELOPMENT      = "development"
	VAL_ENVIRONMENT_PRODUCTION       = "production"
	VAL_ENVIRONMENT_SHORT_CODE_LOCAL = "local"
	VAL_ENVIRONMENT_SHORT_CODE_DEV   = "dev"
	VAL_ENVIRONMENT_SHORT_CODE_PROD  = "prod"
	VAL_GRPC_MAX_PORT                = 50151
	VAL_GRPC_MIN_PORT                = 50051
	VAL_LOCAL_HOST                   = "localhost"
	VAL_NATS                         = "nats"
	VAL_PSQL_PORT                    = 5432
	VAL_TCP                          = "tcp"
)

//goland:noinspection All
const (
	TIME_ONE_HOUR_IN_SECONDS     = 3600
	TIME_FIVE_MINUTES_IN_SECONDS = 300
)

//goland:noinspection All
const (
	VAL_PULL_DAILY      = "daily"
	VAL_PULL_HOURLY     = "hourly"
	VAL_PULL_30_MINUTES = "30-mins"
	VAL_PULL_15_MINUTES = "15-mins"
	VAL_PULL_10_MINUTES = "10-mins"
	VAL_PULL_5_MINUTES  = "5-mins"
)

//goland:noinspection All
const (
	VAL_ZERO         = 0
	VAL_ONE          = 1
	VAL_TWO          = 2
	VAL_THREE        = 3
	VAL_FOUR         = 4
	VAL_FIVE         = 5
	VAL_SIX          = 6
	VAL_SEVEN        = 7
	VAL_EIGHT        = 8
	VAL_NINE         = 9
	VAL_TEN          = 10
	VAL_ELEVEN       = 11
	VAL_TWELVE       = 12
	VAL_THIRTEEN     = 13
	VAL_FOURTEEN     = 14
	VAL_FIFTEEN      = 15
	VAL_SIXTEEN      = 16
	VAL_SEVENTEEN    = 17
	VAL_EIGHTEEN     = 18
	VAL_NINETEEN     = 19
	VAL_TWENTY       = 20
	VAL_TWENTY_ONE   = 21
	VAL_TWENTY_TWO   = 22
	VAL_TWENTY_THREE = 23
	VAL_TWENTY_FOUR  = 24
	VAL_TWENTY_FIVE  = 25
	VAL_TWENTY_SIX   = 26
	VAL_TWENTY_SEVEN = 27
	VAL_TWENTY_EIGHT = 28
	VAL_TWENTY_NINE  = 29
	VAL_THIRTY       = 30
	VAL_THIRTY_ONE   = 31
	VAL_THIRTY_TWO   = 32
	VAL_THIRTY_THREE = 33
	VAL_THIRTY_FOUR  = 34
	VAL_THIRTY_FIVE  = 35
	VAL_THIRTY_SIX   = 36
	VAL_THIRTY_SEVEN = 37
	VAL_THIRTY_EIGHT = 38
	VAL_THIRTY_NINE  = 39
	VAL_FOURTY       = 40
	VAL_FIFTY        = 50
	VAL_SIXTY        = 60
	VAL_SEVENTY      = 70
	VAL_EIGHTY       = 80
	VAL_NINETY       = 90
	VAL_ONE_HUNDRED  = 100
)

//goland:noinspection All
const (
	// Extensions
	VAL_EXTENSION_ADMIN      = "admin"
	VAL_EXTENTION_AGGREGATOR = "aggregator"
	VAL_EXTENSION_DISCOVERY  = "discovery"
	VAL_EXTENSION_DK_CLIENT  = "dk-client"
	VAL_EXTENSION_DK_SIGNAL  = "dk-signal"
	VAL_EXTENSION_EXTRACTOR  = "extractor"
	VAL_EXTENSION_HAL        = "hal"
	VAL_EXTENSION_UTILITY    = "utility"

	// Not an extension, but the owner of the extension.
	VAL_SERVER = "server"
)

//goland:noinspection ALL
const (
	// output Modes
	VAL_MIME_TYPE_HTML  = "text/html"
	VAL_MIME_TYPE_PLAIN = "text/plain"
)

//goland:noinspection All
const (
	// Services
	VAL_SERVICE_AI           = "ai"
	VAL_SERVICE_CLIENT       = "client"
	VAL_SERVICE_DEBUG        = "debug"
	VAL_SERVICE_ERROR        = "error"
	VAL_SERVICE_FIREBASE     = "firebase"
	VAL_SERVICE_FIRESTORE    = "firestore"
	VAL_SERVICE_GCP          = "gcp"
	VAL_SERVICE_GOOGLE_ADS   = "google-ads"
	VAL_SERVICE_GRPC_SERVER  = "grpc-server"
	VAL_SERVICE_GRPC_CLIENT  = "grpc-client"
	VAL_SERVICE_HELPERS      = "helpers"
	VAL_SERVICE_HTTP_SERVER  = "http-server"
	VAL_SERVICE_JWT          = "jwt"
	VAL_SERVICE_LINKEDIN     = "linkedin"
	VAL_SERVICE_OS           = "operating-system"
	VAL_SERVICE_PAYPAL       = "paypal"
	VAL_SERVICE_PROGRAM_INFO = "program-info"
	VAL_SERVICE_PSQL         = "psql"
	VAL_SERVICE_SENDGRID     = "sendgrid"
	VAL_SERVICE_STRIPE       = "stripe"
	VAL_SERVICE_VALIDATORS   = "validators"
)

//goland:noinspection All
const (
	// System Actions
	VAL_SYSTEM_ACTION_CONFIG_NEW_USER = "config-new-user"
	VAL_SYSTEM_ACTION_GET_MY_ANSWERS  = "get-my-answers"
	VAL_SYSTEM_ACTION_PING            = "ping"
	VAL_SYSTEM_ACTION_PULL_DATA       = "pull-data"
	VAL_SYSTEM_ACTION_TURN_DEBUG_OFF  = "turn-debug-off"
	VAL_SYSTEM_ACTION_TURN_DEBUG_ON   = "turn-debug-on"
	//
)

var (
	ExtensionList = []string{
		VAL_EXTENSION_ADMIN,
		VAL_EXTENTION_AGGREGATOR,
		VAL_EXTENSION_DISCOVERY,
		VAL_EXTENSION_DK_CLIENT,
		VAL_EXTENSION_DK_SIGNAL,
		VAL_EXTENSION_EXTRACTOR,
		VAL_EXTENSION_HAL,
		VAL_EXTENSION_UTILITY,
	}

	ServiceList = []string{
		VAL_SERVICE_AI,
		VAL_SERVICE_CLIENT,
		VAL_SERVICE_DEBUG,
		VAL_SERVICE_ERROR,
		VAL_SERVICE_FIREBASE,
		VAL_SERVICE_FIRESTORE,
		VAL_SERVICE_GCP,
		VAL_SERVICE_GOOGLE_ADS,
		VAL_SERVICE_GRPC_SERVER,
		VAL_SERVICE_GRPC_CLIENT,
		VAL_SERVICE_HELPERS,
		VAL_SERVICE_HTTP_SERVER,
		VAL_SERVICE_JWT,
		VAL_SERVICE_LINKEDIN,
		VAL_SERVICE_OS,
		VAL_SERVICE_PAYPAL,
		VAL_SERVICE_PROGRAM_INFO,
		VAL_SERVICE_PSQL,
		VAL_SERVICE_SENDGRID,
		VAL_SERVICE_STRIPE,
		VAL_SERVICE_VALIDATORS,
	}

	SupportedSaaSProviders = []string{
		VAL_SERVICE_GOOGLE_ADS,
		VAL_SERVICE_LINKEDIN,
		VAL_SERVICE_PAYPAL,
		VAL_SERVICE_STRIPE,
	}

	SystemActionList = []string{
		VAL_SYSTEM_ACTION_CONFIG_NEW_USER,
		VAL_SYSTEM_ACTION_GET_MY_ANSWERS,
		VAL_SYSTEM_ACTION_PING,
		VAL_SYSTEM_ACTION_TURN_DEBUG_OFF,
		VAL_SYSTEM_ACTION_TURN_DEBUG_ON,
	}
)
