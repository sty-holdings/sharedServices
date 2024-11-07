package sharedServices

import (
	"runtime"
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2024/constsTypesVars"
)

func TestPrintDebugFunctionInfo(tPtr *testing.T) {

	type arguments struct {
		debugModeOn bool
		outputMode  string
	}

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Debug False - No output.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug true - No output.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Display.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  ctv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Display.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  ctv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Log.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  ctv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Log.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  ctv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				PrintDebugFunctionInfo(ts.arguments.debugModeOn, ts.arguments.outputMode)
			},
		)
	}
}

func TestPrintDebugLine(tPtr *testing.T) {

	type arguments struct {
		message     string
		debugModeOn bool
		outputMode  string
	}

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Debug False - No output.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug true - No output.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Display.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  ctv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Display.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  ctv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Log.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  ctv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Log.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  ctv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				PrintDebugLine(ts.arguments.message, ts.arguments.debugModeOn, ts.arguments.outputMode)
			},
		)
	}
}
