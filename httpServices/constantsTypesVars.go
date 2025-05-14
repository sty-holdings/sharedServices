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
	Host              string       `json:"host" yaml:"host"`
	Ports             []int        `json:"ports" yaml:"ports"`
	TemplateDirectory string       `json:"template_directory" yaml:"template_directory"`
	TLSInfo           jwts.TLSInfo `json:"tls_info" yaml:"tls_info"`
}

type HTTPServerService struct {
	Config       HTTPConfiguration
	GinEnginePtr map[uint]*gin.Engine
	Secure       bool
}

type HTTPRequestService struct {
	clientPtr   *http.Client
	httpRequest *http.Request
	urlPtr      *url.URL
}
