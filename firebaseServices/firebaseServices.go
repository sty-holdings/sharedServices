package sharedServices

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

var (
	CTXBackground = context.Background()
)

// NewFirebaseApp - creates a new Firebase App
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewFirebaseApp(credentialsFQN string) (
	appPtr *firebase.App,
	errorInfo errs.ErrorInfo,
) {

	// DO NOT DELETE
	// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
	// Case 10319546 has been filed with Firebase Support
	// if appPtr, errorInfo.Error = firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)); errorInfo.Error != nil {
	//	errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.ErrServiceFailedFIREBASE.Error())
	// }
	// END DO NOT DELETE
	appPtr, errorInfo.Error = firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN))

	return
}

// FindFirebaseAuthUser - determines if the user exists in the Firebase Auth database. If so, then pointer to the user is return, otherwise, an error.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func FindFirebaseAuthUser(
	authPtr *auth.Client,
	username string,
) (
	userRecordPtr *auth.UserRecord,
	errorInfo errs.ErrorInfo,
) {

	if userRecordPtr, errorInfo.Error = authPtr.GetUser(CTXBackground, username); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
	}

	return
}

// GetFirebaseAppAuthConnection - will create an App and an Auth instance
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseAppAuthConnection(credentialsFQN string) (
	appPtr *firebase.App,
	authPtr *auth.Client,
	errorInfo errs.ErrorInfo,
) {

	// DO NOT DELETE
	// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
	// Case 10319546 has been filed with Firebase Support
	// if appPtr, errorInfo = NewFirebaseApp(credentialsLocation); errorInfo.Error == nil {
	//	authPtr, errorInfo = GetFirebaseAuthConnection(appPtr)
	// }
	// END DO NOT DELETE

	appPtr, errorInfo = NewFirebaseApp(credentialsFQN)
	authPtr, errorInfo = GetFirebaseAuthConnection(appPtr)

	return
}

// GetFirebaseIdTokenPayload
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseIdTokenPayload(
	authPtr *auth.Client,
	idToken string,
) (
	tokenPayload map[any]interface{},
	errorInfo errs.ErrorInfo,
) {

	var (
		tIdTokenPtr *auth.Token
	)

	tokenPayload = make(map[any]interface{})
	if tIdTokenPtr, errorInfo = GetIdTokenPtr(authPtr, idToken); errorInfo.Error == nil {
		tokenPayload[JWT_PAYLOAD_SUBJECT_FN] = tIdTokenPtr.Subject
		tokenPayload[JWT_PAYLOAD_CLAIMS_FN] = tIdTokenPtr.Claims
		tokenPayload[JWT_PAYLOAD_AUDIENCE_FN] = tIdTokenPtr.Audience
		tokenPayload[JWT_PAYLOAD_REQUESTOR_ID_FN] = tIdTokenPtr.UID
		tokenPayload[JWT_PAYLOAD_EXPIRES_FN] = tIdTokenPtr.Expires
		tokenPayload[JWT_PAYLOAD_ISSUER_FN] = tIdTokenPtr.Issuer
		tokenPayload[JWT_PAYLOAD_ISSUED_AT_FN] = tIdTokenPtr.IssuedAt
	} else {
		errorInfo.Error = errors.New(fmt.Sprintf("The provided idTokenPtr is invalid. ERROR: %v", errorInfo.Error.Error()))
	}

	return
}

// GetFirebaseClientInfo - retrieves client information from the Firestore database using the client ID.
//
//	Customer Messages: None
//	Errors: errs.ErrNoFoundDocument, errs.ErrorInfo
//	Verifications: None
func GetFirebaseClientInfo(
	firestoreClientPtr *firestore.Client,
	styhClientID string,
) (
	clientInfo map[string]interface{},
	errorInfo errs.ErrorInfo,
) {

	var (
		tUserDocumentSnapshotPtr *firestore.DocumentSnapshot
	)

	if tUserDocumentSnapshotPtr, errorInfo = GetDocumentById(firestoreClientPtr, DATASTORE_CLIENTS, styhClientID); errorInfo.Error != nil {
		return
	}

	clientInfo = tUserDocumentSnapshotPtr.Data()

	return
}

// GetFirebaseUserInfo - checks if the use exists and returns the user database record when found.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseUserInfo(
	authPtr *auth.Client,
	firestoreClientPtr *firestore.Client,
	username string,
) (
	userInfo map[string]interface{},
	errorInfo errs.ErrorInfo,
) {

	var (
		tUserDocumentSnapshotPtr *firestore.DocumentSnapshot
	)

	if _, errorInfo = FindFirebaseAuthUser(authPtr, username); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%s%s", ctv.LBL_USER_NAME, username))
		return
	}

	if tUserDocumentSnapshotPtr, errorInfo = GetDocumentById(firestoreClientPtr, DATASTORE_USERS, username); errorInfo.Error != nil {
		return
	}

	userInfo = tUserDocumentSnapshotPtr.Data()

	return
}

// GetIdTokenPtr
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIdTokenPtr(
	authPtr *auth.Client,
	idToken string,
) (
	IdTokenPtr *auth.Token,
	errorInfo errs.ErrorInfo,
) {

	if IdTokenPtr, errorInfo.Error = authPtr.VerifyIDToken(CTXBackground, idToken); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error())
	}

	return
}

// IsFirebaseIdTokenValid
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsFirebaseIdTokenValid(
	authPtr *auth.Client,
	idToken string,
) bool {

	if _, err := authPtr.VerifyIDToken(CTXBackground, idToken); err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

// GetFirebaseAuthConnection - creates a new Firebase Auth Connection
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseAuthConnection(appPtr *firebase.App) (
	authPtr *auth.Client,
	errorInfo errs.ErrorInfo,
) {

	if authPtr, errorInfo.Error = appPtr.Auth(CTXBackground); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, "")
	}

	return
}

// SetFirebaseAuthEmailVerified - This will set the Firebase Auth email verify flag to true
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func SetFirebaseAuthEmailVerified(
	authPtr *auth.Client,
	username string,
) (errorInfo errs.ErrorInfo) {

	var (
		tUserRecordPtr *auth.UserRecord
	)

	if tUserRecordPtr, errorInfo = FindFirebaseAuthUser(authPtr, username); tUserRecordPtr != nil {
		params := (&auth.UserToUpdate{}).EmailVerified(true)
		if _, errorInfo.Error = authPtr.UpdateUser(CTXBackground, username, params); errorInfo.Error != nil {
			errorInfo.Error = errors.New(
				fmt.Sprintf(
					"Firebase Auth - Setting email verify to true, failed for Requestor Id: %v Error: %v",
					username,
					errorInfo.Error,
				),
			)
			log.Println(errorInfo.Error.Error())
		}
	}

	return
}

// ValidateFirebaseJWTPayload - Firebase ID Token that is returned when a user logs on successfully
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ValidateFirebaseJWTPayload(
	tokenPayload map[any]interface{},
	audience, issuer string,
) (errorInfo errs.ErrorInfo) {

	var (
		tFindings string
		tSubject  string
		username  string
	)

	if tFindings = vals.AreMapKeysValuesPopulated(tokenPayload); tFindings != ctv.TXT_YES {
		errorInfo = errs.NewErrorInfo(errs.ErrEmptyVariableMapValue, ctv.VAL_EMPTY)
	} else {
		if audience == ctv.VAL_EMPTY || issuer == ctv.VAL_EMPTY {
			errorInfo.Error = errors.New(
				fmt.Sprintf(
					"Require information is missing! %v: '%v' %v: '%v'",
					ctv.FN_AUDIENCE_CAP,
					audience,
					ctv.FN_JWT_ISSUER,
					issuer,
				),
			)
		} else {
			for key, value := range tokenPayload {
				switch strings.ToUpper(key.(string)) {
				case JWT_PAYLOAD_AUDIENCE_FN:
					if value != audience {
						errorInfo.Error = errors.New("The audience of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case JWT_PAYLOAD_ISSUER_FN:
					if value != issuer {
						errorInfo.Error = errors.New("The issuer of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case JWT_PAYLOAD_SUBJECT_FN:
					tSubject = value.(string)
				case JWT_PAYLOAD_REQUESTOR_ID_FN:
					username = value.(string)
				}
			}
			if username != tSubject {
				errorInfo.Error = errors.New("The requestorId/user_id do not match the subject/sub. The ID is invalid.")
				log.Println(errorInfo.Error.Error())
			}
		}
	}

	return
}
