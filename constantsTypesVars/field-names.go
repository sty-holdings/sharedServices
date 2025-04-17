package sharedServices

// Entries here are for JSON, 3rd parties, and internal used names of data. The entries maybe upper, mixed, or lower case as needed.

//goland:noinspection ALL
const (
	FN_ACCOUNTS                                 = "accounts"
	FN_ACCOUNT_ID                               = "account_id"
	FN_ACCOUNT_NAME                             = "account_name"
	FN_AI_MAX_OUTPUT_TOKENS                     = "ai_max_output_tokens"
	FN_AI_MODEL_NAME                            = "ai_model_name"
	FN_AI_OUTPUT_FORMAT                         = "ai_output_format"
	FN_AI_PROMPT                                = "ai_prompt"
	FN_AI_SET_TOP_K                             = "ai_set_top_k"
	FN_AI_SET_TOP_PROBABILITY                   = "ai_set_top_probability"
	FN_AI_SYSTEM_INSTRUCTION                    = "ai_system_instruction"
	FN_AI_TEMPERATURE                           = "ai_temperature"
	FN_AMOUNT_IN_PENNIES                        = "amount_in_pennies"
	FN_ANALYSIS                                 = "analysis"
	FN_ANALYSIS_ID                              = "analysis_id"
	FN_ANALYSIS_STATUS                          = "analysis_status"
	FN_ANALYZE_QUESTION                         = "analyze_question"
	FN_ANSWER                                   = "answer"
	FN_AREA_CODE                                = "area_code"
	FN_AUDIENCE_CAP                             = "AUDIENCE"
	FN_AUTHOR                                   = "author"
	FN_AVERAGE_FLAG                             = "average"
	FN_AWS_ACCESS_KEY_ID                        = "aws_access_key_id"
	FN_AWS_ACCESS_KEY_ID_CAP                    = "AWS_ACCESS_KEY_ID"
	FN_AWS_ACCOUNT_INFO_FILENAME                = "aws_account_info_filename"
	FN_AWS_CLIENT_ID                            = "client_id"
	FN_AWS_COGNITO_USERNAME                     = "cognito:username"
	FN_AWS_IDENTITY_ID                          = "identityId"
	FN_AWS_IDENTITY_POOL_ID                     = "identity_pool_id"
	FN_AWS_PHONE_NUMBER_CUSTOM                  = "custom:phoneNumber"
	FN_AWS_PROFILE                              = "profile"
	FN_AWS_PROFILE_CAP                          = "PROFILE"
	FN_AWS_REGION                               = "region"
	FN_AWS_REGION_CAP                           = "REGION"
	FN_AWS_SECRET_ACCESS_KEY                    = "aws_secret_access_key"
	FN_AWS_SECRET_ACCESS_KEY_CAP                = "AWS_SECRET_ACCESS_KEY"
	FN_AWS_USERPOOL_ID                          = "userpool_id"
	FN_AWS_USERPOOL_ID_CAP                      = "USERPOOL_ID"
	FN_AWS_USERPOOL_NAME                        = "userpool_name"
	FN_AWS_USERPOOL_NAME_CAP                    = "USERPOOL_NAME"
	FN_AWS_USER_CONFIRMED_EMAIL                 = "confirmed_email"
	FN_AWS_USER_CONFIRMED_PHONEL                = "confirmed_phone"
	FN_BALANCE                                  = "Balance"
	FN_BATCH_NAME                               = "batch_name"
	FN_BIRTHDATE                                = "birthdate"
	FN_CANDIDATE_TOKEN_COUNT                    = "candidate_token_count"
	FN_CATEGORY                                 = "category"
	FN_CATEGORY_SENTENCE                        = "category_sentence"
	FN_CATEGORY_SENTENCE_TOTAL_TOKEN_COUNT      = "category_sentence_total_token_count"
	FN_CATEGORY_SENTENCE_PROMPT_TOKEN_COUNT     = "category_sentence_prompt_token_count"
	FN_CATEGORY_SENTENCE_CANDIDATE_TOKEN_COUNT  = "category_sentence_candidate_token_count"
	FN_CERT_KID                                 = "kid"
	FN_CHARGE                                   = "charge"
	FN_CIPHER_MAC                               = "mac"
	FN_CIPHER_NONCE                             = "nonce"
	FN_CIPHER_TEXT                              = "ciphertext"
	FN_CIPHER_TEXT_B64                          = "ciphertextb64"
	FN_CITY                                     = "city"
	FN_CLIENT_ID                                = "client_id"
	FN_CLIENT_KEY                               = "client_key"
	FN_COMPANY_NAME                             = "company_name"
	FN_COMPARISON_FLAG                          = "comparison"
	FN_COMPARISON_QUESTION                      = "comparison_question"
	FN_COMPOUND_FLAG                            = "compound"
	FN_COMPOUND_QUESTION                        = "compound_question"
	FN_CONFIG_FILENAME                          = "config_filename"
	FN_COUNT_BY_SUBJECT                         = "count_by_subject"
	FN_COUNT_FLAG                               = "count"
	FN_CREATE_TIMESTAMP                         = "create_timestamp"
	FN_CREDENTIALS_FILENAME                     = "credentials_filename"
	FN_CURRENT_FLAG                             = "current"
	FN_CUSTOMER                                 = "customer"
	FN_DATASTORE                                = "datastore"
	FN_DATATYPE_BOOL                            = "bool"
	FN_DATATYPE_COMPLEX128                      = "complex128"
	FN_DATATYPE_COMPLEX64                       = "complex64"
	FN_DATATYPE_FLOAT32                         = "float32"
	FN_DATATYPE_FLOAT64                         = "float64"
	FN_DATATYPE_INT                             = "int"
	FN_DATATYPE_INT16                           = "int16"
	FN_DATATYPE_INT32                           = "int32"
	FN_DATATYPE_INT64                           = "int64"
	FN_DATATYPE_INT8                            = "int8"
	FN_DATATYPE_STRING                          = "string"
	FN_DATATYPE_UINT                            = "uint"
	FN_DATATYPE_UINT16                          = "uint16"
	FN_DATATYPE_UINT32                          = "uint32"
	FN_DATATYPE_UINT64                          = "uint64"
	FN_DATATYPE_UINT8                           = "uint8"
	FN_DATATYPE_UINTPTR                         = "uintptr"
	FN_DAY_VALUES                               = "day_values"
	FN_DEBUG_MODE_ON                            = "debug_mode_on"
	FN_DESCRIPTION                              = "description"
	FN_DETAIL_FLAG                              = "detail"
	FN_DISPUTE                                  = "dispute"
	FN_DOCUMENT_ID                              = "document_id"
	FN_DURATION                                 = "Duration"
	FN_ELASPE_TIME_SECONDS                      = "elaspe_time_seconds"
	FN_EMAIL                                    = "email"
	FN_EMAIL_VERIFIED                           = "email_verified"
	FN_END_DATE                                 = "endDate"
	FN_ENVIRONMENT                              = "environment"
	FN_EXPIRY_TIMESTAMP                         = "expiry_timestamp"
	FN_EXTENSION_DECLARATIONS                   = "extension_declarations"
	FN_EXTENSION_DECLARATION_NAME               = "extension_declaration_name"
	FN_EXTENSION_FILENAME                       = "extension_name"
	FN_EXTENSION_NAME                           = "extension_name"
	FN_FAMILY_NAME                              = "family_name"
	FN_FEDERAL_TAX_ID                           = "federal_tax_id"
	FN_FEDERAL_TAX_ID_IS_SET                    = "federal_tax_id_is_set"
	FN_FILENAME                                 = "filename"
	FN_FIRST_NAME                               = "first_name"
	FN_FORECAST_FLAG                            = "forecast"
	FN_FUNCTION_INFO                            = "function_info"
	FN_FUNCTION_NAME                            = "function_name"
	FN_FUTURE_FLAG                              = "coming"
	FN_GCP_CREDENTIAL_FILENAME                  = "gcp_credential_filename"
	FN_GCP_LOCATION                             = "gcp_location"
	FN_GCP_PROJECT_ID                           = "gcp_project_id"
	FN_GENERATE_ANSWER                          = "generate_answer"
	FN_GIN_MODE                                 = "gin_mode"
	FN_GIN_URL_PORT                             = "gin_url_port"
	FN_GIVEN_NAME                               = "given_name"
	FN_GOOGLE_ADS_ACCOUNTS                      = "google_ads_accounts"
	FN_GRPC_SERVER_POINTER                      = "grpc_server_pointer"
	FN_GRPC_TIMEOUT                             = "grpc_timeout"
	FN_HTTP_TLS_INFO                            = "http_tls_info"
	FN_ITERATOR_POINTER                         = "tIterPtr"
	FN_JSON_STRING                              = "json_string"
	FN_JWT                                      = "JWT"
	FN_JWT_ISSUER                               = "issuer"
	FN_KEYSET_URL                               = "keySetURL"
	FN_KEY_B64                                  = "keyB64"
	FN_KEY_DATA                                 = "keyData"
	FN_KEY_PRIVATE                              = "key"
	FN_KEY_PUBLIC                               = "key"
	FN_LAST_FLAG                                = "last"
	FN_LAST_NAME                                = "last_name"
	FN_LAST_REFRESHED                           = "last_refreshed"
	FN_LAST_UPDATE_TIMESTAMP                    = "last_update_timestamp"
	FN_LOAD_EXTENSIONS                          = "load_extensions"
	FN_LOG_DIRECTORY                            = "log_directory"
	FN_MAXIMUM_FLAG                             = "maximum"
	FN_MAX_ALLOCATION                           = "max_allocation"
	FN_MAX_THREADS                              = "max_threads"
	FN_MESSAGE_ENVIRONMENT                      = "message_environment"
	FN_MESSAGE_NAMESPACE                        = "message_namespace"
	FN_MESSAGE_REGISTRY                         = "message_registry"
	FN_MINIMUM_FLAG                             = "minimum"
	FN_MIN_ALLOCATION                           = "min_allocation"
	FN_MONTH                                    = "month"
	FN_MONTH_OVER_MONTH                         = "month_over_month"
	FN_MONTH_TO_DATE                            = "month_to_date"
	FN_MONTH_VALUES                             = "month_values"
	FN_NAVIGATION                               = "navigation"
	FN_NEXT_FLAG                                = "next"
	FN_NICKNAME                                 = "nickname"
	FN_OFFICIAL_NAME                            = "official_name"
	FN_ON_BOARDED                               = "on_boarded"
	FN_PARAMETER_TYPE                           = "parameterType"
	FN_PASSWORD                                 = "password"
	FN_PERCENTAGE_FLAG                          = "percentage"
	FN_PERIOD                                   = "period"
	FN_PREDICTION                               = "predict"
	FN_PHONE_NUMBER                             = "phone_number"
	FN_PHONE_VERIFIED                           = "phone_verified"
	FN_PID_DIRECTORY                            = "pid_directory"
	FN_POINTER                                  = "any_pointer"
	FN_PORT                                     = "port"
	FN_PREVIOUS_FLAG                            = "previous"
	FN_PRIVATE_KEY                              = "PrivateKey"
	FN_PROCESS_DATE                             = "processDate"
	FN_PROGRAM_NAME                             = "program_name"
	FN_PROMPT                                   = "prompt"
	FN_PROMPT_TOKEN_COUNT                       = "prompt_token_count"
	FN_PROVIDED_AML_PHOTO                       = "provided_aml_photo"
	FN_PSQL_DB_NAME                             = "psql_db_name"
	FN_PSQL_DEBUG                               = "psql_debug"
	FN_PSQL_HOST                                = "psql_host"
	FN_PSQL_MAX_CONNECTIONS                     = "psql_max_connections"
	FN_PSQL_PASSWORD                            = "psql_password"
	FN_PSQL_PORT                                = "psql_port"
	FN_PSQL_SSL_MODE                            = "psql_ssl_mode"
	FN_PSQL_TIMESOUT                            = "psql_timeout"
	FN_PSQL_USER_NAME                           = "psql_user_name"
	FN_PURPOSE                                  = "purpose"
	FN_QUARTER                                  = "quarter"
	FN_QUARTER_VALUES                           = "quarter_values"
	FN_QUESTION                                 = "question"
	FN_RECOMMEND                                = "recommend"
	FN_RECORD_NUMBER                            = "record_number"
	FN_REFUND                                   = "refund"
	FN_RELATIVE_TIME                            = "relative_time"
	FN_RELEASE_STATUS                           = "release_status"
	FN_REPORT_FLAG                              = "report"
	FN_SAAS_PROVIDER                            = "saas_provider"
	FN_SAAS_PROVIDERS                           = "saas_providers"
	FN_SAAS_PROVIDER_SELECTED                   = "saas_provider_selected"
	FN_SENTENCE_SUBJECT                         = "sentence_subject"
	FN_SENTENCE_SUBJECT_ADVERB                  = "sentence_subject_adverb"
	FN_SERVER_DEBUG_MODE_ON                     = "server_debug_mode_on"
	FN_SERVER_INSTANCE_NUMBER                   = "server_instance_number"
	FN_SERVER_VERSION                           = "server_version"
	FN_SERVICE_CONFIG_FILENAME                  = "service_config_filename"
	FN_SERVICE_CONFIG_FILENAMES                 = "service_config_filenames"
	FN_SHORT_URL                                = "short_URL"
	FN_SIGNAL                                   = "signal"
	FN_SKELETON_DIRECTORY                       = "skeleton_config_directory"
	FN_SPECIAL_WORD_AVERAGE                     = "average"
	FN_SPECIAL_WORD_COMPARISON                  = "comparison"
	FN_SPECIAL_WORD_COMPOUND                    = "compound"
	FN_SPECIAL_WORD_COUNT                       = "count"
	FN_SPECIAL_WORD_DETAIL                      = "detail"
	FN_SPECIAL_WORD_FORECAST                    = "forecast"
	FN_SPECIAL_WORD_MAXIMUM                     = "maximum"
	FN_SPECIAL_WORD_MINIMUM                     = "minimum"
	FN_SPECIAL_WORD_PERCENTAGE                  = "percentage"
	FN_SPECIAL_WORD_PREDICTION                  = "predict"
	FN_SPECIAL_WORD_RECOMMEND                   = "recommend"
	FN_SPECIAL_WORD_REPORT                      = "report"
	FN_SPECIAL_WORD_SUBTOTAL                    = "subtotal"
	FN_SPECIAL_WORD_SUMMARY                     = "summary"
	FN_SPECIAL_WORD_TOTAL                       = "total"
	FN_SPECIAL_WORD_TRANSACTION                 = "transaction"
	FN_SPECIAL_WORD_TREND                       = "trend"
	FN_SPECIAL_WORDS                            = "special_words"
	FN_SPECIAL_WORDS_TOTAL_TOKEN_COUNT          = "special_words_total_token_count"
	FN_SPECIAL_WORDS_PROMPT_TOKEN_COUNT         = "special_words_prompt_token_count"
	FN_SPECIAL_WORDS_CANDIDATE_TOKEN_COUNT      = "special_words_candidate_token_count"
	FN_START_DATE                               = "startDate"
	FN_STATE                                    = "state"
	FN_STATUS                                   = "status"
	FN_STREET_ADDRESS                           = "street_address"
	FN_STRIPE_KEY                               = "stripe_key"
	FN_STYH_CLIENT_ID                           = "styh_client_id"
	FN_STYH_CUSTOM_SECRET_KEY                   = "custom:secret_key"
	FN_STYH_USER_ID                             = "styh_user_id"
	FN_SUB                                      = "sub"
	FN_SUBJECT                                  = "subject"
	FN_SUB_CATEGORY                             = "sub-category"
	FN_SUB_TOTAL_FLAG                           = "subtotal"
	FN_SUMMARY_FLAG                             = "summary"
	FN_SUNDAY_DATE                              = "sunday_date"
	FN_SYSTEM_ACTION                            = "system_action"
	FN_TEMP_DIRECTORY                           = "temporary_directory_fqd"
	FN_TEST                                     = "test"
	FN_TIME_PERIOD_VALUES                       = "time_period_values"
	FN_TIME_PERIOD_VALUES_TOTAL_TOKEN_COUNT     = "time_period_values_total_token_count"
	FN_TIME_PERIOD_VALUES_PROMPT_TOKEN_COUNT    = "time_period_values_prompt_token_count"
	FN_TIME_PERIOD_VALUES_CANDIDATE_TOKEN_COUNT = "time_period_values_candidate_token_count"
	FN_TIMEZONE                                 = "timezone"
	FN_TLS_CA_BUNDLE                            = "tls_ca_bundle"
	FN_TLS_CA_BUNDLE_FILENAME                   = "tls_ca_bundle_filename"
	FN_TLS_CERTIFICATE                          = "tls_certificate"
	FN_TLS_CERTIFICATE_FILENAME                 = "tls_certificate_filename"
	FN_TLS_INFO                                 = "tls_info"
	FN_TLS_PRIVATE_KEY                          = "tls_private_key"
	FN_TLS_PRIVATE_KEY_FILENAME                 = "tls_private_key_filename"
	FN_TO_DATE_FLAG                             = "to_date"
	FN_TODAY_FLAG                               = "today"
	FN_TOKEN                                    = "token"
	FN_TOKEN_ACCESS                             = "accessToken"
	FN_TOKEN_ID                                 = "idToken"
	FN_TOKEN_PAYLOAD                            = "tokenPayload"
	FN_TOKEN_REFRESH                            = "refreshToken"
	FN_TOKEN_TYPE                               = "tokenType"
	FN_TOTAL_FLAG                               = "total"
	FN_TOTAL_TOKEN_COUNT                        = "total_token_count"
	FN_TRANSACTION_FLAG                         = "transaction"
	FN_TREND_FLAG                               = "trend"
	FN_UID                                      = "UID"
	FN_UPDATED_ADDRESS                          = "updated_address"
	FN_URL                                      = "url"
	FN_USERNAME                                 = "username"
	FN_UUID                                     = "uuid"
	FN_VALUE                                    = "value"
	FN_VALUE_B64                                = "valueB64"
	FN_WEB_ASSETS_URL                           = "Web_Assets_URL"
	FN_WEEK                                     = "week"
	FN_WEEK_ENDING                              = "week_ending"
	FN_WEEK_FLAG                                = "week"
	FN_WEEK_OVER_WEEK                           = "week_over_week"
	FN_WEEK_STARTING                            = "week_starting"
	FN_WEEK_TO_DATE                             = "week_to_date"
	FN_WEEK_VALUES                              = "week_values"
	FN_YEAR                                     = "year"
	FN_YEAR_OVER_YEAR                           = "year_over_year"
	FN_YEAR_TO_DATE                             = "year_to_date"
	FN_YEAR_VALUES                              = "year_values"
	FN_ZIPCODE                                  = "zip_code"
)
