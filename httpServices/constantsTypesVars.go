package sharedServices

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
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
	DebugModeOn       bool         `json:"debug_mode_on" yaml:"debug_mode_on"`
	DeepLinks         []string     `json:"deep_links" yaml:"deep_links"`
	GinMode           string       `json:"gin_mode" yaml:"gin_mode"`
	HTTPDomain        string       `json:"http_domain" yaml:"http_domain"`
	Port              int          `json:"port" yaml:"port"`
	RouteRegistry     []RouteInfo  `json:"route_registry" yaml:"route_registry"`
	TemplateDirectory string       `json:"template_directory" yaml:"template_directory"`
	TLSInfo           jwts.TLSInfo `json:"tls_info" yaml:"tls_info"`
}

type RouteInfo struct {
	Namespace   string `json:"namespace"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type HTTPServerService struct {
	Config       HTTPConfiguration
	GinEnginePtr *gin.Engine
	Secure       bool
}

type HTTPRequestService struct {
	clientPtr   *http.Client
	httpRequest *http.Request
	urlPtr      *url.URL
}
