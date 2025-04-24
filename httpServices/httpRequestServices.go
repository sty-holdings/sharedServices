package sharedServices

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlps "github.com/sty-holdings/sharedServices/v2025/helpers"
	pis "github.com/sty-holdings/sharedServices/v2025/programInfo"
)

// NewHTTPGetRequest - sends an HTTP GET request with specified headers and query parameters.
//
//	Customer Messages: None
//	Errors: errs.ErrEmptyRequiredParameter, errs.ErrFailedHttpRequest, or errors from underlying HTTP functions
//	Verifications: hlps.CheckValueNotEmpty, HTTPRequestService.parseURL, HTTPRequestService.buildRawQuery
func NewHTTPGetRequest(myURL string, headerSettings map[string]string, querySettings map[string]string) (body []byte, errorInfo errs.ErrorInfo) {

	var (
		tHTTPRequestPtr  *http.Request
		tHTTPResponsePtr *http.Response
		tServicePtr      *HTTPRequestService
	)

	if errorInfo = hlps.CheckValueNotEmpty(ctv.LBL_SERVICE_HTTP, myURL, errs.ErrEmptyRequiredParameter, ctv.FN_URL); errorInfo.Error != nil {
		return
	}

	tServicePtr = &HTTPRequestService{
		clientPtr: &http.Client{},
	}

	if errorInfo = tServicePtr.parseURL(myURL); errorInfo.Error != nil {
		return
	}

	tServicePtr.buildRawQuery(querySettings)

	if tHTTPRequestPtr, errorInfo.Error = http.NewRequestWithContext(context.Background(), HTTP_GET, tServicePtr.urlPtr.String(), nil); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), ctv.VAL_EMPTY, ctv.TXT_FAILED))
		return
	}

	for key, value := range headerSettings {
		tHTTPRequestPtr.Header.Set(key, value)
	}

	if tHTTPResponsePtr, errorInfo.Error = tServicePtr.clientPtr.Do(tHTTPRequestPtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), "Do()", ctv.TXT_FAILED))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), "close()", ctv.TXT_FAILED))
			return
		}
	}(tHTTPResponsePtr.Body)

	if tHTTPResponsePtr.StatusCode != http.StatusOK {
		errorInfo = errs.NewErrorInfo(
			errs.ErrFailedHttpRequest,
			errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), fmt.Sprintf("Status Code: %d", tHTTPResponsePtr.StatusCode), ctv.TXT_FAILED),
		)
		return
	}

	if body, errorInfo.Error = io.ReadAll(tHTTPResponsePtr.Body); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), "ReadAll()", ctv.TXT_FAILED))
	}

	return
}

// Private methods below here

// buildRawQuery - appends query parameters to the provided URL and encodes them.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func (servicePtr *HTTPRequestService) buildRawQuery(nameValues map[string]string) {

	var (
		tQuery url.Values
	)

	tQuery = servicePtr.urlPtr.Query()
	for key, value := range nameValues {
		tQuery.Set(key, value)
	}

	servicePtr.urlPtr.RawQuery = tQuery.Encode()

	return
}

// parseURL - parses the given URL string and returns a parsed URL object or an error.
//
// Customer Messages: None
// Errors: errs.ErrorInfo
// Verifications: None
func (servicePtr *HTTPRequestService) parseURL(myURL string) (errorInfo errs.ErrorInfo) {

	if servicePtr.urlPtr, errorInfo.Error = url.Parse(myURL); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HTTP, pis.GetMyFunctionName(true), ctv.VAL_EMPTY, ctv.TXT_FAILED))
	}

	return
}

//  Private Functions
