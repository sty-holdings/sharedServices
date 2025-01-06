package sharedServices

import (
	"fmt"
	"runtime"
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
)

func TestNewErrorInfo(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name      string
		arguments arguments
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				fmt.Println(NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo))
			},
		)
	}
}

func TestPrintError(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied - No Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: "",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied - No Output Mode",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied - Display Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: ctv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied - Display Output Mode",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied - Log Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: ctv.MODE_OUTPUT_LOG,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied - Log Output Mode",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_LOG,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				PrintError(ts.arguments.myError, ts.arguments.additionalInfo)
			},
		)
	}
}

func TestPrintErrorInfo(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: "",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: ctv.MODE_OUTPUT_LOG,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_LOG,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: ctv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_DISPLAY,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				PrintErrorInfo(NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo))
			},
		)
	}
}

func TestOutputError(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_LOG,
		},
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: ctv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: ctv.MODE_OUTPUT_DISPLAY,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				outputError(NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo))
			},
		)
	}
}

func TestNewError(tPtr *testing.T) {

	var (
		buf                = make([]byte, 1024)
		errorInfo          ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if errorInfo = newError(buf, ErrErrorMissing); errorInfo.Error == nil {
				tPtr.Errorf(FORMAT_EXPECTED_ERROR, tFunctionName, ctv.VAL_EMPTY)
			}
			if errorInfo = newError(buf, nil); errorInfo.Error != nil {
				tPtr.Errorf(FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo.Error)
			}
		},
	)
}
