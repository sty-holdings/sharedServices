// Package sty_shared
/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .awsServices/credentials file in the default location.
    * {Enter other restrictions here for AWS

    {Other categories of restrictions}
    * {List of restrictions for the categories

NOTES:
    {Enter any additional notes that you believe will help the next developer.}

COPYRIGHT & WARRANTY:

	Copyright (c) 5/6/24 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by Scott Yacko / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	https://sty-holdings.com

	5/6/24

	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package sty_shared

import (
	"github.com/nats-io/nats.go"

	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
	pi "github.com/sty-holdings/sharedServices/v2024/programInfo"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	METHOD_DASHES      = "dashes"
	METHOD_UNDERSCORES = "underscores"
	METHOD_BLANK       = ""

	CREDENTIAL_FILENAME = "nats-credentials-filename"

	// Test constants
	TEST_MESSAGE_ENVIRONMENT = "local"
	TEST_MESSAGE_NAMESPACE   = "nci"
	TEST_PORT                = 4222
	TEST_PORT_EMPTY          = ""
	//
	TEST_INVALID_URL = "invalid URL"
)

type MessageHandler struct {
	Handler nats.MsgHandler
}

type NATSConfiguration struct {
	NATSCredentialsFilename string `json:"nats_credentials_filename"`
	NATSToken               string
	NATSPort                string       `json:"nats_port"`
	NATSTLSInfo             jwts.TLSInfo `json:"nats_tls_info"`
	NATSURL                 string       `json:"nats_url"`
}

type NATSService struct {
	ConnPtr      *nats.Conn
	InstanceName string
	Secure       bool
	URL          string
}

type NATSReply struct {
	Response  interface{}  `json:"response,omitempty"`
	ErrorInfo pi.ErrorInfo `json:"error,omitempty"`
}
