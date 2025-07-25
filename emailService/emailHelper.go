package sharedServices

import (
	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
)

func validateEmailConfig(config EmailConfig) (errorInfo errs.ErrorInfo) {

	// The config.DebugModeOn is either true or false. No need to check the value.
	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_EMAIL, config.DefaultSenderAddress, ctv.LBL_DEFAULT_SENDER_ADDRESS); errorInfo.Error != nil {
		return
	}
	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_EMAIL, config.DefaultSenderAddress, ctv.LBL_DEFAULT_SENDER_ADDRESS); errorInfo.Error != nil {
		return
	}

	return
}
