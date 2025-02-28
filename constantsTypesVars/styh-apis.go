package sharedServices

//goland:noinspection All
const (
	// Subjects
	NCI_PING           = "ping"
	NCI_TURN_DEBUG_OFF = "turn-debug-off"
	NCI_TURN_DEBUG_ON  = "turn-debug-on"
	//
	API_HAL_GET_MY_ANSWER = "hal.get-my-answer"
	//
	API_GEMINI_ANALYZE_QUESTION = "gemini.analyze-question"
	//
	API_SAAS_PROFILE_ADD    = "saas-profile.add"
	API_SAAS_PROFILE_DELETE = "saas-profile.delete"
	API_SAAS_PROFILE_UPDATE = "saas-profile.update"
	//
	API_STRIPE_BALANCE = "stripe.balance"
	//API_STRIPE_CUSTOMERS                        = "stripe.customers"
	API_STRIPE_LIST_AMOUNT_TRANSACTIONS_BETWEEN = "stripe.list-amount-transactions-between"
	//API_STRIPE_LIST_CHECKOUT_SESSIONS                 = "stripe.list-checkout-sessions"
	API_STRIPE_LIST_DISPUTES_BETWEEN = "stripe.list-disputes-between"
	//API_STRIPE_LIST_INVOICES_BETWEEN                  = "stripe.list-invoices-between"
	//API_STRIPE_LIST_PAYMENT_INTENTS_BETWEEN           = "stripe.list-payment-intents-between"
	API_STRIPE_LIST_PAYMENT_METHODS = "stripe.list-payment-methods"
	//API_STRIPE_LIST_PAYOUTS_BETWEEN                   = "stripe.list-payouts-between"
	API_STRIPE_LIST_REFUNDS_BETWEEN      = "stripe.list-refunds-between"
	API_STRIPE_LIST_TOTAL_AMOUNT_BETWEEN = "stripe.list-total-amount-between"
	//API_STRIPE_TRANSACTION_COUNT_BY_STATUS            = "stripe.count-transactions-by-status"
	//
	API_SENDGRID_SEND_EMAIL = "sendgrid.send-email"
	//
)

var StripeSubjects = []string{
	API_STRIPE_BALANCE,
	//API_STRIPE_CUSTOMERS,
	API_STRIPE_LIST_AMOUNT_TRANSACTIONS_BETWEEN,
	API_STRIPE_LIST_DISPUTES_BETWEEN,
	API_STRIPE_LIST_PAYMENT_METHODS,
	API_STRIPE_LIST_REFUNDS_BETWEEN,
	API_STRIPE_LIST_TOTAL_AMOUNT_BETWEEN,
}

var StripeSubjectDescriptions = map[string]string{
	API_STRIPE_BALANCE:                          "return balance today",
	API_STRIPE_LIST_AMOUNT_TRANSACTIONS_BETWEEN: "return transactions between date",
	API_STRIPE_LIST_DISPUTES_BETWEEN:            "return disputes between date",
	API_STRIPE_LIST_PAYMENT_METHODS:             "return payment methods",
	API_STRIPE_LIST_REFUNDS_BETWEEN:             "return refunds between date",
	API_STRIPE_LIST_TOTAL_AMOUNT_BETWEEN:        "return total amount year quarter month week day",
}
