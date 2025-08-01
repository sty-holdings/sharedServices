package sharedServices

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/goccy/go-yaml"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	pis "github.com/sty-holdings/sharedServices/v2025/programInfo"
	vldts "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewBrevoServer - initializes and returns a new emailService instance.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, error returned by loadConfig, error returned by validateConfig
//	Verifications: vldts.CheckValueNotEmpty, validateConfig
func NewBrevoServer(configFilename string, environment string) (servicePtr *EmailService, errorInfo errs.ErrorInfo) {

	var (
		tConfig EmailConfig
	)

	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadBrevoConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateBrevoConfig(tConfig); errorInfo.Error != nil {
		return
	}

	servicePtr = &EmailService{
		debugModeOn:          tConfig.DebugModeOn,
		DefaultSenderAddress: tConfig.DefaultSenderAddress,
		DefaultSenderName:    tConfig.DefaultSenderName,
		emailProvider:        strings.ToLower(tConfig.EmailProvider),
	}

	switch servicePtr.emailProvider {
	case BREVO_PROVIDER:
		cfg := brevo.NewConfiguration()
		cfg.AddDefaultHeader("api-key", tConfig.Brevo.APIKey)

		servicePtr.brevoClient.clientPtr = brevo.NewAPIClient(cfg)
		servicePtr.brevoClient.transactionalEmailsApiPtr = servicePtr.brevoClient.clientPtr.TransactionalEmailsApi
	case SENDGRID_PROVIDER:
		fallthrough
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_EMAIL_PROVIDER, tConfig.EmailProvider))
		return
	}

	return
}

func (servicePtr *EmailService) SendEmail(emailParams EmailParams) (errorInfo errs.ErrorInfo) {

	var (
		tSendSMTPEmailParams brevo.SendSmtpEmail
	)

	if errorInfo = validateEmailParams(servicePtr.emailProvider, emailParams); errorInfo.Error != nil {
		return
	}

	switch servicePtr.emailProvider {
	case BREVO_PROVIDER:
		if tSendSMTPEmailParams = servicePtr.buildBrevoEmailParams(emailParams); errorInfo.Error != nil {
			return
		}
	case SENDGRID_PROVIDER:
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_EMAIL_PROVIDER, servicePtr.emailProvider))
		return
	}

	if _, _, errorInfo.Error = servicePtr.brevoClient.transactionalEmailsApiPtr.SendTransacEmail(ctxBackground, tSendSMTPEmailParams); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.VAL_EMPTY, ctv.TXT_FAILED))
		return
	}

	return
}

// Private Methods

func (servicePtr *EmailService) buildBrevoEmailParams(emailParams EmailParams) (sendSMTPEmail brevo.SendSmtpEmail) {

	var (
		errorInfo    errs.ErrorInfo
		tSendAddress string
		tSenderName  string
	)

	for _, attachment := range emailParams.Attachments {
		sendSMTPEmail.Attachment = append(
			sendSMTPEmail.Attachment, brevo.SendSmtpEmailAttachment{
				Url:     attachment.URL,
				Content: attachment.Content,
				Name:    attachment.Name,
			},
		)
	}

	for _, bccList := range emailParams.BCCList {
		sendSMTPEmail.Bcc = append(
			sendSMTPEmail.Bcc, brevo.SendSmtpEmailBcc{
				Email: bccList.Address,
				Name:  bccList.Name,
			},
		)
	}

	for _, ccList := range emailParams.CCList {
		sendSMTPEmail.Cc = append(
			sendSMTPEmail.Cc, brevo.SendSmtpEmailCc{
				Email: ccList.Address,
				Name:  ccList.Name,
			},
		)
	}

	if emailParams.Sender.Address == ctv.VAL_EMPTY || emailParams.Sender.Name == ctv.VAL_EMPTY {
		tSendAddress = servicePtr.DefaultSenderAddress
		tSenderName = servicePtr.DefaultSenderName
	}
	sendSMTPEmail.Sender = &brevo.SendSmtpEmailSender{
		Email: tSendAddress,
		Name:  tSenderName,
	}

	sendSMTPEmail.Subject = emailParams.Subject

	for _, toList := range emailParams.ToList {
		sendSMTPEmail.To = append(
			sendSMTPEmail.To, brevo.SendSmtpEmailTo{
				Email: toList.Address,
				Name:  toList.Name,
			},
		)
	}

	switch emailParams.EmailType {
	case SINGLE:
		sendSMTPEmail.HtmlContent = emailParams.HTML
		sendSMTPEmail.TextContent = emailParams.PlainText
	case TEMPLATE:
		switch strings.ToLower(servicePtr.emailProvider) {
		case BREVO_PROVIDER:
			if sendSMTPEmail.TemplateId, errorInfo = convertTemplateIDToInt(emailParams.TemplateID); errorInfo.Error != nil {
				return
			}
			sendSMTPEmail.Params = make(map[string]interface{})
			for k, v := range emailParams.TemplateParams {
				sendSMTPEmail.Params[k] = v
			}
		case SENDGRID_PROVIDER:
			fallthrough
		default:
			errs.PrintErrorInfo(errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), servicePtr.emailProvider)))
		}
	default:
		errs.PrintErrorInfo(errs.NewErrorInfo(errs.ErrInvalidEmailType, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), emailParams.EmailType)))
		return
	}

	return
}

// Private Functions

// loadBrevoConfig - loads and parses the Brevo configuration file.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter if configFilename is empty, or errs.NewErrorInfo for file read/unmarshal errors.
//	Verifications: ctv.
func loadBrevoConfig(configFilename string) (config EmailConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
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
func validateBrevoConfig(config EmailConfig) (errorInfo errs.ErrorInfo) {

	switch strings.ToLower(config.EmailProvider) {
	case BREVO_PROVIDER:
		if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.Brevo.APIKey, ctv.LBL_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	case SENDGRID_PROVIDER:
		fallthrough
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_EMAIL_PROVIDER, config.EmailProvider))
		return
	}
	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.DefaultSenderAddress, ctv.LBL_DEFAULT_SENDER_ADDRESS); errorInfo.Error != nil {
		return
	}
	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.DefaultSenderName, ctv.LBL_DEFAULT_SENDER_NAME); errorInfo.Error != nil {
		return
	}

	return
}

func validateEmailParams(emailProvider string, emailParams EmailParams) (errorInfo errs.ErrorInfo) {

	var (
		dataType interface{}
	)

	for _, attachment := range emailParams.Attachments {
		if attachment.URL == ctv.VAL_EMPTY {
			if attachment.Content == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_ATTACHMENT_CONTENT))
				return
			}
			if attachment.Name == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_ATTACHMENT_NAME))
				return
			}
		}
	}

	for _, bccList := range emailParams.BCCList {
		if bccList.Address == ctv.VAL_EMPTY {
			errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_BCC_ADDRESS))
			return
		}
	}

	for _, ccList := range emailParams.CCList {
		if ccList.Address == ctv.VAL_EMPTY {
			errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_CC_ADDRESS))
			return
		}
	}

	if emailParams.Subject == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_SUBJECT))
		return
	}

	if len(emailParams.ToList) == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(
			errs.ErrEmptyRequiredParameter, errs.BuildLabelSubLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TO_ADDRESS, dataType.(string)),
		)
	} else {
		for _, toList := range emailParams.ToList {
			if toList.Address == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TO_ADDRESS))
			}
		}
	}

	switch strings.ToLower(emailParams.EmailType) {
	case SINGLE:
		if emailParams.HTML == ctv.VAL_EMPTY && emailParams.PlainText == ctv.VAL_EMPTY {
			errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_BODY))
			return
		}
	case TEMPLATE:
		switch emailProvider {
		case BREVO_PROVIDER:
			if _, errorInfo = convertTemplateIDToInt(emailParams.TemplateID); errorInfo.Error != nil {
				return
			}
		case SENDGRID_PROVIDER:
			fallthrough
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), emailProvider))
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailType, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), emailParams.EmailType))
	}

	return
}

func convertTemplateIDToInt(templateID string) (brevoTemplateID int64, errorInfo errs.ErrorInfo) {

	var (
		numericRegex = regexp.MustCompile("^[0-9]*$")
		tTemplateID  int
	)

	if templateID != ctv.VAL_EMPTY && numericRegex.MatchString(templateID) {
		if tTemplateID, errorInfo.Error = strconv.Atoi(templateID); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TEMPLATE_ID))
			return
		}
		brevoTemplateID = int64(tTemplateID)
	} else {
		errorInfo = errs.NewErrorInfo(
			errs.ErrInvalidDataType,
			errs.BuildLabelSubLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TEMPLATE_ID, templateID),
		)
		return
	}

	if brevoTemplateID <= ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(
			errs.ErrGreaterThanZero,
			errs.BuildLabelSubLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TEMPLATE_ID, templateID),
		)
	}

	return
}
