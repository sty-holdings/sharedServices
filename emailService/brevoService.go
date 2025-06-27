package sharedServices

import (
	"os"
	"strings"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/goccy/go-yaml"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	pis "github.com/sty-holdings/sharedServices/v2025/programInfo"
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
			if emailParams.TemplateID != nil {
				sendSMTPEmail.TemplateId = int64(emailParams.TemplateID.(int))
				sendSMTPEmail.Params = emailParams.TemplateParams
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
func validateBrevoConfig(config EmailConfig) (errorInfo errs.ErrorInfo) {

	switch strings.ToLower(config.EmailProvider) {
	case BREVO_PROVIDER:
		if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_BREVO, config.Brevo.APIKey, ctv.LBL_KEY_FILENAME); errorInfo.Error != nil {
			return
		}
	case SENDGRID_PROVIDER:
		fallthrough
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.LBL_SERVICE_BREVO, ctv.LBL_EMAIL_PROVIDER, config.EmailProvider))
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

func validateEmailParams(emailProvider string, emailParams EmailParams) (errorInfo errs.ErrorInfo) {

	var (
		dataType interface{}
	)

	if emailParams.Attachments != nil {
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
	}

	if len(emailParams.BCCList) > ctv.VAL_ZERO {
		for _, bccList := range emailParams.BCCList {
			if bccList.Address == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_BCC_ADDRESS))
				return
			}
		}
	}

	if len(emailParams.CCList) > ctv.VAL_ZERO {
		for _, ccList := range emailParams.CCList {
			if ccList.Address == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrEmptyRequiredParameter, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_CC_ADDRESS))
				return
			}
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
		case "brevo":
			if emailParams.TemplateID != nil {
				switch dataType = emailParams.TemplateID; emailParams.TemplateID.(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
					if dataType.(int) <= ctv.VAL_ZERO {
						errorInfo = errs.NewErrorInfo(
							errs.ErrGreaterThanZero,
							errs.BuildLabelSubLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TEMPLATE_ID, dataType.(string)),
						)
						return
					}
					break
				default:
					errorInfo = errs.NewErrorInfo(
						errs.ErrEmptyRequiredParameter, errs.BuildLabelSubLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), ctv.FN_EMAIL_TEMPLATE_ID, dataType.(string)),
					)
					return
				}
				if errorInfo = hlps.CheckMapLengthGTZero(ctv.VAL_SERVICE_EMAIL, emailParams.TemplateParams, ctv.FN_EMAIL_TEMPLATE_PARAMS); errorInfo.Error != nil {
					return
				}
			}
		case "sendgrid":
			fallthrough
		default:
			errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailProvider, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), emailProvider))
			return
		}
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidEmailType, errs.BuildLabelValue(ctv.VAL_SERVICE_EMAIL, pis.GetMyFunctionName(true), emailParams.EmailType))
		return
	}

	return
}
