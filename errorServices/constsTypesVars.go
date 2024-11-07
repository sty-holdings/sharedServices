// Package sty_shared
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by Scott Yacko / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	www.sty-holdings.com

	01-2024
	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package sty_shared

import (
	// Add imports here

	"errors"
)

//goland:noinspection ALL
const (
	ERROR    = "error"
	NO_ERROR = "no error"
	//
	//
	FORMAT_EXPECTED_ERROR     = "%s Failed: Was expecting an err. Additional Info: %s"
	FORMAT_EXPECTING_NO_ERROR = "%s Failed: Wasn't expecting an err. ERROR: %s"
	FORMAT_UNEXPECTED_ERROR   = "%s Failed: Unexpected err. ERROR: %s"
	//
	// Messages
	ACCESS_TOKEN_MISSING               = "No access token was provided."
	ADDRESS_STREET_MISSING             = "The profile address is missing information. Please review the street, city, state, and zip code."
	ADDRESS_CITY_MISSING               = "The profile address is missing information. Please review the street, city, state, and zip code."
	ADDRESS_STATE_MISSING              = "The profile address is missing information. Please review the street, city, state, and zip code."
	ADDRESS_ZIP_CODE_MISSING           = "The profile address is missing information. Please review the street, city, state, and zip code."
	ATTEMPTS_EXCEEDED                  = "LimitExceededException: Attempt limit exceeded, please try after some time."
	AWS_INVALID_SSM_PARAMETERS         = "GetParameters returned invalid parameter names. Check AWS SSM Parameter Manager."
	BASE64_INVALID                     = "The base64 string is invalid."
	FLOAT_INVALID                      = "The float64 value is invalid."
	INTEGER_INVALID                    = "The integer64 value is invalid."
	BUCKET_NOT_FOUND                   = "The bucket was not found."
	BUFFER_EMPTY                       = "The buffer is empty"
	BUNDLE_ALREADY_EXISTS              = "The bundle already exists in the system."
	BUNDLE_MISSING                     = "The bundle is not in the system."
	CLIENT_SECRET_MISSING              = "The client secret has not been provided.)"
	COGNITO_SECRET_BLOCK_INVALID       = "Unable to decode challenge parameter 'SECRET_BLOCK'."
	COGNITO_USER_NAME_MISSING          = "Username is not in the Cognito user pool."
	COGNITO_USERPOOL_ID_INVALID        = "User Pool ID must be in format: '<region>_<pool name>'"
	CONFIG_FILE_MISSING                = "Not able to read the supplied config file. "
	CURRENCY_INVALID                   = "The curreny type is not supported. See https://github.com/sty-holdings/sharedServices/v2024/constsTypesVars"
	DECODE_STRING_FAILED               = "Unable to decode the string."
	DIRECTORY_MISSING                  = "The directory does not exist."
	DIRECTORY_NOT_FULLY_QUALIFIED      = "The directory doesn't start and end with slash."
	DOCUMENT_NOT_FOUND                 = "The document was not found."
	DOCUMENTS_NONE_FOUND               = "No documents were found."
	DOMAIN_INVALID                     = "The domain value is invalid."
	ENVIRNOMENT_INVALID                = "The environment value is invalid."
	ERROR_MISSING                      = "ERROR MISSING"
	EXTENSION_INVALID                  = "The extensioin name is invalid."
	EXTRACT_KEY_FAILED                 = "Extracting the key has failed."
	FALSE_SHOULD_BE_TRUE               = "The result should have been true."
	FILE_CREATION_FAILED               = "Create the file failed."
	FILE_MISSING                       = "The file doesn't exist."
	FILE_REMOVAL_FAILED                = "The file was not deleted."
	FILE_UNREADABLE                    = "[ERROR} The file is not readable."
	FIREBASE_GC_PROJECT_ID_MISSING     = "No Firebase project id was not provided."
	NAME_FIRST_MISSING                 = "The first name is missing."
	NAME_LAST_MISSING                  = "The last name is missing."
	NAME_FIRST_LAST_MISSING            = "Either the first or the last name is missing."
	GIN_MODE_INVALID                   = "The Gin mode is invalid."
	GREATER_THAN_ZERO                  = "The value must be greater than zero."
	HTTP_REQUEST_FALIED                = "The HTTP request failed with a non-200 status code."
	HTTP_SECURE_SERVER_FAILED          = "The HTTP services secure server failed."
	JSON_GENERATION_FAILED             = "Failed to generate JSON payload"
	JSON_INVALID                       = "The JSON provided is invalid"
	JWT_MISSING                        = "JWT token is missing."
	LESS_THAN_ZERO                     = "The value must be less than zero."
	MAP_IS_EMPTY                       = "Provided map is not populated."
	MAP_MISSING_KEY                    = "Provided map has a nil or empty key."
	MAP_MISSING_VALUE                  = "Provided map has a nil or empty value."
	MAX_THREADS_INVALID                = "The config file max threads value is less than 1."
	NATS_CONNECTION_FAILED             = "Connecting to NATS server failed."
	NATS_URL_INVALID                   = "The NATS URL value is invalid."
	NATS_ZERO                          = "The port value is zero. This is not allowed. Recommended values are 4222 and 9222."
	MESSAGE_JSON_INVALID               = "The message body is not valid JSON."
	MESSAGE_NAMESPACE_INVALID          = "The Message namespace value is invalid."
	NOT_DIVISIBLE_N                    = "Calculate value must not be divisable by N."
	PARSE_BIG_INT_FAILED               = "Unable to parse the value provided."
	RECIPIENTTYPEINVALID               = "Recipient type is invalid."
	REDIRECT_MODE_MISSING              = "The redirect mode is missing."
	REDIRECT_MODE_INVALID              = "The redirect mode is invalid."
	REFRESH_TOO_SOON                   = "Too soon to refresh balances."
	REQUESTOR_ID_MISSING               = "The requestor id is missing."
	REQUIRED_ARGUMENT_MISSING          = "A required argument is empty."
	REQUIRED_FILE_MISSING              = "A required file is missing."
	RETRY_LIMIT_HIT                    = "You have tried too many times. Please try again in 15 mins or contact support@sty-holdings.com."
	PHONE_NUMBER_AREA_CODE_MISSING     = "The area code is missing."
	PHONE_NUMBER_COUNTRY_CODE_MISSING  = "The country code is missing."
	PHONE_NUMBER_MISSING               = "The phone number is missing."
	PID_FILE_EXISTS                    = "A PID file already exists. Delete the 'server.pid' file in '.run' directory and start the server again."
	PLAID_INVALID_PUBLIC_TOKEN         = "INVALID_PUBLIC_TOKEN" // DO NOT change this, it is used to test a condition
	POINTER_MISSING                    = "You must pass a pointer. Nil is not valid!"
	POSTGRES_SSL_MODE                  = "Only disable, allow, prefer and required are supported."
	POSTGRES_CONN_FALIED               = "No database connection has been established"
	POSTGRES_CONN_EMPTY                = "Database connection is empty"
	PROGRAM_NAME_MISSING               = "The program name in main.go is empty."
	SERVER_CONFIGURATION_INVALID       = "The setting in the configuration file are inconsistant."
	SERVER_NAME_MISSING                = "The server name in main.go is empty."
	SERVICE_FAILED_AWS                 = "AWS service has failed. Investigate right away!"
	SERVICE_FAILED_COGNITO             = "Cognito service has failed. Investigate right away!"
	SERVICE_FAILED_FIREBASE            = "FIREBASE service has failed. Investigate right away!"
	SERVICE_FAILED_FIRESTORE           = "FIRESTORE service has failed. Investigate right away!"
	SERVICE_FAILED_NATSCONNECT         = "NATS Connect service has failed. Investigate right away!"
	SERVICE_FAILED_PLAID               = "PLAID service has failed. Investigate right away!"
	SERVICE_FAILED_POSTGRES            = "POSTGRES service has failed. Investigate right away!"
	SERVICE_FAILED_SENDGRID            = "SENDGRID service has failed. Investigate right away!"
	SERVICE_FAILED_STRIPE              = "STRIPE service has failed. Investigate right away!"
	SERVICE_FAILED_SYNADIA             = "SYNADIA service has failed. Investigate right away!"
	SET_STRING_FAILED                  = "Unable to process value using SetString."
	SHORT_URL_ALREADY_EXISTS           = "The short URL already exists in the system."
	SHORT_URL_MISSING                  = "The short URL is not in the system."
	SIGNAL_UNKNOWN                     = "Unknown signal was caught and ignored."
	SRP_A_MOD_N_ZERO                   = "A mod N cannot be 0"
	SRP_B_MOD_N_ZERO                   = "B mod N cannot be 0"
	STRIPE_AMOUNT_INVALID              = "The amount must be a positive number. See https://docs.stripe.com/api/payment_intents."
	STRIPE_CURRENCY_INVALID            = "The curreny type is not supported. See https://docs.stripe.com/api/payment_intents."
	STRIPE_CUSTOMER_FAILED             = "Creating a Stripe customer failed."
	STRIPE_PAYMENT_INTENT_ID_EMPTY     = "An empty payment intent id is not allowed. See https://docs.stripe.com/api/payment_intents."
	STRIPE_PAYMENT_METHOD_EMPTY        = "An empty payment method is not allowed. See https://docs.stripe.com/testing?testing-method=payment-methods#cards."
	STRIPE_PAYMENT_METHOD_INVALID      = "The payment method is not support by NATS Connect. See https://docs.stripe.com/testing?testing-method=payment-methods#cards."
	STRIPE_PAYMENT_METHOD_TYPE_EMPTY   = "An empty payment method type is not allowed. See https://docs.stripe.com/api/payment_methods/object#payment_method_object-type."
	STRIPE_PAYMENT_METHOD_TYPE_INVALID = "The payment method type is not support by NATS Connect. See https://docs.stripe.com/api/payment_methods/object#payment_method_object-type."
	STRIPE_KEY_INVALID                 = "The stripe key is invalid. See https://docs.stripe.com/api/payment_intents source."
	STRIPE_METHOD_TYPE_UNSUPPORTED     = "The payment method is not support. To request support, contact support@sty-holdings.com."
	STRIPE_ONE_TIME_CODE_FAILED        = "Generating the Stripe One Time Use Token failed."
	STRIPE_OUT_NOT_SUPPORTED           = "Transfers out using Stripe are not supported."
	STRIPE_SOURCE_INVALID              = "The provided source is invalid. See https://docs.stripe.com/api/payment_intents."
	STRUCT_INVALID                     = "Provided object is not a struct."
	STYH_CLIENT_ID_EMPTY               = "An empty STYH Client Id is not allowed."
	STYH_USERNAME_EMPTY                = "An empty STYH Username is not allowed."
	SUBJECTS_MISSING                   = "No subject(s) have been defined for the NATS extension."
	SUBJECT_SUBSCRIPTION_FAILED        = "Unable to subscribe to the subject."
	SYNADIA_BASE_URL_EMPTY             = "An empty base URL is not allowed."
	SYNADIA_KEY_INVALID                = "The Synadia access token is invalid. See https://docs.synadia.com/platform or https://cloud.synadia.com/profile/personal-access-tokens"
	SUBJECT_INVALID                    = "The subject is invalid."
	TLS_FILES_MISSING                  = "TLS files are missing."
	TOKEN_CLAIMS_INVALID               = "The token claims are invalid."
	TOKEN_EXPIRED                      = "The token has expired."
	TOKEN_INVALID                      = "The token is invalid."
	TRANSFER_AMOUNT_INVALID            = "The transfer amount is not support for this transfer method!"
	TRANSFER_IN_NOT_ALLOWED            = "Transferring money is not allowed for this transfer method."
	TRANSFER_METHOD_INVALID            = "The transfer method is not support! (Transfer Method is case insensitive)"
	TRANSFER_OUT_NOT_ALLOWED           = "Transferring money out is not allowed for this transfer method."
	TRUE_SHOULD_BE_FALSE               = "The result should have been false."
	UNABLE_READ_FILE                   = "Unable to read file."
	UNAUTHORIZED_REQUEST               = "You are not authorized to use this system."
	UNCONFIRMED_EMAIL                  = "Users email has not been confirmed."
	UNEXPECTED_ERROR                   = "The system has experienced an unexpected issue. Investigate right away!"
	UNSUPPORTED_SUBJECT                = "This subject is not supported."
	UNSUPPORTED_TRANSFER_METHOD        = "The transfer method is not supported."
	UNMARSHAL_FAILED                   = "Unable to unmarshal data"
	USER_MISSING                       = "The user is not in the system."
	USER_ALREADY_EXISTS                = "The user already exists in the system."
	USER_ALREADY_CONFIRMED_EMAIL       = "The user has already been confirmed by email."
	USER_ALREADY_CONFIRMED_PHONE       = "The user has already been confirmed by phone."
	USER_BUNDLE_ALREADY_EXISTS         = "The user bundle already exists in the system."
	USER_BUNDLE_MISSING                = "The user bundle is not in the system."
	VERSION_INVALID                    = "The software version is invalid. Use @env GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.version=$(" +
		"VERSION)\" -o ${ROOT_DIRECTORY}/servers/${SERVER_NAME}/bin/${SERVER_NAME} ${ROOT_DIRECTORY}/servers/${SERVER_NAME}/main.go"
	ZERO_INVALID = "A value of zero is invalid."
	//
	// String that are used to determine third party error messages
	USER_DOES_NOT_EXIST           = "User does not exist."
	NOT_FOUND                     = "not found"
	UNKNOWN                       = "UNKNOWN"
	JWT_TOKEN_SIGNATURE_INVALID   = "token signature is invalid: crypto/rsa: verification error"
	FIREBASE_AUTH_BAD_CREDENTIALS = "cannot read credentials"
	//
	// Testing Strings
	TEST_STRING = "TEST STRING"
)

//goland:noinspection ALL
var (
	ErrAccessTokenMissing             = errors.New(ACCESS_TOKEN_MISSING)
	ErrAddressCityMissing             = errors.New(ADDRESS_CITY_MISSING)
	ErrAddressStateMissing            = errors.New(ADDRESS_STATE_MISSING)
	ErrAddressStreetMissing           = errors.New(ADDRESS_STREET_MISSING)
	ErrAddressZipCodeMissing          = errors.New(ADDRESS_ZIP_CODE_MISSING)
	ErrAlreadyConfirmedEmail          = errors.New(USER_ALREADY_CONFIRMED_EMAIL)
	ErrAlreadyConfirmedPhone          = errors.New(USER_ALREADY_CONFIRMED_PHONE)
	ErrAttemptsExceeded               = errors.New(ATTEMPTS_EXCEEDED)
	ErrAWSInvalidSSMParameters        = errors.New(AWS_INVALID_SSM_PARAMETERS)
	ErrBase64Invalid                  = errors.New(BASE64_INVALID)
	ErrBucketNotFound                 = errors.New(BUCKET_NOT_FOUND)
	ErrBufferEmpty                    = errors.New(BUFFER_EMPTY)
	ErrBundleAlreadyExists            = errors.New(BUNDLE_ALREADY_EXISTS)
	ErrBundleMissing                  = errors.New(BUNDLE_MISSING)
	ErrClientSecretBlockInvalid       = errors.New(COGNITO_SECRET_BLOCK_INVALID)
	ErrClientSecretMissing            = errors.New(CLIENT_SECRET_MISSING)
	ErrCognitoUsernameMissing         = errors.New(COGNITO_USER_NAME_MISSING)
	ErrCognitoUserpoolIdInvalid       = errors.New(COGNITO_USERPOOL_ID_INVALID)
	ErrConfigFileMissing              = errors.New(CONFIG_FILE_MISSING)
	ErrCurrencyInvalid                = errors.New(CURRENCY_INVALID)
	ErrDecodeStringFailed             = errors.New(DECODE_STRING_FAILED)
	ErrDirectoryMissing               = errors.New(DIRECTORY_MISSING)
	ErrDirectoryNotFullyQualified     = errors.New(DIRECTORY_NOT_FULLY_QUALIFIED)
	ErrDocumentNotFound               = errors.New(DOCUMENT_NOT_FOUND)
	ErrDocumentsNoneFound             = errors.New(DOCUMENTS_NONE_FOUND)
	ErrDomainInvalid                  = errors.New(DOMAIN_INVALID)
	ErrEnvironmentInvalid             = errors.New(ENVIRNOMENT_INVALID)
	ErrErrorMissing                   = errors.New(ERROR_MISSING)
	ErrExtensionInvalid               = errors.New(EXTENSION_INVALID)
	ErrExtractKeysFailure             = errors.New(EXTRACT_KEY_FAILED)
	ErrFalseShouldBeTrue              = errors.New(FALSE_SHOULD_BE_TRUE)
	ErrFileCreationFailed             = errors.New(FILE_CREATION_FAILED)
	ErrFileMissing                    = errors.New(FILE_MISSING)
	ErrFileRemovalFailed              = errors.New(FILE_REMOVAL_FAILED)
	ErrFileUnreadable                 = errors.New(FILE_UNREADABLE)
	ErrFirebaseProjectMissing         = errors.New(FIREBASE_GC_PROJECT_ID_MISSING)
	ErrFloatInvalid                   = errors.New(FLOAT_INVALID)
	ErrGinModeInvalid                 = errors.New(GIN_MODE_INVALID)
	ErrGreatThanZero                  = errors.New(GREATER_THAN_ZERO)
	ErrHTTPRequestFalied              = errors.New(HTTP_REQUEST_FALIED)
	ErrHTTPSecureServerFailed         = errors.New(HTTP_SECURE_SERVER_FAILED)
	ErrIntegerInvalid                 = errors.New(INTEGER_INVALID)
	ErrJSONGenerationFailed           = errors.New(JSON_GENERATION_FAILED)
	ErrJSONInvalid                    = errors.New(JSON_INVALID)
	ErrJWTMissing                     = errors.New(JWT_MISSING)
	ErrJWTTokenSignatureInvalid       = errors.New(JWT_TOKEN_SIGNATURE_INVALID)
	ErrLessThanEqualZero              = errors.New(LESS_THAN_ZERO)
	ErrMapIsEmpty                     = errors.New(MAP_IS_EMPTY)
	ErrMapIsMissingKey                = errors.New(MAP_MISSING_KEY)
	ErrMapIsMissingValue              = errors.New(MAP_MISSING_VALUE)
	ErrMaxThreadsInvalid              = errors.New(MAX_THREADS_INVALID)
	ErrMessageJSONInvalid             = errors.New(MESSAGE_JSON_INVALID)
	ErrMessageNamespaceInvalid        = errors.New(MESSAGE_NAMESPACE_INVALID)
	ErrNameFirstLastMissing           = errors.New(NAME_FIRST_LAST_MISSING)
	ErrNameFirstMissing               = errors.New(NAME_FIRST_MISSING)
	ErrNameLastMissing                = errors.New(NAME_LAST_MISSING)
	ErrNATSConnectionFailed           = errors.New(NATS_CONNECTION_FAILED)
	ErrNatsPortInvalid                = errors.New(NATS_ZERO)
	ErrNATSURLInvalid                 = errors.New(NATS_URL_INVALID)
	ErrNotDivisibleN                  = errors.New(NOT_DIVISIBLE_N)
	ErrParseBigIntFailed              = errors.New(PARSE_BIG_INT_FAILED)
	ErrPhoneNumberAreaCodeMissing     = errors.New(PHONE_NUMBER_AREA_CODE_MISSING)
	ErrPhoneNumberCountryCodeMissing  = errors.New(PHONE_NUMBER_COUNTRY_CODE_MISSING)
	ErrPhoneNumberMissing             = errors.New(PHONE_NUMBER_MISSING)
	ErrPIDFileExists                  = errors.New(PID_FILE_EXISTS)
	ErrPlaidInvalidPublicToken        = errors.New(PLAID_INVALID_PUBLIC_TOKEN)
	ErrPointerMissing                 = errors.New(POINTER_MISSING)
	ErrPostgresConnEmpty              = errors.New(POSTGRES_CONN_EMPTY)
	ErrPostgresConnFailed             = errors.New(POSTGRES_CONN_FALIED)
	ErrPostgresSSLMode                = errors.New(POSTGRES_SSL_MODE)
	ErrProgramNameMissing             = errors.New(PROGRAM_NAME_MISSING)
	ErrRecipientTypeInvalid           = errors.New(RECIPIENTTYPEINVALID)
	ErrRedirectModeInvalid            = errors.New(REDIRECT_MODE_INVALID)
	ErrRedirectModeMissing            = errors.New(REDIRECT_MODE_MISSING)
	ErrRefreshTooSoon                 = errors.New(REFRESH_TOO_SOON)
	ErrRequestorIdMissing             = errors.New(REQUESTOR_ID_MISSING)
	ErrRequiredArgumentMissing        = errors.New(REQUIRED_ARGUMENT_MISSING)
	ErrRequiredFileMissing            = errors.New(REQUIRED_FILE_MISSING)
	ErrRetryLimitHit                  = errors.New(RETRY_LIMIT_HIT)
	ErrServerConfigurationInvalid     = errors.New(SERVER_CONFIGURATION_INVALID)
	ErrServerNameMissing              = errors.New(SERVER_NAME_MISSING)
	ErrServiceFailedAWS               = errors.New(SERVICE_FAILED_AWS)
	ErrServiceFailedCognito           = errors.New(SERVICE_FAILED_COGNITO)
	ErrServiceFailedFIREBASE          = errors.New(SERVICE_FAILED_FIREBASE)
	ErrServiceFailedFIRESTORE         = errors.New(SERVICE_FAILED_FIRESTORE)
	ErrServiceFailedNATSCONNECT       = errors.New(SERVICE_FAILED_NATSCONNECT)
	ErrServiceFailedPLAID             = errors.New(SERVICE_FAILED_PLAID)
	ErrServiceFailedPOSTGRES          = errors.New(SERVICE_FAILED_POSTGRES)
	ErrServiceFailedSENDGRID          = errors.New(SERVICE_FAILED_SENDGRID)
	ErrServiceFailedSTRIPE            = errors.New(SERVICE_FAILED_STRIPE)
	ErrServiceFailedSYNADIA           = errors.New(SERVICE_FAILED_SYNADIA)
	ErrSetStringFailed                = errors.New(SET_STRING_FAILED)
	ErrShortURLMissing                = errors.New(SHORT_URL_MISSING)
	ErrSignalUnknown                  = errors.New(SIGNAL_UNKNOWN)
	ErrSRPAModNZero                   = errors.New(SRP_A_MOD_N_ZERO)
	ErrSRPBModNZero                   = errors.New(SRP_B_MOD_N_ZERO)
	ErrStripeAmountInvalid            = errors.New(STRIPE_AMOUNT_INVALID)
	ErrStripeCreateCustomerFailed     = errors.New(STRIPE_CUSTOMER_FAILED)
	ErrStripeCurrencyInvalid          = errors.New(STRIPE_CURRENCY_INVALID)
	ErrStripeKeyInvalid               = errors.New(STRIPE_KEY_INVALID)
	ErrStripeMethodTypeUnsupported    = errors.New(STRIPE_METHOD_TYPE_UNSUPPORTED)
	ErrStripeOneTimeCodeFailed        = errors.New(STRIPE_ONE_TIME_CODE_FAILED)
	ErrStripeOutNotSupported          = errors.New(STRIPE_OUT_NOT_SUPPORTED)
	ErrStripePaymentIntentIdEmpty     = errors.New(STRIPE_PAYMENT_INTENT_ID_EMPTY)
	ErrStripePaymentMethodEmpty       = errors.New(STRIPE_PAYMENT_METHOD_EMPTY)
	ErrStripePaymentMethodInvalid     = errors.New(STRIPE_PAYMENT_METHOD_INVALID)
	ErrStripePaymentMethodTypeEmpty   = errors.New(STRIPE_PAYMENT_METHOD_TYPE_EMPTY)
	ErrStripePaymentMethodTypeInvalid = errors.New(STRIPE_PAYMENT_METHOD_TYPE_INVALID)
	ErrStripeSourceInvalid            = errors.New(STRIPE_SOURCE_INVALID)
	ErrStructInvalid                  = errors.New(STRUCT_INVALID)
	ErrSubjectInvalid                 = errors.New(SUBJECT_INVALID)
	ErrSubjectsMissing                = errors.New(SUBJECTS_MISSING)
	ErrSubjectSubscriptionFailed      = errors.New(SUBJECT_SUBSCRIPTION_FAILED)
	ErrSynadiaBaseURLEmpty            = errors.New(SYNADIA_BASE_URL_EMPTY)
	ErrSynadiaKeyInvalid              = errors.New(SYNADIA_KEY_INVALID)
	ErrSYTHClientIdEmpty              = errors.New(STYH_CLIENT_ID_EMPTY)
	ErrSYTHUserNameEmpty              = errors.New(STYH_USERNAME_EMPTY)
	ErrTLSFilesMissing                = errors.New(TLS_FILES_MISSING)
	ErrTokenClaimsInvalid             = errors.New(TOKEN_CLAIMS_INVALID)
	ErrTokenExpired                   = errors.New(TOKEN_EXPIRED)
	ErrTokenInvalid                   = errors.New(TOKEN_INVALID)
	ErrTransferAmountInvalid          = errors.New(TRANSFER_AMOUNT_INVALID)
	ErrTransferInNotAllowed           = errors.New(TRANSFER_IN_NOT_ALLOWED)
	ErrTransferMethodInvalid          = errors.New(TRANSFER_METHOD_INVALID)
	ErrTransferOutNotAllowed          = errors.New(TRANSFER_OUT_NOT_ALLOWED)
	ErrTrueShouldBeFalse              = errors.New(TRUE_SHOULD_BE_FALSE)
	ErrUnableReadFile                 = errors.New(UNABLE_READ_FILE)
	ErrUnauthorizedRequest            = errors.New(UNAUTHORIZED_REQUEST)
	ErrUnconfirmedEmail               = errors.New(UNCONFIRMED_EMAIL)
	ErrUnexpectedError                = errors.New(UNEXPECTED_ERROR)
	ErrUnmarshalFailed                = errors.New(UNMARSHAL_FAILED)
	ErrUnsupportedSubject             = errors.New(UNSUPPORTED_SUBJECT)
	ErrUnsupportedTransferMethod      = errors.New(UNSUPPORTED_TRANSFER_METHOD)
	ErrUserAccountMissing             = errors.New(USER_MISSING)
	ErrUserAlreadyExists              = errors.New(USER_ALREADY_EXISTS)
	ErrUserBundleAlreadyExists        = errors.New(BUNDLE_ALREADY_EXISTS)
	ErrUserBundleMissing              = errors.New(BUNDLE_MISSING)
	ErrVersionInvalid                 = errors.New(VERSION_INVALID)
	ErrZeroInvalid                    = errors.New(ZERO_INVALID)
)
