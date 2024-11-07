package sharedServices

import (
	"regexp"
)

//goland:noinspection GoSnakeCaseUsage
const (
	TEST_BASE64_STRING          = "VEhpcyBpcyBhIHRlc3Qgc3RyaW5nIDEyMzQxMzQ1MjM1Nl4lKl4mJSYqKCVeKg=="
	TEST_FILE_NAME              = "test_file.txt"
	TEST_DIRECTORY              = "/tmp"
	TEST_DIRECTORY_ENDING_SLASH = "/tmp/"
	TEST_DIRECTORY_NON_ROOT     = "shared-services"
	TEST_STRING                 = "THis is a test string 123413452356^%*^&%&*(%^*"
)

var (
	TestByteArray = []byte(TEST_STRING) // Do not delete
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//goland:noinspection GoSnakeCaseUsage
const (
	TEST_DOMAIN               = "savup.com"
	TEST_FILE_EXISTS_FILENAME = "file_exists.txt"
	TEST_FILE_UNREADABLE      = "unreadable_file.txt"
	TEST_INVALID_DOMAIN       = "tmp"
)

//	type FirebaseFirestoreHelper struct {
//		AppPtr              *firebase.App
//		AuthPtr             *auth.Client
//		FirestoreClientPtr  *firestore.Client
//		CredentialsLocation string
//	}
var (
	testValidJson = []byte("{\"name\": \"Test Name\"}")
)
