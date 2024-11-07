// Package constant_type_vars
/*
This file contains USA states and postal codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package constant_type_vars

import (
	"fmt"
	"strings"

	"github.com/sty-holdings/sharedServices/v2024/jwtServices"
)

//goland:noinspection ALL
const (
	PARAMETER_NATS_TOKEN          = "nats-token"
	PARAMETER_NATS_WEBSOCKET_HOST = "nats-websocket-host"
	PARAMETER_NATS_PORT           = "nats-port"
	PARAMETER_NATS_URL            = "nats-url"
)

var (
	ParameterNameFormat = "%v-%v-%v"
)

// GetParameterName returns the formatted parameter name based on the parameter type, program name, and environment.
// If the parameter type matches one of the predefined constants, it will return the formatted parameter name using the ParameterNameFormat.
// Otherwise, it will return a string indicating that the parameter type is missing.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetParameterName(programName, environment, parameterType string) string {

	switch strings.ToLower(strings.Trim(parameterType, SPACES_ONE)) {
	case PARAMETER_NATS_TOKEN:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_NATS_TOKEN)
	case PARAMETER_NATS_WEBSOCKET_HOST:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_NATS_WEBSOCKET_HOST)
	case PARAMETER_NATS_PORT:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_NATS_PORT)
	case PARAMETER_NATS_URL:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, PARAMETER_NATS_URL)
	case sty_shared.PARAMETER_TLS_CERT:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_CERT)
	case sty_shared.PARAMETER_TLS_CERT_FQN:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_CERT_FQN)
	case sty_shared.PARAMETER_TLS_PRIVATE_KEY:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_PRIVATE_KEY)
	case sty_shared.PARAMETER_TLS_PRIVATE_KEY_FQM:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_PRIVATE_KEY_FQM)
	case sty_shared.PARAMETER_TLS_CA_BUNDLE:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_CA_BUNDLE)
	case sty_shared.PARAMETER_TLS_CA_BUNDLE_FQN:
		return fmt.Sprintf(ParameterNameFormat, programName, environment, sty_shared.PARAMETER_TLS_CA_BUNDLE_FQN)
	default:
		return fmt.Sprintf("%v%v", LBL_MISSING_PARAMETER, FN_PARAMETER_TYPE)
	}
}
