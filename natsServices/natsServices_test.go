package sharedServices

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

func TestGetNATSConnection(tPtr *testing.T) {

	type arguments struct {
		instanceName string
		config       NATSConfiguration
	}

	var (
		errorInfo          errs.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: false,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing Instance Name.",
			arguments: arguments{
				instanceName: ctv.VAL_EMPTY,
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing Credential filename.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: ctv.VAL_EMPTY,
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Port is zero.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                strconv.Itoa(ctv.VAL_ZERO),
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       ctv.VAL_EMPTY,
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: ctv.VAL_EMPTY,
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   ctv.VAL_EMPTY,
					},
					NATSURL: "local.sty-holdings.net",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing URL.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/user-creds/connect-server-us-west",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCertFQN:       "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net.crt",
						TLSPrivateKeyFQN: "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local-sty-holdings-net-private.key",
						TLSCABundleFQN:   "/Volumes/development-share/.keys/sty-holdings.net/local/local_sty-holdings_net/local_sty-holdings_net_CAbundle.crt",
					},
					NATSURL: ctv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = getConnection(ts.arguments.instanceName, ts.arguments.config); errorInfo.Error != nil {
					gotError = true
					errorInfo = errs.ErrorInfo{
						Error: errors.New(fmt.Sprintf("Failed - NATS connection was not created for Test: %v", tFunctionName)),
					}
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

func TestMakeRequestReplyNoHeaderInsecure(tPtr *testing.T) {

	type arguments struct {
		instanceName string
		dkRequest    DKRequest
	}

	var (
		errorInfo       errs.ErrorInfo
		gotError        bool
		tConfig         NATSConfiguration
		tNatsServicePtr *NATSService
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				dkRequest:    []byte(ctv.SYSTEM_ACTION_PING),
			},
			wantError: false,
		},
	}

	tConfig = NATSConfiguration{
		NATSCredentialsFilename: "/Volumes/development-share/.keys/ai-daveknows-dev/nats/creds/connect-server-US-WEST",
		NATSPort:                "4222",
		NATSTLSInfo: jwts.TLSInfo{
			TLSCertFQN:       "/Volumes/development-share/.keys/ai-daveknows-dev/tls/certs/dev_daveknows_net.crt",
			TLSPrivateKeyFQN: "/Volumes/development-share/.keys/ai-daveknows-dev/tls/dk-dev-private.key",
			TLSCABundleFQN:   "/Volumes/development-share/.keys/ai-daveknows-dev/tls/certs/ca-bundle.crt",
		},
		NATSURL: "dev.daveknows.net",
	}

	if tNatsServicePtr, errorInfo = NewNATSService("test", tConfig); errorInfo.Error != nil {
		tPtr.Errorf("Error creating NATS service: %v", errorInfo.Error)
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = makeRequestReplyNoHeaderInsecure(ts.arguments.dkRequest, tNatsServicePtr, ctv.VAL_EMPTY, 5); errorInfo.Error != nil {
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
