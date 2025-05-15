package sharedServices

import (
	"time"
)

type STYHClient struct {
	AccountType     string `json:"account_type"`
	CompanyName     string `firebase:"company_name" json:"company_name"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp"`
	Email                  string   `firebase:"email" json:"email"`
	FirstName              string   `firebase:"first_name" json:"first_name"`
	GoogleAdsAccounts      []string `firebase:"google_ads_accounts" json:"google_ads_accounts"`
	LastName               string   `firebase:"last_name" json:"last_name"`
	LinkedinPageIdList     []int64  `firebase:"linkedin_page_ids" json:"linkedin_page_ids"`
	LocationPtr            *time.Location
	OnBoarded              bool     `firebase:"onboarded" json:"onboarded"`
	PayPalClientId         string   `firebase:"paypal_client_id" json:"paypal_client_id"`
	PayPalClientSecret     string   `firebase:"paypal_client_secret" json:"paypal_client_secret"`
	SaasProviders          []string `firestore:"saas_providers,array"json:"saas_providers"`
	StripeAccessToken      string   `firebase:"stripe_access_token:" json:"stripe_access_token:" yaml:"stripe_access_token"`
	StripeConnectAccountId string   `firebase:"stripe_connect_account_id" json:"stripe_connect_account_id" yaml:"stripe_connect_account_id"`
	StripePullDataStatus   string   `firebase:"stripe_pull_data_status" json:"stripe_pull_data_status" yaml:"stripe_pull_data_status"`
	StripePullFrequency    string   `firebase:"stripe_pull_frequency" json:"stripe_pull_frequency" yaml:"stripe_pull_frequency"`
	StripeRefreshToken     string   `firebase:"stripe_refresh_token" json:"stripe_refresh_token" yaml:"stripe_refresh_token"`
	STYHClientId           string   `firebase:"styh_client_id" json:"styh_client_id"`
	STYHUserId             string   `firebase:"sty_user_id" json:"sty_user_id"`
	Timezone               string   `firebase:"timezone" json:"timezone"`
}

type NewUser struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	STYHUserId  string `json:"sty_user_id,omitempty"`
}
