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

// GetClientStruct - constructs an STYHClient instance from the provided map.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo
//	Verifications: None
func GetClientStruct(userInfo map[string]interface{}) (clientStruct STYHClient, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	if value, ok = userInfo[ctv.FN_COMPANY_NAME]; ok {
		clientStruct.CompanyName = value.(string)
		clientStruct.AccountType = ctv.VAL_BUSINESS
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

	clientStruct.OnBoarded = false // Default unless reset by the userInfo[ctv.FN_SAAS_PROVIDERS] code below. DO NOT REMOVE
	if value, ok = userInfo[ctv.FN_LINKEDIN_PAGE_IDS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.FN_LINKEDIN_PAGE_IDS, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.LinkedinPageIdList); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.FN_LINKEDIN_PAGE_IDS, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = userInfo[ctv.FN_PAYPAL_CLIENT_ID]; ok {
		clientStruct.PayPalClientId = value.(string)
	}
	if value, ok = userInfo[ctv.FN_PAYPAL_CLIENT_SECRET]; ok {
		clientStruct.PayPalClientSecret = value.(string)
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
		if len(clientStruct.SaasProviders) > ctv.VAL_ZERO {
			clientStruct.OnBoarded = true
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

// ProcessConfigureNewUser - configures and saves a new user record in Firestore.
//
//	Customer Messages: None
//	Errors: errs.Err if Firestore document creation fails.
//	Verifications: vlds.AreMapKeysPopulated validates map keys' presence.
func ProcessConfigureNewUser(firestoreClientPtr *firestore.Client, newUser NewUser) {

	var (
		errorInfo errs.ErrorInfo
		tUserInfo = make(map[any]interface{})
	)

	if newUser.CompanyName == ctv.VAL_EMPTY {
		tUserInfo[ctv.FN_ACCOUNT_TYPE] = ctv.VAL_INDIVIDUAL
	} else {
		tUserInfo[ctv.FN_ACCOUNT_TYPE] = ctv.VAL_BUSINESS
		tUserInfo[ctv.FN_COMPANY_NAME] = newUser.CompanyName
	}
	tUserInfo[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	tUserInfo[ctv.FN_EMAIL] = newUser.Email
	tUserInfo[ctv.FN_FIRST_NAME] = newUser.FirstName
	tUserInfo[ctv.FN_LAST_NAME] = newUser.LastName
	tUserInfo[ctv.FN_STYH_CLIENT_ID] = hlp.GenerateUUIDType1(false)
	tUserInfo[ctv.FN_TIMEZONE] = newUser.Timezone
	tUserInfo[ctv.FN_STYH_USER_ID] = newUser.STYHUserId

	if errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_USERS, newUser.STYHUserId, tUserInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}

// ProcessSaaSProviderList - builds saas_provers list for users record in Firestore.
//
//	Customer Messages: None
//	Errors: errs.Err if Firestore document creation fails.
//	Verifications: vlds.AreMapKeysPopulated validates map keys' presence.
func ProcessSaaSProviderList(firestoreClientPtr *firestore.Client, styhClientId string, styhUserId string, saasProviders map[string]bool) {

	var (
		errorInfo      errs.ErrorInfo
		tUserInfo      = make(map[any]interface{})
		tSaasProviders []string
	)

	for provider, checked := range saasProviders {
		if checked {
			tSaasProviders = append(tSaasProviders, provider)
		}
	}
	tUserInfo[ctv.FN_SAAS_PROVIDERS] = tSaasProviders

	if errorInfo = fbs.UpdateDocument(firestoreClientPtr, fbs.DATASTORE_USERS, styhUserId, tUserInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}
