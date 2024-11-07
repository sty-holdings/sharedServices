package sharedServices

import (
	"fmt"
	"strings"
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
