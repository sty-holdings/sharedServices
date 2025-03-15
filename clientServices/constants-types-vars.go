package sharedServices

import (
	"time"
)

type STYHClient struct {
	CompanyName     string `firebase:"company_name" json:"company_name"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp"`
	Email         string          `firebase:"email" json:"email"`
	FirstName     string          `firebase:"first_name" json:"first_name"`
	LastName      string          `firebase:"last_name" json:"last_name"`
	MyGoogleAds   GoogleAps       `firebase:"my_google_ads" json:"my_google_ads"`
	OnBoarded     bool            `firebase:"on_boarded" json:"on_boarded"`
	SaasProfile   UserSaaSProfile `firebase:"saas_profile" json:"saas_profile"`
	SaasProviders []string        `firestore:"saas_providers,array"json:"saas_providers"`
	StyhClientId  string          `firebase:"styh_client_id" json:"styh_client_id"`
	Timezone      string          `firebase:"timezone" json:"timezone"`
	Uid           string          `firebase:"uid" json:"uid"`
}
type GoogleAps struct {
	AccessToken string `firebase:"access_token" json:"access_token"`
	CustomerId  string `firebase:"customer_id" json:"customer_id"`
}

type UserSaaSProfile struct {
	UserSaaSProviders map[string]string
}
