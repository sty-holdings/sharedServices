package sharedServices

import (
	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

type MessageHandler struct {
	Handler nats.MsgHandler
}

type NATSConfiguration struct {
	NATSCredentialsFilename string `json:"nats_credentials_filename" yaml:"nats_credentials_filename"`
	NATSToken               string
	NATSPort                string       `json:"nats_port" yaml:"nats_port"`
	NATSTLSInfo             jwts.TLSInfo `json:"nats_tls_info" yaml:"nats_tls_info"`
	NATSURL                 string       `json:"nats_url" yaml:"nats_url"`
}

type NATSService struct {
	connPtr      *nats.Conn
	instanceName string
	secure       bool
	userInfo     ctv.UserInfo
	url          string
}
