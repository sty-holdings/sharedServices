package sharedServices

import (
	"runtime"
	"testing"

	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

//goland:noinspection ALL
const (
	AI_CONFIG_BAD_FILENAME                    = "/Users/syacko/workspace/sty-holdings/sharedServices/aiServices/test-ai-config-BAD.yaml"
	AI_CONFIG_FULL_FILENAME                   = "/Users/syacko/workspace/sty-holdings/sharedServices/aiServices/test-ai-config-full.yaml"
	AI_CONFIG_NO_SYSTEM_INSTRUCTIONS_FILENAME = "/Users/syacko/workspace/sty-holdings/sharedServices/aiServices/test-ai-config-no-system-instructions.yaml"
	GCP_CREDENTIALS_FILE                      = "/Volumes/development-share/.keys/dev-daveknows-net/google/service-accounts/dkanswers-development.json"
	GCP_LOCATION                              = "us-central1"
	GCP_PROJECT_ID                            = "dkanswers-development"
)

var (
//goland:noinspection ALL
)

func TestLoadAIConfig(tPtr *testing.T) {

	var (
		errorInfo          errs.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(t *testing.T) {
			if _, errorInfo = loadAIConfig(AI_CONFIG_BAD_FILENAME); errorInfo.Error == nil {
				tPtr.Errorf(errs.FORMAT_EXPECTED_ERROR, tFunctionName, errorInfo)
			}
		},
	)
	tPtr.Run(
		tFunctionName, func(t *testing.T) {
			if _, errorInfo = loadAIConfig(AI_CONFIG_NO_SYSTEM_INSTRUCTIONS_FILENAME); errorInfo.Error != nil {
				tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo)
			}
		},
	)
	tPtr.Run(
		tFunctionName, func(t *testing.T) {
			if _, errorInfo = loadAIConfig(AI_CONFIG_FULL_FILENAME); errorInfo.Error != nil {
				tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo)
			}
		},
	)
}

func TestNewAIService(tPtr *testing.T) {

	var (
		errorInfo          errs.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(t *testing.T) {
			if _, errorInfo = NewAIService(GCP_CREDENTIALS_FILE, GCP_PROJECT_ID, GCP_LOCATION, AI_CONFIG_FULL_FILENAME, false); errorInfo.Error != nil {
				tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo)
			}
		},
	)
}
