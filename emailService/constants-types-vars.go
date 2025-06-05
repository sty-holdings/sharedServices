package sharedServices

import (
	"context"
	"regexp"

	brevo "github.com/getbrevo/brevo-go/lib"
	sendgrid "github.com/sendgrid/sendgrid-go"
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

type Brevo struct {
	KeyFQN string `json:"key_fqn" yaml:"key_fqn"`
}

type EmailService struct {
	clientBrevoPtr       *brevo.APIClient
	clientSendGridPtr    *sendgrid.Client
	debugModeOn          bool
	defaultSenderAddress string
	defaultSenderName    string
}

type EmailConfig struct {
	Brevo                Brevo    `json:"brevo" yaml:"brevo"`
	DebugModeOn          bool     `json:"debug_mode_on" yaml:"debug_mode_on"`
	DefaultSenderName    string   `json:"default_sender_name" yaml:"sender_name"`
	DefaultSenderAddress string   `json:"default_address" yaml:"sender_address"`
	SendGrid             SendGrid `json:"sendgrid" yaml:"sendgrid"`
}

type EmailServer struct {
	emailInfo Email
}

type Email struct {
	DefaultSender  EmailSender
	ToList         []EmailToCCBCC
	CCList         []EmailToCCBCC
	BCCList        []EmailToCCBCC
	Subject        string
	PlainText      string
	HTML           string
	Attachments    []EmailAttachment
	TemplateID     string
	TemplateParams map[string]string
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
	Filepath    string
	ContentType string
	Buffer      []byte
}

type SendGrid struct {
	KeyFQN string `json:"key_fqn" yaml:"key_fqn"`
}

var (
	ctxBackground = context.Background()
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
