package sharedServices

import (
	// Add imports here

	"os"
	"runtime"
)

// GetFunctionInfo - returns information about the function based on the level provided.
//
// 0 will always return information about GetFunctionInfo
//
// 1 will return the caller of GetFunctionInfo
//
// 2+ will return the corresponding caller back up the chain.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFunctionInfo(level int) (functionInfo FunctionInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	functionInfo.Name = runtime.FuncForPC(tFunction).Name()
	_, functionInfo.FileName, functionInfo.LineNumber, _ = runtime.Caller(level)

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
