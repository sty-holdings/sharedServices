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

func FindClientByDomain(firestoreClientPtr *firestore.Client, domain string) (clientInfo InternalClient) {

	var (
		errorInfo    errs.ErrorInfo
		found        bool
		tDocumentPtr *firestore.DocumentSnapshot
	)

	if found, tDocumentPtr, errorInfo = fbs.FindDocument(
		firestoreClientPtr, fbs.DATASTORE_CLIENTS, fbs.NameValueQuery{
			FieldName:  ctv.FN_DOMAIN,
			FieldValue: domain,
		},
	); found {
		if clientInfo, errorInfo = getClientStruct(tDocumentPtr.Data(), tDocumentPtr.Ref.ID); errorInfo.Error != nil {
			errs.PrintErrorInfo(errorInfo)
		}
	}

	return
}

// GetAllClientsInfo - retrieves and constructs a list of all client details from the datastore.
//
//	Customer Messages: None
//	Errors: errs.ErrNoDataFound
//	Verifications: None
func GetAllClientsInfo(firestoreClientPtr *firestore.Client) (internalClientsInfo []InternalClient, errorInfo errs.ErrorInfo) {

	var (
		tDocumentSnapshotsPtrs []*firestore.DocumentSnapshot
		tInternalClient        InternalClient
	)

	if tDocumentSnapshotsPtrs, errorInfo = fbs.GetAllDocuments(firestoreClientPtr, fbs.DATASTORE_CLIENTS); errorInfo.Error != nil {
		return
	}
	if len(tDocumentSnapshotsPtrs) == 0 {
		errorInfo = errs.NewErrorInfo(errs.ErrNoDataFound, errs.BuildLabelValueMessage(ctv.VAL_SERVICE_CLIENT, ctv.LBL_DOCUMENT_ID, ctv.TXT_ARE_EMPTY, ctv.TXT_NO_DATA_FOUND))
		return
	}

	for _, document := range tDocumentSnapshotsPtrs {
		if tInternalClient, errorInfo = getClientStruct(document.Data(), document.Ref.ID); errorInfo.Error != nil {
			return
		}
		internalClientsInfo = append(internalClientsInfo, tInternalClient)
	}

	return
}

// GetAllUsersInfo - retrieves and constructs a list of all user details from the datastore.
//
//	Customer Messages: None
//	Errors: errs.ErrNoDataFound
//	Verifications: None
func GetAllUsersInfo(firestoreClientPtr *firestore.Client) (internalUsersInfo []InternalUser, errorInfo errs.ErrorInfo) {

	var (
		tDocumentSnapshotsPtrs []*firestore.DocumentSnapshot
		tInternalUser          InternalUser
	)

	if tDocumentSnapshotsPtrs, errorInfo = fbs.GetAllDocuments(firestoreClientPtr, fbs.DATASTORE_USERS); errorInfo.Error != nil {
		return
	}
	if len(tDocumentSnapshotsPtrs) == 0 {
		errorInfo = errs.NewErrorInfo(errs.ErrNoDataFound, errs.BuildLabelValueMessage(ctv.VAL_SERVICE_CLIENT, ctv.LBL_DOCUMENT_ID, ctv.TXT_ARE_EMPTY, ctv.TXT_NO_DATA_FOUND))
		return
	}

	for _, document := range tDocumentSnapshotsPtrs {
		if tInternalUser, errorInfo = getUserStruct(document.Data(), document.Ref.ID); errorInfo.Error != nil {
			return
		}
		internalUsersInfo = append(internalUsersInfo, tInternalUser)
	}

	return
}

// GetClientInfo - retrieves client details from Clients datastore, populating an InternalClient struct or returning an error if any issues occur.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetClientInfo(firestoreClientPtr *firestore.Client, internalClientID string) (clientInfo InternalClient, errorInfo errs.ErrorInfo) {

	var (
		tDocumentSnapshotPtr *firestore.DocumentSnapshot
	)

	if tDocumentSnapshotPtr, errorInfo = fbs.GetDocumentById(firestoreClientPtr, fbs.DATASTORE_CLIENTS, internalClientID); errorInfo.Error != nil {
		return
	}

	clientInfo, errorInfo = getClientStruct(tDocumentSnapshotPtr.Data(), tDocumentSnapshotPtr.Ref.ID)

	return
}

// GetClientUserInfo - retrieves combined client and user details based on the provided InternalUserID, using Firebase and Firestore services.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo
//	Verifications: None
func GetClientUserInfo(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, internalUserID string) (clientUserInfo InternalClientUser, errorInfo errs.ErrorInfo) {

	if clientUserInfo.MyInternalUser, errorInfo = GetUserInfo(firebaseAuthPtr, firestoreClientPtr, internalUserID); errorInfo.Error != nil {
		return
	}

	clientUserInfo.MyInternalClient, errorInfo = GetClientInfo(firestoreClientPtr, clientUserInfo.MyInternalUser.InternalClientID)

	return
}

// GetUserInfo - retrieves user information from Firebase and Firestore and constructs a InternalUser object.
//
//	Customer Messages: None
//	Errors: errs.NewErrorInfo
//	Verifications: None
func GetUserInfo(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, internalUserID string) (userInfo InternalUser, errorInfo errs.ErrorInfo) {

	var (
		tUserInfo  map[string]interface{}
		tUserRefID string
	)

	if tUserInfo, tUserRefID, errorInfo = fbs.GetFirebaseUserInfo(
		firebaseAuthPtr,
		firestoreClientPtr,
		internalUserID,
	); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildInternalUserIDLabelValue(ctv.LBL_SERVICE_CLIENT, internalUserID, ctv.LBL_FIREBASE_AUTH, ctv.TXT_FAILED))
		return
	}

	userInfo, errorInfo = getUserStruct(tUserInfo, tUserRefID)

	return
}

// CheckClientExists - verifies if a client already exists in the datastore based on provided parameters. if so, returns the InternalClientID populated.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrNoFoundDocument, errs.ErrFailedServiceFirestore
//	Verifications: None
func CheckClientExists(firestoreClientPtr *firestore.Client, companyName string, phoneAreaCode string, phoneNumber string, userEmail string, websiteURL string) (internalClientID string) {

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
		internalClientID = tDocumentSnapshotPtr.Ref.ID
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
		internalClientID = tDocumentSnapshotPtr.Ref.ID
	}

	return
}

// ProcessNewClient - creates a new client record in the datastore.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrFailedServiceFirebase
//	Verifications: vals.
func ProcessNewClient(firestoreClientPtr *firestore.Client, checkExists bool, newClient NewClient, userEmail string) (internalClientId string, errorInfo errs.ErrorInfo) {

	var (
		tClientStruct = make(map[any]interface{})
	)

	if checkExists {
		internalClientId = CheckClientExists(
			firestoreClientPtr,
			newClient.CompanyName,
			newClient.PhoneAreaCode,
			newClient.PhoneNumber,
			userEmail,
			newClient.WebSiteURL,
		)
		if internalClientId != ctv.VAL_EMPTY {
			return
		}
	}

	internalClientId = hlps.GenerateUUIDType1(true)
	//
	tClientStruct[ctv.FN_COMPANY_NAME] = newClient.CompanyName
	tClientStruct[ctv.FN_DOMAIN] = newClient.Domain
	tClientStruct[ctv.FN_CREATE_TIMESTAMP] = time.Now()
	tClientStruct[ctv.FN_FORMATION_TYPE] = newClient.FormationType
	tClientStruct[ctv.FN_PHONE_COUNTRY_CODE] = newClient.PhoneCountryCode
	tClientStruct[ctv.FN_PHONE_AREA_CODE] = newClient.PhoneAreaCode
	tClientStruct[ctv.FN_PHONE_NUMBER] = newClient.PhoneNumber
	tClientStruct[ctv.FN_INTERNAL_CLIENT_ID] = internalClientId
	tClientStruct[ctv.FN_TIMEZONE_HQ] = newClient.TimezoneHQ
	tClientStruct[ctv.FN_WEBSITE_URL] = newClient.WebSiteURL
	errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_CLIENTS, internalClientId, tClientStruct)

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
	tUserInfo[ctv.FN_INTERNAL_USER_ID] = newUser.InternalUserID
	tUserInfo[ctv.FN_INTERNAL_CLIENT_ID] = newUser.InternalClientID

	if errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_USERS, newUser.InternalUserID, tUserInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}

// ProcessSaaSProviderList - builds saas_client_providers list for clients record in Firestore.
//
//	Customer Messages: None
//	Errors: errs.Err if Firestore document creation fails.
//	Verifications: vlds.AreMapKeysPopulated validates map keys' presence.
func ProcessSaaSProviderList(firestoreClientPtr *firestore.Client, internalClientID string, saasClientProviders map[string]bool) {

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

	if errorInfo = fbs.UpdateDocument(firestoreClientPtr, fbs.DATASTORE_CLIENTS, internalClientID, tClientInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}

func RedeemAccessCode(firestoreClientPtr *firestore.Client, internalClientID string, accessCode string) (errorInfo errs.ErrorInfo) {

	// var (
	// 	tDocumentPtr *firestore.DocumentSnapshot
	// )
	//
	// if tDocumentPtr, errorInfo = fbs.GetDocumentById(firestoreClientPtr, fbs.DATASTORE_ACCESS_CODE, accessCode); errorInfo.Error != nil {
	// 	errs.PrintErrorInfo(errorInfo)
	// 	return
	// }
	//
	// if errorInfo = fbs.RemoveDocument(
	// 	firestoreClientPtr, fbs.DATASTORE_ACCESS_CODE, fbs.NameValueQuery{
	// 		FieldName:  ctv.FN_ACCESS_CODE,
	// 		FieldValue: accessCode,
	// 	},
	// ); errorInfo.Error != nil {
	// 	errs.PrintErrorInfo(errorInfo)
	// }

	return
}

func SetAccessCode(firestoreClientPtr *firestore.Client, internalClientID string) (accessCode string) {

	var (
		errorInfo             errs.ErrorInfo
		tAccessCodes          []string
		tAccessCodeClientInfo = make(map[any]interface{})
	)

	tAccessCodes = append(tAccessCodes, hlps.GenerateUUIDType1(true))
	tAccessCodeClientInfo[ctv.FN_INTERNAL_CLIENT_ID] = internalClientID

	if errorInfo = fbs.SetDocument(firestoreClientPtr, fbs.DATASTORE_ACCESS_CODE, hlps.GenerateUUIDType1(true), tAccessCodeClientInfo); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
	}

	return
}

// Private methods below here

// GetClientStruct - constructs and returns a populated InternalClient struct using data from the provided clientInfo map.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo for JSON marshal/unmarshal, time location loading, and other data-processing issues.
//	Verifications: None
func getClientStruct(clientInfo map[string]interface{}, clientInfoRefID string) (clientStruct InternalClient, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	clientStruct.InternalClientID = clientInfoRefID

	if value, ok = clientInfo[ctv.FN_COMPANY_NAME]; ok {
		clientStruct.CompanyName = value.(string)
	}

	if value, ok = clientInfo[ctv.FN_DEMO_ACCOUNT]; ok {
		clientStruct.DemoAccount = value.(bool)
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

// GetUserStruct - constructs and returns a populated InternalUser struct using data from the provided userInfo map.
//
//	Customer Messages: None
//	Errors: errs.NewErrorInfo
//	Verifications: None
func getUserStruct(userInfo map[string]interface{}, userInfoRefID string) (userStruct InternalUser, errorInfo errs.ErrorInfo) {

	var (
		jsonData []byte
		ok       bool
		value    interface{}
	)

	userStruct.InternalUserID = userInfoRefID

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

	if value, ok = userInfo[ctv.FN_POSTAL_CODE]; ok {
		userStruct.PostalCode = value.(string)
	}

	if value, ok = userInfo[ctv.FN_INTERNAL_CLIENT_ID]; ok {
		userStruct.InternalClientID = value.(string)
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
