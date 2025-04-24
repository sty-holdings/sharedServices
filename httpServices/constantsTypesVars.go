package sharedServices

import (
	"net/http"
	"net/url"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

//goland:noinspection ALL
const (
	PAYPAL_API_BASE_URL = "https://api-m.paypal.com"
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
