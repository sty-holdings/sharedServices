package sharedServices

import (
	// Add imports here

	"log"
	"runtime"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
)

type ErrorInfo struct {
	AdditionalInfo string `json:"error_additional_info"`
	Error          error  `json:"-"`
	FileName       string `json:"error_filename"`
	FunctionName   string `json:"error_function_name"`
	LineNumber     int    `json:"error_line_number"`
	Message        string `json:"error_message"`
}

// NewErrorInfo - will return an ErrorInfo object.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func NewErrorInfo(
	myError error,
	additionalInfo string,
) (errorInfo ErrorInfo) {

	if myError == nil {
		return
	} else {
		errorInfo = newError(myError)
	}

	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}
	errorInfo.Message = myError.Error()

	return
}

// PrintError - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func PrintError(
	myError error,
	additionalInfo string,
) {

	var (
		errorInfo ErrorInfo
	)

	if myError == nil {
		errorInfo = newError(ErrErrorMissing)
	} else {
		errorInfo = newError(myError)
	}
	if additionalInfo == ctv.VAL_EMPTY {
		errorInfo.AdditionalInfo = ctv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}

	outputError(errorInfo)
}

// PrintErrorInfo - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: ErrErrorMissing
//	Verifications: None
func PrintErrorInfo(errorInfo ErrorInfo) {

	if errorInfo.Error == nil {
		errorInfo = newError(ErrErrorMissing)
	}

	outputError(errorInfo)
}

// Private Functions
func outputError(errorInfo ErrorInfo) {

	log.Printf(
		"[ERROR] %v Additional Info: '%v' File: %v Near Line Number: %v\n",
		errorInfo.Error.Error(),
		errorInfo.AdditionalInfo,
		errorInfo.FileName,
		errorInfo.LineNumber,
	)
}

func newError(myError error) (errorInfo ErrorInfo) {

	errorInfo = getErrorFunctionFileNameLineNumber(3)
	errorInfo.Error = myError

	return
}

func getErrorFunctionFileNameLineNumber(level int) (errorInfo ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	errorInfo.FunctionName = runtime.FuncForPC(tFunction).Name()
	_, errorInfo.FileName, errorInfo.LineNumber, _ = runtime.Caller(level)

	return
}

// DumpErrorInfos - outputs multiple error messages
//
//	func DumpErrorInfos(ErrorInfos []ErrorInfo) {
//		for _, info := range ErrorInfos {
//			PrintError(info)
//		}
//	}

// ToDo move to validation which should return the errors
// func GetMapKeyPopulatedError(finding string) (errorInfo ErrorInfo) {
//
// 	GetFunctionInfo(1)
//
// 	switch strings.ToLower(finding) {
// 	case ctv.TXT_EMPTY:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsEmpty,
// 			Message: ErrMapIsEmpty.Error(),
// 		}
// 	case ctv.TXT_MISSING_KEY:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsMissingKey,
// 			Message: ErrMapIsMissingKey.Error(),
// 		}
// 	case ctv.TXT_MISSING_VALUE:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsMissingValue,
// 			Message: ErrMapIsMissingValue.Error(),
// 		}
// 	case ctv.VAL_EMPTY:
// 		fallthrough
// 	default:
// 		errorInfo.Error = ErrRequiredArgumentMissing
// 		errorInfo.Message = ErrRequiredArgumentMissing.Error()
// 		errorInfo.AdditionalInfo = "The 'finding' argument is empty."
// 	}
//
// 	return
// }
