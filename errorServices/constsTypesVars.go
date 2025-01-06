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
	StackTrace     string `json:"error_stack_trace"`
	Message        string `json:"error_message"`
}
