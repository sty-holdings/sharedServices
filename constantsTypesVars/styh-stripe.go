//package sharedServices
/*
NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package sharedServices

import (
	"time"
)

//goland:noinspection ALL
const (
	// 	Payment Method Types Supported
	// Source of values: https://docs.stripe.com/api/payment_methods/object#payment_method_object-type
	PAYMENT_METHOD_TYPE_ACSSDEBIT        = "acss_debit"
	PAYMENT_METHOD_TYPE_AFFIRM           = "affirm"
	PAYMENT_METHOD_TYPE_AFTERPAYCLEARPAY = "afterpay_clearpay"
	PAYMENT_METHOD_TYPE_ALIPAY           = "alipay"
	PAYMENT_METHOD_TYPE_AUBECSDEBIT      = "au_becs_debit"
	PAYMENT_METHOD_TYPE_BACSDEBIT        = "bacs_debit"
	PAYMENT_METHOD_TYPE_BANCONTACT       = "bancontact"
	PAYMENT_METHOD_TYPE_BLIK             = "blik"
	PAYMENT_METHOD_TYPE_BOLETO           = "boleto"
	PAYMENT_METHOD_TYPE_CARD             = "card"
	PAYMENT_METHOD_TYPE_CARD_PRESENT     = "card_present"
	PAYMENT_METHOD_TYPE_CASHAPP          = "cashapp"
	PAYMENT_METHOD_TYPE_CUSTOMER_BALANCE = "customer_balance"
	PAYMENT_METHOD_TYPE_EPS              = "eps"
	PAYMENT_METHOD_TYPE_FPX              = "fpx"
	PAYMENT_METHOD_TYPE_GIROPAY          = "giropay"
	PAYMENT_METHOD_TYPE_GRABPAY          = "grabpay"
	PAYMENT_METHOD_TYPE_IDEAL            = "ideal"
	PAYMENT_METHOD_TYPE_INTERAC_PRESENT  = "interac_present"
	PAYMENT_METHOD_TYPE_KLARNA           = "klarna"
	PAYMENT_METHOD_TYPE_KONBINI          = "konbini"
	PAYMENT_METHOD_TYPE_LINK             = "link"
	PAYMENT_METHOD_TYPE_OXXO             = "oxxo"
	PAYMENT_METHOD_TYPE_P24              = "p24"
	PAYMENT_METHOD_TYPE_PAYNOW           = "paynow"
	PAYMENT_METHOD_TYPE_PAYPAL           = "paypal"
	PAYMENT_METHOD_TYPE_PIX              = "pix"
	PAYMENT_METHOD_TYPE_PROMPTPAY        = "promptpay"
	PAYMENT_METHOD_TYPE_REVOLUTPAY       = "revolut_pay"
	PAYMENT_METHOD_TYPE_SEPADEBIT        = "sepa_debit"
	PAYMENT_METHOD_TYPE_SOFORT           = "sofort"
	PAYMENT_METHOD_TYPE_SWISH            = "swish"
	PAYMENT_METHOD_TYPE_USBANKACCOUNT    = "us_bank_account"
	PAYMENT_METHOD_TYPE_WECHATPAY        = "wechat_pay"
	PAYMENT_METHOD_TYPE_ZIP              = "zip"
	//

	// Cards by brand
	// Source: "https://docs.stripe.com/testing?testing-method=payment-methods#cards"
	CARD_BRAND_VISA                        = "pm_card_visa"
	CARD_BRAND_VISA_DEBIT                  = "pm_card_visa_debit"
	CARD_BRAND_MASTERCARD                  = "pm_card_mastercard"
	CARD_BRAND_MASTERCARD_DEBIT            = "pm_card_mastercard_debit"
	CARD_BRAND_MASTERCARD_PREPAID          = "pm_card_mastercard_prepaid"
	CARD_BRAND_AMEX                        = "pm_card_amex"
	CARD_BRAND_DISCOVER                    = "pm_card_discover"
	CARD_BRAND_DINERS_CLUB                 = "pm_card_diners"
	CARD_BRAND_JCB                         = "pm_card_jcb"
	CARD_BRAND_UNION_PAY                   = "pm_card_unionpay"
	CARD_BRAND_CARTES_BANCAIRES_VISA       = "pm_card_visa_cartesBancaires"
	CARD_BRAND_CARTES_BANCAIRES_MASTERCARD = "pm_card_mastercard_cartesBancaires"
	CARD_BRAND_EFTPOS_AUSTRALIA_VISA       = "pm_card_visa_debit_eftposAuCoBranded"
	CARD_BRAND_EFTPOS_AUSTRALIA_MASTERCARD = "pm_card_mastercard_debit_eftposAuCoBranded"

	// Cards by Country
	// Source: "https://docs.stripe.com/testing?testing-method=payment-methods#international-cards"
	CARD_USA_VISA                        = "pm_card_us"
	CARD_ARGENTINA_VISA                  = "pm_card_ar"
	CARD_BRAZIL_VISA                     = "pm_card_br"
	CARD_CANADA_VISA                     = "pm_card_ca"
	CARD_MEXICO_VISA                     = "pm_card_mx"
	CARD_UNITED_ARAB_EMIRATES_VISA       = "pm_card_ae"
	CARD_UNITED_ARAB_EMIRATES_MASTERCARD = "pm_card_ae_mastercard"
	CARD_AUSTRIA_VISA                    = "pm_card_at"
	CARD_BELGIUM_VISA                    = "pm_card_be"
	CARD_BULGARIA_VISA                   = "pm_card_bg"
	CARD_BELARUS_VISA                    = "pm_card_by"
	CARD_CROATIA_VISA                    = "pm_card_hr"
	CARD_CYPRUS_VISA                     = "pm_card_cy"
	CARD_CZECH_REPUBLIC_VISA             = "pm_card_cz"
	CARD_DENMARK_VISA                    = "pm_card_dk"
	CARD_ESTONIA_VISA                    = "pm_card_ee"
	CARD_FINLAND_VISA                    = "pm_card_fi"
	CARD_FRANCE_VISA                     = "pm_card_fr"
	CARD_GERMANY_VISA                    = "pm_card_de"
	CARD_GIBRALTAR_VISA                  = "pm_card_gi"
	CARD_GREECE_VISA                     = "pm_card_gr"
	CARD_HUNGARY_VISA                    = "pm_card_hu"
	CARD_IRELAND_VISA                    = "pm_card_ie"
	CARD_ITALY_VISA                      = "pm_card_it"
	CARD_LATVIA_VISA                     = "pm_card_lv"
	CARD_LIECHTENSTEIN_VISA              = "pm_card_li"
	CARD_LITHUANIA_VISA                  = "pm_card_lt"
	CARD_LUXEMBOURG_VISA                 = "pm_card_lu"
	CARD_MALTA_VISA                      = "pm_card_mt"
	CARD_NETHERLANDS_VISA                = "pm_card_nl"
	CARD_NORWAY_VISA                     = "pm_card_no"
	CARD_POLAND_VISA                     = "pm_card_pl"
	CARD_PORTUGAL_VISA                   = "pm_card_pt"
	CARD_ROMANIA_VISA                    = "pm_card_ro"
	CARD_SLOVENIA_VISA                   = "pm_card_si"
	CARD_SLOVAKIA_VISA                   = "pm_card_sk"
	CARD_SPAIN_VISA                      = "pm_card_es"
	CARD_SWEDEN_VISA                     = "pm_card_se"
	CARD_SWITZERLAND_VISA                = "pm_card_ch"
	CARD_UNITED_KINGDOM_VISA             = "pm_card_gb"
	CARD_UNITED_KINGDOM_DEBIT            = "pm_card_gb_debit"
	CARD_UNITED_KINGDOM_MASTERCARD       = "pm_card_gb_mastercard"
	CARD_AUSTRALIA_VISA                  = "pm_card_au"
	CARD_CHINA_VISA                      = "pm_card_cn"
	CARD_HONG_KONG_VISA                  = "pm_card_hk"
	CARD_INDIA_VISA                      = "pm_card_in"
	CARD_JAPAN_VISA                      = "pm_card_jp"
	CARD_JAPAN_JCB                       = "pm_card_jcb"
	CARD_MALAYSIA_VISA                   = "pm_card_my"
	CARD_NEW_ZEALAND_VISA                = "pm_card_nz"
	CARD_SINGAPORE_VISA                  = "pm_card_sg"
	CARD_THAILAND_VISA_CREDIT            = "pm_card_th_credit"
	CARD_THAILAND_DEBIT_VISA_DEBIT       = "pm_card_th_debit"

	// Testing Declined Payment
	// Source: "https://docs.stripe.com/testing?testing-method=payment-methods#declined-payments"
	DECLINE_GENERIC                 = "pm_card_visa_chargeDeclined"
	DECLINE_INSUFFICIENT_FUNDS      = "pm_card_visa_chargeDeclinedInsufficientFunds"
	DECLINE_LOST_CARD               = "pm_card_visa_chargeDeclinedLostCard"
	DECLINE_STOLEN_CARD             = "pm_card_visa_chargeDeclinedStolenCard"
	DECLINE_EXPIRED_CARD            = "pm_card_chargeDeclinedExpiredCard"
	DECLINE_INCORRECT_CVC           = "pm_card_chargeDeclinedIncorrectCvc"
	DECLINE_PROCESSING_ERROR        = "pm_card_chargeDeclinedProcessingError"
	DECLINE_VELOCITY_LIMIT_EXCEEDED = "pm_card_visa_chargeDeclinedVelocityLimitExceeded"
	DECLINE_CUSTOMER_FAIL           = "pm_card_chargeCustomerFail"

	// Testing Fraud Provention
	// Source: "https://docs.stripe.com/testing?testing-method=payment-methods#fraud-prevention"
	FRAUD_RADAR_BLOCK         = "pm_card_radarBlock"        // Always blocked
	FRAUD_RISK_LEVEL_HIGHEST  = "pm_card_riskLevelHighest"  // Highest risk
	FRAUD_RISK_LEVEL_ELEVATED = "pm_card_riskLevelElevated" // Elevated risk
	FRAUD_CVC_CHECK_FAIL      = "pm_card_cvcCheckFail"      // CVC check fails
	FRAUD_AVS_ZIP_FAIL        = "pm_card_avsZipFail"        // Postal code check fails
	FRAUD_AVS_LINE1_FAIL      = "pm_card_avsLine1Fail"      // Line1 check fails
	FRAUD_AVS_FAIL            = "pm_card_avsFail"           // Address checks fail
	FRAUD_AVS_UNCHECKED       = "pm_card_avsUnchecked"      // Address unavailable

	// Testing Invalid Data
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#invalid-data
	CARD_INVALID_EXPIRY_MONTH = 13
	CARD_INVALID_CVC          = 99
	CARD_INVALID_NUMBER       = 4242424242424241

	// Testing Disputes
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#disputes
	DISPUTES_FRAUDULENT   = "pm_card_createDispute"
	DISPUTES_NOT_RECEIVED = "pm_card_createDisputeProductNotReceived"
	DISPUTES_INQUIRY      = "pm_card_createDisputeInquiry"
	DISPUTES_MULTIPLE     = "pm_card_createMultipleDisputes"

	// Testing Evidence
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#refunds
	EVIDENCE_WINNING = "winning_evidence"
	EVIDENCE_LOSING  = "losing_evidence"

	// Testing Refunds
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#evidence
	REFUND_REFUND_PENDING = "pm_card_pendingRefund"
	REFUND_PENDING_FAIL   = "pm_card_refundFail"

	// Testing Available Balance
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#available-balance
	AVAILABLE_BALANCE_BYPASS_PENDING               = "pm_card_bypassPending"
	AVAILABLE_BALANCE_BYPASS_PENDING_INTERNATIONAL = "pm_card_bypassPendingInternational"

	// Testing 3D Secure Authentication
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#regulatory-cards
	THREE_D_AUTHENTICATION_REQUIRED_ON_SETUP  = "pm_card_authenticationRequiredOnSetup"
	THREE_D_AUTHENTICATION_REQUIRED           = "pm_card_authenticationRequired"
	THREE_D_AUTHENTICATION_ALREADY_SETUP      = "pm_card_authenticationRequiredSetupForOffSession"
	THREE_D_AUTHENTICATION_INSUFFICIENT_FUNDS = "pm_card_authenticationRequiredChargeDeclinedInsufficientFunds"

	// 3D Support and Availability
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#three-ds-cards
	// ToDo Implement Stripe Support and Availability

	// 3D Secure mobile challenge flows
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#3d-secure-mobile-challenge-flows
	// ToDo Implement 3D Secure mobile challenge flows

	// Payments with PINs
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#terminal
	// ToDo Implement Payments with PINs

	// Test account numbers
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#test-account-numbers
	// ToDo Implement Test account numbers

	// Test microdeposit amounts and descriptor codes
	// Source https://docs.stripe.com/testing?testing-method=payment-methods#test-microdeposit-amounts-and-descriptor-codes
	// ToDo Implement Test microdeposit amounts and descriptor codes
)

var (
	CARD_INVALID_EXPIRY_YEAR = time.Now().AddDate(-50, 0, 0)
)
