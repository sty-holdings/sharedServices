package sharedServices

import (
	// Add imports here

	"runtime"
	"time"

	"cloud.google.com/go/firestore"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
)

// RecordSubjectTiming - stores a timing record for a subject
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RecordSubjectTiming(dkElapsedTime float64, environment string, extensionName string, subject string, firestoreClientPtr *firestore.Client, testMode bool) {

	var (
		tFields = make(map[any]interface{})
	)

	tFields[ctv.FN_ELASPE_TIME_SECONDS] = dkElapsedTime
	tFields[ctv.FN_ENVIRONMENT] = environment
	tFields[ctv.FN_EXTENSION_NAME] = extensionName
	tFields[ctv.FN_SUBJECT] = subject
	tFields[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	if testMode == false {
		fbs.SetDocument(firestoreClientPtr, ctv.DATASTORE_TIMINGS, hlp.GenerateUUIDType1(true), tFields)
	}
}

// RecordSubjectTiming - stores a timing record for a function
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RecordFunctionTiming(dkElapsedTime float64, environment string, extensionName string, firestoreClientPtr *firestore.Client, testMode bool) {

	var (
		tFields            = make(map[any]interface{})
		tFunction, _, _, _ = runtime.Caller(1)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tFields[ctv.FN_ELASPE_TIME_SECONDS] = dkElapsedTime
	tFields[ctv.FN_ENVIRONMENT] = environment
	tFields[ctv.FN_EXTENSION_NAME] = extensionName
	tFields[ctv.FN_FUNCTION_NAME] = tFunctionName
	tFields[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	if testMode == false {
		fbs.SetDocument(firestoreClientPtr, ctv.DATASTORE_TIMINGS, hlp.GenerateUUIDType1(true), tFields)
	}
}
