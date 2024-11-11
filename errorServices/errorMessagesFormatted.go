package sharedServices

//goland:noinspection ALL
const (
	FORMAT_EXPECTED_ERROR              = "%s Failed: Was expecting an err. Additional Info: %s"
	FORMAT_EXPECTING_NO_ERROR          = "%s Failed: Wasn't expecting an err. ERROR: %s"
	FORMAT_UNEXPECTED_ERROR            = "%s Failed: Unexpected err. ERROR: %s"
	FORMAT_FIRESTORE_ARGUMENTS_MISSING = "Require information is missing! Firestore Client Pointer or Datastore: '%v' Document Id: '%v'"
)
