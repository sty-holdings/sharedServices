package sharedServices

import (
	"context"
	"regexp"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/sendgrid/sendgrid-go"
)

//goland:noinspection ALL
const (
	//
	// Addresses
	SUPPORT_ADDRESS        = "support@sty-holdings.com"
	SUPPORT_NAME           = "SavUp Support By STY Holdings"
	DEFAULT_SENDER_ADDRESS = SUPPORT_ADDRESS
	DEFAULT_SENDER_NAME    = SUPPORT_NAME
	DEVELOPMENT_ADDRESS    = "developer@sty-holdings.com"
	VERIFY_ADDRESS         = "verify@sty-holdings.com"
	VERIFY_NAME            = "SavUp By STY Holdings Verification"
	//
	// Settings
	MINE_PLAIN_TEXT    = "text/plain"
	MINE_HTML          = "text/html"
	RECIPIENT_TO       = "SEND_TO"
	RECIPIENT_CC       = "SEND_CC"
	RECIPIENT_BCC      = "SEND_BCC"
	SENDGRID_HOST      = "https://api.sendgrid.com"
	SENDGRID_ENDPOINT  = "/v3/mail/send"
	BANK_ID            = "bank"
	TRANSFER_IN_ID     = "transferIn"
	TRANSFER_OUT_ID    = "transferOut"
	VERIFY_EMAIL       = "verify"
	FORGOT_USERNAME_ID = "forgotUsername"
	TEMPLATE_ID_COUNT  = 4
	//
	// Subjects
	VERIFY_SUBJECT           = "Verification of SavUp Account"
	BANK_REGISTERED_SUBJECT  = "Bank Linked to SavUp Account"
	TRANSFER_REQUEST_SUBJECT = "Transfer Request"
)

type BrevoClient struct {
	clientPtr                 *brevo.APIClient
	transactionalEmailsApiPtr *brevo.TransactionalEmailsApiService
}

type BrevoConfig struct {
	APIKey string `json:"api_key" yaml:"api_key"`
}

type EmailService struct {
	brevoClient          BrevoClient
	sendGridClient       SendGridClient
	debugModeOn          bool
	defaultSenderAddress string
	defaultSenderName    string
}

type EmailConfig struct {
	Brevo                BrevoConfig    `json:"brevo" yaml:"brevo"`
	DebugModeOn          bool           `json:"debug_mode_on" yaml:"debug_mode_on"`
	DefaultSenderName    string         `json:"default_sender_name" yaml:"default_sender_name"`
	DefaultSenderAddress string         `json:"default_sender_address" yaml:"default_sender_address"`
	SendGrid             SendGridConfig `json:"sendgrid" yaml:"sendgrid"`
}

type EmailParams struct {
	Attachments    []EmailAttachment
	BCCList        []EmailToCCBCC
	CCList         []EmailToCCBCC
	Sender         EmailSender
	HTML           string
	PlainText      string
	Subject        string
	TemplateID     interface{}
	TemplateParams map[string]interface{}
	ToList         []EmailToCCBCC
}

type EmailSender struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
}

type EmailToCCBCC struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
}

type EmailAttachment struct {
	Buffer      []byte
	Content     string // Base64 encoded check of data
	ContentType string
	Name        string // Required with Content is used.
	URL         string
}

type SendGridClient struct {
	clientPtr *sendgrid.Client
}

type SendGridConfig struct {
	APIKey string `json:"api_key" yaml:"api_key"`
}

var (
	ctxBackground = context.Background()
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
