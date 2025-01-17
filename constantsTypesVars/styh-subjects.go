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
	SUB_STRIPE_BALANCE = "stripe.balance"
	//SUB_STRIPE_CUSTOMERS                        = "stripe.customers"
	SUB_STRIPE_LIST_AMOUNT_TRANSACTIONS_BETWEEN = "stripe.list-amount-transactions-between"
	//SUB_STRIPE_LIST_CHECKOUT_SESSIONS                 = "stripe.list-checkout-sessions"
	SUB_STRIPE_LIST_DISPUTES_BETWEEN = "stripe.list-disputes-between"
	//SUB_STRIPE_LIST_INVOICES_BETWEEN                  = "stripe.list-invoices-between"
	//SUB_STRIPE_LIST_PAYMENT_INTENTS_BETWEEN           = "stripe.list-payment-intents-between"
	//SUB_STRIPE_LIST_PAYMENT_METHODS                   = "stripe.list-payment-methods"
	//SUB_STRIPE_LIST_PAYOUTS_BETWEEN                   = "stripe.list-payouts-between"
	SUB_STRIPE_LIST_REFUNDS_BETWEEN                   = "stripe.list-refunds-between"
	SUB_STRIPE_LIST_TOTAL_AMOUNT_TRANSACTIONS_BETWEEN = "stripe.list-total-amount-transactions-between"
	//SUB_STRIPE_TRANSACTION_COUNT_BY_STATUS            = "stripe.count-transactions-by-status"
	//
	SUB_SENDGRID_SEND_EMAIL = "sendgrid.send-email"
	//
)

var StripeSubjects = []string{
	SUB_STRIPE_BALANCE,
	//SUB_STRIPE_CUSTOMERS,
	SUB_STRIPE_LIST_AMOUNT_TRANSACTIONS_BETWEEN,
	SUB_STRIPE_LIST_DISPUTES_BETWEEN,
	SUB_STRIPE_LIST_REFUNDS_BETWEEN,
	SUB_STRIPE_LIST_TOTAL_AMOUNT_TRANSACTIONS_BETWEEN,
}
