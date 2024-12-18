package sharedServices

//==============================
// Analyze Question
//==============================

type AnalyzeQuestionRequest struct {
	Question string `json:"question"`
}

//==============================
// Generate Answer
//==============================

type GenerateAnswerRequest struct {
	Prompt string
}

//==============================
// HAL
//==============================

type GetMyAnswerRequest struct {
	Question string `json:"question"`
}

//==============================
// SaaS Profile
//==============================

type SaaSProfileRequest struct {
	Provider        string `json:"provider"`
	Action          string `json:"action"`
	ProviderKeyInfo string `json:"providerKeyInfo"`
}

//==============================
// SendGrid - Twilio
//==============================

type SendEmailRequest struct {
	BodyPlain          string           `json:"body_plain,omitempty"`
	BodyHTML           string           `json:"body_html,omitempty"`
	EmailToRecipient   []EmailRecipient `json:"email_to_recipient"`
	SaaSKey            string           `json:"saas_key"`
	SenderEmailAddress string           `json:"sender_email_address,omitempty"`
	SenderName         string           `json:"sender_name,omitempty"`
	Subject            string           `json:"subject,omitempty"`
}

type EmailRecipient struct {
	Name    string
	Address string
}

//==============================
// Stripe
//==============================

type BalanceRequest struct {
	SaaSKey string `json:"saas_key"`
}

type ListAllChargeRequest struct {
	SaaSKey string `json:"saas_key"`
	StartAt string `json:"start_at"`
	EndBy   string `json:"end_by"`
}

type ListPaymentMethodRequest struct {
	SaaSKey string `json:"saas_key"`
}

type ListPaymentIntentRequest struct {
	CustomerId    string `json:"customer_id,omitempty"`
	Limit         int64  `json:"limit,omitempty"`
	SaaSKey       string `json:"saas_key"`
	StartingAfter string `json:"starting_after,omitempty"`
}

type PaymentIntentRequest struct {
	Amount                  float64 `json:"amount"`
	AutomaticPaymentMethods bool    `json:"automatic_payment_methods,omitempty"`
	Currency                string  `json:"currency"`
	Description             string  `json:"description,omitempty"`
	ReceiptEmail            string  `json:"receipt_email"`
	ReturnURL               string  `json:"return_url,omitempty"`
	SaaSKey                 string  `json:"saas_key"`
	// Confirm            bool     `json:"confirm,omitempty"`
	// PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
}