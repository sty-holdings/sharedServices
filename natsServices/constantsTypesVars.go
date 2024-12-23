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
	connPtr      *nats.Conn
	instanceName string
	secure       bool
	userInfo     natsUserInfo
	url          string
}

type natsUserInfo struct {
	keyB64       string
	styhClientId string
	uId          string
}
