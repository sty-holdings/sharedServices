package sharedServices

import (
	"time"
)

type STYHUser struct {
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp" yaml:"create_timestamp"`
	Email                string `firebase:"email" json:"email"`
	FirstName            string `firebase:"first_name" json:"first_name"`
	LastName             string `firebase:"last_name" json:"last_name"`
	LocationPtr          *time.Location
	OnBoarded            bool   `firebase:"onboarded" json:"onboarded"`
	STYHInternalClientID string `firebase:"styh_internal_client_id" json:"styh_internal_client_id" yaml:"styh_internal_client_id"`
	STYHInternalUserID   string `firebase:"styh_internal_user_id" json:"styh_user_id"yaml:"styh_internal_user_id"`
	UserTimezone         string `firebase:"user_timezone" json:"user_timezone" yaml:"user_timezone"`
}

type STYHClient struct {
	AccountType         string   `firestore:"account_type" json:"account_type" yaml:"account_type"`
	ClientSaasProviders []string `firestore:"saas_providers,array" json:"saas_providers" yaml:"saas_providers"`
	CompanyName         string   `firebase:"company_name" json:"company_name" yaml:"company_name"`
	CreateTimestamp     struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp" yaml:"create_timestamp"`
	GoogleAdsAccounts            []string `firebase:"google_ads_accounts" json:"google_ads_accounts" yaml:"google_ads_accounts"`
	HQLocationPtr                *time.Location
	HQTimezone                   string  `firebase:"hq_timezone" json:"hq_timezone" yaml:"hq_timezone"`
	LinkedinPageIdList           []int64 `firebase:"linkedin_page_ids" json:"linkedin_page_ids" yaml:"linkedin_page_ids"`
	OnBoarded                    bool    `firebase:"onboarded" json:"onboarded" yaml:"onboarded"`
	PayPalClientId               string  `firebase:"paypal_client_id" json:"paypal_client_id" yaml:"paypal_client_id"`
	PayPalClientSecret           string  `firebase:"paypal_client_secret" json:"paypal_client_secret" yaml:"paypal_client_secret"`
	StripeClientAccessToken      string  `firebase:"stripe_client_access_token:" json:"stripe_client_access_token:" yaml:"stripe_client_access_token"`
	StripeClientConnectAccountId string  `firebase:"stripe_client_connect_account_id" json:"stripe_client_connect_account_id" yaml:"stripe_client_connect_account_id"`
	StripeInitialPullDataStatus  string  `firebase:"stripe_pull_data_status" json:"stripe_pull_data_status" yaml:"stripe_pull_data_status"`
	StripePullFrequency          string  `firebase:"stripe_pull_frequency" json:"stripe_pull_frequency" yaml:"stripe_pull_frequency"`
	StripeClientRefreshToken     string  `firebase:"stripe_client_refresh_token" json:"stripe_client_refresh_token" yaml:"stripe_client_refresh_token"`
	StripeStartDate              string  `firebase:"stripe_start_date" json:"stripe_start_date" yaml:"stripe_start_date"`
	STYHInternalClientID         string  `firebase:"styh_internal_client_id" json:"styh_internal_client_id" yaml:"styh_internal_client_id"`
	STYHInternalUserID           string  `firebase:"styh_internal_user_id" json:"styh_user_id" yaml:"styh_internal_user_id"`
}

type NewUser struct {
	Email              string `firebase:"email" json:"email,omitempty" yaml:"email,omitempty"`
	FirstName          string `firebase:"firstName" json:"firstName,omitempty" yaml:"firstName,omitempty"`
	LastName           string `firebase:"lastName" json:"lastName,omitempty" yaml:"lastName,omitempty"`
	CompanyName        string `firebase:"companyName" json:"companyName,omitempty" yaml:"companyName,omitempty"`
	Timezone           string `firebase:"timezone" json:"timezone,omitempty" yaml:"timezone,omitempty"`
	STYHInternalUserID string `firebase:"styh_internal_user_id" json:"styh_internal_user_id,omitempty" yaml:"styh_internal_user_id,omitempty"`
}
