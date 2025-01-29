package sharedServices

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// NewNATSService - builds a reusable NATS Service that creates an instance name, builds a connection, and has HandleRequestWithHeader,
// MakeRequestReplyWithHeader, SendReplyWithHeader, and Subscribe as methods.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewNATSService(
	extensionName string,
	config NATSConfiguration,
) (natsServicePtr *NATSService, errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckValueNotEmpty(extensionName, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.NATSURL, errs.ErrRequiredParameterMissing, ctv.LBL_NATS_URL); errorInfo.Error != nil {
		return
	}

	natsServicePtr = &NATSService{
		secure: true,
		url:    config.NATSURL,
	}
	if natsServicePtr.instanceName, errorInfo = buildInstanceName(extensionName, config.NATSURL); errorInfo.Error != nil {
		return
	}
	natsServicePtr.connPtr, errorInfo = getConnection(natsServicePtr.instanceName, config)

	return
}

// HandleRequestWithHeader - accepts a NATS message pointer, decrypts request message data, and return a DKRequest string. The provided requestMessagePtr
// must be retained by the caller, so it can be used to send a reply.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (natsServicePtr *NATSService) GetStatus() string {

	return natsServicePtr.connPtr.Status().String()
}

// HandleRequestWithHeader - accepts a NATS message pointer, decrypts request message data, and return a DKRequest string. The provided requestMessagePtr
// must be retained by the caller, so it can be used to send a reply.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (natsServicePtr *NATSService) HandleRequestWithHeader(
	keyB64 string,
	requestMessagePtr *nats.Msg,
) (
	dkRequest DKRequest,
	errorInfo errs.ErrorInfo,
) {

	dkRequest, errorInfo = handleRequestWithHeader(requestMessagePtr, keyB64)

	return
}

// makeRequestReplyWithHeader - submits a Base64 DK Request and wait for a DK Reply. The function will validate inputs,
// build a NATS message pointer, adjust the time-out in seconds as needed, make the request, wait for the reply, unmarshal the
// reply, and decrypt the DKReply.Reply string.
//
// The caller must create the DKRequest []byte and handling any errors returned.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (natsServicePtr *NATSService) MakeRequestReplyWithHeader(
	dkRequest []byte,
	keyB64 string,
	styhClientId string,
	subject string,
	uId string,
	timeOutInSec int,
) (
	dkReply DKReply,
	errorInfo errs.ErrorInfo,
) {

	natsServicePtr.userInfo.UId = uId
	natsServicePtr.userInfo.STYHClientId = styhClientId
	natsServicePtr.userInfo.KeyB64 = keyB64
	dkReply, errorInfo = makeRequestReplyWithHeader(dkRequest, natsServicePtr, subject, timeOutInSec)

	return
}

// MakeRequestReplyWithMessage - submits a NATS message and wait for a DK Reply. The function will validate inputs,
// adjust the time-out in seconds as needed, update the subject, make the request, wait for the reply, unmarshal the
// reply, and decrypt the DKReply.Reply string.
//
// The caller must provide the requestMessagePtr *nats.MSG and handling any errors returned.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (natsServicePtr *NATSService) MakeRequestReplyWithMessage(
	keyB64 string,
	requestMessagePtr *nats.Msg,
	subject string,
	timeOutInSec int,
) (
	dkReply DKReply,
	errorInfo errs.ErrorInfo,
) {

	natsServicePtr.userInfo.KeyB64 = keyB64
	dkReply, errorInfo = makeRequestReplyWithMessage(natsServicePtr, requestMessagePtr, subject, timeOutInSec)

	return
}

// SendReplyWithHeader - will reply to a request.
// The DKReply.Reply will be encrypted into a []byte. The DKReply will then be marshalled and sent out as a response
// using the original message (requestMessagePtr).
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (natsServicePtr *NATSService) SendReplyWithHeader(
	dkReply DKReply,
	keyB64 string,
	requestMessagePtr *nats.Msg,
) (
	errorInfo errs.ErrorInfo,
) {

	errorInfo = sendReplyWithHeader(dkReply, keyB64, requestMessagePtr)

	return
}

// Subscribe - will create a NATS subscription to a subject.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (natsServicePtr *NATSService) Subscribe(
	handler nats.MsgHandler,
	subject string,
) (
	subscriptionPtr *nats.Subscription,
	errorInfo errs.ErrorInfo,
) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	if errorInfo = hlps.CheckPointerNotNil(natsServicePtr.connPtr, errs.ErrPointerMissing, ctv.LBL_NATS); errorInfo.Error != nil {
		return
	}

	if subscriptionPtr, errorInfo.Error = natsServicePtr.connPtr.Subscribe(subject, handler); errorInfo.Error != nil {
		log.Printf("ALERT %v: Subscribe failed on subject: %v", natsServicePtr.instanceName, subject)
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(tFunctionName, ctv.TXT_SUBSCRIPTION_FAILED))
		return
	}
	log.Printf("%v Subscribed to subject: %v", natsServicePtr.instanceName, subject)

	return
}

//  Private Functions

// buildInstanceName - will create the NATS connection name with the delimiter between nodes.
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func buildInstanceName(
	extensionName string,
	natsURL string,
) (
	instanceName string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tHostName string
	)

	if errorInfo = hlps.CheckValueNotEmpty(extensionName, errs.ErrRequiredParameterMissing, ctv.LBL_EXTENSION_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(natsURL, errs.ErrRequiredParameterMissing, ctv.LBL_NATS_URL); errorInfo.Error != nil {
		return
	}

	tHostName, _ = os.Hostname()

	instanceName = fmt.Sprintf("%s-%s-%s", tHostName, extensionName, natsURL)

	return
}

// buildURLWithPort - will create the NATS URL with the port.
//
//	Customer Messages: None
//	Errors: error returned by natsServices.Connect
//	Verifications: None
func buildURLWithPort(
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

// getConnection - will connect to a NATS leaf server with either a ssl or non-ssl connection.
// This connection function requires natsServices.NATSConfiguration be populated. The following fields
// do not have to be at this time: TLSCert, TLSPrivateKey, TLSCABundle. The fields TLSCertFQN, TLSPrivateKeyFQN,
// TLSCABundleFQN must be populated.
//
// Notes:
//
//	MaxReconnects is set to 5
//	ReconnectWait is set to 2 seconds
//
// Customer Messages: None
// Errors: error returned by natsServices.Connect
// Verifications: None
func getConnection(
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

	if errorInfo = hlps.CheckValueNotEmpty(instanceName, errs.ErrRequiredParameterMissing, ctv.LBL_INSTANCE_NAME); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.NATSURL, errs.ErrRequiredParameterMissing, ctv.LBL_NATS_URL); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(config.NATSPort, errs.ErrRequiredParameterMissing, ctv.LBL_NATS_PORT); errorInfo.Error != nil {
		return
	}
	if vals.DoesFileExist(config.NATSCredentialsFilename) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, errs.BuildLabelValue(ctv.LBL_FILENAME, config.NATSCredentialsFilename))
	}
	if vals.DoesFileExist(config.NATSTLSInfo.TLSCertFQN) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, errs.BuildLabelValue(ctv.LBL_FILENAME, config.NATSTLSInfo.TLSCertFQN))
		return
	}
	if vals.DoesFileExist(config.NATSTLSInfo.TLSPrivateKeyFQN) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, errs.BuildLabelValue(ctv.LBL_FILENAME, config.NATSTLSInfo.TLSPrivateKeyFQN))
		return
	}
	if vals.DoesFileExist(config.NATSTLSInfo.TLSCABundleFQN) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, errs.BuildLabelValue(ctv.LBL_FILENAME, config.NATSTLSInfo.TLSCABundleFQN))
		return
	}

	opts = []nats.Option{
		nats.Name(instanceName),             // Set a client name
		nats.MaxReconnects(5),               // Set maximum reconnection attempts
		nats.ReconnectWait(2 * time.Second), // Set reconnection wait time
		nats.UserCredentials(config.NATSCredentialsFilename),
		nats.RootCAs(config.NATSTLSInfo.TLSCABundleFQN),
		nats.ClientCert(config.NATSTLSInfo.TLSCertFQN, config.NATSTLSInfo.TLSPrivateKeyFQN),
	}

	if tURL, errorInfo = buildURLWithPort(config.NATSURL, config.NATSPort); errorInfo.Error != nil {
		return
	}
	if connPtr, errorInfo.Error = nats.Connect(tURL, opts...); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v: %v", instanceName, ctv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	log.Printf("%v: A connection has been established with the NATS server at %v.", instanceName, config.NATSURL)
	log.Printf(
		"%v: URL: %v CLuster/Server Name: %v Server Id: %v Address: %v",
		instanceName,
		connPtr.ConnectedUrl(),
		connPtr.ConnectedClusterName(),
		connPtr.ConnectedServerId(),
		connPtr.ConnectedAddr(),
	)

	return
}

// handleRequestWithHeader - accepts a NATS message pointer, decrypts request message data, and return a DKRequest string. The provided requestMessagePtr
// must be retained by the caller, so it can be used to send a reply.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func handleRequestWithHeader(requestMessagePtr *nats.Msg, keyB64 string) (dkRequest DKRequest, errorInfo errs.ErrorInfo) {

	if errorInfo = hlps.CheckPointerNotNil(requestMessagePtr, errs.ErrPointerMissing, ctv.LBL_MESSAGE_REQUEST_POINTER); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(keyB64, errs.ErrPointerMissing, ctv.FN_KEY_B64); errorInfo.Error != nil {
		return
	}

	if dkRequest, errorInfo = jwts.DecryptToByte(requestMessagePtr.Header.Get(ctv.FN_UID), keyB64, string(requestMessagePtr.Data)); errorInfo.Error != nil {
		return
	}

	return
}

// makeRequestReplyWithHeader - submits a Base64 DK Request and wait for a DK Reply. The function will validate inputs,
// build a NATS message pointer, adjust the time-out in seconds as needed, make the request, wait for the reply, unmarshal the
// reply, and decrypt the DKReply.Reply string.
//
// The caller must create the DKRequest []byte and handling any errors returned.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func makeRequestReplyWithHeader(
	dkRequest []byte,
	natsServicePtr *NATSService,
	subject string,
	timeOutInSec int,
) (
	dkReply DKReply,
	errorInfo errs.ErrorInfo,
) {

	var (
		tActualTimeOut     time.Duration
		tReplyMessagePtr   *nats.Msg
		tRequestMessagePtr *nats.Msg
	)

	if errorInfo = hlps.CheckValueNotEmpty(string(dkRequest), errs.ErrRequiredParameterMissing, ctv.LBL_DK_REQEST); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckPointerNotNil(natsServicePtr, errs.ErrPointerMissing, ctv.LBL_NATS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckPointerNotNil(natsServicePtr.connPtr, errs.ErrPointerMissing, ctv.LBL_NATS_CONN_POINTER); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(subject, errs.ErrRequiredParameterMissing, ctv.LBL_SUBJECT); errorInfo.Error != nil {
		return
	}

	tRequestMessagePtr = &nats.Msg{
		Header:  make(nats.Header),
		Subject: subject,
	}
	tRequestMessagePtr.Header.Add(ctv.FN_UID, natsServicePtr.userInfo.UId)
	tRequestMessagePtr.Header.Add(ctv.FN_STYH_CLIENT_ID, natsServicePtr.userInfo.STYHClientId)
	if tRequestMessagePtr.Data, errorInfo = jwts.EncryptByteToByte(natsServicePtr.userInfo.UId, natsServicePtr.userInfo.KeyB64, dkRequest); errorInfo.Error != nil {
		return
	}

	tActualTimeOut = validateAdjustTimeOut(timeOutInSec)
	if tReplyMessagePtr, errorInfo.Error = natsServicePtr.connPtr.RequestMsg(tRequestMessagePtr, tActualTimeOut); errorInfo.Error != nil {
		log.Printf(
			"ALERT %s: RequestWithHeader failed on %s %s for %s: %s",
			natsServicePtr.instanceName,
			ctv.LBL_SUBJECT,
			subject,
			ctv.LBL_UID,
			tRequestMessagePtr.Header.Get(ctv.FN_UID),
		)
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdLabelValue(tRequestMessagePtr.Header.Get(ctv.FN_UID), natsServicePtr.instanceName, ctv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	if errorInfo.Error = json.Unmarshal(tReplyMessagePtr.Data, &dkReply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdLabelValue(tRequestMessagePtr.Header.Get(ctv.FN_UID), ctv.LBL_MESSAGE_REPLY, ctv.TXT_UNMARSHALL_FAILED))
		return
	}

	if errorInfo.Error == nil {
		dkReply.Reply, errorInfo = jwts.DecryptByteToByte(tRequestMessagePtr.Header.Get(ctv.FN_UID), natsServicePtr.userInfo.KeyB64, dkReply.Reply)
	}

	return
}

// makeRequestReplyWithMessage - submits a NATS message and wait for a DK Reply. The function will validate inputs,
// adjust the time-out in seconds as needed, update the subject, make the request, wait for the reply, unmarshal the
// reply, and decrypt the DKReply.Reply string.
//
// The caller must provide the requestMessagePtr *nats.MSG and handling any errors returned.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func makeRequestReplyWithMessage(
	natsServicePtr *NATSService,
	requestMessagePtr *nats.Msg,
	subject string,
	timeOutInSec int,
) (
	dkReply DKReply,
	errorInfo errs.ErrorInfo,
) {

	var (
		tActualTimeOut   time.Duration
		tReplyMessagePtr *nats.Msg
	)

	if errorInfo = hlps.CheckPointerNotNil(natsServicePtr, errs.ErrPointerMissing, ctv.LBL_NATS); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckPointerNotNil(natsServicePtr.connPtr, errs.ErrPointerMissing, ctv.LBL_NATS_CONN_POINTER); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckPointerNotNil(requestMessagePtr, errs.ErrRequiredParameterMissing, ctv.LBL_MESSAGE_REQUEST_POINTER); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckValueNotEmpty(subject, errs.ErrRequiredParameterMissing, ctv.LBL_SUBJECT); errorInfo.Error != nil {
		return
	}

	requestMessagePtr.Subject = subject
	tActualTimeOut = validateAdjustTimeOut(timeOutInSec)
	if tReplyMessagePtr, errorInfo.Error = natsServicePtr.connPtr.RequestMsg(requestMessagePtr, tActualTimeOut); errorInfo.Error != nil {
		log.Printf(
			"ALERT %s: RequestWithHeader failed on %s %s for %s: %s",
			natsServicePtr.instanceName,
			ctv.LBL_SUBJECT,
			subject,
			ctv.LBL_UID,
			requestMessagePtr.Header.Get(ctv.FN_UID),
		)
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdLabelValue(requestMessagePtr.Header.Get(ctv.FN_UID), natsServicePtr.instanceName, ctv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	if errorInfo.Error = json.Unmarshal(tReplyMessagePtr.Data, &dkReply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdLabelValue(requestMessagePtr.Header.Get(ctv.FN_UID), ctv.LBL_MESSAGE_REPLY, ctv.TXT_UNMARSHALL_FAILED))
		return
	}

	if errorInfo.Error == nil {
		dkReply.Reply, errorInfo = jwts.DecryptByteToByte(requestMessagePtr.Header.Get(ctv.FN_UID), natsServicePtr.userInfo.KeyB64, dkReply.Reply)
	}

	return
}

// sendReplyWithHeader - will take in an object, build a json object out of it, and send out the reply.
// The DKReply.Reply will be encrypted into a []byte. The DKReply will then be marshalled and sent out as a response
// using the original message (requestMessagePtr).
//
// The caller must create the DKReply.Reply string and handling any errors returned.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func sendReplyWithHeader(
	dkReply DKReply,
	keyB64 string,
	requestMessagePtr *nats.Msg,
) (errorInfo errs.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tReplyJSON         []byte
	)

	if dkReply.ErrorInfo.Error == nil {
		if errorInfo = hlps.CheckValueNotEmpty(string(dkReply.Reply), errs.ErrRequiredParameterMissing, ctv.LBL_DK_REPLY); errorInfo.Error != nil {
			return
		}
	}
	if dkReply.Reply == nil {
		if errorInfo = hlps.CheckValueNotEmpty(dkReply.ErrorInfo.Message, errs.ErrRequiredParameterMissing, ctv.LBL_ERROR_MESSAGE); errorInfo.Error != nil {
			return
		}
	}
	if errorInfo = hlps.CheckValueNotEmpty(keyB64, errs.ErrRequiredParameterMissing, ctv.LBL_KEY_B64); errorInfo.Error != nil {
		return
	}
	if errorInfo = hlps.CheckPointerNotNil(requestMessagePtr, errs.ErrRequiredParameterMissing, ctv.LBL_MESSAGE_REQUEST_POINTER); errorInfo.Error != nil {
		return
	}

	if len(dkReply.Reply) > ctv.VAL_ZERO {
		if dkReply.Reply, errorInfo = jwts.EncryptByteToByte(requestMessagePtr.Header.Get(ctv.FN_UID), keyB64, dkReply.Reply); errorInfo.Error != nil {
			return
		}
	}

	if tReplyJSON, errorInfo.Error = json.Marshal(dkReply); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdSubjectLabelValue(requestMessagePtr.Header.Get(ctv.FN_UID), requestMessagePtr.Subject, ctv.LBL_DK_REPLY, ctv.TXT_UNMARSHALL_FAILED))
		return
	}

	if errorInfo.Error = requestMessagePtr.Respond(tReplyJSON); errorInfo.Error != nil {
		log.Printf(
			"ALERT %s %s for %s%s %s%s",
			tFunctionName,
			ctv.TXT_FAILED,
			ctv.LBL_UID,
			requestMessagePtr.Header.Get(ctv.FN_UID),
			ctv.LBL_SUBJECT,
			requestMessagePtr.Subject,
		)
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildUIdSubjectLabelValue(requestMessagePtr.Header.Get(ctv.FN_UID), requestMessagePtr.Subject, ctv.LBL_NATS, ctv.TXT_FAILED))
	}

	return
}

// validateAdjustTimeOut - will check the timeout (Seconds) is between 2 and 30. If not it, will adjust the value
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func validateAdjustTimeOut(timeOutInSeconds int) (actualTimeOut time.Duration) {

	if timeOutInSeconds < 2 {
		actualTimeOut = 2 * time.Second
		return
	}

	if timeOutInSeconds > 30 {
		actualTimeOut = 30 * time.Second
		return
	}

	actualTimeOut = time.Duration(timeOutInSeconds) * time.Second

	return
}
