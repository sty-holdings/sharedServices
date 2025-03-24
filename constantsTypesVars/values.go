package sharedServices

// Entries are for comparison or setting arguments, parameters, and variables. If the file gets to large, split out larger sections to topic specific files.

//goland:noinspection All
const (
	VAL_DETERMINE_SUBJECTS   = "determine-subjects"
	VAL_EMPTY                = ""
	VAL_END_DAY              = "23:59:59"
	VAL_INVALID_FLAG_SETTING = "00000"
	VAL_MID_NIGHT            = "23:59:59"
	VAL_NIL                  = "nil"
	VAL_NONE                 = "none"
	VAL_NOT_FOUND            = "not found"
	VAL_NULL                 = "null"
	VAL_START_DAY            = "00:00:00"
	VAL_START_MONTH          = "-01 00:00:00"
	VAL_TEST                 = "TEST"
	VAL_UNKNOWN              = "UNKNOWN"
	VAL_USER_DOES_NOT_EXIST  = "User does not exist."
	VAL_VALID                = "VALID"
)

//goland:noinspection All
const (
	VAL_GRPC_MAX_PORT = 50151
	VAL_GRPC_MIN_PORT = 50051
	VAL_LOCAL_HOST    = "localhost"
	VAL_NATS          = "nats"
	VAL_PSQL_PORT     = 5432
	VAL_TCP           = "tcp"
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
	VAL_EXTENSION_ADMIN     = "admin"
	VAL_EXTENSION_DIGITS    = "digits"
	VAL_EXTENSION_DK_CLIENT = "dk-client"
	VAL_EXTENSION_HAL       = "hal"
	VAL_EXTENSION_QTESTER   = "qtester"
)

//goland:noinspection ALL
const (
	// output Modes
	VAL_MIME_TYPE_HTML  = "text/html"
	VAL_MIME_TYPE_PLAIN = "text/plain"
)

//goland:noinspection All
const (
	// Extensions
	VAL_SERVICE_GCP         = "gcp"
	VAL_SERVICE_AI          = "ai"
	VAL_SERVICE_GOOGLE_ADS  = "google-ads"
	VAL_SERVICE_GRPC_SERVER = "grpc-server"
	VAL_SERVICE_GRPC_CLIENT = "grpc-client"
	VAL_SERVICE_PSQL        = "psql"
	VAL_SERVICE_SENDGRID    = "sendgrid"
	VAL_SERVICE_STRIPE      = "stripe"
)

//goland:noinspection All
const (
	// System Actions
	VAL_SYSTEM_ACTION_CONFIG_NEW_USER = "config-new-user"
	VAL_SYSTEM_ACTION_GET_MY_ANSWERS  = "get-my-answers"
	VAL_SYSTEM_ACTION_PING            = "ping"
	VAL_SYSTEM_ACTION_TURN_DEBUG_OFF  = "turn-debug-off"
	VAL_SYSTEM_ACTION_TURN_DEBUG_ON   = "turn-debug-on"
	//
)

var (
	ExtensionList = []string{
		VAL_EXTENSION_ADMIN,
		VAL_EXTENSION_DIGITS,
		VAL_EXTENSION_DK_CLIENT,
		VAL_EXTENSION_HAL,
		VAL_EXTENSION_QTESTER,
	}

	ServiceList = []string{
		VAL_SERVICE_GCP,
		VAL_SERVICE_AI,
		VAL_SERVICE_GOOGLE_ADS,
		VAL_SERVICE_GRPC_SERVER,
		VAL_SERVICE_GRPC_CLIENT,
		VAL_SERVICE_PSQL,
		VAL_SERVICE_SENDGRID,
	}

	SupportedSaaSProviders = []string{
		VAL_SERVICE_GOOGLE_ADS,
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
