package sharedServices

import (
	"encoding/json"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
)

// GetClientStruct - will retrieve the client record and return a struct.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetClientStruct(firebaseAuthPtr *auth.Client, firestoreClientPtr *firestore.Client, uId string, testMode bool) (clientStruct STYHClient, errorInfo errs.ErrorInfo) {

	var (
		jsonData  []byte
		ok        bool
		tUserInfo map[string]interface{}
		value     interface{}
	)

	if tUserInfo, errorInfo = fbs.GetFirebaseUserInfo(
		firebaseAuthPtr,
		firestoreClientPtr,
		uId,
	); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdLabelValue(uId, ctv.LBL_FIREBASE_AUTH, ctv.TXT_FAILED))
		return
	}

	if value, ok = tUserInfo[ctv.FN_COMPANY_NAME]; ok {
		clientStruct.CompanyName = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_EMAIL]; ok {
		clientStruct.Email = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_FIRST_NAME]; ok {
		clientStruct.FirstName = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_LAST_NAME]; ok {
		clientStruct.LastName = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_ON_BOARDED]; ok {
		clientStruct.OnBoarded = value.(bool)
	}

	if value, ok = tUserInfo[ctv.FN_SAAS_PROFILE]; ok {
		if jsonData, errorInfo.Error = json.Marshal(value); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SAAS_PROFILE, ctv.TXT_MARSHALL_FAILED))
			return
		}
		if errorInfo.Error = json.Unmarshal(jsonData, &clientStruct.SaasProfile); errorInfo.Error != nil {
			errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SAAS_PROFILE, ctv.TXT_UNMARSHALL_FAILED))
			return
		}
	}

	if value, ok = tUserInfo[ctv.FN_STYH_CLIENT_ID]; ok {
		clientStruct.StyhClientId = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_TIMEZONE]; ok {
		clientStruct.Timezone = value.(string)
	}

	if value, ok = tUserInfo[ctv.FN_UID]; ok {
		clientStruct.Uid = value.(string)
	}

	return
}

// SaaSProfilePopulated - determines is a SaaS provider exists.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func SaaSProfilePopulated(clientStruct STYHClient) bool {

	if len(clientStruct.SaasProfile.UserSaaSProviders) > ctv.VAL_ZERO {
		return true
	}

	return false
}
