package sharedServices

import (
	"testing"
	// ctv "github.com/sty-holdings/GriesPikeThomp/shared-services"
)

func TestNewHTTP(tPtr *testing.T) {

	// type arguments struct {
	// 	hostname       string
	// 	configFilename string
	// }
	//
	// var (
	// 	errorInfo errs.ErrorInfo
	// 	gotError  bool
	// )
	//
	// tests := []struct {
	// 	name      string
	// 	arguments arguments
	// 	wantError bool
	// }{
	// 	{
	// 		name: ctv.TEST_POSITIVE_SUCCESS + "Secure connection.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: false,
	// 	},
	// 	{
	// 		name: ctv.TEST_POSITIVE_SUCCESS + "Bad URL.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: true,
	// 	},
	// 	{
	// 		name: ctv.TEST_NEGATIVE_SUCCESS + "Missing credentials location.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: true,
	// 	},
	// 	{
	// 		name: ctv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: true,
	// 	},
	// 	{
	// 		name: ctv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: true,
	// 	},
	// 	{
	// 		name: ctv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
	// 		arguments: arguments{
	// 			hostname:       "localhost",
	// 			configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/natsServices-connect/config/local/httpServices-inbound-config.json",
	// 		},
	// 		wantError: true,
	// 	},
	// }
	//
	// for _, ts := range tests {
	// 	tPtr.Run(
	// 		ts.name, func(t *testing.T) {
	// 			if _, errorInfo = NewHTTP(ts.arguments.configFilename); errorInfo.Error != nil {
	// 				gotError = true
	// 			} else {
	// 				gotError = false
	// 			}
	// 			if gotError != ts.wantError {
	// 				tPtr.Error(ts.name)
	// 				tPtr.Error(errorInfo)
	// 			}
	// 		},
	// 	)
	// }
}
