package sharedServices

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
)

//goland:noinspection ALL

const (
	NOT_FOUND_MAYBE_CORRECT = "Getting the 'The document was found ' error maybe correct. Review code logic."
)

type NameValueQuery struct {
	FieldName  string
	FieldValue interface{}
}

var (
	CTXBackground = context.Background()
)

// BuildFirestoreUpdate - while the nameValues is a map[any], the function using a string assertion on the key.
// func BuildFirestoreUpdate(nameValues map[any]interface{}) (firestoreUpdateFields []firestoreServices.Update, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUpdate            firestoreServices.Update
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding == ctv.GOOD {
// 		for field, value := range nameValues {
// 			tUpdate.Path = field.(string)
// 			tUpdate.Value = value
// 			firestoreUpdateFields = append(firestoreUpdateFields, tUpdate)
// 		}
// 	} else {
// 		errorInfo.Error = errs.GetMapKeyPopulatedError(tFinding)
// 	}
//
// 	return
// }

// DoesDocumentExist - checks the document Reference pointer exists
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func doesDocumentExist(documentReferencePtr *firestore.DocumentRef) bool {

	if _, err := documentReferencePtr.Get(CTXBackground); err == nil {
		return true
	}

	return false
}

// FindDocument - Returns an error for documents not found, but it doesn't print the error to the log.
//
//	Customer Messages: None
//	Errors: errs.ErrRequiredArgumentMissing, errs.ErrDocumentNotFound, errs.ErrServiceFailedFIRESTORE
//	Verifications: None
func FindDocument(firestoreClientPtr *firestore.Client, datastore string, queryParameters ...NameValueQuery) (found bool, documentSnapshotPtr *firestore.DocumentSnapshot, errorInfo errs.ErrorInfo) {

	var (
		tQuery firestore.Query
	)

	if datastore == ctv.VAL_EMPTY || len(queryParameters) < 1 {
		errorInfo.Error = errs.ErrRequiredArgumentMissing
	} else {
		tQuery = firestoreClientPtr.Collection(datastore).Query
		for _, parameter := range queryParameters {
			if parameter.FieldName == ctv.VAL_EMPTY || parameter.FieldValue == ctv.VAL_EMPTY {
				errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, ctv.VAL_EMPTY)
				break
			} else {
				tQuery = tQuery.Where(parameter.FieldName, ctv.OPER_EQUAL_SIGN, parameter.FieldValue)
			}
		}
	}

	if errorInfo.Error == nil {
		tDocuments := tQuery.Documents(CTXBackground)
		for {
			documentSnapshotPtr, errorInfo.Error = tDocuments.Next()
			if errorInfo.Error != nil {
				if errors.Is(errorInfo.Error, iterator.Done) {
					errorInfo = errs.NewErrorInfo(errs.ErrDocumentNotFound, NOT_FOUND_MAYBE_CORRECT)
					break
				} else {
					errorInfo = errs.NewErrorInfo(errs.ErrServiceFailedFIRESTORE, ctv.VAL_EMPTY)
					break
				}
			}
			if len(documentSnapshotPtr.Data()) > 0 {
				found = true
				break
			}
		}
	}

	return
}

// GetAllDocuments will return snapshot pointers to each document in the datastore.
// If no documents are found, the documents will have a count of zero.
//
//	Customer Messages: None
//	Errors: errs.ErrRequiredArgumentMissing
//	Verifications: None
func GetAllDocuments(firestoreClientPtr *firestore.Client, datastore string) (documents []*firestore.DocumentSnapshot, errorInfo errs.ErrorInfo) {

	var (
		tCollectionReferencePtr *firestore.CollectionRef
	)

	if firestoreClientPtr == nil || datastore == ctv.VAL_EMPTY {
		errorInfo.Error = errs.ErrRequiredArgumentMissing
		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v", firestoreClientPtr, datastore)
		errs.PrintErrorInfo(errorInfo)
	} else {
		tCollectionReferencePtr = firestoreClientPtr.Collection(datastore)
		documents, errorInfo.Error = tCollectionReferencePtr.Documents(CTXBackground).GetAll()
		if documents == nil && errorInfo.Error == nil {
			errorInfo.Error = errs.ErrDocumentsNoneFound
		}
	}

	return
}

// GetAllDocumentsWhere will return snapshot pointers to each document in the datastore that meet the where condition.
// If no documents are found, the documents will have a count of zero.
//
//	Customer Messages: None
//	Errors: errs.ErrRequiredArgumentMissing, errs.ErrDocumentsNoneFound, errs.ErrServiceFailedFIRESTORE
//	Verifications: None
// func GetAllDocumentsWhere(firestoreClientPtr *firestoreServices.Client, datastore, fieldName string, fieldValue interface{}) (documents []*firestoreServices.DocumentSnapshot, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tQuery             firestoreServices.Query
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == ctv.VAL_EMPTY || fieldName == ctv.VAL_EMPTY || fieldValue == nil {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v Field Name: %v Field Value: %v", firestoreClientPtr, datastore, fieldName, fieldValue)
// 		errs.PrintError(errorInfo)
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Where(fieldName, "==", fieldValue)
// 		if documents, errorInfo.Error = tQuery.Documents(CTXBackground).GetAll(); len(documents) == 0 {
// 			if errorInfo.Error == nil {
// 				errorInfo.AdditionalInfo = ctv.NOT_FOUND + ctv.IS_OK
// 				errorInfo.Error = errs.ErrDocumentsNoneFound
// 				errs.PrintError(errorInfo)
// 			} else {
// 				errorInfo.AdditionalInfo = errorInfo.Error.Error()
// 				errorInfo.Error = errs.ErrServiceFailedFIRESTORE
// 				errs.PrintError(errorInfo)
// 			}
// 		}
// 	}
//
// 	return
// }

// GetSomeDocumentsWhere provides snapshot pointers to documents in the datastore that meet the specified 'where' condition, limited by the record count and starting from the offset position.
// If no documents are found, the documents variable will have a zero length.
//
//	Customer Messages: None
//	Errors: errs.ErrRequiredArgumentMissing
//	Verifications: None
// func GetSomeDocumentsWhere(firestoreClientPtr *firestoreServices.Client, datastore, fieldName string, fieldValue interface{}, offset, recordCount int) (documents []*firestoreServices.DocumentSnapshot, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tQuery             firestoreServices.Query
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == ctv.VAL_EMPTY || fieldName == ctv.VAL_EMPTY || fieldValue == nil {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v Field Name: %v Field Value: %v", firestoreClientPtr, datastore, fieldName, fieldValue)
// 		errs.PrintError(errorInfo)
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Where(fieldName, ctv.EQUALS, fieldValue).Offset(offset).Limit(recordCount)
// 		documents, errorInfo.Error = tQuery.Documents(CTXBackground).GetAll()
// 	}
//
// 	return
// }

// GetDocumentById - will return a non-nil documentSnapshotPtr if the document is found.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetDocumentById(firestoreClientPtr *firestore.Client, datastore string, documentId string) (documentSnapshotPtr *firestore.DocumentSnapshot, errorInfo errs.ErrorInfo) {

	if firestoreClientPtr == nil || datastore == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf(errs.FORMAT_FIRESTORE_ARGUMENTS_MISSING, datastore, documentId))
	} else {
		if documentSnapshotPtr, errorInfo.Error = firestoreClientPtr.Doc(datastore + "/" + documentId).Get(CTXBackground); documentSnapshotPtr == nil || errorInfo.Error != nil {
			if strings.Contains(errorInfo.Error.Error(), ctv.TXT_NOT_FOUND) {
				errorInfo.Error = errs.ErrDocumentNotFound
			}
			documentSnapshotPtr = nil
		}
	}

	return
}

// getDocumentRef
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getDocumentRef(firestoreClientPtr *firestore.Client, datastore, documentId string) (documentReferencePtr *firestore.DocumentRef, errorInfo errs.ErrorInfo) {

	if datastore == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf(errs.FORMAT_FIRESTORE_ARGUMENTS_MISSING, datastore, documentId))
		log.Println(errorInfo.Error.Error())
	} else {
		documentReferencePtr = firestoreClientPtr.Collection(datastore).Doc(documentId)
		if doesDocumentExist(documentReferencePtr) == false {
			errorInfo.Error = errors.New(fmt.Sprintf("The document was not found. %v: '%v'", ctv.FN_DOCUMENT_ID, documentId))
			log.Println(errorInfo.Error.Error())
			documentReferencePtr = nil
		}
	}

	return
}

// GetDocumentIdsWithSubCollections
// func GetDocumentIdsWithSubCollections(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName string) (documentRefIds []string, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tPath              string
// 		tDocumentPtr       []*firestoreServices.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v", datastore, parentDocumentId, subCollectionName)
// 		tDocumentPtr, errorInfo.Error = firestoreClientPtr.Collection(tPath).Documents(CTXBackground).GetAll()
// 		for _, snapshot := range tDocumentPtr {
// 			documentRefIds = append(documentRefIds, snapshot.Ref.ID)
// 		}
// 	}
//
// 	return
// }

// GetDocumentFromSubCollectionByDocumentId
//
//	If the document is not found, an error will be returned.
// func GetDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName, documentId string) (data map[string]interface{}, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tDocumentRefPtr    *firestoreServices.DocumentRef
// 		tDocumentPtr       *firestoreServices.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tPath              string
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v/%v", datastore, parentDocumentId, subCollectionName, documentId)
// 		if tDocumentRefPtr = firestoreClientPtr.Doc(tPath); errorInfo.Error == nil {
// 			if tDocumentPtr, errorInfo.Error = tDocumentRefPtr.Get(CTXBackground); errorInfo.Error == nil {
// 				data = tDocumentPtr.Data()
// 			}
// 		}
// 	}
//
// 	return
// }

// GetFirestoreClientConnection - will connect to Firestore service using Firebase Auth.
//
//	Customer Messages: None
//	Errors: ErrServiceFailedFIREBASE,
//	Verifications: None
func GetFirestoreClientConnection(appPtr *firebase.App) (firestoreClientPtr *firestore.Client, errorInfo errs.ErrorInfo) {

	if appPtr == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFirebaseAppConnectionFailed, "Firebase appPtr is nil.")
		return
	}

	// firestoreClientPtr is in the function definition because error is passed up the stack by Firebase/Firestore
	if firestoreClientPtr, errorInfo.Error = appPtr.Firestore(context.Background()); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFirestoreClientFailed, ctv.VAL_EMPTY)
		return
	}

	log.Printf("The Firestore client has been created successfully.")

	return
}

// RemoveDocument
// func RemoveDocument(firestoreClientPtr *firestoreServices.Client, datastore string, queryParameters ...NameValueQuery) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tDocument          *firestoreServices.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tQuery             firestoreServices.Query
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || len(queryParameters) < 1 {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' nameValueQuery argument is '%v'", datastore, ctv.VAL_EMPTY))
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Query
// 		for _, parameter := range queryParameters {
// 			if parameter.FieldName == ctv.VAL_EMPTY || parameter.FieldValue == ctv.VAL_EMPTY {
// 				errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' nameValueQuery parameter is '%v' Field Name: %v, Field Value: %v", datastore, ctv.VAL_EMPTY,
// 					parameter.FieldName, parameter.FieldValue))
// 				break
// 			} else {
// 				tQuery = tQuery.Where(parameter.FieldName, ctv.EQUALS, parameter.FieldValue)
// 			}
// 		}
// 	}
//
// 	if errorInfo.Error == nil {
// 		tDocuments := tQuery.Documents(CTXBackground)
// 		for {
// 			tDocument, errorInfo.Error = tDocuments.Next()
// 			if errors.Is(errorInfo.Error, iterator.Done) {
// 				errorInfo.Error = nil
// 				break
// 			}
// 			if errorInfo.Error != nil {
// 				errorInfo.AdditionalInfo = fmt.Sprintf("An error occurred trying to remove a document. Error: %v", errorInfo.Error)
// 				errorInfo.Error = errs.ErrServiceFailedFIRESTORE
// 				errs.PrintError(errorInfo)
// 				// todo handle error & notification
// 			}
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(tDocument.Ref.ID).Delete(CTXBackground); errorInfo.Error != nil {
// 				errorInfo.AdditionalInfo = fmt.Sprintf("%v Failed: Investigate, there is something wrong! Error: %v", tFunctionName, errorInfo.Error.Error())
// 				errorInfo.Error = errs.ErrServiceFailedFIRESTORE
// 				errs.PrintError(errorInfo)
// 				// todo Handle error and Notification
// 			}
// 		}
// 	}
//
// 	return
// }

// RemoveDocumentById
// func RemoveDocumentById(firestoreClientPtr *firestoreServices.Client, datastore, documentId string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Document Id: '%v'", datastore, documentId))
// 	} else {
// 		_, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Delete(CTXBackground)
// 	}
//
// 	return
// }

// RemoveDocumentFromSubCollectionByDocumentId
// func RemoveDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName, documentId string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Parent Document Id: '%v' Sub-Collection Name: '%v' Document Id: '%v'", datastore, parentDocumentId,
// 			subCollectionName, documentId))
// 	} else {
// 		if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).Doc(documentId).Delete(CTXBackground); errorInfo.Error != nil {
// 			errorInfo.Error = errors.New(fmt.Sprintf("%v Failed: Investigate, there is something wrong! Error: %v", "removeDocument", errorInfo.Error.Error()))
// 			log.Println(errorInfo.Error.Error())
// 			// todo Handle error and Notification
// 		}
// 	}
//
// 	return
// }

// RemoveDocumentFromSubCollection
//
//	Customer Messages: None
//	Errors: errs.ErrRequiredArgumentMissing
//	Verification: Check datastore, parentDocumentId, and subCollectionName are populated
// func RemoveDocumentFromSubCollection(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName string) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tDocumentRefIterPtr *firestoreServices.DocumentRefIterator
// 		tDocumentRefPtr     *firestoreServices.DocumentRef
// 		tFunction, _, _, _  = runtime.Caller(0)
// 		tFunctionName       = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 	} else {
// 		tDocumentRefIterPtr = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).DocumentRefs(CTXBackground)
// 		for {
// 			tDocumentRefPtr, errorInfo.Error = tDocumentRefIterPtr.Next()
// 			if errors.Is(errorInfo.Error, iterator.Done) {
// 				errorInfo.Error = nil
// 				break
// 			}
// 			if errorInfo.Error != nil {
// 				break
// 			}
// 			_, _ = tDocumentRefPtr.Delete(CTXBackground)
// 		}
// 	}
//
// 	return
// }

// SetDocument - This will create or overwrite the record. While nameValues is a map[any], this function will apply a string assertion on the key.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func SetDocument(firestoreClientPtr *firestoreServices.Client, datastore, documentId string, nameValues map[any]interface{}) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if coreValidators.AreMapKeysPopulated(nameValues) == false {
// 		errorInfo.Error = errs.GetMapKeyPopulatedError(tFinding)
// 	} else {
// 		if firestoreClientPtr == nil || datastore == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 			errorInfo.Error = errs.ErrRequiredArgumentMissing
// 			errs.PrintError(errorInfo)
// 			// todo Handle errors and Notifications
// 		} else {
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Set(CTXBackground, coreHelpers.ConvertMapAnyToMapString(nameValues)); errorInfo.Error != nil {
// 				errorInfo.Error = errs.ErrServiceFailedFIRESTORE
// 				errs.PrintError(errorInfo)
// 				// todo Handle errors and Notifications
// 			}
// 		}
// 	}
//
// 	return
// }

// SetDocumentWithSubCollection - This will create or overwrite the existing record that is in a sub-collection. While nameValues is a map[any], this function will apply a string assertion on the key.
// func SetDocumentWithSubCollection(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName, documentId string, nameValues map[any]interface{}) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding != ctv.GOOD {
// 		errorInfo.Error = errs.GetMapKeyPopulatedError(tFinding)
// 	} else {
// 		// if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 		if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Parent Document Id: '%v' Sub-collection Name: '%v' Document Id: '%v' Function Name: %v", datastore, parentDocumentId, subCollectionName, documentId, tFunctionName))
// 			log.Println(errorInfo.Error.Error())
// 			// todo Handle errors and Notifications
// 		} else {
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).Doc(documentId).Set(CTXBackground, coreHelpers.ConvertMapAnyToMapString(nameValues)); errorInfo.Error != nil {
// 				errorInfo.Error = errors.New(fmt.Sprintf("An error has occurred creating Document Id: %v for Datastore: %v Parent Document Id: '%v' Subcollection Name: '%v' Error: %v", documentId, datastore,
// 					parentDocumentId, subCollectionName, errorInfo.Error.Error()))
// 				log.Println(errorInfo.Error.Error())
// 				// todo Handle errors and Notifications
// 			}
// 		}
// 	}
//
// 	return
// }

// UpdateDocument- will return an error of nil when successful. If the document is not found, shared_services.ErrDocumentNotFound will be returned, otherwise the error from Firestore will be returned.
// func UpdateDocument(firestoreClientPtr *firestoreServices.Client, datastore, documentId string, nameValues map[any]interface{}) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUpdateFields      []firestoreServices.Update
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	errorInfo.AdditionalInfo = fmt.Sprintf("Datastore: %v Document Id: %v", datastore, documentId)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding != ctv.GOOD {
// 		errorInfo.Error = errs.GetMapKeyPopulatedError(tFinding)
// 		errs.PrintError(errorInfo)
// 	} else {
// 		if datastore == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 			errorInfo.Error = errs.ErrRequiredArgumentMissing
// 			errs.PrintError(errorInfo)
// 			// todo Handle errors and Notifications
// 		} else {
// 			if tUpdateFields, errorInfo = BuildFirestoreUpdate(nameValues); errorInfo.Error == nil {
// 				if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Update(CTXBackground, tUpdateFields); errorInfo.Error != nil {
// 					errs.PrintError(errorInfo)
// 				}
// 			}
// 		}
// 	}
//
// 	return
// }

// UpdateDocumentFromSubCollectionByDocumentId
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, Any error from Firestore
//	Verifications: None
// func UpdateDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestoreServices.Client, datastore, parentDocumentId, subCollectionName, documentId string, updateFields []firestoreServices.Update) (errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tPath              string
// 	)
//
// 	errs.PrintDebugTrail(tFunctionName)
//
// 	if datastore == ctv.VAL_EMPTY || parentDocumentId == ctv.VAL_EMPTY || subCollectionName == ctv.VAL_EMPTY || documentId == ctv.VAL_EMPTY {
// 		errorInfo.Error = errs.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v/%v", datastore, parentDocumentId, subCollectionName, documentId)
// 		_, errorInfo.Error = firestoreClientPtr.Doc(tPath).Update(CTXBackground, updateFields)
// 	}
//
// 	return
// }
