package sharedServices

import (
	"github.com/nats-io/nats.go"

	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	METHOD_DASHES      = "dashes"
	METHOD_UNDERSCORES = "underscores"
	METHOD_BLANK       = ""

	CREDENTIAL_FILENAME = "nats-credentials-filename"

	// Test constants
	TEST_MESSAGE_ENVIRONMENT = "local"
	TEST_MESSAGE_NAMESPACE   = "nci"
	TEST_PORT                = 4222
	TEST_PORT_EMPTY          = ""
	//
	TEST_INVALID_URL = "invalid URL"
)

type MessageHandler struct {
	Handler nats.MsgHandler
}

type NATSConfiguration struct {
	NATSCredentialsFilename string `json:"nats_credentials_filename"`
	NATSToken               string
	NATSPort                string       `json:"nats_port"`
	NATSTLSInfo             jwts.TLSInfo `json:"nats_tls_info"`
	NATSURL                 string       `json:"nats_url"`
}

type NATSService struct {
	ConnPtr      *nats.Conn
	InstanceName string
	Secure       bool
	URL          string
}
