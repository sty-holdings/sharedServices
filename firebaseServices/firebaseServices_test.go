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

	Written by <Replace with FULL_NAME> / syacko
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
	"runtime"
	"testing"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

//goland:noinspection ALL
const (
	FIREBASE_CREDENTIALS_FILENAME     = "/Volumes/development-share/.keys/com.styholdings.dkanswers/google/service-account-key/dkanswers-key.json"
	BAD_FIREBASE_CREDENTIALS_FILENAME = "/Volumes/development-share/.keys/com.styholdings.dkanswers/google/service-account-key/dkanswers-key.txt"
	TEST_LOCAL_USERNAME               = "U7NjH4JilwcRmUJK8aBBeoUigzw2"
	TEST_BAD_LOCAL_USERNAME           = "U7NjH4JilwcRmUJK8aBBeogzw2"
)

var (
//goland:noinspection ALL
)

func TestFindFirebaseAuthUser(tPtr *testing.T) {

	var (
		errorInfo          pi.ErrorInfo
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(t *testing.T) {
			if _, tAuthPtr, errorInfo = GetFirebaseAppAuthConnection(FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error != nil {
				tPtr.Fatal(errorInfo.Error.Error())
			}
			if _, errorInfo = FindFirebaseAuthUser(tAuthPtr, TEST_BAD_LOCAL_USERNAME); errorInfo.Error == nil {
				tPtr.Errorf(pi.FORMAT_EXPECTED_ERROR, tFunctionName, ctv.STATUS_COMPLETED)
			}
			if _, errorInfo = FindFirebaseAuthUser(tAuthPtr, TEST_LOCAL_USERNAME); errorInfo.Error != nil {
				tPtr.Errorf(pi.FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo.Error.Error())
			}
		},
	)

}

func TestGetIdTokenPayload(tPtr *testing.T) {

	var (
	//errorInfo          pi.ErrorInfo
	//tFileData          []byte
	//tAuthPtr           *auth.Client
	//tFunction, _, _, _ = runtime.Caller(0)
	//tFunctionName      = runtime.FuncForPC(tFunction).Name()
	//tTokenPayload      = make(map[any]interface{})
	)

	//if tFileData, errorInfo.Error = os.ReadFile(FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error == nil {
	//_, tAuthPtr, _ = GetFirebaseAppAuthConnection(string(tFileData))
	//} else {
	//	tPtr.Error(errorInfo.Error)
	//}

	//tPtr.Run(
	//	tFunctionName, func(t *testing.T) {
	//		if tTokenPayload, _ = GetFirebaseIdTokenPayload(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); len(tTokenPayload) == 0 {
	//			tPtr.Errorf("%v Failed: Was expecting the JWT payload to be populated.", tFunctionName)
	//		}
	//	},
	//)
}

func TestGetIdTokenPtr(tPtr *testing.T) {

	var (
	//errorInfo          pi.ErrorInfo
	//tAuthPtr           *auth.Client
	//tFileData          []byte
	//tFunction, _, _, _ = runtime.Caller(0)
	//tFunctionName      = runtime.FuncForPC(tFunction).Name()
	//tIdTokenPtr        *auth.Token
	)

	//if tFileData, errorInfo.Error = os.ReadFile(FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error == nil {
	//	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(string(tFileData))
	//} else {
	//	tPtr.Error(errorInfo.Error)
	//}

	//tPtr.Run(
	//	tFunctionName, func(t *testing.T) {
	//		if tIdTokenPtr, _ = GetIdTokenPtr(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); tIdTokenPtr == nil {
	//			tPtr.Errorf("%v Failed: No Token was return. Make sure the tIdTokenValid is a valid and recent JWT.", tFunctionName)
	//		}
	//	},
	//)
}

func TestIsFirebaseIdTokenValid(tPtr *testing.T) {

	var (
	//errorInfo          pi.ErrorInfo
	//tAuthPtr           *auth.Client
	//tFileData          []byte
	//tFunction, _, _, _ = runtime.Caller(0)
	//tFunctionName      = runtime.FuncForPC(tFunction).Name()
	//tValid             bool
	)

	//if tFileData, errorInfo.Error = os.ReadFile(FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error == nil {
	//	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(string(tFileData))
	//} else {
	//	tPtr.Error(errorInfo.Error)
	//}

	//tPtr.Run(
	//	tFunctionName, func(tPtr *testing.T) {
	//		if tValid = IsFirebaseIdTokenValid(tAuthPtr, TEST_FIREBASE_IDTOKEN_INVALID); tValid == true {
	//			tPtr.Errorf("%v Failed: Token is should be invalid. Valid: %v", tFunctionName, tValid)
	//		}
	//		if tValid = IsFirebaseIdTokenValid(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); tValid == false {
	//			tPtr.Errorf("%v Failed: Token is should be valid. Valid: %v", tFunctionName, tValid)
	//		}
	//	},
	//)
}

func TestNewFirebaseApp(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// DO NOT DELETE
			// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
			// Case 10319546 has been filed with Firebase Support
			//if tAppPtr, errorInfo = NewFirebaseApp(BAD_FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error == nil {
			//	tPtr.Errorf(pi.EXPECTED_ERROR_FORMAT, tFunctionName)
			//}
			// END DO NOT DELETE
			NewFirebaseApp(FIREBASE_CREDENTIALS_FILENAME)
		},
	)
}

func TestGetFirebaseAuthConnection(tPtr *testing.T) {

	var (
		errorInfo          pi.ErrorInfo
		tAppPtr            *firebase.App
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// DO NOT DELETE
			// This code will not work because the underlying firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsFQN)) will always return an object.
			// Case 10319546 has been filed with Firebase Support
			//if tAppPtr, errorInfo = NewFirebaseApp(BAD_FIREBASE_CREDENTIALS_FILENAME); errorInfo.Error == nil {
			//	tPtr.Errorf(pi.EXPECTED_ERROR_FORMAT, tFunctionName)
			//}
			// END DO NOT DELETE
			tAppPtr, _ = NewFirebaseApp(BAD_FIREBASE_CREDENTIALS_FILENAME)
			if _, errorInfo = GetFirebaseAuthConnection(tAppPtr); errorInfo.Error == nil {
				tPtr.Errorf(pi.FORMAT_EXPECTED_ERROR, tFunctionName, pi.FIREBASE_AUTH_BAD_CREDENTIALS)
			}
			tAppPtr, _ = NewFirebaseApp(FIREBASE_CREDENTIALS_FILENAME)
			if _, errorInfo = GetFirebaseAuthConnection(tAppPtr); errorInfo.Error != nil {
				tPtr.Errorf(pi.FORMAT_EXPECTING_NO_ERROR, tFunctionName, errorInfo.Error.Error())
			}
		},
	)
}

func TestValidateFirebaseJWTPayload(tPtr *testing.T) {

	var (
	//errorInfo          pi.ErrorInfo
	//tAuthPtr           *auth.Client
	//tFileData          []byte
	//tFunction, _, _, _ = runtime.Caller(0)
	//tFunctionName      = runtime.FuncForPC(tFunction).Name()
	//tTokenPayload      = make(map[any]interface{})
	//tValid             bool
	)

	//if tFileData, errorInfo.Error = os.ReadFile(FIREBASE_CREDENTIALS_FILENAME); err == nil {
	//	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(string(tFileData))
	//} else {
	//	tPtr.Error(errorInfo.Error)
	//}
	//tTokenPayload, _ = GetFirebaseIdTokenPayload(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID)

	//tPtr.Run(
	//	tFunctionName, func(tPtr *testing.T) {
	//		if errorInfo = ValidateFirebaseJWTPayload(
	//			tTokenPayload,
	//			ctv.CERT_SAVUPDEV_AUDIENCE,
	//			ctv.CERT_DEV_ID_TOEKN_ISSUER,
	//		); errorInfo.Error != nil {
	//			tPtr.Errorf("%v Failed: Token payload should be valid. Valid: %v", tFunctionName, tValid)
	//		}
	//	},
	//)
}
