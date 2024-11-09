package sharedServices

import (
	"strings"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
)

// IsDirectoryFullyQualified - checks to see if the directory starts and ends with a slash.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsDirectoryFullyQualified(directory string) bool {

	if strings.HasPrefix(directory, ctv.FORWARD_SLASH) {
		if strings.HasSuffix(directory, ctv.FORWARD_SLASH) {
			return true
		}
	}

	return false

}
