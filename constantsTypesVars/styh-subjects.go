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
	SUB_GEMINI_GENERATE_ANSWER  = "gemini.generate-answer"
	//
	SUB_SAAS_PROFILE_ADD    = "saas-profile.add"
	SUB_SAAS_PROFILE_DELETE = "saas-profile.delete"
	SUB_SAAS_PROFILE_UPDATE = "saas-profile.update"
	//
	SUB_STRIPE_LIST_ALL_CHARGES           = "stripe.list-all-charges"
	SUB_STRIPE_LIST_ALL_CHECKOUT_SESSIONS = "stripe.list-all-checkout-sessions"
	SUB_STRIPE_LIST_ALL_DISPUTES          = "stripe.list-all-disputes"
	SUB_STRIPE_LIST_ALL_INVOICES          = "stripe.list-all-invoices"
	SUB_STRIPE_LIST_ALL_PAYOUTS           = "stripe.list-all-payouts"
	SUB_STRIPE_LIST_ALL_REFUNDS           = "stripe.list-all-refunds"
	SUB_STRIPE_LIST_PAYMENT_INTENTS       = "stripe.list-payment-intents"
	SUB_STRIPE_LIST_PAYMENT_METHODS       = "stripe.list-payment-methods"
	SUB_STRIPE_BALANCE                    = "stripe.balance"
	SUB_STRIPE_CUSTOMERS                  = "stripe.customers"
	//
	SUB_SENDGRID_SEND_EMAIL = "sendgrid.send-email"
	//
)
