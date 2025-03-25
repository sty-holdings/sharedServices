package sharedServices

import (
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

func TestNewGRPCServer(tPtr *testing.T) {

	type arguments struct {
		config GRPCConfig
	}

	var (
		errorInfo errs.ErrorInfo
		gotError  bool
		// tFunction, _, _, _ = runtime.Caller(0)
		// tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Insecure connection.",
			arguments: arguments{
				config: GRPCConfig{
					GRPCHost: "localhost",
					GRPCPort: 50051,
					GRPCSecure: SecureSettings{
						ServerSide: true,
						Mutual:     false,
					},
					GRPCTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
				},
			},
			wantError: false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				config: GRPCConfig{
					GRPCHost: "localhost",
					GRPCPort: 50051,
					GRPCSecure: SecureSettings{
						ServerSide: true,
						Mutual:     false,
					},
					GRPCTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
				},
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				// if _, errorInfo = NewGRPCServer(ts.arguments.config); errorInfo.Error != nil {
				// 	gotError = true
				// 	errorInfo = errs.ErrorInfo{
				// 		Error: errors.New(fmt.Sprintf("Failed - gRPC connection was not created for Test: %v", tFunctionName)),
				// 	}
				// } else {
				// 	gotError = false
				// }
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}

func TestNewGRPCClient(tPtr *testing.T) {

	type arguments struct {
		filename string
	}

	var (
		errorInfo errs.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Insecure connection.",
			arguments: arguments{
				filename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/connect-server/config/hal/run-on-mac/test-local-grpc-client-config.json",
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = NewGRPCClient(ts.arguments.filename); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}
