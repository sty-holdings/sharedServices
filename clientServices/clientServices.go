package sharedServices

import (
	"encoding/json"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
)

// GetClientStruct parses userInfo map to populate and return an STYHClient struct and an errorInfo object if any issue arises during the process.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetClientStruct(userInfo map[string]interface{}) (clientStruct STYHClient, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	if value, ok = userInfo[ctv.FN_COMPANY_NAME]; ok {
		clientStruct.CompanyName = value.(string)
	}

	if value, ok = userInfo[ctv.FN_EMAIL]; ok {
		clientStruct.Email = value.(string)
	}

	if value, ok = userInfo[ctv.FN_FIRST_NAME]; ok {
		clientStruct.FirstName = value.(string)
	}

	if value, ok = userInfo[ctv.FN_GOOGLE_ADS_ACCOUNTS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_GOOGLE_ADS_ACCOUNTS, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.GoogleAdsAccounts); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_GOOGLE_ADS_ACCOUNTS, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = userInfo[ctv.FN_LAST_NAME]; ok {
		clientStruct.LastName = value.(string)
	}

	if value, ok = userInfo[ctv.FN_ON_BOARDED]; ok {
		clientStruct.OnBoarded = value.(bool)
	}

	if value, ok = userInfo[ctv.FN_SAAS_PROVIDERS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_PROVIDER, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.SaasProviders); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_PROVIDER, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = userInfo[ctv.FN_STRIPE_KEY]; ok {
		clientStruct.StripeKey = value.(string)
	}

	if value, ok = userInfo[ctv.FN_STYH_CLIENT_ID]; ok {
		clientStruct.STYHClientId = value.(string)
	}

	if value, ok = userInfo[ctv.FN_TIMEZONE]; ok {
		clientStruct.Timezone = value.(string)
		if clientStruct.LocationPtr, errorInfo.Error = time.LoadLocation(clientStruct.Timezone); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_TIMEZONE, clientStruct.Timezone))
			return
		}
	}

	if value, ok = userInfo[ctv.FN_STYH_USER_ID]; ok {
		clientStruct.STYHUserId = value.(string)
	}

	return
}

// GetClientStructUsingFirebase - retrieves user's client details from Firebase Auth and Firestore by styhUserId, populating an STYHClient struct or returning an error if any issues occur.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetClientUsingFirebase(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, styhUserId string) (clientStruct STYHClient, errorInfo errs.ErrorInfo) {

	var (
		tUserInfo map[string]interface{}
	)

	if tUserInfo, errorInfo = fbs.GetFirebaseUserInfo(
		firebaseAuthPtr,
		firestoreClientPtr,
		styhUserId,
	); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildSTYHUserIdLabelValue(ctv.LBL_SERVICE_CLIENT, styhUserId, ctv.LBL_FIREBASE_AUTH, ctv.TXT_FAILED))
		return
	}

	clientStruct, errorInfo = GetClientStruct(tUserInfo)

	return
}

// ProcessConfigureNewUser - add a new user to the users datastore
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ProcessConfigureNewUser(firestoreClientPtr *firestore.Client, newUser NewUser) {

	var (
		errorInfo errs.ErrorInfo
		tUserInfo = make(map[any]interface{})
	)

	tUserInfo[ctv.FN_COMPANY_NAME] = newUser.CompanyName
	tUserInfo[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	tUserInfo[ctv.FN_EMAIL] = newUser.Email
	tUserInfo[ctv.FN_FIRST_NAME] = newUser.FirstName
	tUserInfo[ctv.FN_GOOGLE_ADS_ACCOUNTS] = []string{} // initialize Google Ads Accounts field
	tUserInfo[ctv.FN_LAST_NAME] = newUser.LastName
	tUserInfo[ctv.FN_SAAS_PROVIDERS] = []string{} // initialize SaaS Providers field
	tUserInfo[ctv.FN_STRIPE_KEY] = ctv.VAL_EMPTY  // initialize Stripe Key field
	tUserInfo[ctv.FN_STYH_CLIENT_ID] = hlp.GenerateUUIDType1(false)
	tUserInfo[ctv.FN_TIMEZONE] = newUser.Timezone
	tUserInfo[ctv.FN_STYH_USER_ID] = newUser.STYHUserId
	tUserInfo[ctv.FN_ON_BOARDED] = false

	if errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_USERS, newUser.STYHUserId, tUserInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}
