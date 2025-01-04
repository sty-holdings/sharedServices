package sharedServices

import (
	"fmt"
	"runtime"
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

func TestGetFunctionInfo(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tFunctionInfo      FunctionInfo
	)

	type arguments struct {
		level         int
		filenameValue string
	}

	var (
		gotError bool
	)

	tests := []struct {
		name               string
		arguments          arguments
		wantError          bool
		errorMessageFormat string
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Level 0 Filename Empty",
			arguments: arguments{
				level:         0,
				filenameValue: ctv.VAL_EMPTY,
			},
			errorMessageFormat: errs.FORMAT_EXPECTING_NO_ERROR,
			wantError:          false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Level 1 Filename Empty",
			arguments: arguments{
				level:         1,
				filenameValue: ctv.VAL_EMPTY,
			},
			errorMessageFormat: errs.FORMAT_EXPECTING_NO_ERROR,
			wantError:          false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Level 2 Filename Empty",
			arguments: arguments{
				level:         2,
				filenameValue: ctv.VAL_EMPTY,
			},
			errorMessageFormat: errs.FORMAT_EXPECTING_NO_ERROR,
			wantError:          false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Level 3 Filename Empty",
			arguments: arguments{
				level:         3,
				filenameValue: ctv.VAL_EMPTY,
			},
			errorMessageFormat: errs.FORMAT_EXPECTING_NO_ERROR,
			wantError:          false,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Level 4 Filename Empty",
			arguments: arguments{
				level:         4,
				filenameValue: ctv.VAL_EMPTY,
			},
			errorMessageFormat: errs.FORMAT_EXPECTING_NO_ERROR,
			wantError:          true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				tFunctionInfo = GetFunctionInfo(ts.arguments.level)
				if tFunctionInfo.FileName == ts.arguments.filenameValue ||
					tFunctionInfo.Name == ctv.VAL_EMPTY {
					gotError = true
				} else {
					fmt.Println("FileName: ", tFunctionInfo.FileName)
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, tFunctionName, ctv.STATUS_UNKNOWN)
				}
			},
		)
	}
}

func TestGetProgramInfo(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tProgramInfo       ProgramInfo
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			tProgramInfo = GetProgramInfo()
			if tProgramInfo.FileName == ctv.VAL_EMPTY ||
				tProgramInfo.NumberCPUs == ctv.VAL_ZERO ||
				tProgramInfo.GoVersion == ctv.VAL_EMPTY {
				tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, GetFunctionInfo(1).Name, ctv.STATUS_UNKNOWN)
			}
		},
	)
}

func TestGetWorkingDirectory(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tProgramInfo       ProgramInfo
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			tProgramInfo.WorkingDirectory = GetWorkingDirectory()
			if tProgramInfo.WorkingDirectory == ctv.VAL_EMPTY {
				tPtr.Errorf(errs.FORMAT_EXPECTING_NO_ERROR, GetFunctionInfo(1).Name, ctv.STATUS_UNKNOWN)
			}
		},
	)
}
