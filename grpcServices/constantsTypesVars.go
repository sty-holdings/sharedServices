package sharedServices

import (
	"net"

	"google.golang.org/grpc"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

type GRPCConfiguration struct {
	GRPCHost    string       `json:"grpc_host"`
	GRPCPort    int          `json:"grpc_port"`
	GRPCSecure  bool         `json:"grpc_secure"`
	GRPCTLSInfo jwts.TLSInfo `json:"grpc_tls_info"`
}

type GRPCService struct {
	gRPCListenerPtr *net.Listener
	gRPCServerPtr   *grpc.Server
	secure          bool
	host            string
	userInfo        ctv.UserInfo
}
