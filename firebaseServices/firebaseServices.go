// Package coreFirebase
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by Scott Yacko / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	www.sty-holdings.com

	01-2024
	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package coreFirebase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	CTXBackground = context.Background()
)

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
	errorInfo pi.ErrorInfo,
) {

	if userRecordPtr, errorInfo.Error = authPtr.GetUser(CTXBackground, username); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, ctv.VAL_EMPTY)
	}

	return
}

// GetFirebaseFirestoreConnection
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseAppAuthConnection(credentialsFQN string) (
	appPtr *firebase.App,
	authPtr *auth.Client,
	errorInfo pi.ErrorInfo,
) {

	// DO NOT DELETE
	// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
	// Case 10319546 has been filed with Firebase Support
	//if appPtr, errorInfo = NewFirebaseApp(credentialsLocation); errorInfo.Error == nil {
	//	authPtr, errorInfo = GetFirebaseAuthConnection(appPtr)
	//}
	// END DO NOT DELETE

	appPtr, errorInfo = NewFirebaseApp(credentialsFQN)
	authPtr, errorInfo = GetFirebaseAuthConnection(appPtr)

	return
}

// GetFirebaseIdTokenPayload
func GetFirebaseIdTokenPayload(
	authPtr *auth.Client,
	idToken string,
) (
	tokenPayload map[any]interface{},
	errorInfo pi.ErrorInfo,
) {

	var (
		tIdTokenPtr *auth.Token
	)

	tokenPayload = make(map[any]interface{})
	if tIdTokenPtr, errorInfo = GetIdTokenPtr(authPtr, idToken); errorInfo.Error == nil {
		tokenPayload[ctv.PAYLOAD_SUBJECT_FN] = tIdTokenPtr.Subject
		tokenPayload[ctv.PAYLOAD_CLAIMS_FN] = tIdTokenPtr.Claims
		tokenPayload[ctv.PAYLOAD_AUDIENCE_FN] = tIdTokenPtr.Audience
		tokenPayload[ctv.PAYLOAD_REQUESTOR_ID_FN] = tIdTokenPtr.UID
		tokenPayload[ctv.PAYLOAD_EXPIRES_FN] = tIdTokenPtr.Expires
		tokenPayload[ctv.PAYLOAD_ISSUER_FN] = tIdTokenPtr.Issuer
		tokenPayload[ctv.PAYLOAD_ISSUED_AT_FN] = tIdTokenPtr.IssuedAt
	} else {
		errorInfo.Error = errors.New(fmt.Sprintf("The provided idTokenPtr is invalid. ERROR: %v", errorInfo.Error.Error()))
	}

	return
}

// GetIdTokenPtr
func GetIdTokenPtr(
	authPtr *auth.Client,
	idToken string,
) (
	IdTokenPtr *auth.Token,
	errorInfo pi.ErrorInfo,
) {

	if IdTokenPtr, errorInfo.Error = authPtr.VerifyIDToken(CTXBackground, idToken); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error())
	}

	return
}

// IsFirebaseIdTokenValid
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

// NewFirebaseApp - creates a new Firebase App
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewFirebaseApp(credentialsFQN string) (
	appPtr *firebase.App,
	errorInfo pi.ErrorInfo,
) {

	// DO NOT DELETE
	// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
	// Case 10319546 has been filed with Firebase Support
	//if appPtr, errorInfo.Error = firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)); errorInfo.Error != nil {
	//	errorInfo = pi.NewErrorInfo(errorInfo.Error, pi.ErrServiceFailedFIREBASE.Error())
	//}
	// END DO NOT DELETE
	appPtr, errorInfo.Error = firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN))

	return
}

// GetFirebaseAuthConnection - creates a new Firebase Auth Connection
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFirebaseAuthConnection(appPtr *firebase.App) (
	authPtr *auth.Client,
	errorInfo pi.ErrorInfo,
) {

	if authPtr, errorInfo.Error = appPtr.Auth(CTXBackground); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, "")
	} else {
		log.Println("The Firebase Auth client has been created.")
	}

	return
}

// SetFirebaseAuthEmailVerified - This will set the Firebase Auth email verify flag to true
func SetFirebaseAuthEmailVerified(
	authPtr *auth.Client,
	username string,
) (errorInfo pi.ErrorInfo) {

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
func ValidateFirebaseJWTPayload(
	tokenPayload map[any]interface{},
	audience, issuer string,
) (errorInfo pi.ErrorInfo) {

	var (
		tFindings string
		tSubject  string
		username  string
	)

	if tFindings = hvs.AreMapKeysValuesPopulated(tokenPayload); tFindings != ctv.TXT_YES {
		errorInfo = pi.NewErrorInfo(pi.ErrMapIsMissingKey, ctv.VAL_EMPTY)
	} else {
		if audience == ctv.VAL_EMPTY || issuer == ctv.VAL_EMPTY {
			errorInfo.Error = errors.New(
				fmt.Sprintf(
					"Require information is missing! %v: '%v' %v: '%v'",
					ctv.FN_AUDIENCE,
					audience,
					ctv.FN_ISSUER,
					issuer,
				),
			)
		} else {
			for key, value := range tokenPayload {
				switch strings.ToUpper(key.(string)) {
				case ctv.PAYLOAD_AUDIENCE_FN:
					if value != audience {
						errorInfo.Error = errors.New("The audience of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case ctv.PAYLOAD_ISSUER_FN:
					if value != issuer {
						errorInfo.Error = errors.New("The issuer of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case ctv.PAYLOAD_SUBJECT_FN:
					tSubject = value.(string)
				case ctv.PAYLOAD_REQUESTOR_ID_FN:
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
