package sharedServices

//==============================
// Analyze Question
//==============================

type AnalyzeQuestionRequest struct {
	Question string `json:"question"`
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

type CancelPaymentIntentRequest struct {
	CancellationReason string `json:"cancellation_reason"`
	PaymentIntentId    string `json:"id"`
	SaaSKey            string `json:"saas_key"`
}

type ConfirmPaymentIntentRequest struct {
	CaptureMethod   string `json:"capture_method,omitempty"`
	PaymentIntentId string `json:"id"`
	PaymentMethod   string `json:"payment_method,omitempty"`
	ReceiptEmail    string `json:"receipt_email,omitempty"`
	ReturnURL       string `json:"return_url,omitempty,omitempty"`
	SaaSKey         string `json:"saas_key"`
}

//==============================
// Generate Answer
//==============================

type GenerateAnswerRequest struct {
	Prompt string
}

//==============================
// Synadia Cloud
//==============================

type GetPersonalAccessTokenRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TokenId string `json:"token_id"`
}

type GetPrometheusMetricsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetSystemLimitsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetSystemRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetTeamRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type GetTeamLimitsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type GetVersionRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
}

type ListAccountsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListInfoAppUserTeamRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListNATSUsersRequest struct {
	SaaSKey   string `json:"saas_key"`
	BaseURL   string `json:"base_url"`
	AccountId string `json:"account_id"`
}

type ListPersonalAccessTokensRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListSystemsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListSystemAccountInfoRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListSystemServerInfoRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListTeamsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
}

type ListTeamServerAccountsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}
