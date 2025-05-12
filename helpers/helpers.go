package sharedServices

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	vals "github.com/sty-holdings/sharedServices/v2025/validators"
)

// AddMonths - will add the number of months and adjust the year.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func AddMonths(year int, month int, monthsToAdd int) (int, int) {
	month += monthsToAdd
	year += month / 12 // Add any full years
	month = month % 12 // Get the remaining months (modulo)
	if month == 0 {    // If month is 0 (divisible by 12), it was December
		month = 12
	}
	return year, month
}

// AdjustDateAdjustDateByDays - will modify the year, month and day when adding days
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func AdjustDateByDays(year int, month int, day int, addDays int) (int, int, int) {

	day += addDays

	for day > DetermineDaysInMonth(year, month) {
		day -= DetermineDaysInMonth(year, month)
		month++
		if month > 12 {
			month = 1
			year++
		}
	}

	return year, month, day
}

// function_name - builds ...
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func AppendToByteArray(byteArrayIn []byte, data ...any) (byteArrayOut []byte, errorInfo errs.ErrorInfo) {

	var (
		tSubjectPromptPart []byte
	)

	if tSubjectPromptPart, errorInfo.Error = json.Marshal(&data); errorInfo.Error != nil {
		return
	}
	byteArrayOut = append(byteArrayIn, tSubjectPromptPart...)

	return
}

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

// CheckArrayLengthGTZero - validates that the array length is greater than zero. If the values length is zero, then an error message is returned. The field label starts with ctv.LBL_ or ctv.FN_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckArrayLengthGTZero[T any](extensionName string, value []T, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if len(value) == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(extensionName, fmt.Sprintf("%s ", fieldLabel), ctv.TXT_IS_EMPTY))
	}

	return
}

// CheckMapLengthGTZero - validates that the map length is greater than zero. If the values length is zero, then an error message is returned. The field label starts with ctv.LBL_ or ctv.FN_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckMapLengthGTZero[K comparable, V any](extensionName string, value map[K]V, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if len(value) == ctv.VAL_ZERO {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(extensionName, fmt.Sprintf("%s ", fieldLabel), ctv.TXT_IS_EMPTY))
	}

	return
}

// CheckMissingFieldsInMap - compares two lists and returns missing fields
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckMissingFieldsInMap(data map[string]any, requiredFields []string) (missingFields []string) {

	for _, field := range requiredFields {
		if _, exists := data[field]; !exists || data[field] == nil || data[field] == "" {
			missingFields = append(missingFields, field)
		}
	}

	return missingFields
}

// CheckPointerNotNil - validates that the pointer is not nil. If the pointer is nil, then an error message is returned. The field label starts with ctv.LBL_ or ctv.FN_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckPointerNotNil(extensionName string, value interface{}, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if value == nil {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(extensionName, fmt.Sprintf("%s ", fieldLabel), ctv.TXT_IS_NIL))
	}

	return
}

// CheckValueNotEmpty - validates that the value is not empty. If the value is empty, then an error message is returned. The field label starts with ctv.LBL_.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func CheckValueNotEmpty(extensionName string, value string, err error, fieldLabel string) (errorInfo errs.ErrorInfo) {

	if value == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(err, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, fmt.Sprintf("%s ", fieldLabel), ctv.TXT_IS_EMPTY))
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

// ConvertMapStringToString - takes a map of strings and returns a string using the format of key: value
// with a space.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertMapStringToString(mapIn map[string]string) (mapString string) {

	for key, value := range mapIn {
		mapString += key + ": " + value + " "
	}

	return
}

// ConvertStringArrayToPSQLInList - takes an array of strings and returns a string for using by PSQL IN LIST
// with a space.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertStringArrayToPSQLInList(values []string) (inList string) {

	for _, value := range values {
		inList += "'" + value + "', "
	}

	inList = strings.TrimSuffix(inList, ", ")

	return
}

// ConvertStringDateToTime - converts a string date to a time.Time object.
//
//	Customer Messages: None
//	Errors: if parsing fails or invalid input is provided.
//	Verifications: None
func ConvertStringDateToTime(dateStringFormat string) (dateTimeFormat time.Time, errorInfo errs.ErrorInfo) {

	if dateTimeFormat, errorInfo.Error = time.Parse("2006-01-02", dateStringFormat); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dateStringFormat, ctv.TXT_IS_INVALID))
	}

	return
}

// ConvertStringMoneyToFloat - converts a string representation of money to a float value.
//
//	Customer Messages: None
//	Errors: errs.ErrorInfo returned if the string cannot be parsed.
//	Verifications: None
func ConvertStringMoneyToFloat(moneyString string) (money float64, errorInfo errs.ErrorInfo) {

	if moneyString == "" {
		return
	}

	if money, errorInfo.Error = strconv.ParseFloat(moneyString, 64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValueMessage(ctv.LBL_SERVICE_HELPERS, ctv.LBL_MONEY, moneyString, ctv.TXT_IS_INVALID))
	}

	return
}

// ConvertIntArrayToPSQLInList - takes an array of integers and returns a string for using by PSQL IN LIST
// with a space.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertIntArrayToPSQLInList(values []int64) (inList string) {

	for _, value := range values {
		inList += "'" + strconv.Itoa(int(value)) + "', "
	}

	inList = strings.TrimSuffix(inList, ", ")

	return
}

// ConvertMapAnyToMapString
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertMapAnyToMapString(mapIn map[any]interface{}) (mapOut map[string]interface{}) {

	mapOut = make(map[string]interface{})

	if vals.IsMapPopulated(mapIn) {
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

// ConvertStringToUUID - checks if string is a valid UUID and then converts it.
//
//	Customer Messages: None
//	Errors: returned by uuid.Parse
//	Verifications: None
func ConvertStringToUUID(stringIn string) (uuidOut uuid.UUID, errorInfo errs.ErrorInfo) {

	if uuidOut, errorInfo.Error = uuid.Parse(stringIn); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, stringIn, ctv.TXT_IS_INVALID))
	}

	return
}

// ConvertStructToMap - converts a given struct to a map by marshaling it into JSON and then unmarshalling it into a map.
// If there is an error during the marshaling or unmarshalling process, the error information is returned.
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

// ConvertDateTimeToTimestamp - converts a date and time string (2025-01-01 00:00:00) to a timestamp.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ConvertDateTimeToTimestamp(dateTime string, timezone string) (timestamp time.Time, errorInfo errs.ErrorInfo) {

	var (
		tLocationPtr *time.Location
	)

	if tLocationPtr, errorInfo.Error = time.LoadLocation(timezone); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, ctv.LBL_TIMEZONE, timezone))
		return
	}

	if timestamp, errorInfo.Error = time.ParseInLocation("2006-01-02 15:04:05", dateTime, tLocationPtr); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dateTime, ctv.TXT_IS_INVALID))
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
		errorInfo = errs.NewErrorInfo(errs.ErrEmptyProgramName, fmt.Sprintf("%v%v", ctv.LBL_REDIRECT, redirectTo))
	}

	return
}

// DetermineDaysInMonth - returns the number of days in the month
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DetermineDaysInMonth(year, month int) int {
	switch month {
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29
		} else {
			return 28
		}
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
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

// DollarsToPennies - multiples the value by 100. Called pennies because we did for the US first.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DollarsToPennies(amount float64) (pennies int64) {

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

// GetBatchName - returns a batch name for grouping actions. It uses the function name truncating everything before the first '('.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetBatchName(extensionName string, additionalLabel string) string {

	return fmt.Sprintf("%s%s %s", extensionName, additionalLabel[strings.Index(additionalLabel, ctv.PARENTHESE_LEFT):], fmt.Sprintf("%s", time.Now().Format("2006-01-02T15:04:05Z07:00")))
}

// GetDateParts splits a date string in "YYYY-MM-DD" format into its individual components; returns an error if the format is invalid.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetDateParts(dateString string) (dateParts []string, errorInfo errs.ErrorInfo) {

	if dateParts = strings.Split(dateString, "-"); len(dateParts) != 3 {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidDate, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, ctv.LBL_DATE, dateString))
		return
	}

	return
}

// GetDay - returns the current day of the month as an integer
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetDay() int {

	return int(time.Now().Day())
}

// GetDaveKnowsNetDomain - returns the host and port for the given environment.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetDaveKnowsNetDomain(environment string, port uint) (host string) {

	switch environment {
	case ctv.VAL_ENVIRONMENT_LOCAL:
		host = fmt.Sprintf("%s:%d", ctv.VAL_LOCAL_HOST, port)
		return
	case ctv.VAL_ENVIRONMENT_DEVELOPMENT:
		host = fmt.Sprintf("%s.%s:%d", ctv.VAL_ENVIRONMENT_SHORT_CODE_DEV, ctv.VAL_DAVEKNOWS_NET, port)
		return
	case ctv.VAL_ENVIRONMENT_PRODUCTION:
		host = fmt.Sprintf("%s.%s:%d", ctv.VAL_ENVIRONMENT_SHORT_CODE_PROD, ctv.VAL_DAVEKNOWS_NET, port)
	}

	return
}

// GetMonth - returns the current month as an integer
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetMonth() int {

	return int(time.Now().Month())
}

// GetEnvironmentShortCode - returns a short code representation of the given environment.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetEnvironmentShortCode(environment string) string {

	switch environment {
	case ctv.VAL_ENVIRONMENT_PRODUCTION:
		return ctv.VAL_ENVIRONMENT_SHORT_PROD
	case ctv.VAL_ENVIRONMENT_DEVELOPMENT:
		return ctv.VAL_ENVIRONMENT_SHORT_DEV
	case ctv.VAL_ENVIRONMENT_LOCAL:
		return ctv.VAL_ENVIRONMENT_LOCAL
	}

	return ctv.TXT_UNKNOWN
}

// getMonthFromDateString takes date parts and returns the month as an integer.
// It returns an error if the input string is not in the expected format or month is invalid.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getMonthFromDateParts(dateParts []string) (month int, errorInfo errs.ErrorInfo) {

	if month, errorInfo.Error = strconv.Atoi(dateParts[1]); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dateParts[1], ctv.TXT_NOT_CONVERTIBLE))
		return
	}

	if month < 1 || month > 12 {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidDate, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dateParts[1], ctv.TXT_IS_INVALID))
	}

	return
}

// GetQuarter determines the quarter of the year based on the provided month (1-12).
// Returns the quarter as an integer or an error if the month is invalid.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetQuarter(month int) (quarter int, errorInfo errs.ErrorInfo) {

	switch month {
	case 1, 2, 3:
		quarter = 1
	case 4, 5, 6:
		quarter = 2
	case 7, 8, 9:
		quarter = 3
	case 10, 11, 12:
		quarter = 4
	default:
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidDate, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, strconv.Itoa(month), ctv.TXT_IS_INVALID))
	}

	return
}

// GetQuarterStartEndDate - determines the start and end dates of the given quarter in a specified year.
// The function accepts the year and quarter as inputs and returns the respective dates as strings.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetQuarterStartEndDate(year int, quarter int) (quarterStart string, quarterEnd string) {

	switch quarter {
	case ctv.VAL_ONE:
		quarterStart = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_ONE_START_DATE)
		quarterEnd = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_ONE_END_DATE)
	case ctv.VAL_TWO:
		quarterStart = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_TWO_START_DATE)
		quarterEnd = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_TWO_END_DATE)
	case ctv.VAL_THREE:
		quarterStart = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_THREE_START_DATE)
		quarterEnd = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_THREE_END_DATE)
	case ctv.VAL_FOUR:
		quarterStart = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_FOUR_START_DATE)
		quarterEnd = fmt.Sprintf("%d-%s", year, ctv.VAL_QUARTER_FOUR_END_DATE)
	}

	return
}

// GetSundaySaturdayFromYearMonthDay takes the year, month, and day as int values
// and returns string for the Sunday of that week.
// It assumes the input date parts represent a valid date.
//
//	Customer Messages: None
//	Errors: Returns an error if the input string parts cannot be converted to integers.
//	Verifications: None
func GetSundaySaturdayFromYearMonthDay(year int, month int, day int) (sundayDate string, saturdayDate string, errorInfo errs.ErrorInfo) {

	var (
		tInputDate time.Time
		tWeekday   time.Weekday
	)

	tInputDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	tWeekday = tInputDate.Weekday()

	daysToSubtract := 0
	if tWeekday != time.Sunday {
		daysToSubtract = int(tWeekday)
	}

	tInputDate = tInputDate.AddDate(0, 0, -daysToSubtract)

	sundayDate = tInputDate.Format("2006-01-02")

	tInputDate = tInputDate.AddDate(0, 0, +6)

	saturdayDate = tInputDate.Format("2006-01-02")

	return
}

// GetSundayDateThisWeek returns the year, month, and day of the Sunday for the current week.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetSundayDateThisWeek() (year, month, day int) {
	var (
		today = time.Now()
	)

	// Calculate the weekday of today.
	weekday := today.Weekday()

	// Calculate the number of days to subtract to reach the Sunday.
	daysToSubtract := 0
	if weekday != time.Sunday {
		daysToSubtract = int(weekday)
	}

	// Subtract the days from today to get the Sunday.
	thisSunday := today.AddDate(0, 0, -daysToSubtract)

	// Extract the year, month, and day.
	year = thisSunday.Year()
	month = int(thisSunday.Month())
	day = thisSunday.Day()

	return year, month, day
}

// GetSundayDateWeeksAgo returns the year, month, and day of a Sunday, weeks_ago weeks from today.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetSundayDateWeeksAgo(weeksAgo int) (year, month, day int) {

	var (
		today = time.Now()
	)

	// Calculate the weekday of the given time.
	weekday := today.Weekday()

	// Calculate the number of days to subtract to reach the Sunday of the current week.
	daysToSubtractCurrentWeek := int(weekday)

	// Calculate the total number of days to subtract.
	totalDaysToSubtract := daysToSubtractCurrentWeek + (weeksAgo * 7)

	// Subtract the days from the given time.
	targetSunday := today.AddDate(0, 0, -totalDaysToSubtract)

	// Extract the year, month, and day from the target Sunday.
	year = targetSunday.Year()
	month = int(targetSunday.Month())
	day = targetSunday.Day()

	return year, month, day
}

// GetUnixDateFromValues - takes year, month, day, hour, minute, and second as integers and returns the corresponding Unix timestamp (seconds since epoch).
// For epoch date, hours, minutes and seconds must be zero.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetUnixDateFromValues(year int, month int, day int, hour int, minute int, second int) int64 {

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC).Unix()
}

// GetYear - returns the current year integer
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYear() int {

	return time.Now().Year()
}

// GetYearEndDateTime - return the last date/time of the supplied year
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearEndDateTime(year int) string {

	t := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")
}

// getYearFromDateParts returns the year as an integer, along with an error if the input is invalid or the year cannot be parsed.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getYearFromDateParts(dateParts []string) (year int, errorInfo errs.ErrorInfo) {

	if year, errorInfo.Error = strconv.Atoi(dateParts[0]); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dateParts[0], ctv.TXT_NOT_CONVERTIBLE))
	}

	return
}

// GetYearQuarterMonthWeekDayFromString takes a date string in "yyyy-mm-dd" format and returns the year, quarter, month,
// and day as integers. The function returns the date for Sunday for that week. It returns an error if the input string is not in the expected format or year/month is invalid.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearQuarterMonthWeekDayFromString(dateString string) (year int, quarter int, month int, weekStart string, weekEnd string, day int, errorInfo errs.ErrorInfo) {

	var (
		tParts []string
	)

	if tParts, errorInfo = GetDateParts(dateString); errorInfo.Error != nil {
		return
	}

	if year, errorInfo = getYearFromDateParts(tParts); errorInfo.Error != nil {
		return
	}
	if month, errorInfo = getMonthFromDateParts(tParts); errorInfo.Error != nil {
		return
	}
	if day, errorInfo = getDayFromDateParts(year, month, tParts[2]); errorInfo.Error != nil {
		return
	}
	if quarter, errorInfo = GetQuarter(month); errorInfo.Error != nil {
		return
	}
	if weekStart, weekEnd, errorInfo = GetSundaySaturdayFromYearMonthDay(year, month, day); errorInfo.Error != nil {
		return
	}

	return
}

// GetYearStartDateTime - return the first date/time of the supplied year
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearStartDateTime(year int) string {

	t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	return t.Format("2006-01-02 15:04:05")
}

// GetYearQuarterStartDateTime - return the last date/time of the supplied year and quarter
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearQuarterEndDateTime(year int, quarter int) string {

	var (
		tTime time.Time
	)

	switch quarter {
	case 1:
		tTime = time.Date(year, 3, 31, 23, 59, 59, 0, time.UTC)
	case 2:
		tTime = time.Date(year, 6, 30, 23, 59, 59, 0, time.UTC)
	case 3:
		tTime = time.Date(year, 9, 30, 23, 59, 59, 0, time.UTC)
	case 4:
		tTime = time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	default:
		return ctv.TXT_UNKNOWN
	}

	return tTime.Format("2006-01-02 15:04:05")
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

// GetYearMonthEndDateTime - return the last date/time of the supplied year and month
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetYearMonthEndDateTime(year int, month int) string {

	var (
		tTime time.Time
	)

	tTime = time.Date(year, time.Month(month+1), 1, 23, 59, 59, 0, time.UTC)
	return tTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
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
// 	case fmt.Sprintf("%v,%v", ctv.VAL_ENVIRONMENT_PRODUCTION, true):
// 		url = formatURL(ctv.HTTP_PROTOCOL_SECURE, ctv.HTTP_DOMAIN_API_PROD, ctv.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", ctv.VAL_ENVIRONMENT_PRODUCTION, false):
// 		url = formatURL(ctv.HTTP_PROTOCOL_NON_SECURE, ctv.HTTP_DOMAIN_API_PROD, ctv.HTTP_PORT_NON_SECURE)
// 	}
//
// 	return
// }

// CalculateFlagCombination takes a JSON string representing flags
// and uses bitwise to build a string. String order is
// day week month quarter year
//
//	 1    0    0      0      0
//	 0    1    0      0      0
//	 0    0    1      0      0
//	 0    0    0      1      0
//	 0    0    0      0      1
//
//		Customer Messages: None
//		Errors: None
//		Verifications: None
// func CalculateTimePeriodWordsFlagCombination() string {

// var (
//	tFlagCombination uint8
// )
//
// if wordsPresent.Year {
//	tFlagCombination |= 1 << ctv.FLAG_YEARS
// }
// if wordsPresent.Quarter {
//	tFlagCombination |= 1 << ctv.FLAG_QUARTERS
// }
// if wordsPresent.Month {
//	tFlagCombination |= 1 << ctv.FLAG_MONTHS
// }
// if wordsPresent.Week {
//	tFlagCombination |= 1 << ctv.FLAG_WEEKS
// }
// if wordsPresent.Day {
//	tFlagCombination |= 1 << ctv.FLAG_DAYS
// }
//
// return fmt.Sprintf("%05b")
//
// }

// FixFloat64ToDecimalPlaces - takes a float and truncates it to the number of places.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func FixFloat64ToDecimalPlaces(num float64, places int) float64 {
	format := fmt.Sprintf("%%.%df", places)
	fixedStr := fmt.Sprintf(format, num)
	var fixedNum float64
	_, _ = fmt.Sscanf(fixedStr, format, &fixedNum)
	return fixedNum
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
func GetDate() string {

	return time.Now().Format("2006-01-02")
}

// GetFormattedDate - builds a formatted date (yyyy-mm-dd)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFormattedDate(year int, month int, day int, locationPtr *time.Location) string {

	return time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		locationPtr,
	).Format("2006-01-02")
}

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
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, ctv.LBL_TIMEZONE, timezone))
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
		errorInfo = errs.NewErrorInfo(errs.ErrEmptyConfigFilename, tAdditionalInfo)
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

// GetTimeSinceEpoch -  converts an int64 representing seconds since the epoch to a formatted string in the format "yyyy-mm-dd hh:mm:ss AM/PM".
//
//	Customer Message: None
//	Errors: None
//	Verification: None
func GetTimeSinceEpoch(seconds int64) string {

	var (
		tTime time.Time
	)

	tTime = time.Unix(seconds, 0)

	return tTime.Format("2006-01-02 03:04:05 PM")
}

// IntSliceToInt32Slice - Helper function to convert []int to []int32
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IntSliceToInt32Slice(intSlice []int) (int32Slice []int32) {

	int32Slice = make([]int32, len(intSlice))

	for i, v := range intSlice {
		int32Slice[i] = int32(v)
	}

	return
}

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

// PenniesToDollars
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PenniesToDollars(pennies int64) float64 {

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
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidRedirectMode, fmt.Sprintf("%v%v", ctv.LBL_REDIRECT, redirectTo))
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
		errorInfo = errs.NewErrorInfo(errs.ErrOSFileDoesntExist, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_HELPERS, ctv.VAL_EMPTY, ctv.LBL_FILENAME, fqn, ctv.TXT_DELETE_FAILED))
		return
	}

	if errorInfo.Error = os.Remove(fqn); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrOSFileRemoval, fmt.Sprintf("%v%v", ctv.LBL_FILENAME, fqn))
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

// SubtractMonths - subtracts months and adjusted the year as needed. ...
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func SubtractMonths(year int, month int, monthsToSubtract int) (int, int) {

	month -= monthsToSubtract
	if month <= 0 {
		yearsToSubtract := int(math.Abs(float64(month)/12)) + 1 // Use math.Abs() and convert to int
		year -= yearsToSubtract
		month += 12 * yearsToSubtract
	}

	return year, month
}

// TrimString - takes a search value, finds it in the string, and trims everything to the left of the search value.
// If the search value is not found, the data is returned.
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
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelSubLabelValueMessage(ctv.LBL_SERVICE_HELPERS, ctv.VAL_EMPTY, ctv.LBL_FILENAME, fqn, ctv.TXT_CREATE_FAILED))
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

	if vals.IsDirectoryFullyQualified(logFQD) == false {
		errorInfo = errs.NewErrorInfo(errs.ErrOSDirectoryNotFullyQualified, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, ctv.LBL_LOG_DIRECTORY, logFQD))
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

// getDayFromDateParts takes date parts and returns the day as an integer.
// It returns an error if the input string is not in the expected format or day is invalid for the given month and year.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func getDayFromDateParts(year int, month int, dayIn string) (day int, errorInfo errs.ErrorInfo) {

	if day, errorInfo.Error = strconv.Atoi(dayIn); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errs.ErrInvalidDate, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dayIn, ctv.TXT_IS_INVALID))
		return
	}

	if vals.IsDayValid(year, month, day) == false {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_HELPERS, dayIn, ctv.TXT_IS_INVALID))
	}

	return
}

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
