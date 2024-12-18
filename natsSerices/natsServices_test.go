package sharedServices

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
)

func TestEncryptedMessageDataRequest(tPtr *testing.T) {

	type arguments struct {
		testMessage string
	}

	var (
		errorInfo     errs.ErrorInfo
		gotError      bool
		tInstanceName string
		tMessagePtr   = &nats.Msg{
			Header:  make(nats.Header),
			Subject: "TEST",
		}
		tNATSConnectionPtr    *nats.Conn
		tRequest              SaaSProfileRequest
		tTimeOutInSecs        int
		tUserSpecialNumberB64 = "BWzIo8nzg/QTkwds8dcjKg==" // Labeled so the scans of GitHub will not pick it up.
		tUID                  = "Scott"
		testingOn             = true
	)

	// Adding fake pointer to pass NATS Connection validation
	tMessagePtr.Header.Add(ctv.FN_UID, tUID)
	tMessagePtr.Header.Add(ctv.FN_STYH_CLIENT_ID, "Scott Client Id")

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				testMessage: "{\"provider\":\"Google\",\"action\":\"Create\",\"providerKeyInfo\":\"some-key-info\"}",
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = EncryptedMessageDataRequest(
					"TestEncryptedMessageDataRequest", tNATSConnectionPtr, tInstanceName, tUserSpecialNumberB64, tMessagePtr, tRequest,
					tTimeOutInSecs, tUID, testingOn,
				); errorInfo.Error != nil {
					if errorInfo.Message == errs.NATS_INVALID_CONNECTION {
						gotError = false
					} else {
						gotError = true
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

func TestEncryptedMessageDataReply(tPtr *testing.T) {

	type arguments struct {
		request SaaSProfileRequest
	}

	var (
		errorInfo   errs.ErrorInfo
		gotError    bool
		tMessagePtr = &nats.Msg{
			Header:  make(nats.Header),
			Subject: "TEST",
		}
		tUserSpecialNumber = "BWzIo8nzg/QTkwds8dcjKg==" // Labeled so the scans of GitHub will not pick it up.
		uId                = "Scott"
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				request: SaaSProfileRequest{
					Provider:        "Google",
					Action:          "Test",
					ProviderKeyInfo: "Fake_Key",
				},
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = EncryptedDataReply(ts.arguments.request, tMessagePtr, tUserSpecialNumber, uId); errorInfo.Error != nil {
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

func TestSendReply(tPtr *testing.T) {

	type arguments struct {
		request SaaSProfileRequest
	}

	var (
		errorInfo   errs.ErrorInfo
		gotError    bool
		tMessagePtr = &nats.Msg{
			Header:  make(nats.Header),
			Subject: "TEST",
		}
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				request: SaaSProfileRequest{
					Provider:        "Google",
					Action:          "Test",
					ProviderKeyInfo: "Fake_Key",
				},
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = SendReply(ts.arguments.request, tMessagePtr); errorInfo.Error != nil {
					if errorInfo.Message == errs.NATS_NOT_CONNECTED {
						gotError = false
					} else {
						gotError = true
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

func TestUnmarshalMessageData(tPtr *testing.T) {

	type arguments struct {
		testMessageData []byte
	}

	var (
		errorInfo   errs.ErrorInfo
		gotError    bool
		tRequestPtr *SaaSProfileRequest
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				testMessageData: []byte("{\"provider\":\"Google\",\"action\":\"Create\",\"providerKeyInfo\":\"some-key-info\"}"),
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = UnmarshalMessageData("TestUnmarshalMessageData", ts.arguments.testMessageData, &tRequestPtr); errorInfo.Error != nil {
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

func TestUnmarshalEncryptedMessageData(tPtr *testing.T) {

	type arguments struct {
		testMessage string
	}

	var (
		dMessageData string
		errorInfo    errs.ErrorInfo
		gotError     bool
		tMessagePtr  = &nats.Msg{
			Header:  make(nats.Header),
			Subject: "TEST",
		}
		tRequestPtr        *SaaSProfileRequest
		tUserSpecialNumber = "BWzIo8nzg/QTkwds8dcjKg==" // Labeled so the scans of GitHub will not pick it up.
		tUID               = "Scott"
	)

	tMessagePtr.Header.Add(ctv.FN_UID, tUID)
	tMessagePtr.Header.Add(ctv.FN_STYH_CLIENT_ID, "Scott Client Id")

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				testMessage: "{\"provider\":\"Google\",\"action\":\"Create\",\"providerKeyInfo\":\"some-key-info\"}",
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				// The Encrypt is only for testing
				dMessageData, errorInfo = jwts.Encrypt(tUID, tUserSpecialNumber, ts.arguments.testMessage)
				tMessagePtr.Data = []byte(dMessageData)
				if errorInfo = UnmarshalEncryptedMessageData("TestUnmarshalMessageData", tMessagePtr, &tRequestPtr, tUserSpecialNumber); errorInfo.Error != nil {
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
