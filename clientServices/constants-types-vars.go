package sharedServices

import (
	"time"
)

type InternalClientUser struct {
	MyInternalClient InternalClient
	MyInternalUser   InternalUser
}

type InternalClient struct {
	CompanyName     string `firebase:"company_name" json:"company_name" yaml:"company_name"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp" yaml:"create_timestamp"`
	Domain                       string   `firebase:"domain" json:"domain" yaml:"domain"`
	DemoAccount                  bool     `firebase:"demo_account" json:"demo_account" yaml:"demo_account"`
	FormationType                string   `firestore:"formation_type" json:"formation_type" yaml:"formation_type"`
	GoogleAdsAccounts            []string `firebase:"google_ads_accounts" json:"google_ads_accounts" yaml:"google_ads_accounts"`
	InternalClientID             string   `firebase:"internal_client_id" json:"internal_client_id" yaml:"internal_client_id"`
	LinkedinPageIds              []int64  `firebase:"linkedin_page_ids" json:"linkedin_page_ids" yaml:"linkedin_page_ids"`
	OnBoarded                    bool     `firebase:"onboarded" json:"onboarded" yaml:"onboarded"`
	Owners                       []string `firebase:"owners" json:"owners" yaml:"owners"`
	PayPalClientID               string   `firebase:"paypal_client_id" json:"paypal_client_id" yaml:"paypal_client_id"`
	PayPalClientSecret           string   `firebase:"paypal_client_secret" json:"paypal_client_secret" yaml:"paypal_client_secret"`
	PayPalClientRefreshToken     string   `firebase:"paypal_client_refresh_token" json:"paypal_client_refresh_token" yaml:"paypal_client_refresh_token"`
	PhoneCountryCode             string   `firebase:"phone_country_code" json:"phone_country_code" yaml:"phone_country_code"`
	PhoneAreaCode                string   `firebase:"phone_area_code" json:"phone_area_code" yaml:"phone_area_code"`
	PhoneNumber                  string   `firebase:"phone_number" json:"phone_number" yaml:"phone_number"`
	SaaSClientProviders          []string `firestore:"saas_client_providers,array" json:"saas_client_providers" yaml:"saas_client_providers"`
	StripeClientConnectAccountId string   `firebase:"stripe_client_connect_account_id" json:"stripe_client_connect_account_id" yaml:"stripe_client_connect_account_id"`
	StripeClientRefreshToken     string   `firebase:"stripe_client_refresh_token" json:"stripe_client_refresh_token" yaml:"stripe_client_refresh_token"`
	StripeInitialPullDataStatus  string   `firebase:"stripe_pull_data_status" json:"stripe_pull_data_status" yaml:"stripe_pull_data_status"`
	StripePullFrequency          string   `firebase:"stripe_pull_frequency" json:"stripe_pull_frequency" yaml:"stripe_pull_frequency"`
	StripeStartDate              string   `firebase:"stripe_start_date" json:"stripe_start_date" yaml:"stripe_start_date"`
	TimezoneHQ                   string   `firebase:"timezone_hq" json:"timezone_hq" yaml:"timezone_hq"`
	TimezoneHQLocationPtr        *time.Location
	WebsiteURL                   string `firebase:"website_url" json:"website_url" yaml:"website_url"`
}

type InternalUser struct {
	ApprovedBy      string `firebase:"approved_by" json:"approved_by" yaml:"approved_by"`
	ApprovedByDate  string `firebase:"approved_by_date" json:"approved_by_date" yaml:"approved_by_date"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp" yaml:"create_timestamp"`
	Email                   string   `firebase:"email" json:"email" yaml:"email"`
	FirstName               string   `firebase:"first_name" json:"first_name" yaml:"first_name"`
	LastName                string   `firebase:"last_name" json:"last_name" yaml:"last_name"`
	Permissions             []string `firebase:"permissions" json:"permissions" yaml:"permissions"`
	InternalClientID        string   `firebase:"internal_client_id" json:"internal_client_id" yaml:"internal_client_id"`
	InternalUserID          string   `firebase:"internal_user_id" json:"styh_user_id"yaml:"internal_user_id"`
	TimezoneUser            string   `firebase:"timezone_user" json:"user_timezone_user" yaml:"user_timezone_user"`
	TimezoneUserLocationPtr *time.Location
}

type NewClient struct {
	CompanyName           string `firebase:"company_name" json:"company_name,omitempty" yaml:"company_name,omitempty"`
	Domain                string `firebase:"domain" json:"domain" yaml:"domain"`
	FormationType         string `firebase:"formation_type" json:"formation_type,omitempty" yaml:"formation_type,omitempty"`
	PhoneCountryCode      string `firebase:"phone_country_code" json:"phone_country_code,omitempty" yaml:"phone_country_code,omitempty"`
	PhoneAreaCode         string `firebase:"phone_area_code" json:"phone_area_code,omitempty" yaml:"phone_area_code,omitempty"`
	PhoneNumber           string `firebase:"phone_number" json:"phone_number,omitempty" yaml:"phone_number,omitempty"`
	TimezoneHQ            string `firebase:"timezone_hq" json:"timezone_hq,omitempty" yaml:"timezone_hq,omitempty"`
	TimezoneHQLocationPtr *time.Location
	WebSiteURL            string `firebase:"web_site_url" json:"web_site_url,omitempty" yaml:"web_site_url,omitempty"`
}

type NewUser struct {
	Email                   string `firebase:"email" json:"email,omitempty" yaml:"email,omitempty"`
	FirstName               string `firebase:"firstName" json:"firstName,omitempty" yaml:"firstName,omitempty"`
	LastName                string `firebase:"lastName" json:"lastName,omitempty" yaml:"lastName,omitempty"`
	TimezoneUser            string `firebase:"timezone_user" json:"timezone_user,omitempty" yaml:"timezone_user,omitempty"`
	TimezoneUserLocationPtr *time.Location
	InternalClientID        string `firebase:"internal_client_id" json:"internal_client_id,omitempty" yaml:"internal_client_id,omitempty"`
	InternalUserID          string `firebase:"internal_user_id" json:"internal_user_id,omitempty" yaml:"internal_user_id,omitempty"`
}
