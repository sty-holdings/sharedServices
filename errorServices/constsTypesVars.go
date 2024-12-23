package sharedServices

//goland:noinspection ALL
const (
	ERROR    = "error"
	NO_ERROR = "no error"
	//
	TEST_STRING = "TEST STRING"
)

type ErrorInfo struct {
	AdditionalInfo string `json:"error_additional_info"`
	Error          error
	FileName       string `json:"error_filename"`
	FunctionName   string `json:"error_function_name"`
	LineNumber     int    `json:"error_line_number"`
	Message        string `json:"error_message"`
}
