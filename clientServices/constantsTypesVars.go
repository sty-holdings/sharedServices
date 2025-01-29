package sharedServices

import (
	"time"
)

type STYHClient struct {
	CompanyName     string `firebase:"company_name" json:"company_name"`
	CreateTimestamp struct {
		Time time.Time `json:"__time__"`
	} `firebase:"create_timestamp" json:"create_timestamp"`
	Email        string          `firebase:"email" json:"email"`
	FirstName    string          `firebase:"first_name" json:"first_name"`
	LastName     string          `firebase:"last_name" json:"last_name"`
	SaasProfile  UserSaaSProfile `firebase:"saas_profile" json:"saas_profile"`
	SecretKey    string          `firebase:"secret_key" json:"secret_key"`
	StyhClientId string          `firebase:"styh_client_id" json:"styh_client_id"`
	Timezone     string          `firebase:"timezone" json:"timezone"`
	Uid          string          `firebase:"uid" json:"uid"`
	OnBoarded    bool            `firebase:"on_boarded" json:"on_boarded"`
}

type UserSaaSProfile struct {
	UserSaaSProviders map[string]string
}
