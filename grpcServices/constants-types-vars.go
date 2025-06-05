package sharedServices

import (
	"net"
	"time"

	"google.golang.org/grpc"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

// For Reference: https://pkg.go.dev/google.golang.org/grpc/keepalive@v1.71.0#pkg-types
type GRPCConfig struct {
	ClientKeepAlive ClientKeepAlive `json:"client_keep_alive,omitempty" yaml:"client_keep_alive,omitempty"`
	DebugModeOn     bool            `json:"debug_mode_on" yaml:"debug_mode_on"`
	Host            string          `json:"host,omitempty" yaml:"host,omitempty"` // This is only used on the client side. The server side is set to 0.0.0.0.
	Port            int             `json:"port" yaml:"port"`
	Secure          SecureSettings  `json:"secure" yaml:"secure"`                                           // If you are using Keep Alive, both the client and server must have keep alive enabled.
	ServerKeepAlive ServerKeepAlive `json:"server_keep_alive,omitempty" yaml:"server_keep_alive,omitempty"` // If you are using Keep Alive, both the client and server must have keep alive enabled.
	TLSInfo         jwts.TLSInfo    `json:"tls_info" yaml:"tls_info"`
	Timeout         int             `json:"timeout" yaml:"timeout"`
}

// If you are using Keep Alive, both the client and server must have keep alive enabled.
type ClientKeepAlive struct {
	PingIntervalSec     int  `json:"ping_interval_sec" yaml:"ping_interval_sec"`         // The Go default property name is Time. Default: none Recommended: 120 secs
	PingTimeoutSec      int  `json:"ping_timeout_sec" yaml:"ping_timeout_sec"`           // The Go default property name is Timeout.
	PermitWithoutStream bool `json:"permit_without_stream" yaml:"permit_without_stream"` // The Go default property name is PermitWithoutStream. Default: false Recommended: true
}

// If you are using Keep Alive, both the client and server must have keep alive enabled.
type ServerKeepAlive struct {
	// If the client does not receive an acknowledgment within this time, it will close the connection. The recommended value is 60 secs. Default: 20 secs. The Go default property name is Timeout.
	ServerEnforcementPolicy ServerEnforcementPolicy `json:"server_enforcement_policy" yaml:"server_enforcement_policy"`
	ServerParameters        ServerParameters        `json:"server_parameters" yaml:"server_parameters"`
}

type ServerEnforcementPolicy struct {
	MinTimeClientPingsSec int `json:"min_time_client_pings_sec" yaml:"min_time_client_pings_sec"` // The minimum amount of time a client should wait before sending a keepalive ping.
	// Default: 300 secs Recommended: 120 secs (MinTime)
	PermitWithoutStream bool `json:"permit_without_stream" yaml:"permit_without_stream"` // If true, the server allows keepalive pings even when there are no active streams(RPCs). If false, and client sends ping when there are no active
	// streams, the server will send GOAWAY and close the connection. Default: false Recommended: true (PermitWithoutStream)
}

type ServerParameters struct {
	MaxConnectionIdleSec int `json:"max_connection_idle_sec,omitempty" yaml:"max_connection_idle_sec,omitempty"` // Maximum time that a channel may have no outstanding RPC calls
	// after which the server will close the connection (GOAWAY).  Default: Infinite Recommended: Infinite (MaxConnectionIdle)
	MaxConnectionAgeSec      int `json:"max_connection_age_sec,omitempty" yaml:"max_connection_age_sec,omitempty"`             // Maximum time that a channel may exist. Default: Infinite Recommended: Infinite (MaxConnectionAge)
	MaxConnectionAgeGraceSec int `json:"max_connection_age_grace_sec,omitempty" yaml:"max_connection_age_grace_sec,omitempty"` // Grace period after the channel reaches its max age.
	// Default: Infinite Recommended: Infinite (MaxConnectionAgeGrace)
	PingIntervalSec int `json:"ping_interval_sec" yaml:"ping_interval_sec"` // The interval in seconds between PING frames. The recommended value is 120 secs. The Go default property name is Time.
	PingTimeoutSec  int `json:"ping_timeout_sec" yaml:"ping_timeout_sec"`   // The timeout in seconds for a PING frame to be acknowledged. The Go default property name is Timeout.
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
