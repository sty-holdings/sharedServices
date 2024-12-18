package sharedServices

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
	hlp "github.com/sty-holdings/sharedServices/v2024/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2024/jwtServices"
)

// BuildInstanceName - will create the NATS connection name with dashes, underscores between nodes or as provided.
// The method can be natsServices.METHOD_DASHES, natsServices.METHOD_UNDERSCORES, ctv.VAL_EMPTY, "dashes", "underscores" or ""
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func BuildInstanceName(
	method string,
	nodes ...string,
) (
	instanceName string,
	errorInfo errs.ErrorInfo,
) {

	if len(nodes) == 1 {
		method = METHOD_BLANK
	}
	switch strings.Trim(strings.ToLower(method), ctv.SPACES_ONE) {
	case METHOD_DASHES:
		instanceName, errorInfo = buildInstanceName(ctv.DASH, nodes...)
	case METHOD_UNDERSCORES:
		instanceName, errorInfo = buildInstanceName(ctv.UNDERSCORE, nodes...)
	default:
		instanceName, errorInfo = buildInstanceName(ctv.VAL_EMPTY, nodes...)
	}

	return
}

// GetConnection - will connect to a NATS leaf server with either a ssl or non-ssl connection.
// This connection function requires natsServices.NATSConfiguration be populated. The following fields
// do not have to be at this time: TLSCert, TLSPrivateKey, TLSCABundle. The fields TLSCertFQN, TLSPrivateKeyFQN,
// TLSCABundleFQN must be populated.
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func GetConnection(
	instanceName string,
	config NATSConfiguration,
) (
	connPtr *nats.Conn,
	errorInfo errs.ErrorInfo,
) {

	var (
		opts []nats.Option
		tURL string
	)

	opts = []nats.Option{
		nats.Name(instanceName),             // Set a client name
		nats.MaxReconnects(5),               // Set maximum reconnection attempts
		nats.ReconnectWait(5 * time.Second), // Set reconnection wait time
		nats.UserCredentials(config.NATSCredentialsFilename),
		nats.RootCAs(config.NATSTLSInfo.TLSCABundleFQN),
		nats.ClientCert(config.NATSTLSInfo.TLSCertFQN, config.NATSTLSInfo.TLSPrivateKeyFQN),
	}

	if tURL, errorInfo = buildURLPort(config.NATSURL, config.NATSPort); errorInfo.Error != nil {
		return
	}
	if connPtr, errorInfo.Error = nats.Connect(tURL, opts...); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v: %v", instanceName, ctv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	log.Printf("%v: A connection has been established with the NATS server at %v.", instanceName, config.NATSURL)
	log.Printf(
		"%v: URL: %v Server Name: %v Server Id: %v Address: %v",
		instanceName,
		connPtr.ConnectedUrl(),
		connPtr.ConnectedClusterName(),
		connPtr.ConnectedServerId(),
		connPtr.ConnectedAddr(),
	)

	return
}

// RequestWithHeader - will submit a request and wait for a response.
// Min timeOut is 2 seconds and the max is 5 seconds.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func RequestWithHeader(
	connectionPtr *nats.Conn,
	instanceName string,
	messagePtr *nats.Msg,
	timeOut time.Duration,
) (
	responsePtr *nats.Msg,
	errorInfo errs.ErrorInfo,
) {

	if connectionPtr == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrPointerMissing, fmt.Sprintf("%s%s", ctv.LBL_POINTER, ctv.TXT_NATS))
	}
	if timeOut < 2*time.Second {
		timeOut = 2 * time.Second
	}
	if timeOut > 5*time.Second {
		timeOut = 5 * time.Second
	}
	if responsePtr, errorInfo.Error = connectionPtr.RequestMsg(messagePtr, timeOut); errorInfo.Error != nil {
		log.Printf("%v: RequestWithHeader failed on %v %v for %v: %v", instanceName, ctv.LBL_SUBJECT, messagePtr.Subject, ctv.FN_CLIENT_ID, messagePtr.Header.Get(ctv.FN_CLIENT_ID))
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v %v", instanceName, ctv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	responsePtr.Header = messagePtr.Header

	return
}

// EncryptedDataReply - takes a structure, marshals to a []byte, encrypts the data and responds.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func EncryptedDataReply(
	response interface{},
	msg *nats.Msg,
	keyB64 string,
	uId string,
) (errorInfo errs.ErrorInfo) {

	var (
		eMessageB64 []byte
		tJSON       []byte
	)

	if tJSON, errorInfo.Error = json.Marshal(response); errorInfo.Error != nil {
		// todo correct error handling
		return
	}

	if eMessageB64, errorInfo = jwts.EncryptFromByteToByte(uId, keyB64, tJSON); errorInfo.Error != nil {
		// todo add error handling
		return
	}

	if errorInfo.Error = msg.Respond(eMessageB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v%v%v", ctv.LBL_SUBJECT, msg.Subject, ctv.LBL_MESSAGE_HEADER, msg.Header))
	}

	return
}

// EncryptedMessageDataRequest - takes a structure, marshals to a []byte, encrypts the data and makes a request.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func EncryptedMessageDataRequest(
	functionName string,
	natsConnectionPtr *nats.Conn,
	instanceName string,
	keyB64 string,
	msg *nats.Msg,
	request interface{},
	timeOutInSecs int,
	uId string,
	testingOn bool,
) (
	replyMsgPtr *nats.Msg,
	errorInfo errs.ErrorInfo,
) {

	var (
		eMessageDataB64    []byte
		tAdditionalInfo    string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tJSON              []byte
	)

	if functionName == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_FUNCTION_NAME, ctv.TXT_IS_MISSING))
	}
	if natsConnectionPtr == nil && testingOn == false {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_NATS_CONNECT_SERVER, ctv.TXT_IS_MISSING))
	}
	if instanceName == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_INSTANCE_NAME, ctv.TXT_IS_MISSING))
	}
	if keyB64 == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_SECRET_KEY, ctv.TXT_IS_MISSING))
	}
	if msg == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_MESSAGE, ctv.TXT_IS_MISSING))
	}
	if request == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_REQUEST, ctv.TXT_IS_MISSING))
	}
	if uId == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredParameterMissing, errs.BuildAdditionalInfo(ctv.LBL_UID, ctv.TXT_IS_MISSING))
	}

	tAdditionalInfo = fmt.Sprintf("%s%s %s @ %s", ctv.LBL_FUNCTION_NAME, functionName, ctv.TXT_FAILED, tFunctionName)

	if tJSON, errorInfo.Error = json.Marshal(request); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if eMessageDataB64, errorInfo = jwts.EncryptToByte(uId, keyB64, string(tJSON)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	msg.Data = eMessageDataB64

	if replyMsgPtr, errorInfo = RequestWithHeader(natsConnectionPtr, instanceName, msg, validateAdjustTimeOut(timeOutInSecs)); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
	}

	return
}

// SendReply - will take in an object, build a json object out of it, and send out the reply
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func SendReply(
	reply interface{},
	msg *nats.Msg,
) (errorInfo errs.ErrorInfo) {

	var (
		tJSONReply []byte
	)

	if tJSONReply, errorInfo = buildJSONReply(reply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v%v%v", ctv.LBL_SUBJECT, msg.Subject, ctv.LBL_MESSAGE_HEADER, msg.Header))
		return
	}

	if errorInfo.Error = msg.Respond(tJSONReply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v%v%v", ctv.LBL_SUBJECT, msg.Subject, ctv.LBL_MESSAGE_HEADER, msg.Header))
	}

	return
}

// Subscribe - will create a NATS subscription
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func Subscribe(
	connectionPtr *nats.Conn,
	instanceName, subject string,
	handler nats.MsgHandler,
) (
	subscriptionPtr *nats.Subscription,
	errorInfo errs.ErrorInfo,
) {

	if connectionPtr == nil {
		errorInfo = errs.NewErrorInfo(errs.ErrPointerMissing, fmt.Sprintf("%s%s", ctv.LBL_POINTER, ctv.TXT_NATS))
	}

	if subscriptionPtr, errorInfo.Error = connectionPtr.Subscribe(subject, handler); errorInfo.Error != nil {
		log.Printf("%v: Subscribe failed on subject: %v", instanceName, subject)
		return
	}
	log.Printf("%v Subscribed to subject: %v", instanceName, subject)

	return
}

//  Private Functions

// buildInstanceName - will create the NATS connection name with the delimiter between nodes.
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func buildInstanceName(
	delimiter string,
	nodes ...string,
) (
	instanceName string,
	errorInfo errs.ErrorInfo,
) {

	if len(nodes) == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprint(ctv.TXT_AT_LEAST_ONE))
		return
	}
	for index, node := range nodes {
		if index == 0 {
			instanceName = strings.Trim(node, ctv.SPACES_ONE)
		} else {
			instanceName = fmt.Sprintf("%v%v%v", instanceName, delimiter, strings.Trim(node, ctv.SPACES_ONE))
		}
	}

	return
}

// buildJSONReply - return a JSON reply object
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func buildJSONReply(reply interface{}) (
	jsonReply []byte,
	errorInfo errs.ErrorInfo,
) {

	if jsonReply, errorInfo.Error = json.Marshal(reply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_REPLY_TYPE, reflect.ValueOf(reply).Type().String()))
		return
	}

	return
}

// BuildTemporaryFiles - creates temporary files for Token.
// The function checks if the  NATSCredentialsFilename is provided. If the value is empty,
// the function returns an error.
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, returned from WriteFile
//	Verifications: None
func BuildTemporaryFiles(
	tempDirectory string,
	config NATSConfiguration,
) (
	errorInfo errs.ErrorInfo,
) {

	if config.NATSCredentialsFilename == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.LBL_MISSING_PARAMETER, ctv.FN_TOKEN))
		return
	} else {
		if errorInfo = hlp.WriteFile(fmt.Sprintf("%v/%v", tempDirectory, CREDENTIAL_FILENAME), []byte(config.NATSCredentialsFilename), 0744); errorInfo.Error != nil {
			return
		}
	}

	return
}

// buildURLPort - will create the NATS URL with the port.
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func buildURLPort(
	url string,
	port string,
) (
	natsURL string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tNATSPort, _ = strconv.Atoi(port)
	)

	if url == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprint(ctv.FN_URL))
		return
	}
	if tNATSPort == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(errs.ErrGreatThanZero, fmt.Sprint(ctv.FN_PORT))
		return
	}

	return fmt.Sprintf("%v:%d", url, tNATSPort), errs.ErrorInfo{}
}

// validateAdjustTimeOut - will check the timeout (Seconds) is between 2 and 5. If not it, will adjust the value
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func validateAdjustTimeOut(timeOutInSeconds int) (actualTimeOut time.Duration) {

	if timeOutInSeconds < 2 {
		actualTimeOut = 2 * time.Second
		return
	}

	if timeOutInSeconds > 5 {
		actualTimeOut = 5 * time.Second
		return
	}

	actualTimeOut = time.Duration(timeOutInSeconds) * time.Second

	return
}
