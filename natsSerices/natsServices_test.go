package sharedServices

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
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
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				instanceName: ctv.VAL_EMPTY,
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: false,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing Credential filename.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: ctv.VAL_EMPTY,
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Port is zero.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "0",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       ctv.VAL_EMPTY,
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: ctv.VAL_EMPTY,
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   ctv.VAL_EMPTY,
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Missing URL.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: NATSConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/natsServices-savup-backend.key",
					NATSPort:                "4222",
					NATSTLSInfo: jwts.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
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
				if _, errorInfo = GetConnection(ts.arguments.instanceName, ts.arguments.config); errorInfo.Error != nil {
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
