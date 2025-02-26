package sharedServices

//goland:noinspection ALL
const ()

type GCPService struct {
	GCPConfig GCPConfig
}

type GCPConfig struct {
	GCPCredentialFilename string `json:"gcp_credential_filename"`
	GCPLocation           string `json:"gcp_location"`
	GCPProjectID          string `json:"gcp_project_id"`
}
