package sharedServices

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/plaid/plaid-go/v9/plaid"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

// NewSendGridServer - initializes and returns a new SendGridService instance.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, error returned by loadConfig, error returned by validateConfig
//	Verifications: hlps.CheckValueNotEmpty, validateConfig
func NewSendGridServer(configFilename string, environment string) (servicePtr *SendGridService, errorInfo errs.ErrorInfo) {

	var (
		tConfig SendGridConfig
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_SENDGRID, configFilename, ctv.LBL_CONFIG_EXTENSION_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfig, errorInfo = loadConfig(configFilename); errorInfo.Error != nil {
		return
	}

	if errorInfo = validateSendGridConfig(tConfig, environment); errorInfo.Error != nil {
		return
	}

	servicePtr = &SendGridService{
		debugModeOn:          tConfig.DebugModeOn,
		defaultSenderAddress: tConfig.DefaultSenderAddress,
		defaultSenderName:    tConfig.DefaultSenderName,
	}

	servicePtr.clientPtr = sendgrid.NewSendClient(tConfig.KeyFQN)

	return
}

// NewSendGridServer - initialize the SendGrid service for use. When the mode is production of demo, the defaultSenderAddress is used. For other modes, developer@sty-holdings.com is used.
func xNewSendGridServer(defaultSenderAddress, defaultSenderName, environment, sendgridKeyFQN string) (emailServerPtr *EmailServer, errorInfo errs.ErrorInfo) {

	// var (
	// 	tSendGrid       SendGridHelper
	// 	tEmailServer    EmailServer
	// 	tEmailServerPtr = &tEmailServer
	// )
	//
	// if tSendGrid, errorInfo = sendGridGetKey(sendgridKeyFQN); errorInfo.Error == nil {
	// 	if errorInfo = validateEmailAddress(defaultSenderAddress); errorInfo.Error == nil {
	// 		if coreValidators.IsEnvironmentValid(environment) {
	// 			tEmailServerPtr.emailInfo = Email{
	// 				Host:           SENDGRID_HOST,
	// 				Key:            tSendGrid.Key,
	// 				Environment:    environment,
	// 				ProviderKeyFQN: sendgridKeyFQN,
	// 			}
	// 			switch strings.ToUpper(environment) {
	// 			case constants.VAL_ENVIRONMENT_PRODUCTION:
	// 				tEmailServerPtr.emailInfo.DefaultSender.Name = defaultSenderName
	// 				if errorInfo = validateEmailAddress(defaultSenderAddress); errorInfo.Error == nil {
	// 					tEmailServerPtr.emailInfo.DefaultSender.Address = defaultSenderAddress
	// 				}
	// 			default:
	// 				tEmailServerPtr.emailInfo.DefaultSender.Name = defaultSenderName
	// 				tEmailServerPtr.emailInfo.DefaultSender.Address = DEVELOPMENT_ADDRESS
	// 			}
	// 			emailServerPtr = tEmailServerPtr
	// 		}
	// 	}
	// }

	return
}

// addTemplateData
func (emailServerPtr *EmailServer) addTemplateData(personalizationPtr *mail.Personalization, templateData map[string]interface{}) {

	personalizationPtr.DynamicTemplateData = templateData
}

// NewPersonalization - adds the 'from' address if valid, otherwise it uses the default sender. The toList must be populated, while the ccList and bccList are optional.
func (emailServerPtr *EmailServer) newPersonalization(personalizationPtr *mail.Personalization, toList, ccList, bccList []EmailItem) (errorInfo errs.ErrorInfo) {

	if isRecipientListPopulated(toList) {
		if errorInfo = addRecipientList(personalizationPtr, toList, RECIPIENT_TO); errorInfo.Error == nil {
			if isRecipientListPopulated(ccList) {
				if errorInfo = addRecipientList(personalizationPtr, ccList, RECIPIENT_CC); errorInfo.Error == nil {
					if isRecipientListPopulated(bccList) && errorInfo.Error == nil {
						errorInfo = addRecipientList(personalizationPtr, bccList, RECIPIENT_BCC)
					}
				}
			}
		}
	} else {
		errorInfo.Error = errors.New("Require information is missing! toList is not populated:")
		log.Println(errorInfo.Error.Error())
	}

	return
}

// SendEmailUsingPlainText - The toList must have non-blank address to send an email. The ccList and bccList parameters can have empty addresses.
func (emailServerPtr *EmailServer) sendEmailUsingPlainText(from EmailItem, subject, body string, toList, ccList, bccList []EmailItem, replyTo EmailItem) (
	response *rest.Response,
	errorInfo errs.ErrorInfo,
) {

	// var (
	// 	tEmailPtr           = mail.NewV3Mail()
	// 	tPersonalizationPtr = mail.NewPersonalization()
	// )
	//
	// if subject == constants.EMPTY || body == constants.EMPTY || isRecipientListPopulated(toList) == false {
	// 	errorInfo.Error = errs.ErrRequiredArgumentMissing
	// 	errorInfo.AdditionalInfo = fmt.Sprintf("Subject: '%v' Body: '%v' and/or the 'To List'.", subject, body)
	// 	errs.PrintError(errorInfo)
	// } else {
	// 	addFrom(emailServerPtr, tEmailPtr, from)
	// 	if errorInfo = validateSubject(subject); errorInfo.Error == nil {
	// 		tEmailPtr.Subject = subject
	// 		if errorInfo = emailServerPtr.newPersonalization(tPersonalizationPtr, toList, ccList, bccList); errorInfo.Error == nil {
	// 			tEmailPtr.AddPersonalizations(tPersonalizationPtr)
	// 			addContent(tEmailPtr, MINE_PLAIN_TEXT, body)
	// 			if errorInfo = addReplyTo(tEmailPtr, replyTo); errorInfo.Error == nil {
	// 				response, errorInfo = sendEmail(tEmailPtr, emailServerPtr.emailInfo.Key, emailServerPtr.emailInfo.Host)
	// 			}
	// 		}
	// 	}
	// }

	return
}

// SendEmailUsingPlainText - The toList, template id, and the template data must be populated to send an email. The ccList and bccList parameters can have empty addresses.
func (emailServerPtr *EmailServer) SendEmailUsingTemplate(
	from EmailItem,
	subject string,
	toList, ccList, bccList []EmailItem,
	replyTo EmailItem,
	templateId string,
	templateData map[any]interface{},
	test bool,
) (response *rest.Response, errorInfo errs.ErrorInfo) {

	// var (
	// 	tEmailPtr           = mail.NewV3Mail()
	// 	tFindings           string
	// 	tPersonalizationPtr = mail.NewPersonalization()
	// )

	// if tFindings = coreValidators.AreMapKeysValuesPopulated(templateData); tFindings != constants.GOOD {
	// 	errorInfo.Error = errs.GetMapKeyPopulatedError(tFindings)
	// } else {
	// 	if subject == constants.EMPTY || isRecipientListPopulated(toList) == false || templateId == constants.EMPTY {
	// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
	// 		errorInfo.AdditionalInfo = fmt.Sprintf("Subject: '%v' Template Id: '%v' and/or the 'To List'.", subject, templateId)
	// 		errs.PrintError(errorInfo)
	// 	} else {
	// 		addFrom(emailServerPtr, tEmailPtr, from)
	// 		if errorInfo = validateSubject(subject); errorInfo.Error == nil {
	// 			tEmailPtr.Subject = subject
	// 			if errorInfo = emailServerPtr.newPersonalization(tPersonalizationPtr, toList, ccList, bccList); errorInfo.Error == nil {
	// 				tPersonalizationPtr.DynamicTemplateData = coreHelpers.ConvertMapAnyToMapString(templateData)
	// 				tEmailPtr.SetTemplateID(templateId)
	// 				tEmailPtr.AddPersonalizations(tPersonalizationPtr)
	// 				if errorInfo = addReplyTo(tEmailPtr, replyTo); errorInfo.Error == nil {
	// 					if test == false {
	// 						response, errorInfo = sendEmail(tEmailPtr, emailServerPtr.emailInfo.Key, emailServerPtr.emailInfo.Host)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return
}

// validateAddress - checks the length, the domain and the format of the address
func (emailServerPtr *EmailServer) validateAddress(emailAddress string) (errorInfo errs.ErrorInfo) {

	return validateEmailAddress(emailAddress)
}

// addContent
// ToDo Add profanity checking service for subject line
func addContent(emailPtr *mail.SGMailV3, mineType, body string) {

	emailPtr.AddContent(mail.NewContent(mineType, body))
}

// addFrom - populates the email from with the supplied from or the default sender if the 'from' is empty.
func addFrom(emailServerPtr *EmailServer, emailPtr *mail.SGMailV3, from EmailItem) {

	var (
		errorInfo errs.ErrorInfo
	)

	// If the supplied 'from' email address is invalid, then the default email address and name is used.
	if errorInfo = validateEmailAddress(from.Address); errorInfo.Error == nil {
		emailPtr.SetFrom(mail.NewEmail(from.Name, from.Address))
	} else {
		emailPtr.SetFrom(mail.NewEmail(emailServerPtr.emailInfo.DefaultSender.Name, emailServerPtr.emailInfo.DefaultSender.Address))
	}
}

// addRecipientList
// ToDo Add profanity checking service for subject line
func addRecipientList(personalizationPtr *mail.Personalization, recipientList []EmailItem, recipientType string) (errorInfo errs.ErrorInfo) {

	for _, recipient := range recipientList {
		if errorInfo = validateEmailAddress(recipient.Address); errorInfo.Error == nil {
			tNameAddress := []*mail.Email{
				mail.NewEmail(recipient.Name, recipient.Address),
			}
			switch strings.ToUpper(recipientType) {
			case RECIPIENT_TO:
				personalizationPtr.AddTos(tNameAddress...)
			case RECIPIENT_CC:
				personalizationPtr.AddCCs(tNameAddress...)
			case RECIPIENT_BCC:
				personalizationPtr.AddBCCs(tNameAddress...)
			}
		}
	}

	return
}

// addReplyTo
func addReplyTo(myEmailPtr *mail.SGMailV3, replyTo EmailItem) (errorInfo errs.ErrorInfo) {

	if errorInfo = validateEmailAddress(replyTo.Address); errorInfo.Error == nil {
		myEmailPtr.SetReplyTo(mail.NewEmail(replyTo.Name, replyTo.Address))
	}

	return
}

// isRecipientListPopulated - checks if all the entries in the recipient list for an empty address.
func isRecipientListPopulated(recipientList []EmailItem) bool {

	// for _, recipient := range recipientList {
	// 	if recipient.Address == constants.EMPTY {
	// 		return false
	// 	}
	// }

	return true
}

// GenerateVerifyEmail - will format and send the verification email for a newly created user
func GenerateVerifyEmail(emailServerPtr *EmailServer, templateId string, firstName, lastName, email, shortURL string, test bool) (errorInfo errs.ErrorInfo) {

	var (
		tBCCList []EmailItem
		tCCList  []EmailItem
		tFrom    = EmailItem{
			Name:    VERIFY_NAME,
			Address: VERIFY_ADDRESS,
		}
		tReplyTo = EmailItem{
			Name:    SUPPORT_NAME,
			Address: SUPPORT_ADDRESS,
		}
		tTemplateData = make(map[any]interface{})
		tToList       []EmailItem
	)

	tToList = []EmailItem{
		{
			Name:    fmt.Sprintf("%v %v", firstName, lastName),
			Address: email,
			// ToDo Add logging for the response and error handling
		},
	}
	tTemplateData["su_first_name"] = firstName
	tTemplateData["shorturl"] = shortURL
	_, errorInfo = emailServerPtr.SendEmailUsingTemplate(tFrom, VERIFY_SUBJECT, tToList, tCCList, tBCCList, tReplyTo, templateId, tTemplateData, test)

	return
}

// GenerateBankRegisteredEmail - will format and send an email for the linked bank account
//
//	Customer Messages: None
//	Errors: Any error returned from emailServerPtr.SendEmailUsingTemplate
//	Verifications: None
func GenerateBankRegisteredEmail(
	emailServerPtr *EmailServer,
	templateId string,
	firstName, lastName, email, institutionName string,
	accountData []plaid.AccountBase,
	test bool,
) (errorInfo errs.ErrorInfo) {

	var (
		tBCCList []EmailItem
		tCCList  []EmailItem
		tFrom    = EmailItem{
			Name:    VERIFY_NAME,
			Address: VERIFY_ADDRESS,
		}
		tReplyTo = EmailItem{
			Name:    SUPPORT_NAME,
			Address: SUPPORT_ADDRESS,
		}
		tTemplateData = make(map[any]interface{})
		tToList       []EmailItem
	)

	tToList = []EmailItem{
		{
			Name:    fmt.Sprintf("%v %v", firstName, lastName),
			Address: email,
			// ToDo Add logging for the response and error handling
		},
	}
	tTemplateData["su_first_name"] = firstName
	tTemplateData["su_institution_name"] = institutionName
	for i := 0; i < len(accountData); i++ {
		tTemplateData[fmt.Sprintf("su_institution_account_label_%v", i)] = "Account:"
		tTemplateData[fmt.Sprintf("su_institution_account_name_%v", i)] = accountData[i].OfficialName
	}
	_, errorInfo = emailServerPtr.SendEmailUsingTemplate(tFrom, BANK_REGISTERED_SUBJECT, tToList, tCCList, tBCCList, tReplyTo, templateId, tTemplateData, test)

	return
}

// GenerateTransferRequestEmail - will format and send an email for a transfer request
// The map[string]string must have the following keys to generate the email correctly:
// Keys: direction, amount, method, and completion where direction is either 'into' or 'out of'
//
//	Customer Messages: None
//	Errors: Any error returned from emailServerPtr.SendEmailUsingTemplate
//	Verifications: None
func GenerateTransferRequestEmail(emailServerPtr *EmailServer, templateId string, firstName, lastName, email string, transferData map[string]string, test bool) (errorInfo errs.ErrorInfo) {

	var (
		tBCCList []EmailItem
		tCCList  []EmailItem
		tFrom    = EmailItem{
			Name:    VERIFY_NAME,
			Address: VERIFY_ADDRESS,
		}
		tReplyTo = EmailItem{
			Name:    SUPPORT_NAME,
			Address: SUPPORT_ADDRESS,
		}
		tTemplateData = make(map[any]interface{})
		tToList       []EmailItem
	)

	tToList = []EmailItem{
		{
			Name:    fmt.Sprintf("%v %v", firstName, lastName),
			Address: email,
			// ToDo Add logging for the response and error handling
		},
	}
	tBCCList = []EmailItem{
		{
			Name:    SUPPORT_NAME,
			Address: SUPPORT_ADDRESS,
			// ToDo Add logging for the response and error handling
		},
	}
	tTemplateData["su_first_name"] = firstName
	tTemplateData["su_transfer_amount"] = transferData["amount"]
	// switch strings.ToUpper(transferData["method"]) {
	// case constants.TRANFER_CHECK:
	// 	tTemplateData["su_transfer_method"] = constants.CHECK
	// case constants.TRANFER_WIRE:
	// 	tTemplateData["su_transfer_method"] = constants.WIRE
	// 	tTemplateData["su_institution_lbl"] = constants.TRANSFER_INSTITUTION_NAME
	// 	tTemplateData["su_institution_name"] = transferData["institution"]
	// case constants.TRANFER_ZELLE:
	// 	tTemplateData["su_transfer_method"] = constants.ZELLE
	// case constants.TRANFER_STRIPE:
	// 	tTemplateData["su_transfer_method"] = constants.STRIPE
	// }
	tTemplateData["su_estimated_completion"] = transferData["completion"]
	_, errorInfo = emailServerPtr.SendEmailUsingTemplate(tFrom, TRANSFER_REQUEST_SUBJECT, tToList, tCCList, tBCCList, tReplyTo, templateId, tTemplateData, test)

	return
}

func sendEmail(emailPtr *mail.SGMailV3, key, host string) (response *rest.Response, errorInfo errs.ErrorInfo) {

	request := sendgrid.GetRequest(key, SENDGRID_ENDPOINT, host)
	// request.Method = constants.HTTP_POST
	request.Body = mail.GetRequestBody(emailPtr)
	response, errorInfo.Error = sendgrid.API(request)

	return
}

// sendGridGetKey
// NOTE: This is a critical start-up function that enforce having the SendGrid key file available.
// This retrieves the Stripe key and sets the 'stripe.Key' variable.
func sendGridGetKey(sendgridFQN string) (sendGrid SendGridHelper, errorInfo errs.ErrorInfo) {

	var (
		tStripe []byte
	)

	if tStripe, errorInfo.Error = os.ReadFile(sendgridFQN); errorInfo.Error != nil {
		// errorInfo.Error = errs.ErrServiceFailedSendGrid
		// errorInfo.AdditionalInfo = fmt.Sprintf("SendGrid key file: %v", sendgridFQN)
		// errs.PrintError(errorInfo)
	} else {
		if errorInfo.Error = json.Unmarshal(tStripe, &sendGrid); errorInfo.Error != nil {
			// errorInfo.Error = errs.ErrJSONInvalid
			// errorInfo.AdditionalInfo = fmt.Sprintf("SendGrid JSON file: %v", sendgridFQN)
			// errs.PrintError(errorInfo)
		}
	}

	return
}

// validateEmailAddress
func validateEmailAddress(emailAddress string) (errorInfo errs.ErrorInfo) {

	var (
		mx []*net.MX
	)

	if len(emailAddress) < 3 || len(emailAddress) > 254 {
		errorInfo.Error = errors.New("The email address length must be greater than 2 and less than 255.")
		log.Println(errorInfo.Error.Error())
	} else {
		if emailRegex.MatchString(emailAddress) {
			parts := strings.Split(emailAddress, "@")
			if mx, errorInfo.Error = net.LookupMX(parts[1]); errorInfo.Error != nil || len(mx) == 0 {
				errorInfo.Error = errors.New(fmt.Sprintf("The email address failed the Domain: '%v' lookup.", parts[1]))
				log.Println(errorInfo.Error.Error())
			}
		} else {
			errorInfo.Error = errors.New(fmt.Sprintf("The email address '%v' is invalid.", emailAddress))
			log.Println(errorInfo.Error.Error())
		}
	}

	return
}

// validateSubject
// ToDo Add profanity checking service for subject line
func validateSubject(subject string) (errorInfo errs.ErrorInfo) {

	if len(subject) < 5 || len(subject) > 78 {
		errorInfo.Error = errors.New("The email subject length must be greater than 4 and less than 79 characters.")
		log.Println(errorInfo.Error.Error())
	}

	return
}

// Private Functions

// loadConfig - loads the SendGrid service configuration from a YAML file.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.NewErrorInfo
//	Verifications: hlps.CheckValueNotEmpty, yaml.Unmarshal
func loadConfig(configFilename string) (config SendGridConfig, errorInfo errs.ErrorInfo) {

	var (
		tConfigData []byte
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_SENDGRID, configFilename, ctv.FN_CONFIG_FILENAME); errorInfo.Error != nil {
		return
	}

	if tConfigData, errorInfo.Error = os.ReadFile(hlps.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_SENDGRID, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	if errorInfo.Error = yaml.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_SENDGRID, ctv.LBL_CONFIG_EXTENSION_FILENAME, configFilename))
		return
	}

	return
}

// validateSendGridConfig - validates the SendGridConfig and checks that required fields are not empty.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter
//	Verifications: hlps.CheckValueNotEmpty
func validateSendGridConfig(config SendGridConfig, environment string) (errorInfo errs.ErrorInfo) {

	// The config.DebugModeOn is either true or false. No need to check the value.
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_SENDGRID, config.KeyFQN, ctv.LBL_KEY_FILENAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_SENDGRID, config.DefaultSenderAddress, ctv.LBL_DEFAULT_SENDER_ADDRESS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_SENDGRID, config.DefaultSenderName, ctv.LBL_DEFAULT_SENDER_NAME); errorInfo.Error != nil {
		return
	}

	return
}
