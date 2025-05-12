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
	Email              string   `firebase:"email" json:"email"`
	FirstName          string   `firebase:"first_name" json:"first_name"`
	GoogleAdsAccounts  []string `firebase:"google_ads_accounts" json:"google_ads_accounts"`
	LastName           string   `firebase:"last_name" json:"last_name"`
	LinkedinPageIdList []int64  `firebase:"linkedin_page_ids" json:"linkedin_page_ids"`
	LocationPtr        *time.Location
	OnBoarded          bool     `firebase:"onboarded" json:"onboarded"`
	PayPalClientId     string   `firebase:"paypal_client_id" json:"paypal_client_id"`
	PayPalClientSecret string   `firebase:"paypal_client_secret" json:"paypal_client_secret"`
	SaasProviders      []string `firestore:"saas_providers,array"json:"saas_providers"`
	STYHClientId       string   `firebase:"styh_client_id" json:"styh_client_id"`
	Timezone           string   `firebase:"timezone" json:"timezone"`
	StripeKey          string   `firebase:"stripe_key" json:"stripe_key"`
	STYHUserId         string   `firebase:"sty_user_id" json:"sty_user_id"`
}

type NewUser struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	STYHUserId  string `json:"sty_user_id,omitempty"`
}
