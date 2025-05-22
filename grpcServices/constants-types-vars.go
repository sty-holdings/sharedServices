package sharedServices

import (
	"net"
	"time"

	"google.golang.org/grpc"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

type GRPCConfig struct {
	GRPCDebug   bool           `json:"grpc_debug" yaml:"grpc_debug"`
	GRPCHost    string         `json:"grpc_host" yaml:"grpc_host"` // This is only used on the client side. Server side is set to 0.0.0.0.
	GRPCPort    int            `json:"grpc_port" yaml:"grpc_port"`
	GRPCSecure  SecureSettings `json:"grpc_secure" yaml:"grpc_secure"`
	GRPCTLSInfo jwts.TLSInfo   `json:"grpc_tls_info" yaml:"grpc_tls_info"`
	GRPCTimeout int            `json:"grpc_timeout" yaml:"grpc_timeout"`
}

type GRPCService struct {
	debugModeOn     bool
	GRPCListenerPtr *net.Listener
	GRPCServerPtr   *grpc.Server
	GRPCClientPtr   *grpc.ClientConn
	Host            string
	Port            uint
	Secure          SecureSettings
	Timeout         time.Duration
}

// If both ServerSide and Mutual are false, then it is the default NoClient.
type SecureSettings struct {
	ServerSide bool `json:"server_side" yaml:"server_side"`
	Mutual     bool `json:"mutual" yaml:"mutual"`
}
