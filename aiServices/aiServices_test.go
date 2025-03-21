package sharedServices

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/vertexai/genai"
	"github.com/nats-io/nats.go"
	"google.golang.org/api/option"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	fbs "github.com/sty-holdings/sharedServices/v2025/firebaseServices"
	ns "github.com/sty-holdings/sharedServices/v2025/natsServices"
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

func buildTestaiInstance() (errorInfo errs.ErrorInfo) {

	var (
		tUniqueSettingsData []byte
	)

	if tUniqueSettingsData, errorInfo.Error = os.ReadFile(UNIQUE_SETTING_FQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_FILENAME, UNIQUE_SETTING_FQN))
		return
	}

	halInstancePtr = &halInstance{
		extensionName:                      ctv.EXTENSION_ANALYZE_QUESTION,
		identityProvider:                   ctv.IDP_FIREBASE,
		identityProviderCredentialFilename: GCP_CREDENTIALS_FQN,
		subscriptionPtrs:                   make(map[string]*nats.Subscription),
		testingOn:                          true,
		waitGroup:                          sync.WaitGroup{},
	}

	if errorInfo.Error = json.Unmarshal(tUniqueSettingsData, &halInstancePtr.extensionUniqueSettings); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_UNIQUE_SETTINGS, ctv.TXT_MARSHAL_FAILED))
		return
	}

	if errorInfo = validateUniqueSettings(halInstancePtr.extensionUniqueSettings.UniqueSettings); errorInfo.Error != nil {
		return
	}

	if halInstancePtr.natsServicePtr, errorInfo = ns.NewNATSService(ctv.EXTENSION_ANALYZE_QUESTION, NATSConfig); errorInfo.Error != nil {
		log.Fatalln("NATS connection failed")
	}
	if errorInfo = getFirebaseConnection(); errorInfo.Error != nil {
		log.Fatalln("FB connection failed")
	}
	if errorInfo = getaiConnection(); errorInfo.Error != nil {
		log.Fatalln("ai connection failed")
	}
	if errorInfo = getFBUser(); errorInfo.Error != nil {
		log.Fatalln("FB Auth User retrieval failed")
	}

	return
}

func getFirebaseConnection() (errorInfo errs.ErrorInfo) {

	if halInstancePtr.firebaseAppPtr, halInstancePtr.firebaseAuthPtr, errorInfo = fbs.GetFirebaseAppAuthConnection(GCP_CREDENTIALS_FQN); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
		return
	}

	if halInstancePtr.firestoreClientPtr, errorInfo = fbs.GetFirestoreClientConnection(halInstancePtr.firebaseAppPtr); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
		return
	}

	return
}

func getaiConnection() (errorInfo errs.ErrorInfo) {

	var (
		tUniqueSettingsData []byte
	)

	if tUniqueSettingsData, errorInfo.Error = os.ReadFile(UNIQUE_SETTING_FQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_FILENAME, UNIQUE_SETTING_FQN))
		return
	}

	if errorInfo.Error = json.Unmarshal(tUniqueSettingsData, &halInstancePtr.extensionUniqueSettings); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_UNIQUE_SETTINGS, ctv.TXT_MARSHAL_FAILED))
		return
	}

	if errorInfo = validateUniqueSettings(halInstancePtr.extensionUniqueSettings.UniqueSettings); errorInfo.Error != nil {
		return
	}

	if halInstancePtr.aiConnectionPtr, errorInfo.Error = genai.NewClient(
		context.Background(), halInstancePtr.extensionUniqueSettings.UniqueSettings.GCPProjectId, halInstancePtr.extensionUniqueSettings.UniqueSettings.GCPLocation,
		option.WithCredentialsFile(halInstancePtr.extensionUniqueSettings.UniqueSettings.GCPCredentialFilename),
	); errorInfo.Error != nil {
		errs.PrintErrorInfo(errorInfo)
		return
	}

	return
}
