// package shared
/*
NOTES:
    None

COPYRIGHT:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package sendGridService

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"albert/constants"
	"albert/core/coreError"
	"github.com/plaid/plaid-go/v9/plaid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	//
	//  Data
	TestEmptyList = EmailItem{}
	TestFromEmail = EmailItem{
		Name:    "From First and Last Name",
		Address: "from_developer@sty-holdings.com",
	}
	TestToListOne = EmailItem{
		Name:    "TO1 First and Last Name",
		Address: "TO1_developer@sty-holdings.com",
	}
	TestToListTwo = EmailItem{
		Name:    "TO2 First and Last Name",
		Address: "TO2_developer@example.com",
	}
	TestCCListOne = EmailItem{
		Name:    "CC1 First and Last Name",
		Address: "CC1_developer@example.com",
	}
	TestCCListTwo = EmailItem{
		Name:    "CC2 First and Last Name",
		Address: "CC2_developer@example.com",
	}
	TestBCCListOne = EmailItem{
		Name:    "BCC1 First and Last Name",
		Address: "BCC1_developer@example.com",
	}
	TestBCCListTwo = EmailItem{
		Name:    "BCC2 First and Last Name",
		Address: "BCC2_developer@example.com",
	}
	TestMultipleEmptyRecipientList = []EmailItem{TestEmptyList, TestEmptyList}
	TestMultipleTORecipientList    = []EmailItem{TestToListOne, TestToListTwo}
	TestMultipleCCRecipientList    = []EmailItem{TestCCListOne, TestCCListTwo}
	TestMultipleBCCRecipientList   = []EmailItem{TestBCCListOne, TestBCCListTwo}
	//
	TestAccountBaseOne = plaid.AccountBase{
		AccountId:            constants.TEST_USER_BANK_ACCOUNT_ID_1,
		Balances:             plaid.AccountBalance{},
		Mask:                 plaid.NullableString{},
		Name:                 "Plaid Checking",
		OfficialName:         plaid.NullableString{},
		Type:                 "depository",
		Subtype:              plaid.NullableAccountSubtype{},
		VerificationStatus:   nil,
		AdditionalProperties: nil,
	}
	//
	TestAccountBaseTwo = plaid.AccountBase{
		AccountId:            constants.TEST_USER_BANK_ACCOUNT_ID_2,
		Balances:             plaid.AccountBalance{},
		Mask:                 plaid.NullableString{},
		Name:                 "Plaid Savings",
		OfficialName:         plaid.NullableString{},
		Type:                 "depository",
		Subtype:              plaid.NullableAccountSubtype{},
		VerificationStatus:   nil,
		AdditionalProperties: nil,
	}
	//
	TestAccounts = []plaid.AccountBase{
		TestAccountBaseOne,
		TestAccountBaseTwo,
	}
)

func TestNewSendGridServer(tPtr *testing.T) {
	var (
		errorInfo          coreError.ErrorInfo
		tEmailServerPtr    *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Success
			if tEmailServerPtr, errorInfo = NewSendGridServer(
				constants.TEST_EMAIL_ADDRESS,
				constants.TEST_STRING,
				constants.ENVIRONMENT_DEVELOPMENT,
				constants.TEST_STRING,
			); errorInfo.Error != nil && tEmailServerPtr != nil {
				tPtr.Errorf("%v Failed: Was not expecting no error but got %v.", tFunctionName, errorInfo)
			}
			// Missing default sender email
			if _, errorInfo = NewSendGridServer(constants.EMPTY, constants.TEST_STRING, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_STRING); errorInfo.Error == nil {
				tPtr.Errorf("%v Failed: Was not expecting an error but got %v.", tFunctionName, errorInfo)
			}
		},
	)
}

func TestEmail_newPersonalization(tPtr *testing.T) {

	var (
		errorInfo           coreError.ErrorInfo
		tEmailServer        *EmailServer
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		tPersonalizationPtr = mail.NewPersonalization()
	)

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_STRING, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Missing Tolist
			if errorInfo = tEmailServer.newPersonalization(tPersonalizationPtr, TestMultipleEmptyRecipientList, TestMultipleCCRecipientList, TestMultipleBCCRecipientList); errorInfo.Error == nil {
				tPtr.Errorf("%v Failed: Was not expecting an error but got %v.", tFunctionName, errorInfo.Error)
			}
			// Success with missing CClist
			if errorInfo = tEmailServer.newPersonalization(tPersonalizationPtr, TestMultipleTORecipientList, TestMultipleEmptyRecipientList, TestMultipleBCCRecipientList); errorInfo.Error != nil {
				tPtr.Errorf("%v Failed: Was not expecting an error but got %v.", tFunctionName, errorInfo.Error)
			}
			// Success with missing BCClist
			if errorInfo = tEmailServer.newPersonalization(tPersonalizationPtr, TestMultipleTORecipientList, TestMultipleCCRecipientList, TestMultipleEmptyRecipientList); errorInfo.Error != nil {
				tPtr.Errorf("%v Failed: Was not expecting an error but got %v.", tFunctionName, errorInfo.Error)
			}
			// Success
			if errorInfo = tEmailServer.newPersonalization(tPersonalizationPtr, TestMultipleTORecipientList, TestMultipleCCRecipientList, TestMultipleBCCRecipientList); errorInfo.Error != nil {
				tPtr.Errorf("%v Failed: Was not expecting an error but got %v.", tFunctionName, errorInfo.Error)
			}
		},
	)
}

func TestEmail_sendEmailUsingPlainText(tPtr *testing.T) {

	type arguments struct {
		from    EmailItem
		subject string
		body    string
		toList  []EmailItem
		ccList  []EmailItem
		bccList []EmailItem
		replyTo EmailItem
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tEmailServer       *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleCCRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing subject!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.EMPTY,
				body:    constants.TEST_STRING,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleCCRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing body!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.EMPTY,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleCCRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing toList!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleEmptyRecipientList,
				ccList:  TestMultipleCCRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Positive Case: Missing ccList!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleEmptyRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Positive Case: Missing bccList!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleBCCRecipientList,
				bccList: TestMultipleEmptyRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Positive Case: Missing from!",
			arguments: arguments{
				from:    TestEmptyList,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleTORecipientList,
				ccList:  TestMultipleBCCRecipientList,
				bccList: TestMultipleEmptyRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing relyTo!",
			arguments: arguments{
				from:    TestFromEmail,
				subject: constants.TEST_EMAIL_SUBJECT,
				body:    constants.TEST_STRING,
				toList:  TestMultipleEmptyRecipientList,
				ccList:  TestMultipleCCRecipientList,
				bccList: TestMultipleBCCRecipientList,
				replyTo: EmailItem{},
			},
			wantError: true,
		},
	}

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_STRING, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				if _, errorInfo = tEmailServer.sendEmailUsingPlainText(
					ts.arguments.from,
					ts.arguments.subject,
					ts.arguments.body,
					ts.arguments.toList,
					ts.arguments.ccList,
					ts.arguments.bccList,
					ts.arguments.replyTo,
				); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}

}

func TestEmail_sendEmailUsingTemplate(tPtr *testing.T) {

	type arguments struct {
		from         EmailItem
		subject      string
		templateId   string
		templateData map[any]interface{}
		toList       []EmailItem
		ccList       []EmailItem
		bccList      []EmailItem
		replyTo      EmailItem
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tEmailServer       *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing subject!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.EMPTY,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing template id!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.EMPTY,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing template data!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.EMPTY,
				templateData: map[any]interface{}{"su_first_name": nil},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing toList!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleEmptyRecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: true,
		},
		{
			name: "Positive Case: Missing ccList!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleEmptyRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Positive Case: Missing bccList!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleBCCRecipientList,
				bccList:      TestMultipleEmptyRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Positive Case: Missing from!",
			arguments: arguments{
				from:         TestEmptyList,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleTORecipientList,
				ccList:       TestMultipleBCCRecipientList,
				bccList:      TestMultipleEmptyRecipientList,
				replyTo: EmailItem{
					Name:    constants.TEST_EMAIL_NAME,
					Address: constants.TEST_EMAIL_ADDRESS,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing relyTo!",
			arguments: arguments{
				from:         TestFromEmail,
				subject:      constants.TEST_EMAIL_SUBJECT,
				templateId:   constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
				templateData: map[any]interface{}{"su_first_name": "Scott"},
				toList:       TestMultipleEmptyRecipientList,
				ccList:       TestMultipleCCRecipientList,
				bccList:      TestMultipleBCCRecipientList,
				replyTo:      EmailItem{},
			},
			wantError: true,
		},
	}

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_STRING, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				if _, errorInfo = tEmailServer.SendEmailUsingTemplate(
					ts.arguments.from, ts.arguments.subject, ts.arguments.toList, ts.arguments.ccList, ts.arguments.bccList,
					ts.arguments.replyTo, ts.arguments.templateId, ts.arguments.templateData, true,
				); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}

}

func TestGenerateBankRegisteredEmail(tPtr *testing.T) {

	type arguments struct {
		firstName      string
		lastName       string
		email          string
		accountDetails []plaid.AccountBase
	}

	var (
		err                error
		gotError           bool
		errorInfo          coreError.ErrorInfo
		tEmailServer       *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				firstName:      constants.TEST_USER_FIRST_NAME,
				lastName:       constants.TEST_USER_LAST_NAME,
				email:          constants.TEST_USER_EMAIL,
				accountDetails: TestAccounts,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Empty Account Details!",
			arguments: arguments{
				firstName:      constants.TEST_USER_FIRST_NAME,
				lastName:       constants.TEST_USER_LAST_NAME,
				email:          constants.TEST_USER_EMAIL,
				accountDetails: []plaid.AccountBase{},
			},
			wantError: false,
		},
	}

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_SENDER_NAME, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = GenerateBankRegisteredEmail(
					tEmailServer,
					constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
					ts.arguments.firstName,
					ts.arguments.lastName,
					ts.arguments.email,
					constants.TEST_INSTITUTION_CITIZEN_BANK,
					ts.arguments.accountDetails,
					false,
				); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, err.Error())
				}
			},
		)
	}
}

func TestGenerateTransferRequestEmail(tPtr *testing.T) {

	type arguments struct {
		firstName       string
		lastName        string
		email           string
		transferRequest map[string]string
	}

	var (
		err                error
		gotError           bool
		errorInfo          coreError.ErrorInfo
		tEmailServer       *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTestMapData       = make(map[string]string)
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				firstName:       constants.TEST_USER_FIRST_NAME,
				lastName:        constants.TEST_USER_LAST_NAME,
				email:           constants.TEST_USER_EMAIL,
				transferRequest: tTestMapData,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Empty Account Details!",
			arguments: arguments{
				firstName:       constants.TEST_USER_FIRST_NAME,
				lastName:        constants.TEST_USER_LAST_NAME,
				email:           constants.TEST_USER_EMAIL,
				transferRequest: make(map[string]string),
			},
			wantError: true,
		},
	}

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_SENDER_NAME, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)
	tTestMapData["direction"] = "into"
	tTestMapData["amount"] = strconv.FormatFloat(123.45, 'g', 5, 64)
	tTestMapData["method"] = strings.ToTitle(constants.TRANFER_WIRE)
	tTestMapData["completion"] = time.Now().AddDate(0, 2, 0).Format("Mon Jan 2, 2006")

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = GenerateTransferRequestEmail(
					tEmailServer,
					constants.TEST_EMAIL_TEMPLATE_VERIFACTION_ID,
					ts.arguments.firstName,
					ts.arguments.lastName,
					ts.arguments.email,
					ts.arguments.transferRequest,
					true,
				); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, err.Error())
				}
			},
		)
	}
}

func TestGenerateVerifyEmail(tPtr *testing.T) {

	type arguments struct {
		firstName string
		lastName  string
		email     string
		shortURL  string
	}

	var (
		err                error
		gotError           bool
		errorInfo          coreError.ErrorInfo
		tEmailServer       *EmailServer
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTemplateId        = ""
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				firstName: constants.TEST_USER_FIRST_NAME,
				lastName:  constants.TEST_USER_LAST_NAME,
				email:     constants.TEST_USER_EMAIL,
				shortURL:  constants.TEST_URL_VALID,
			},
			wantError: false,
		},
	}

	tEmailServer, _ = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_SENDER_NAME, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_SENDGRID_KEY_FILE)

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = GenerateVerifyEmail(tEmailServer, tTemplateId, ts.arguments.firstName, ts.arguments.lastName, ts.arguments.email, ts.arguments.shortURL, true); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, err.Error())
				}
			},
		)
	}
}

func TestAddRecipientList(tPtr *testing.T) {

	type arguments struct {
		myList   []EmailItem
		listType string
	}

	var (
		errorInfo           coreError.ErrorInfo
		gotError            bool
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		tPersonalizationPtr = mail.NewPersonalization()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Good To addresses!.",
			arguments: arguments{
				myList: []EmailItem{
					{
						Name:    "Scott",
						Address: "test@example.com",
					}, {
						Name:    "Scott 2",
						Address: "test_2@example.com",
					},
				},
				listType: coreSendGrid.RECIPIENT_TO,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Good CC addresses!.",
			arguments: arguments{
				myList: []EmailItem{
					{
						Name:    "Scott",
						Address: "test@example.com",
					}, {
						Name:    "Scott 2",
						Address: "test_2@example.com",
					},
				},
				listType: coreSendGrid.RECIPIENT_CC,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Good BCC addresses!.",
			arguments: arguments{
				myList: []EmailItem{
					{
						Name:    "Scott",
						Address: "test@example.com",
					}, {
						Name:    "Scott 2",
						Address: "test_2@example.com",
					},
				},
				listType: coreSendGrid.RECIPIENT_BCC,
			},
			wantError: false,
		},
		{
			name: "Negative Case: To address is empty",
			arguments: arguments{
				myList: []EmailItem{
					{
						Name:    "Scott",
						Address: constants.EMPTY,
					},
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				if errorInfo = addRecipientList(tPersonalizationPtr, ts.arguments.myList, ts.arguments.listType); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}

func TestEmailAddress(tPtr *testing.T) {

	type arguments struct {
		address string
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Good address!.",
			arguments: arguments{
				address: "example@sty-holdings.com",
			},
			wantError: false,
		},
		{
			name: "Negative Case: Address is empty",
			arguments: arguments{
				address: constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Too short",
			arguments: arguments{
				address: "X",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Too long",
			arguments: arguments{
				address: "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Invalid format",
			arguments: arguments{
				address: "123456789012345678901234567890",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Invalid domain",
			arguments: arguments{
				address: "example@123123456456789789.com",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				if errorInfo = validateEmailAddress(ts.arguments.address); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}

func TestValidateSubject(tPtr *testing.T) {

	type arguments struct {
		subject string
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Good subject!.",
			arguments: arguments{
				subject: "This is a test",
			},
			wantError: false,
		},
		{
			name: "Negative Case: Subject is empty",
			arguments: arguments{
				subject: constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Too short",
			arguments: arguments{
				subject: "X",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Too long",
			arguments: arguments{
				subject: "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			tFunctionName, func(t *testing.T) {
				if errorInfo = validateSubject(ts.arguments.subject); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}
