package sharedServices

import (
	"time"
)

type STYHClient struct {
	CompanyName     string `firebase:"company_name" json:"company_name"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp"`
	Email         string `firebase:"email" json:"email"`
	FirstName     string `firebase:"first_name" json:"first_name"`
	LastName      string `firebase:"last_name" json:"last_name"`
	LocationPtr   *time.Location
	OnBoarded     bool     `firebase:"on_boarded" json:"on_boarded"`
	SaasProviders []string `firestore:"saas_providers,array"json:"saas_providers"`
	StyhClientId  string   `firebase:"styh_client_id" json:"styh_client_id"`
	Timezone      string   `firebase:"timezone" json:"timezone"`
	UID           string   `firebase:"UID" json:"UID"`
}

type NewUser struct {
	Email       string `json:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	UID         string `json:"uId,omitempty"` // The json uId is outside the STYH control, so don't change it.
}
