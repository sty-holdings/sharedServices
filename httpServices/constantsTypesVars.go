package sharedServices

import (
	"net/http"
	"net/url"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

//goland:noinspection GoSnakeCaseUsage
const (
	TEST_CREDENTIALS_FILENAME = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key"
	TEST_MESSAGE_ENVIRONMENT  = "local"
	TEST_MESSAGE_NAMESPACE    = "nci"
	TEST_URL                  = "savup-local-0030.savup.com"
	TEST_PORT                 = 4222
	TEST_PORT_EMPTY           = ""
	TEST_TLS_CERT             = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt"
	TEST_TLS_PRIVATE_KEY      = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key"
	TEST_TLS_CA_BUNDLE        = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt"
	//
	TEST_INVALID_URL = "invalid URL"
)

//goland:noinspection ALL
const (
	//
	//  HTTP Ports
	HTTP_PORT_SECURE     = 8443
	HTTP_PORT_NON_SECURE = 8080
	//
	// HTTP Methods
	HTTP_POST = "POST"
	HTTP_GET  = "GET"
)

type HTTPConfiguration struct {
	CredentialsFilename string       `json:"credentials_filename"`
	GinMode             string       `json:"gin_mode"`
	HTTPDomain          string       `json:"http_domain"`
	MessageEnvironment  string       `json:"message_environment"`
	Port                int          `json:"port"`
	RequestedThreads    uint         `json:"requested_threads"`
	RouteRegistry       []RouteInfo  `json:"route_registry"`
	TLSInfo             jwts.TLSInfo `json:"tls_info"`
}

type RouteInfo struct {
	Namespace   string `json:"namespace"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type HTTPServerService struct {
	Config         HTTPConfiguration
	CredentialsFQN string
	HTTPServerPtr  *http.Server
	Secure         bool
}

type HTTPRequestService struct {
	clientPtr   *http.Client
	httpRequest *http.Request
	urlPtr      *url.URL
}
