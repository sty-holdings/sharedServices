package sharedServices

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
)

// Base64Decode - will decode a base64 string to a string. If there is an error,
// the first 20 characters of the base64 string are logged.
// REMINDER: If the base64 string has sensitivity information, empty out the
// ErrorInfo.AdditionalInfo field before logging or outputting the error.
//
//	Customer Messages: None
//	Errors: error returned by StdEncoding.DecodeString
//	Verifications: None
func Base64Decode(base64Value string) (
	value []byte,
	errorInfo errs.ErrorInfo,
) {

	if value, errorInfo.Error = b64.StdEncoding.DecodeString(base64Value); errorInfo.Error != nil {
		errorInfo.AdditionalInfo = fmt.Sprintf("%v%v", ctv.LBL_BASE64, base64Value[:20])
	}

	return
}

// PrependWorkingDirectory - will add the working directory.
// if the filename first character is a /, the passed value will be returned
// unmodified.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrependWorkingDirectory(filename string) string {

	var (
		tWorkingDirectory, _ = os.Getwd()
	)

	if filepath.IsAbs(filename) {
		return filename
	}

	return fmt.Sprintf("%v/%v", tWorkingDirectory, filename)
}
