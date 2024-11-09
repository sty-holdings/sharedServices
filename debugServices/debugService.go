package sharedServices

import (
	// Add imports here

	"fmt"
	"log"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
)

//goland:noinspection GoSnakeCaseUsage
const (
	DEBUG_FUNCTION_FORMAT = "[DEBUG_FUNCTION] File: %v Function Name: '%v' Near Line Number: %v\n"
	DEBUG_MESSAGE_FORMAT  = "[DEBUG_MESSAGE] File: %v Function Name: '%v' Message: %v\n"
)

// PrintDebugFunctionInfo - if debugMode is true the function info of the caller will be output.
// The format of the messages is
// "[DEBUG] File: {Filename} Function: {Function Name} Near Line Number: {Line Number}\n"
// Set shared_services.ProgramInfo.DebugModeOn to true to turn on debug mode. The default is false.
// This function uses the shared_services.ProgramInfo.OutputModeOn. The default is log.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrintDebugFunctionInfo(
	debugModeOn bool,
	outputMode string,
) {

	var (
		tFunctionInfo = GetFunctionInfo(1)
	)

	if debugModeOn {
		if outputMode == ctv.MODE_OUTPUT_DISPLAY {
			fmt.Printf(DEBUG_FUNCTION_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, tFunctionInfo.LineNumber)
		} else {
			log.Printf(DEBUG_FUNCTION_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, tFunctionInfo.LineNumber)
		}
	}
}

// PrintDebugLine - if debugMode is true the function info of the caller will be output.
// The format of the messages is
// "[DEBUG] File: {Filename} Function: {Function Name} Near Line Number: {Line Number}\n"
// Set shared_services.ProgramInfo.DebugModeOn to true to turn on debug mode. The default is false.
// This function uses the shared_services.ProgramInfo.OutputModeOn. The default is log.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrintDebugLine(
	message string,
	debugModeOn bool,
	outputMode string,
) {

	var (
		tFunctionInfo = GetFunctionInfo(1)
	)

	if message == ctv.VAL_EMPTY {
		message = ctv.TXT_MISSING
	}

	if debugModeOn {
		if outputMode == ctv.MODE_OUTPUT_DISPLAY {
			fmt.Printf(DEBUG_MESSAGE_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, message)
		} else {
			log.Printf(DEBUG_MESSAGE_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, message)
		}
	}
}
