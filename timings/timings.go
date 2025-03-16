package sharedServices

import (
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
func RecordSystemActionTiming(dkElapsedTime float64, environment string, extensionName string, systemAction string, firestoreClientPtr *firestore.Client, testMode bool) {

	var (
		tFields = make(map[any]interface{})
	)

	tFields[ctv.FN_ELASPE_TIME_SECONDS] = dkElapsedTime
	tFields[ctv.FN_ENVIRONMENT] = environment
	tFields[ctv.FN_EXTENSION_NAME] = extensionName
	tFields[ctv.FN_SYSTEM_ACTION] = systemAction
	tFields[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	if testMode == false {
		fbs.SetDocument(firestoreClientPtr, ctv.DATASTORE_STATS_FUNCTION_TIMINGS, hlp.GenerateUUIDType1(true), tFields)
	}
}

// RecordFunctionTiming - stores a timing record for a function. This can not be used by Firebase Services.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RecordFunctionTiming(dkElapsedTime float64, functionName string, firestoreClientPtr *firestore.Client, testMode bool) {

	var (
		tFields = make(map[any]interface{})
	)

	tFields[ctv.FN_ELASPE_TIME_SECONDS] = dkElapsedTime
	tFields[ctv.FN_FUNCTION_NAME] = functionName
	tFields[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	if testMode == false {
		fbs.SetDocument(firestoreClientPtr, ctv.DATASTORE_STATS_FUNCTION_TIMINGS, hlp.GenerateUUIDType1(true), tFields)
	}
}
