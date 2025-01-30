package sharedServices

import (
	"net"

	"google.golang.org/grpc"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

type GRPCConfiguration struct {
	GRPCHost    string         `json:"grpc_host" yaml:"grpc_host"` // This is only used on the client side. Server side is set to localhost.
	GRPCPort    int            `json:"grpc_port" yaml:"grpc_port"`
	GRPCSecure  SecureSettings `json:"grpc_secure" yaml:"grpc_secure"`
	GRPCTLSInfo jwts.TLSInfo   `json:"grpc_tls_info" yaml:"grpc_tls_info"`
}
type GRPCService struct {
	GRPCListenerPtr *net.Listener
	GRPCServerPtr   *grpc.Server
	secure          SecureSettings
	host            string
}

// If both ServerSide and Mutual are false, then it is the default NoClient.
type SecureSettings struct {
	ServerSide bool `json:"server_side" yaml:"server_side"`
	Mutual     bool `json:"mutual" yaml:"mutual"`
}
