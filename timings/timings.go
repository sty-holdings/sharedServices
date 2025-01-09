package sharedServices

import (
	// Add imports here

	"time"

	"cloud.google.com/go/firestore"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
)

// RecordSubjectTimings - stores a timing record for a subject
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RecordSubjectTimings(dkElapsedTime float64, environment string, extensionName string, subject string, firestoreClientPtr *firestore.Client, testMode bool) {

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
