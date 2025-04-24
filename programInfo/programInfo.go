package sharedServices

import (
	// Add imports here

	"os"
	"runtime"
	"strings"
)

// GetMyFunctionInfo - retrieves information about the function that executed this function.
// The level is set to 1, so it will always return information about the caller.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetMyFunctionInfo(trimPath bool) (functionInfo FunctionInfo) {

	return GetFunctionInfo(1, trimPath)
}

// GetMyFunctionName - retrieves the name of the function that executed this function.
// The level is set to 1, so it will always return information about the caller.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetMyFunctionName(trimPath bool) (functionName string) {

	return GetFunctionName(2, trimPath)
}

// GetFunctionInfo - returns information about the function based on the level provided.
//
// 0 will always return information about GetFunctionInfo
//
// 1 will return the caller of GetFunctionInfo
//
// 2+ will return the corresponding caller moving up the caller chain.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFunctionInfo(level int, trimPath bool) (functionInfo FunctionInfo) {

	functionInfo.Name = GetFunctionName(level, trimPath)
	_, functionInfo.FileName, functionInfo.LineNumber, _ = runtime.Caller(level)

	if trimPath {
		functionInfo.FileName = functionInfo.FileName[strings.LastIndex(functionInfo.FileName, "/")+1:]
	}

	return
}

// GetFunctionName - retrieves the name of the function at the specified stack level.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFunctionName(level int, trimPath bool) (functionName string) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	functionName = runtime.FuncForPC(tFunction).Name()

	if trimPath {
		functionName = functionName[strings.LastIndex(functionName, "/")+1:]
	}

	return
}

// GetWorkingDirectory - return the working directory for the program
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetWorkingDirectory() (path string) {

	path, _ = os.Getwd()

	return
}

// GetProgramInfo - returns information about the program and the system where it is executing.
// The level is set to 1, so it will always return information about the caller. It is recommended that
// you use this when you initialize your program.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetProgramInfo() (programInfo ProgramInfo) {

	_, programInfo.FileName, _, _ = runtime.Caller(1)
	programInfo.GoVersion = runtime.Version()
	programInfo.NumberCPUs = runtime.NumCPU()

	return
}
