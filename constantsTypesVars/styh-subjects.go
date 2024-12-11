package sharedServices

//goland:noinspection All
const (
	// Subjects
	NCI_PING           = "ping"
	NCI_TURN_DEBUG_OFF = "turn-debug-off"
	NCI_TURN_DEBUG_ON  = "turn-debug-on"
	//
	SUB_HAL_GET_MY_ANSWER = "hal.get-my-answer"
	//
	SUB_GEMINI_ANALYZE_QUESTION = "gemini.analyze-question"
	SUB_GEMINI_SUMMARIZE_ANSWER = "gemini.analyze-answer"
	//
	SUB_SAAS_PROFILE_ADD    = "saas-profile.add"
	SUB_SAAS_PROFILE_DELETE = "saas-profile.delete"
	SUB_SAAS_PROFILE_UPDATE = "saas-profile.update"
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
