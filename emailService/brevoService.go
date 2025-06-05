package sharedServices

import (
	"fmt"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/goccy/go-yaml"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

// NewBrevoServer - initializes and returns a new emailService instance.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, error returned by loadConfig, error returned by validateConfig
//	Verifications: hlps.CheckValueNotEmpty, validateConfig
func NewBrevoServer(configFilename string, environment string) (servicePtr *EmailService, errorInfo errs.ErrorInfo) {

	var (
		tConfig EmailConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadBrevoConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateBrevoConfig(tConfig, environment); errorInfo.Error != nil {
		return
	}

	servicePtr = &EmailService{
		debugModeOn:          tConfig.DebugModeOn,
		defaultSenderAddress: tConfig.DefaultSenderAddress,
		defaultSenderName:    tConfig.DefaultSenderName,
	}

	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", tConfig.Brevo.KeyFQN)

	servicePtr.clientBrevoPtr = brevo.NewAPIClient(cfg)

	return
}

func (servicePtr *EmailService) BuildSendSMTPEmail(sender EmailSender, toList []EmailToCCBCC, ccList []EmailToCCBCC, bcc []EmailToCCBCC) (sendSMTPEmail brevo.SendSmtpEmail) {

	return
}

func (servicePtr *EmailService) SendEmail() {

	var (
		tSendSMTPEmail brevo.SendSmtpEmail
	)

	tSendSMTPEmail = brevo.SendSmtpEmail{}
	tSendSMTPEmail.Sender = &brevo.SendSmtpEmailSender{
		Name:  "Your Name",
		Email: "your.sender@example.com", // Must be a verified sender in Brevo
	}
	tSendSMTPEmail.To = []brevo.SendSmtpEmailTo{
		{
			Name:  "Recipient Name",
			Email: "recipient@example.com",
		},
	}
	tSendSMTPEmail.Subject = "Test Email from Go SDK"
	tSendSMTPEmail.HtmlContent = "<h1>This is a test email!</h1><p>Sent from the Brevo Go SDK.</p>"

	_, emailResp, emailErr := servicePtr.clientBrevoPtr.TransactionalEmailsApi.SendTransacEmail(ctxBackground, tSendSMTPEmail)
	if emailErr != nil {
		fmt.Printf("Error sending email: %v\n", emailErr)
		if emailResp != nil {
			fmt.Printf("Email send response status: %s\n", emailResp.Status)
		}
	} else {
		fmt.Printf("Email sent successfully! Response status: %s\n", emailResp.Status)
	}

	return
}

// Private Functions

func buildSendSMTPEmail(sender EmailSender) (sendSMTPEmail brevo.SendSmtpEmail) {

	sendSMTPEmail.Sender = &brevo.SendSmtpEmailSender{
		Name:  "Your Name",
		Email: "your.sender@example.com", // Must be a verified sender in Brevo
	}

	sendSMTPEmail.To = []brevo.SendSmtpEmailTo{
		{
			Name:  "Recipient Name",
			Email: "recipient@example.com",
		},
	}

	sendSMTPEmail.Subject = "Test Email from Go SDK"
	sendSMTPEmail.HtmlContent = "<h1>This is a test email!</h1><p>Sent from the Brevo Go SDK.</p>"

	return
}

// loadBrevoConfig - loads and parses the Brevo configuration file.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter if configFilename is empty, or errs.NewErrorInfo for file read/unmarshal errors.
//	Verifications: ctv.
func loadBrevoConfig(configFilename string) (config EmailConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	return
}

// validateBrevoConfig - validates the Brevo configuration parameters.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter
//	Verifications: ctv.
func validateBrevoConfig(config EmailConfig, environment string) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.Brevo.KeyFQN, ctv.LBL_KEY_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.DefaultSenderAddress, ctv.LBL_DEFAULT_SENDER_ADDRESS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.DefaultSenderName, ctv.LBL_DEFAULT_SENDER_NAME); errorInfo.Error != nil {
		return
	}

	return
}
