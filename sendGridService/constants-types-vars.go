package sendGridService

import (
	"regexp"
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

type EmailServer struct {
	emailInfo Email
}

type Email struct {
	DefaultSender  EmailItem
	Host           string
	Key            string
	Environment    string
	ProviderKeyFQN string
}

type EmailItem struct {
	Name    string
	Address string
}

type EmailAttachment struct {
	Filepath    string
	ContentType string
	Buffer      []byte
}

type SendGridHelper struct {
	SendGridCredentialsFQN string
	Key                    string `json:"sendgrid_key"`
	EmailServerPtr         *EmailServer
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
