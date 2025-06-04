package sharedServices

import (
	"net"
	"time"

	"google.golang.org/grpc"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

type GRPCConfig struct {
	DebugModeOn bool           `json:"debug_mode_on" yaml:"debug_mode_on"`
	Host        string         `json:"host,omitempty" yaml:"host,omitempty"`             // This is only used on the client side. The server side is set to 0.0.0.0.
	KeepAlive   KeepAlive      `json:"keep_alive,omitempty" yaml:"keep_alive,omitempty"` // This is only used on the client side.
	Port        int            `json:"port" yaml:"port"`
	Secure      SecureSettings `json:"secure" yaml:"secure"` // These settings must match for client and server.
	TLSInfo     jwts.TLSInfo   `json:"tls_info" yaml:"tls_info"`
	Timeout     int            `json:"timeout" yaml:"timeout"`
}

type KeepAlive struct {
	PingInternalSec     int  `json:"ping_internal_sec" yaml:"ping_internal_sec"`
	PingTimeoutSec      int  `json:"ping_timeout_sec" yaml:"ping_timeout_sec"`
	PermitWithoutStream bool `json:"permit_without_stream" yaml:"permit_without_stream"`
	ServerMinPingTime   int  `json:"server_min_ping_time,omitempty" yaml:"server_min_ping_time,omitempty"`
}

type GRPCService struct {
	debugModeOn       bool
	GRPCListenerPtr   *net.Listener
	GRPCServerPtr     *grpc.Server
	GRPCClientPtr     *grpc.ClientConn
	Host              string
	Port              uint
	Secure            SecureSettings
	ServerMinPingTime uint
	Timeout           time.Duration
}

// If both ServerSide and Mutual are false, then it is the default NoClient.
type SecureSettings struct {
	ServerSide bool `json:"server_side" yaml:"server_side"`
	Mutual     bool `json:"mutual" yaml:"mutual"`
}
