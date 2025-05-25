package sharedServices

import (
	"encoding/json"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// GetClientUserStruct - retrieves and constructs both client/user structures from the provided client/user information.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo
//	Verifications: None
func GetClientUserStruct(clientInfo map[string]interface{}, userInfo map[string]interface{}) (clientUserStruct STYHClientUser, errorInfo errs.ErrorInfo) {

	if clientUserStruct.MySTYHClient, errorInfo = GetClientStruct(clientInfo); errorInfo.Error != nil {
		return
	}
	clientUserStruct.MySTYHUser, errorInfo = GetUserStruct(userInfo)

	return
}

// GetClientStruct - constructs and returns a populated STYHClient struct using data from the provided clientInfo map.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo for JSON marshal/unmarshal, time location loading, and other data-processing issues.
//	Verifications: None
func GetClientStruct(clientInfo map[string]interface{}) (clientStruct STYHClient, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	if value, ok = clientInfo[ctv.FN_COMPANY_NAME]; ok {
		clientStruct.CompanyName = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_FORMATION_TYPE]; ok {
		clientStruct.FormationType = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_GOOGLE_ADS_ACCOUNTS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_GOOGLE_ADS_ACCOUNTS, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.GoogleAdsAccounts); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_GOOGLE_ADS_ACCOUNTS, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = clientInfo[ctv.FN_LINKEDIN_PAGE_IDS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_LINKEDIN_PAGE_IDS, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.LinkedinPageIds); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_LINKEDIN_PAGE_IDS, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	clientStruct.OnBoarded = false // Default unless reset by the clientInfo[ctv.FN_SAAS_CLIENT_PROVIDERS] code below. DO NOT REMOVE

	if value, ok = clientInfo[ctv.FN_OWNERS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_OWNERS, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.Owners); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_OWNERS, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = clientInfo[ctv.FN_PAYPAL_CLIENT_ID]; ok {
		clientStruct.PayPalClientID = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_PAYPAL_CLIENT_SECRET]; ok {
		clientStruct.PayPalClientSecret = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_PHONE_COUNTRY_CODE]; ok {
		clientStruct.PhoneCountryCode = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_PHONE_AREA_CODE]; ok {
		clientStruct.PhoneAreaCode = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_PHONE_NUMBER]; ok {
		clientStruct.PhoneNumber = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_SAAS_CLIENT_PROVIDERS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_CLIENT_PROVIDER, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.SaaSClientProviders); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_CLIENT_PROVIDER, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
		if len(clientStruct.SaaSClientProviders) > ctv.VAL_ZERO {
			clientStruct.OnBoarded = true
		}
	}

	if value, ok = clientInfo[ctv.FN_STRIPE_CLIENT_CONNECT_ACCOUNT_ID]; ok {
		clientStruct.StripeClientConnectAccountId = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_STRIPE_CLIENT_REFRESH_TOKEN]; ok {
		clientStruct.StripeClientRefreshToken = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_STRIPE_INITIAL_PULL_DATA_STATUS]; ok {
		clientStruct.StripeInitialPullDataStatus = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_STRIPE_PULL_FREQUENCY]; ok {
		clientStruct.StripePullFrequency = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_STRIPE_START_DATE]; ok {
		clientStruct.StripeStartDate = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_STYH_INTERNAL_CLIENT_ID]; ok {
		clientStruct.STYHInternalClientID = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_TIMEZONE_HQ]; ok {
		clientStruct.TimezoneHQ = value.(string)
		if clientStruct.TimezoneHQLocationPtr, errorInfo.Error = time.LoadLocation(clientStruct.TimezoneHQ); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_TIMEZONE, clientStruct.TimezoneHQ))
			return
		}
	}

	if value, ok = clientInfo[ctv.FN_WEBSITE_URL]; ok {
		clientStruct.WebsiteURL = value.(string)
	}

	return
}

// GetUserStruct - constructs and returns a populated STYHUser struct using data from the provided userInfo map.
//
//	Customer Messages: None
//	Errors: errs.NewErrorInfo
//	Verifications: None
func GetUserStruct(userInfo map[string]interface{}) (userStruct STYHUser, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	if value, ok = userInfo[ctv.FN_APPROVED_BY]; ok {
		userStruct.ApprovedBy = value.(string)
	}

	if value, ok = userInfo[ctv.FN_APPROVED_BY_DATE]; ok {
		userStruct.ApprovedByDate = value.(string)
	}

	if value, ok = userInfo[ctv.FN_EMAIL]; ok {
		userStruct.ApprovedByDate = value.(string)
	}

	if value, ok = userInfo[ctv.FN_FIRST_NAME]; ok {
		userStruct.FirstName = value.(string)
	}

	if value, ok = userInfo[ctv.FN_LAST_NAME]; ok {
		userStruct.LastName = value.(string)
	}

	if value, ok = userInfo[ctv.FN_PERMISSIONS]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_CLIENT_PROVIDER, ctv.TXT_MARSHAL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &userStruct.Permissions); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_SAAS_CLIENT_PROVIDER, ctv.TXT_UNMARSHAL_FAILED))
			return
		}
	}

	if value, ok = userInfo[ctv.FN_STYH_INTERNAL_CLIENT_ID]; ok {
		userStruct.STYHInternalClientID = value.(string)
	}

	if value, ok = userInfo[ctv.FN_STYH_INTERNAL_USER_ID]; ok {
		userStruct.STYHInternalUserID = value.(string)
	}

	if value, ok = userInfo[ctv.FN_TIMEZONE_USER]; ok {
		userStruct.TimezoneUser = value.(string)
		if userStruct.TimezoneUserLocationPtr, errorInfo.Error = time.LoadLocation(userStruct.TimezoneUser); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_CLIENT, ctv.LBL_TIMEZONE_USER, userStruct.TimezoneUser))
			return
		}
	}

	return
}

// GetClientInfo - retrieves client details from Clients datastore, populating an STYHClient struct or returning an error if any issues occur.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetClientInfo(firestoreClientPtr *firestore.Client, styhInternalClientID string) (clientInfo STYHClient, errorInfo errs.ErrorInfo) {

	var (
		tDocumentSnapshotPtr *firestore.DocumentSnapshot
	)

	if tDocumentSnapshotPtr, errorInfo = fbs.GetDocumentById(firestoreClientPtr, fbs.DATASTORE_CLIENTS, styhInternalClientID); errorInfo.Error != nil {
		return
	}

	clientInfo, errorInfo = GetClientStruct(tDocumentSnapshotPtr.Data())

	return
}

// GetClientUserInfo - retrieves combined client and user details based on the provided STYHInternalUserID, using Firebase and Firestore services.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo
//	Verifications: None
func GetClientUserInfo(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, styhInternalUserID string) (clientUserInfo STYHClientUser, errorInfo errs.ErrorInfo) {

	if clientUserInfo.MySTYHUser, errorInfo = GetUserInfo(firebaseAuthPtr, firestoreClientPtr, styhInternalUserID); errorInfo.Error != nil {
		return
	}

	clientUserInfo.MySTYHClient, errorInfo = GetClientInfo(firestoreClientPtr, clientUserInfo.MySTYHUser.STYHInternalClientID)

	return
}

// GetUserInfo - retrieves user information from Firebase and Firestore and constructs a STYHUser object.
//
//	Customer Messages: None
//	Errors: errs.NewErrorInfo
//	Verifications: None
func GetUserInfo(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, styhInternalUserID string) (userInfo STYHUser, errorInfo errs.ErrorInfo) {

	var (
		tUserInfo map[string]interface{}
	)

	if tUserInfo, errorInfo = fbs.GetFirebaseUserInfo(
		firebaseAuthPtr,
		firestoreClientPtr,
		styhInternalUserID,
	); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildSTYHInternalUserIDLabelValue(ctv.LBL_SERVICE_CLIENT, styhInternalUserID, ctv.LBL_FIREBASE_AUTH, ctv.TXT_FAILED))
		return
	}

	userInfo, errorInfo = GetUserStruct(tUserInfo)

	return
}

// CheckClientExists - verifies if a client already exists in the datastore based on provided parameters. if so, returns the STYHInternalClientID populated.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrNoFoundDocument, errs.ErrFailedServiceFirestore
//	Verifications: None
func CheckClientExists(firestoreClientPtr *firestore.Client, companyName string, phoneAreaCode string, phoneNumber string, userEmail string, websiteURL string) (styhInternalClientID string) {

	var (
		errorInfo            errs.ErrorInfo
		found                bool
		tConfirmedCount      uint
		tDocumentSnapshotPtr *firestore.DocumentSnapshot
	)

	if errorInfo = hlps.CheckPointerNotNil(ctv.VAL_SERVICE_CLIENT, firestoreClientPtr, ctv.LBL_SERVICE_FIREBASE); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_CLIENT, companyName, ctv.LBL_COMPANY_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_CLIENT, phoneAreaCode, ctv.LBL_PHONE_AREA_CODE); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_CLIENT, phoneNumber, ctv.LBL_PHONE_NUMBER); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_CLIENT, userEmail, ctv.LBL_EMAIL); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(ctv.VAL_SERVICE_CLIENT, websiteURL, ctv.LBL_WEBSITE_URL); errorInfo.Error != nil {
		return
	}

	if vals.DoesWebsiteEmailMatch(userEmail, websiteURL) {
		if found, tDocumentSnapshotPtr, errorInfo = fbs.FindDocument(
			firestoreClientPtr, fbs.DATASTORE_CLIENTS, fbs.NameValueQuery{
				FieldName:  ctv.FN_WEBSITE_URL,
				FieldValue: websiteURL,
			},
		); errorInfo.Error != nil {
			return
		}
		styhInternalClientID = tDocumentSnapshotPtr.Ref.ID
		return
	}

	// Company Name Match
	if found, tDocumentSnapshotPtr, errorInfo = fbs.FindDocument(
		firestoreClientPtr, fbs.DATASTORE_CLIENTS, fbs.NameValueQuery{
			FieldName:  ctv.FN_COMPANY_NAME,
			FieldValue: companyName,
		},
	); errorInfo.Error != nil {
		return
	}
	if found {
		tConfirmedCount++
	}

	// Company Phone Match
	if found, tDocumentSnapshotPtr, errorInfo = fbs.FindDocument(
		firestoreClientPtr, fbs.DATASTORE_CLIENTS, fbs.NameValueQuery{
			FieldName:  ctv.FN_PHONE_AREA_CODE,
			FieldValue: phoneAreaCode,
		},
		fbs.NameValueQuery{
			FieldName:  ctv.FN_PHONE_NUMBER,
			FieldValue: phoneNumber,
		},
	); errorInfo.Error != nil {
		return
	}
	if found {
		tConfirmedCount++
	}

	if tConfirmedCount > ctv.VAL_ZERO {
		styhInternalClientID = tDocumentSnapshotPtr.Ref.ID
	}

	return
}

// ProcessNewClient - processes the creation of a new client in the datastore.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrNoFoundDocument, errs.ErrFailedServiceFirestore
//	Verifications: vals.
func ProcessNewClient(firestoreClientPtr *firestore.Client, newClient NewClient, userEmail string) (styhInternalClientId string, errorInfo errs.ErrorInfo) {

	var (
		tClientStruct = make(map[any]interface{})
	)

	if styhInternalClientId = CheckClientExists(
		firestoreClientPtr,
		newClient.CompanyName,
		newClient.PhoneAreaCode,
		newClient.PhoneNumber,
		userEmail,
		newClient.WebSiteURL,
	); styhInternalClientId == ctv.VAL_EMPTY {
		styhInternalClientId = hlps.GenerateUUIDType1(true)
		//
		tClientStruct[ctv.FN_COMPANY_NAME] = newClient.CompanyName
		tClientStruct[ctv.FN_CREATE_TIMESTAMP] = time.Now()
		tClientStruct[ctv.FN_FORMATION_TYPE] = newClient.FormationType
		tClientStruct[ctv.FN_PHONE_COUNTRY_CODE] = newClient.PhoneCountryCode
		tClientStruct[ctv.FN_PHONE_AREA_CODE] = newClient.PhoneAreaCode
		tClientStruct[ctv.FN_PHONE_NUMBER] = newClient.PhoneNumber
		tClientStruct[ctv.FN_STYH_INTERNAL_CLIENT_ID] = styhInternalClientId
		tClientStruct[ctv.FN_TIMEZONE_HQ] = newClient.TimezoneHQ
		tClientStruct[ctv.FN_WEBSITE_URL] = newClient.WebSiteURL
		errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_CLIENTS, styhInternalClientId, tClientStruct)
	}

	return
}

// ProcessNewUser - configures and saves a new user record in the Users datastore.
//
//	Customer Messages: None
//	Errors: errs.Err if Firestore document creation fails.
//	Verifications: vlds.AreMapKeysPopulated validates map keys' presence.
func ProcessNewUser(firestoreClientPtr *firestore.Client, newUser NewUser) {

	var (
		errorInfo errs.ErrorInfo
		tUserInfo = make(map[any]interface{})
	)

	tUserInfo[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	tUserInfo[ctv.FN_EMAIL] = newUser.Email
	tUserInfo[ctv.FN_FIRST_NAME] = newUser.FirstName
	tUserInfo[ctv.FN_LAST_NAME] = newUser.LastName
	tUserInfo[ctv.FN_TIMEZONE_USER] = newUser.TimezoneUser
	tUserInfo[ctv.FN_STYH_INTERNAL_USER_ID] = newUser.STYHInternalUserID
	tUserInfo[ctv.FN_STYH_INTERNAL_CLIENT_ID] = newUser.STYHInternalClientID

	if errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_USERS, newUser.STYHInternalUserID, tUserInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}

// ProcessSaaSProviderList - builds saas_client_providers list for clients record in Firestore.
//
//	Customer Messages: None
//	Errors: errs.Err if Firestore document creation fails.
//	Verifications: vlds.AreMapKeysPopulated validates map keys' presence.
func ProcessSaaSProviderList(firestoreClientPtr *firestore.Client, styhInternalClientID string, saasClientProviders map[string]bool) {

	var (
		errorInfo            errs.ErrorInfo
		tClientInfo          = make(map[any]interface{})
		tSaasClientProviders []string
	)

	for provider, checked := range saasClientProviders {
		if checked {
			tSaasClientProviders = append(tSaasClientProviders, provider)
		}
	}
	tClientInfo[ctv.FN_SAAS_CLIENT_PROVIDERS] = tSaasClientProviders

	if errorInfo = fbs.UpdateDocument(firestoreClientPtr, fbs.DATASTORE_USERS, styhInternalClientID, tClientInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}
