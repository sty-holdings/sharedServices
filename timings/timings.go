package sharedServices

import (
	// Add imports here

	"time"

	"cloud.google.com/go/firestore"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
	pi "github.com/sty-holdings/sharedServices/v2025/programInfo"
)

// RecordFunctionTimings - stores a timing record
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RecordFunctionTimings(dkElapsedTime time.Duration, firestoreClientPtr *firestore.Client, testMode bool) {

	var (
		tFields       = make(map[any]interface{})
		tFunctionInfo pi.FunctionInfo
	)

	tFunctionInfo = pi.GetFunctionInfo(2)

	tFields[ctv.FN_ELASPE_TIME_SECONDS] = dkElapsedTime
	tFields[ctv.FN_FUNCTION_NAME] = tFunctionInfo.Name
	tFields[ctv.FN_FILENAME] = tFunctionInfo.FileName
	tFields[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	if testMode == false {
		fbs.SetDocument(firestoreClientPtr, ctv.DATASTORE_TIMINGS, hlp.GenerateUUIDType1(true), tFields)
	}

}
