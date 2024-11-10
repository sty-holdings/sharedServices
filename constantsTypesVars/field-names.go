package sharedServices

//goland:noinspection ALL
const (
	FN_ACCOUNTS                           = "accounts"
	FN_ACCOUNT_ID                         = "account_id"
	FN_ACCOUNTS_STRIPE_BANK_ACCOUNT_TOKEN = "accounts.stripe_bank_account_tokens."
	FN_ADDED_PHONE                        = "added_phone"
	FN_ADDED_SAVINGS_GOAL                 = "added_savings_goal"
	FN_ADDITIONAL_TRANSFER_INFO           = "additional_transfer_info"
	FN_AGG_TOTAL_LOAN_AMOUNT              = "agg_total_loan_amount"
	FN_AGG_TWI_SCORE                      = "agg_twi_score"
	FN_AMOUNT_IN_PENNIES                  = "amount_in_pennies"
	FN_AREA_CODE                          = "area_code"
	FN_AREA_CODE_CUSTOM                   = "custom:areaCode"
	FN_AUDIENCE_CAP                       = "AUDIENCE"
	FN_AVAILABLE_FUNDS                    = "available_funds"
	FN_AWS_ACCESS_KEY_ID                  = "aws_access_key_id"
	FN_AWS_ACCESS_KEY_ID_CAP              = "AWS_ACCESS_KEY_ID"
	FN_AWS_ACCOUNT_INFO_FILENAME          = "aws_account_info_filename"
	FN_AWS_CLIENT_ID                      = "client_id"
	FN_AWS_IDENTITY_ID                    = "identityId"
	FN_AWS_IDENTITY_POOL_ID               = "identity_pool_id"
	FN_AWS_PROFILE                        = "profile"
	FN_AWS_PROFILE_CAP                    = "PROFILE"
	FN_AWS_REGION                         = "region"
	FN_AWS_REGION_CAP                     = "REGION"
	FN_AWS_SECRET_ACCESS_KEY              = "aws_secret_access_key"
	FN_AWS_SECRET_ACCESS_KEY_CAP          = "AWS_SECRET_ACCESS_KEY"
	FN_AWS_USERPOOL_ID                    = "userpool_id"
	FN_AWS_USERPOOL_ID_CAP                = "USERPOOL_ID"
	FN_AWS_USERPOOL_NAME                  = "userpool_name"
	FN_AWS_USERPOOL_NAME_CAP              = "USERPOOL_NAME"
	FN_BALANCE                            = "Balance"
	FN_BIRTHDATE                          = "birthdate"
	FN_BUNDLE_ID                          = "bundle_id"
	FN_BUNDLE_TITLE                       = "bundle_title"
	FN_CERT_KID                           = "kid"
	FN_CITY                               = "city"
	FN_CLIENT_ID                          = "client_id"
	FN_COGNITO_USERNAME                   = "cognito:username"
	FN_CREATE_TIMESTAMP                   = "create_timestamp"
	FN_CREDENTIALS_FILENAME               = "credentials_filename"
	FN_DATASTORE                          = "datastore"
	FN_DATATYPE_INT8                      = "int8"
	FN_DATATYPE_INT16                     = "int16"
	FN_DATATYPE_INT32                     = "int32"
	FN_DATATYPE_INT64                     = "int64"
	FN_DATATYPE_UINT8                     = "uint8"
	FN_DATATYPE_UINT16                    = "uint16"
	FN_DATATYPE_UINT32                    = "uint32"
	FN_DATATYPE_UINT64                    = "uint64"
	FN_DATATYPE_INT                       = "int"
	FN_DATATYPE_UINT                      = "uint"
	FN_DATATYPE_UINTPTR                   = "uintptr"
	FN_DATATYPE_FLOAT32                   = "float32"
	FN_DATATYPE_FLOAT64                   = "float64"
	FN_DATATYPE_COMPLEX64                 = "complex64"
	FN_DATATYPE_COMPLEX128                = "complex128"
	FN_DATATYPE_BOOL                      = "bool"
	FN_DATATYPE_STRING                    = "string"
	FN_DEBUG_MODE_ON                      = "debug_mode_on"
	FN_DESCRIPTION                        = "description"
	FN_DOCUMENT_ID                        = "document_id"
	FN_DURATION                           = "Duration"
	FN_EMAIL                              = "email"
	FN_EMAIL_VERIFIED                     = "email_verified"
	FN_ENVIRONMENT                        = "environment"
	FN_EXPIRY_TIMESTAMP                   = "expiry_timestamp"
	FN_FAMILY_NAME                        = "family_name"
	FN_FILENAME                           = "filename"
	FN_FIRST_NAME                         = "first_name"
	FN_FEDERAL_TAX_ID                     = "federal_tax_id"
	FN_FEDERAL_TAX_ID_IS_SET              = "federal_tax_id_is_set"
	FN_GCP_CREDENTIAL_FILENAME            = "gcp_credential_filename"
	FN_GIVEN_NAME                         = "given_name"
	FN_GEMINI_MAX_OUTPUT_TOKENS           = "gemini_max_output_tokens"
	FN_GEMINI_MODEL_NAME                  = "gemini_model_name"
	FN_GEMINI_OUTPUT_FORMAT               = "gemini_output_format"
	FN_GEMINI_SET_TOP_K                   = "gemini_set_top_k"
	FN_GEMINI_SET_TOP_PROBABILITY         = "gemini_set_top_probability"
	FN_GEMINI_SYSTEM_INSTRUCTION          = "gemini_system_instruction"
	FN_GEMINI_TEMPERATURE                 = "gemini_temperature"
	FN_GIN_MODE                           = "gin_mode"
	FN_GIN_URL_PORT                       = "gin_url_port"
	FN_GOOGLE_LOCATION                    = "google_location"
	FN_GOOGLE_PROJECT_ID                  = "google_project_id"
	FN_HTTP_TLS_INFO                      = "http_tls_info"
	FN_INSTITUTIONS                       = "institutions"
	FN_INSTITUTION_NAME                   = "institution_name"
	FN_INSTITUTION_ACCOUNT                = "institution_account"
	FN_IS_BUNDLE_LOCKED                   = "is_bundle_locked"
	FN_ISSUER                             = "issuer"
	FN_JWT                                = "JWT"
	FN_KEY                                = "key"
	FN_KEY_DATA                           = "keyData"
	FN_KEYSET_URL                         = "keySetURL"
	FN_LAST_NAME                          = "last_name"
	FN_LAST_REFRESHED                     = "last_refreshed"
	FN_LAST_UPDATE_TIMESTAMP              = "last_update_timestamp"
	FN_LINKED_BANK                        = "linked_bank"
	FN_LOAD_EXTENSIONS                    = "load_extensions"
	FN_LOAN_TYPE                          = "loan_type"
	FN_LOANED_AMOUNT_INVESTED             = "loaned_amount_invested"
	FN_LOANED_AMOUNT_RETURNED             = "loaned_amount_returned"
	FN_LOANS                              = "loans"
	FN_LOCKUP_END_DATE                    = "lockup_end_date"
	FN_LOCKUP_MONTHS                      = "lockup_months"
	FN_LOCKUP_START_DATE                  = "lockups_start_date"
	FN_LOG_DIRECTORY                      = "log_directory"
	FN_MAX_ALLOCATION                     = "max_allocation"
	FN_MAX_THREADS                        = "max_threads"
	FN_MESSAGE_ENVIRONMENT                = "message_environment"
	FN_MESSAGE_NAMESPACE                  = "message_namespace"
	FN_MESSAGE_REGISTRY                   = "message_registry"
	FN_MIN_ALLOCATION                     = "min_allocation"
	FN_NAVIGATION                         = "navigation"
	FN_NICKNAME                           = "nickname"
	FN_OFFERED_INTEREST_RATE              = "offered_interest_rate"
	FN_OFFICIAL_NAME                      = "official_name"
	FN_PASSWORD                           = "password"
	FN_PAYMENT_FREQENCY                   = "payment_frequency"
	FN_PERIOD                             = "peroid"
	FN_PLAID_ACCESS_TOKEN                 = "plaid_access_token"
	FN_PLAID_ACCOUNT                      = "plaid_account"
	FN_PLAID_ACCOUNTS                     = "plaid_accounts"
	FN_PLAID_INFO_FQN                     = "Plaid_Key_FQN"
	FN_PLAID_ITEM_ID                      = "plaid_item_id"
	FN_PHONE_NUMBER                       = "phone_number"
	FN_PHONE_NUMBER_CUSTOM                = "custom:phoneNumber"
	FN_PHONE_VERIFIED                     = "phone_verified"
	FN_PID_DIRECTORY                      = "pid_directory"
	FN_PORT                               = "port"
	FN_PARAMETER_TYPE                     = "parameterType"
	FN_PRIVATE_KEY                        = "PrivateKey"
	FN_PROVIDED_AML_PHOTO                 = "provided_aml_photo"
	FN_PURPOSE                            = "purpose"
	FN_REPORT_BALANCE                     = "report_balance"
	FN_REPORT_BALANCE_SOURCE              = "report_balance_source"
	FN_REQUESTOR_ID                       = "requestor_id"
	FN_RISK_RATING                        = "risk_rating"
	FN_SAVUP_TAKE                         = "savup_take"
	FN_SECRET_KEY                         = "secret_key"
	FN_SET_BANKER_PREFERENCES             = "set_banker_preferences"
	FN_SHORT_URL                          = "short_URL"
	FN_SIGNAL                             = "signal"
	FN_SKELETON_DIRECTORY                 = "skeleton_config_directory"
	FN_SERVER_INSTANCE_NUMBER             = "server_instance_number"
	FN_STATE                              = "state"
	FN_STATUS                             = "status"
	FN_STREET_ADDRESS                     = "street_address"
	FN_STRIPE_ACCESS_TOKEN                = "stripe_access_token"
	FN_STRIPE_CUSTOMER_ACCOUNT_ID         = "StripeCustomerAccountId"
	FN_STRIPE_LOCK                        = "stripe_lock"
	FN_STYH_CLIENT_ID                     = "styh_client_id"
	FN_STYH_CUSTOM_SECRET_KEY             = "custom:secret_key"
	FN_SUB                                = "sub"
	FN_SUBJECT                            = "subject"
	FN_TEMP_DIRECTORY                     = "temporary_directory_fqd"
	FN_TLS_CA_BUNDLE                      = "tls_ca_bundle"
	FN_TLS_CA_BUNDLE_FILENAME             = "tls_ca_bundle_filename"
	FN_TLS_CERTIFICATE                    = "tls_certificate"
	FN_TLS_CERTIFICATE_FILENAME           = "tls_certificate_filename"
	FN_TLS_PRIVATE_KEY                    = "tls_private_key"
	FN_TLS_PRIVATE_KEY_FILENAME           = "tls_private_key_filename"
	FN_TLS_INFO                           = "tls_info"
	FN_TOKEN                              = "token"
	FN_TOKEN_ID                           = "idToken"
	FN_TOKEN_ACCESS                       = "accessToken"
	FN_TOKEN_REFRESH                      = "refreshToken"
	FN_TOKEN_PAYLOAD                      = "tokenPayload"
	FN_TOKEN_TYPE                         = "tokenType"
	FN_TOTAL_OFFERING_SIZE                = "total_offering_size"
	FN_TRANSFER_DIRECTION                 = "transfer_direction"
	FN_TRANSFER_INSTITUTION_NAME          = "Transfer Bank:"
	FN_TRANSFER_METHOD                    = "transfer_method"
	FN_TRANSFER_STATUS                    = "transfer_status"
	FN_TRANSFERRED_FUNDS                  = "transferred_funds"
	FN_TWI_RATE                           = "twi_rate"
	FN_UPDATED_ADDRESS                    = "updated_address"
	FN_USER_BUNDLE_ALLOCATED_AMOUNT       = "user_bundle_allocated_amount"
	FN_USER_BUNDLE_ALLOCATION_DATE        = "user_bundle_allocation_date"
	FN_USER_BUNDLE_LOCKED                 = "user_bundle_locked"
	FN_USER_CONFIRMED_EMAIL               = "confirmed_email"
	FN_USER_CONFIRMED_PHONEL              = "confirmed_phone"
	FN_USERNAME                           = "username"
	FN_URL                                = "url"
	FN_UUID                               = "uuid"
	FN_SERVER_VERSION                     = "server_version"
	FN_WEB_ASSETS_URL                     = "Web_Assets_URL"
	FN_ZELLE_REQUEST_METHOD               = "zelle_request_method"
	FN_ZIPCODE                            = "zip_code"
)
