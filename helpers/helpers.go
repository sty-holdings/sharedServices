package sharedServices

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	vlds "github.com/sty-holdings/sharedServices/v2025/validators"
)

// Base64Decode - will decode a base64 string to a string. If there is an error,
// the first 20 characters of the base64 string are logged.
// REMINDER: If the base64 string has sensitivity information, empty out the
// ErrorInfo.AdditionalInfo field before logging or outputting the error.
//
//	Customer Messages: None
//	Errors: error returned by StdEncoding.DecodeString
//	Verifications: None
func Base64Decode(base64Value string) (
	value []byte,
	errorInfo errs.ErrorInfo,
) {

	if value, errorInfo.Error = b64.StdEncoding.DecodeString(base64Value); errorInfo.Error != nil {
		errorInfo.AdditionalInfo = fmt.Sprintf("%v%v", ctv.LBL_BASE64, base64Value[:20])
	}

	return
}

// Base64Encode - will encode a string to a base64 string
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func Base64Encode(value string) string {

	return b64.StdEncoding.EncodeToString([]byte(value))
}

// CheckLengthGTZero - validates that the value length is greater than zero. If the values length is zero, then an error message is returned. The field label starts with ctv.LBL_ or ctv.FN_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckLengthGTZero(value []interface{}, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if len(value) == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(fieldLabel, ctv.TXT_IS_EMPTY))
	}

	return
}

// CheckPointerNotNil - validates that the pointer is not nil. If the pointer is nil, then an error message is returned. The field label starts with ctv.LBL_ or ctv.FN_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckPointerNotNil(value interface{}, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if value == nil {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(fieldLabel, ctv.TXT_IS_NIL))
	}

	return
}

// CheckValueNotEmpty - validates that the value is not empty. If the value is empty, then an error message is returned. The field label starts with ctv.LBL_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckValueNotEmpty(value string, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if value == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(fieldLabel, ctv.TXT_IS_EMPTY))
	}

	return
}

// BuildJSONRequest
// func BuildJSONRequest(request interface{}) (jsonRequest []byte) {
//
// 	var (
// 		err error
// 	)
//
// 	if jsonRequest, err = json.Marshal(request); err != nil {
// 		err = errors.New(fmt.Sprintf("Failed to generate JSON payload. Error: %v", err.Error()))
// 		log.Println(err.Error())
// 		// 	todo Error Handling & Notification
// 	}
//
// 	if coreValidators.IsJSONValid(jsonRequest) == false {
// 		jsonRequest = nil
// 		err = errors.New(fmt.Sprintf("Was not able to generate valid json for request %v", request))
// 		log.Println(err.Error())
// 		// 	todo Error Handling & Notification
// 	}
//
// 	return
// }

// BuildLegalName
// func BuildLegalName(firstName, lastName string) (legalName string) {
//
// 	if firstName != ctv.EMPTY && lastName != ctv.EMPTY {
// 		legalName = fmt.Sprintf("%v %v", firstName, lastName)
// 	}
//
// 	return
// }

// CapitalizeFirstLetter - will make the first letter of the string to upper case and the other letters to lower
// func CapitalizeFirstLetter(stringIn string) string {
//
// 	if stringIn == ctv.EMPTY {
// 		return ctv.EMPTY
// 	}
//
// 	x := []byte(stringIn)
// 	y := bytes.ToUpper([]byte{x[0]})
// 	z := bytes.ToLower(x[1:])
//
// 	return string(bytes.Join([][]byte{y, z}, nil))
// }

// ConvertMapAnyToMapString
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertMapAnyToMapString(mapIn map[any]interface{}) (mapOut map[string]interface{}) {

	mapOut = make(map[string]interface{})

	if vlds.IsMapPopulated(mapIn) {
		for key, value := range mapIn {
			mapOut[key.(string)] = value
		}
	}

	return
}

// ConvertMapStringToMapAny
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertMapStringToMapAny(mapIn map[string]interface{}) (mapOut map[any]interface{}) {

	mapOut = make(map[any]interface{})

	for key, value := range mapIn {
		mapOut[interface{}(key)] = value
	}

	return
}

// ConvertStructToMap - converts a given struct to a map by marshaling it into JSON and then unmarshaling it into a map.
// If there is an error during the marshaling or unmarshaling process, the error information is returned.
//
// Customer Messages: None
// Errors: error returned by json.Marshal and json.Unmarshal
// Verifications: None
func ConvertStructToMap(structIn interface{}) (
	mapOut map[string]interface{},
	errorInfo errs.ErrorInfo,
) {

	var (
		data []byte
	)

	if data, errorInfo.Error = json.Marshal(structIn); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(data, &mapOut); errorInfo.Error != nil {
		return
	}

	return
}

// ConvertDateTimeToTimestamp - converts a date and time string (2024-01-01 00:00:00 or 2024-01-01 00:00 or 2024-01-01) to timestamp.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertDateTimeToTimestamp(fieldName string, dateString string, timezone string) (timestamp time.Time, errorInfo errs.ErrorInfo) {

	var (
		tLocationPtr *time.Location
	)

	if tLocationPtr, errorInfo.Error = time.LoadLocation(timezone); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TIMEZONE, timezone))
		return
	}

	if timestamp, errorInfo.Error = time.ParseInLocation("2006-01-02 15:04:05", dateString, tLocationPtr); errorInfo.Error == nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(fieldName, ctv.TXT_IS_INVALID))
	}

	return
}

// ConvertSliceToSliceOfPtrs - takes a slice and returns a slice of pointers to the items in the slice.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertStringSliceToSliceOfPtrs(inbound []string) (outbound []*string) {

	for _, tSlice := range inbound {
		x := tSlice // required to get different pointers for each value.
		outbound = append(outbound, &x)
	}

	return
}

// CreateAndRedirectLogOutput - will create the fully qualified config file log directory.
// The log output is based on the redirectTo value, [MODE_OUTPUT_LOG | MODE_OUTPUT_LOG_DISPLAY].
// The log file name uses this format: 2006-01-02 15:04:05.000 Z0700. All spaces, colons, and periods
// are replaced with underscores.
//
//	Customer Messages: None
//	Errors: ErrDirectoryNotFullyQualified, any error from os.OpenFile
//	Verifications: IsDirectoryFullyQualified
func CreateAndRedirectLogOutput(logDirectory, redirectTo string) (
	logFileHandlerPtr *os.File,
	logFQN string,
	errorInfo errs.ErrorInfo,
) {

	switch redirectTo {
	case ctv.MODE_OUTPUT_LOG:
		logFileHandlerPtr, logFQN, errorInfo = createLogFile(logDirectory)
		log.SetOutput(io.MultiWriter(logFileHandlerPtr))
	case ctv.MODE_OUTPUT_LOG_DISPLAY:
		logFileHandlerPtr, logFQN, errorInfo = createLogFile(logDirectory)
		log.SetOutput(io.MultiWriter(os.Stdout, logFileHandlerPtr))
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrServerNameMissing, fmt.Sprintf("%v%v", ctv.LBL_REDIRECT, redirectTo))
	}

	return
}

// DetermineStartTime - will set the start time for the start.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DetermineStartTime(
	startDate string,
	timezone string,
) (
	startAt string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tLocationPtr *time.Location
		tStart       time.Time
	)

	if errorInfo = CheckValueNotEmpty(startDate, errs.ErrTimezoneNotSupported, ctv.LBL_END_DATE); errorInfo.Error != nil {
		return
	}
	if errorInfo = CheckValueNotEmpty(timezone, errs.ErrTimezoneNotSupported, ctv.LBL_TIMEZONE); errorInfo.Error != nil {
		return
	}

	if tLocationPtr, errorInfo.Error = time.LoadLocation(timezone); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TIMEZONE, timezone))
		return
	}
	tStart, errorInfo.Error = time.ParseInLocation("2006-01-02", startDate, tLocationPtr)

	startAt = fmt.Sprintf("%s %s", tStart.Format("2006-01-02"), ctv.TXT_START_DAY)

	return
}

// DetermineEndTime - will set the end time for the start.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DetermineEndEndTime(
	endDate string,
	timezone string,
) (
	endBy string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tLocationPtr *time.Location
		tEndBy       time.Time
	)

	if errorInfo = CheckValueNotEmpty(endDate, errs.ErrTimezoneNotSupported, ctv.LBL_END_DATE); errorInfo.Error != nil {
		return
	}
	if errorInfo = CheckValueNotEmpty(timezone, errs.ErrTimezoneNotSupported, ctv.LBL_TIMEZONE); errorInfo.Error != nil {
		return
	}

	if tLocationPtr, errorInfo.Error = time.LoadLocation(timezone); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TIMEZONE, timezone))
		return
	}
	tEndBy, errorInfo.Error = time.ParseInLocation("2006-01-02", endDate, tLocationPtr)

	endBy = fmt.Sprintf("%s %s", tEndBy.Format("2006-01-02"), ctv.TXT_START_DAY)

	return
}

// DoesFieldExist - tests the struct for the field name.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DoesFieldExist(
	structType interface{},
	fieldName string,
) bool {

	var (
		found bool
	)

	_, found = reflect.TypeOf(structType).FieldByName(fieldName)

	return found
}

// FloatToPennies - multiples the value by 100. Called pennies because we did for the US first.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func FloatToPennies(amount float64) (pennies int64) {

	return int64(amount * 100)
}

// formatURL - will return a formatted url with the protocol, domain, and port.
//
//	Validation: none
//	Format: "%v://%v:%v"
//	Example: http://verifyemail.savup.com:2134, https://verifyemail.savup.com:2134, http://localhost:2134, https://localhost:2134
// func formatURL(protocol, domain string, port uint) (url string) {
//
// 	if domain == ctv.ENVIRONMENT_LOCAL {
// 		url = fmt.Sprintf("%v://%v:%v", protocol, ctv.HTTP_DOMAIN_LOCALHOST, port)
// 	} else {
// 		url = fmt.Sprintf("%v://%v:%v", protocol, domain, port)
// 	}
//
// 	return
// }

// GetYearStartDateTime - return the first date/time of the supplied year
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearStartDateTime(year int) string {

	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")
}

// GetYearEndDateTime - return the last date/time of the supplied year
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearEndDateTime(year int) string {

	t := time.Date(year, 1, 1, 23, 59, 59, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")
}

// GetYearQuarterStartDateTime - return the first date/time of the supplied year and quarter
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearQuarterStartDateTime(year int, quarter int) string {

	var (
		tTime time.Time
	)
	switch quarter {
	case 1:
		tTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	case 2:
		tTime = time.Date(year, 4, 1, 0, 0, 0, 0, time.UTC)
	case 3:
		tTime = time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)
	case 4:
		tTime = time.Date(year, 10, 1, 0, 0, 0, 0, time.UTC)
	default:
		return ctv.TXT_UNKNOWN
	}

	return tTime.Format("2006-01-02 15:04:05")
}

// GetYearMonthStartDateTime - return the first date/time of the supplied year and month
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearMonthStartDateTime(year int, month int) string {

	t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")
}

// GetLastWeekStartDateTime - return the first date/time of last week depending on the weekStartDay value.
// weekStartDay of zero is Sunday and one is Monday. Anything other than 0 or 1, 0 will be used.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetLastWeekStartDateTime(weekStartDay int, today time.Time) time.Time {

	var (
		tDaysToSubtract int
		tWeekday        time.Weekday
	)

	tWeekday = today.Weekday()

	if weekStartDay == 1 {
		tDaysToSubtract = int(tWeekday) + 6
	} else {
		tDaysToSubtract = int(tWeekday) + 7
	}

	return today.AddDate(0, 0, -tDaysToSubtract)
}

// GenerateURL - will return the protocol, domain, and port. Using HTTP_PROTOCOL_SECURE or HTTP_PROTOCOL_NON_SECURE,
// ENDPOINT_VERIFY_EMAIL and HTTP_SECURE_PORT or HTTP_NON_SECURE_PORT based on the arguments.
//
//	Customer Messages: None
//	Errors: None
//	Verification: none
//	Example: http://localhost:1234, https://localhost:1234, http://api-dev.savup.com:1234, https://api-dev.savup.com:1234
//
// ToDo Change the Environment_local domain to local host once we have resolved the handshake issue happening when savup-httpServices is run locally.
// func GenerateURL(environment string, secure bool) (url string) {
//
// 	switch fmt.Sprintf("%v,%v", strings.ToUpper(environment), secure) {
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_LOCAL, true):
// 		url = formatURL(ctv.HTTP_PROTOCOL_SECURE, ctv.HTTP_DOMAIN_API_LOCAL, ctv.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_LOCAL, false):
// 		url = formatURL(ctv.HTTP_PROTOCOL_NON_SECURE, ctv.HTTP_DOMAIN_API_LOCAL, ctv.HTTP_PORT_NON_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_DEVELOPMENT, true):
// 		url = formatURL(ctv.HTTP_PROTOCOL_SECURE, ctv.HTTP_DOMAIN_API_DEV, ctv.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_DEVELOPMENT, false):
// 		url = formatURL(ctv.HTTP_PROTOCOL_NON_SECURE, ctv.HTTP_DOMAIN_API_DEV, ctv.HTTP_PORT_NON_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_PRODUCTION, true):
// 		url = formatURL(ctv.HTTP_PROTOCOL_SECURE, ctv.HTTP_DOMAIN_API_PROD, ctv.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.ENVIRONMENT_PRODUCTION, false):
// 		url = formatURL(ctv.HTTP_PROTOCOL_NON_SECURE, ctv.HTTP_DOMAIN_API_PROD, ctv.HTTP_PORT_NON_SECURE)
// 	}
//
// 	return
// }

// CalculateFlagCombination takes a JSON string representing flags
// and returns the bitwise flag combination as a uint8.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CalculateTimePeriodWordsFlagCombination(wordsPresent ctv.TimePeriodWordsPresent) (flagCombination uint8) {

	if wordsPresent.Year {
		flagCombination |= 1 << ctv.FLAG_YEARS
	}
	if wordsPresent.Quarter {
		flagCombination |= 1 << ctv.FLAG_QUARTERS
	}
	if wordsPresent.Month {
		flagCombination |= 1 << ctv.FLAG_MONTHS
	}
	if wordsPresent.Week {
		flagCombination |= 1 << ctv.FLAG_WEEKS
	}
	if wordsPresent.Day {
		flagCombination |= 1 << ctv.FLAG_DAYS
	}

	return
}

// GenerateUUIDType1 - provides the high level of uniqueness for UUIDs.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GenerateUUIDType1(removeDashes bool) (myUUID string) {

	_UUID, _ := uuid.NewUUID()
	myUUID = fmt.Sprintf("%v", _UUID)

	if removeDashes {
		myUUID = strings.Replace(myUUID, "-", "", -1)
	}

	return
}

// GenerateUUIDType4 - generates a 128-bit random UUID. It is highly likely to be unique but not guaranteed. If uniqueness is needed
// used UUID type 1 (GenerateUUIDType1)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GenerateUUIDType4(removeDashes bool) (myUUID string) {

	_UUID, _ := uuid.NewRandom()
	myUUID = fmt.Sprintf("%v", _UUID)

	if removeDashes {
		myUUID = strings.Replace(myUUID, "-", "", -1)
	}

	return
}

// GenerateVerifyEmailURLWithUUID - return the url and uuid for the Verify Email.
// func GenerateVerifyEmailURLWithUUID(environment string, secure bool) (url, uuid string) {
//
// 	uuid = GenerateUUIDType4(false)
// 	url = fmt.Sprintf("%v?uuid=%v", GenerateVerifyEmailURL(environment, secure), uuid)
//
// 	return
// }

// GenerateVerifyEmailURLWithUUIDUsername - return the url, uuid and the username for the Verify Email.
// func GenerateVerifyEmailURLWithUUIDUsername(username, environment string, secure bool) (url, uuid string) {
//
// 	uuid = GenerateUUIDType4(false)
// 	url = fmt.Sprintf("%v?%v=%v&%v=%v", GenerateVerifyEmailURL(environment, secure), ctv.FN_UUID, uuid, ctv.FN_USERNAME, username)
//
// 	return
// }

// GenerateVerifyEmailURL - return the url.
// func GenerateVerifyEmailURL(environment string, secure bool) (url string) {
//
// 	url = fmt.Sprintf("%v/%v", GenerateURL(environment, secure), ctv.ENDPOINT_VERIFY_EMAIL)
//
// 	return
// }

// GetDate - return the current date in YYYY-MM-DD format
//
//	Customer Message: None
//	Errors: None
//	Verification: None
// func GetDate() string {
// 	return time.Now().Format("2006-01-02")
// }

// GetDateTimeWithLocation - returns the date/time (month day, year hh:mm:ss AM/PM format based on
// the IANA Time Zone Database value.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetDateTimeWithLocation(timezone string) (myDateTime string, errorInfo errs.ErrorInfo) {

	var (
		tLocationPtr *time.Location
	)

	if tLocationPtr, errorInfo = GetLocationTimePtr(timezone); errorInfo.Error != nil {
		return
	}

	return time.Now().In(tLocationPtr).Format("January 2, 2006 3:04:05 PM"), errs.ErrorInfo{}
}

// GetLocationTimePtr - using the IANA Time Zone Database, this will return a locationPtr for date/time.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetLocationTimePtr(timezone string) (locationPtr *time.Location, errorInfo errs.ErrorInfo) {

	if locationPtr, errorInfo.Error = time.LoadLocation(timezone); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_TIMEZONE, timezone))
		return
	}

	return
}

// GetUnixTimestamp - gets date/time in unix format (Mon Jan _2 15:04:05 MST 2006)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetUnixTimestamp() (timestamp string) {

	return time.Now().Local().Format(time.UnixDate)
}

// GetJSONFile - reads and unmarshal the fully qualified file into the object pointer.
//
//	Customer Messages: None
//	Errors: ErrJSONInvalid
//	Verifications: None
func GetJSONFile(
	jsonFileFQN string,
	jsonFilePtr *interface{},
) (
	errorInfo errs.ErrorInfo,
) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", ctv.LBL_FILENAME, jsonFileFQN)
		tJSONFileData   []byte
	)

	if tJSONFileData, errorInfo.Error = os.ReadFile(jsonFileFQN); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrConfigFileMissing, tAdditionalInfo)
	}

	if errorInfo.Error = json.Unmarshal(tJSONFileData, &jsonFilePtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	return
}

// GetAWSTimestamp - gets date/time in AWS format (Mon Jan _2 15:04:05 MST 2006)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetAWSTimestamp() (timestamp string) {

	x := time.Now().Local()
	weekday := x.Format("Mon")
	day := x.Format("2")
	month := x.Format("Jan")
	year := x.Year()
	hour := x.Hour()
	minutes := x.Minute()
	seconds := x.Second()
	timezone := x.Format("MST")

	return fmt.Sprintf("%v %v %v %v:%v:%v %v %v", weekday, month, day, hour, minutes, seconds, timezone, year)
}

// GetUnixTimestampByte - gets date/time in unix format (Mon Jan _2 15:04:05 MST 2006) as []byte
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetUnixTimestampByte() (timestamp []byte) {

	return []byte(time.Now().Local().Format(time.UnixDate))
}

// GetTime - return the current time in HH-mm-ss.00000 format, where hour is in military time.
//
//	Customer Message: None
//	Errors: None
//	Verification: None
// func GetTime() string {
// 	return time.Now().Format("15-04-05.00000")
// }

// GetFieldsNames - will return a list fields in a struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFieldsNames(unknownStruct interface{}) (
	fields map[string]interface{},
	errorInfo errs.ErrorInfo,
) {

	fields = make(map[string]interface{})

	tStruct := reflect.ValueOf(unknownStruct)
	tType := tStruct.Type()

	for i := 0; i < tType.NumField(); i++ {
		if tType.Field(i).IsExported() {
			fields[tType.Field(i).Name] = tStruct.FieldByName(tType.Field(i).Name).Interface()
		}
	}

	return
}

// PenniesToFloat
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PenniesToFloat(pennies int64) float64 {

	return float64(pennies) / 100
}

// PrintAndDie - is exported for access in other packages. Not going to test
// func PrintAndDie(msg string) {
//
// 	_, _ = fmt.Fprintln(os.Stderr, msg)
// 	os.Exit(1)
//
// }

// PrependWorkingDirectory - will add the working directory.
// if the filename first character is a /, the passed value will be returned
// unmodified.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrependWorkingDirectory(filename string) string {

	var (
		tWorkingDirectory, _ = os.Getwd()
	)

	if filepath.IsAbs(filename) {
		return filename
	}

	return fmt.Sprintf("%v/%v", tWorkingDirectory, filename)
}

// PrependWorkingDirectoryWithEndingSlash - will add the working directory, a slash, the directory
// provided, and an ending slash. If the directory first character is a slash, the passed value will
// be returned unmodified. The last character is not checked, so you could end up with two slashes.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrependWorkingDirectoryWithEndingSlash(directory string) string {

	var (
		tWorkingDirectory, _ = os.Getwd()
	)

	if filepath.IsAbs(directory) {
		return directory
	}

	return fmt.Sprintf("%v/%v/", tWorkingDirectory, directory)
}

// printDashLine - will output a given number of dashed lines based on the outputMode.
// The default is to output to the log
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func printDashLines(lines int, outputMode string) {
//
// 	for i := 0; i < lines; i++ {
// 		if strings.ToLower(outputMode) == ctv.MODE_OUTPUT_DISPLAY {
// 			fmt.Println("------------------------------------------")
// 		} else {
// 			log.Println("------------------------------------------")
// 		}
// 	}
// }

// PrintLinesAtStartOfRequest - will output dashed lines when a new request arrives.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func PrintLinesAtStartOfRequest(lines int, outputMode string) {
//
// 	printDashLines(lines, outputMode)
//
// }

// RedirectLogOutput - will redirect log output based on the redirectTo value, [MODE_OUTPUT_LOG | MODE_OUTPUT_LOG_DISPLAY].
//
//	Customer Messages: None
//	Errors: ErrDirectoryNotFullyQualified, any error from os.OpenFile
//	Verifications: IsDirectoryFullyQualified
func RedirectLogOutput(
	inLogFileHandlerPtr *os.File,
	redirectTo string,
) (errorInfo errs.ErrorInfo) {

	switch redirectTo {
	case ctv.MODE_OUTPUT_LOG:
		log.SetOutput(io.MultiWriter(inLogFileHandlerPtr))
	case ctv.MODE_OUTPUT_LOG_DISPLAY:
		log.SetOutput(io.MultiWriter(os.Stdout, inLogFileHandlerPtr))
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrRedirectModeInvalid, fmt.Sprintf("%v%v", ctv.LBL_REDIRECT, redirectTo))
	}

	return
}

// RemoveFile - removes a file for the file system.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RemoveFile(fqn string) (errorInfo errs.ErrorInfo) {

	// This doesn't use the coreValidator.DoesFileExist by design.
	if _, err := os.Stat(fqn); err != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFileDoesntExist, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, fqn))
		return
	}

	if errorInfo.Error = os.Remove(fqn); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrFileRemovalFailed, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, fqn))
		return
	}

	return
}

// RemovePidFile - removes the pid file for the running instance
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RemovePidFile(pidFQN string) (errorInfo errs.ErrorInfo) {

	if errorInfo = RemoveFile(pidFQN); errorInfo.Error != nil {
		return
	}

	return
}

// TrimString - takes a search value, finds it in the string, and trims everything to the left of the search value.
// If the search value is not found the data is returned.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func TrimString(searchValue string, data string) (trimmedPath string) {

	index := strings.Index(data, searchValue)
	if index == -1 {
		trimmedPath = data
		return
	}
	trimmedPath = data[index:]

	return
}

// WriteFile - will create and write to a fully qualified file.
//
//	Customer Messages: None
//	Errors: ErrFileCreationFailed
//	Verifications: None
func WriteFile(
	fqn string,
	fileData []byte,
	filePermissions os.FileMode,
) (errorInfo errs.ErrorInfo) {

	if errorInfo.Error = os.WriteFile(fqn, fileData, filePermissions); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v %v%v", errs.ErrFileCreationFailed.Error(), ctv.LBL_FILENAME, fqn))
	}

	return
}

// WritePidFile - will create and write the server pid file.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func WritePidFile(
	pidFQN string,
	pid int,
) (errorInfo errs.ErrorInfo) {

	if errorInfo = WriteFile(pidFQN, []byte(strconv.Itoa(pid)), 0766); errorInfo.Error == nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, pidFQN))
	}

	return
}

// Private Functions

// createLogFile - will create and open the  log file using the fully qualified directory.
//
//	Customer Messages: None
//	Errors: ErrDirectoryNotFullyQualified, any error from os.OpenFile
//	Verifications: IsDirectoryFullyQualified
func createLogFile(logFQD string) (
	logFileHandlerPtr *os.File,
	logFQN string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tLogFileName string
	)

	if vlds.IsDirectoryFullyQualified(logFQD) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrDirectoryNotFullyQualified, fmt.Sprintf("%v%v", ctv.LBL_DIRECTORY, logFQD))
		return
	}

	tDateTime := time.Now().Format("2006-01-02 15:04:05.000 Z0700")
	tLogFileName = strings.Replace(
		strings.Replace(strings.Replace(tDateTime, ctv.SPACES_ONE, ctv.UNDERSCORE, -1), ctv.COLON, ctv.UNDERSCORE, -1),
		ctv.PERIOD,
		ctv.UNDERSCORE,
		-1,
	)
	logFQN = fmt.Sprintf("%v%v.log", logFQD, tLogFileName)

	// Set log file output
	if logFileHandlerPtr, errorInfo.Error = os.OpenFile(logFQN, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, logFQN))
		return
	}

	return
}

// getType
// func getType(myVar interface{}) (myType string) {
//
// 	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
// 		myType = "*" + t.Elem().Name()
// 	} else {
// 		myType = t.Name()
// 	}
//
// 	return
// }

// getUTCOffsetSeconds - Helper function to convert UTC offset string to seconds.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getUTCOffsetSeconds(userUTCOffsetting string) (utcOffset int, errorInfo errs.ErrorInfo) {

	var (
		hours   int
		minutes int
	)

	if _, errorInfo.Error = fmt.Sscanf(userUTCOffsetting, "%d:%d", &hours, &minutes); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, userUTCOffsetting))
		return
	}

	utcOffset = (hours*60 + minutes) * 60

	return
}
