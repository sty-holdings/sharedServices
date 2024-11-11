package sharedServices

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
)

var (
	ctx = context.Background()
)

// CreateStorageClient - connect to Google Cloud Platform services
//
//	Customer Messages: None
//	Errors: Return an errors generated by GCP with Ending Execution appended
//	Verifications: None
func CreateStorageClient(
	credentialsFile string,
	test bool,
) (
	client *storage.Client,
	errorInfo errs.ErrorInfo,
) {

	if client, errorInfo.Error = storage.NewClient(ctx, option.WithCredentialsJSON(getGCPKey(credentialsFile, test))); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error(), ctv.TXT_ENDING_EXECUTION)
	}

	return
}

// getBucket - return a pointer to a storage bucket
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: storageClientPtr is not nil
func getBucket(
	storageClientPtr *storage.Client,
	bucketName string,
) (
	bucketPtr *storage.BucketHandle,
	errorInfo errs.ErrorInfo,
) {

	if storageClientPtr == nil || bucketName == ctv.VAL_EMPTY {
		errorInfo.Error = errs.ErrRequiredArgumentMissing
	} else {
		// Create a bucket object for the specified bucket.
		bucketPtr = storageClientPtr.Bucket(bucketName)
	}

	return
}

// getGCPKey - will read the JSON key file. If either fail, exit is called.
//
//	Customer Messages: None
//	Errors: ErrUnableReadFile
//	Validations: File readable
func getGCPKey(
	GCPCredentialsFQN string,
	test bool,
) (GCPCredentials []byte) {

	var (
		errorInfo errs.ErrorInfo
	)

	if GCPCredentials, errorInfo.Error = os.ReadFile(GCPCredentialsFQN); errorInfo.Error != nil {
		errorInfo.Error = errs.ErrUnableReadFile
		log.Println(errorInfo.Error.Error())
	}

	if errorInfo.Error != nil && test == false {
		os.Exit(1)
	}

	return
}

// ListObjectsInBucket - return all the object names in a bucket, folders and files
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: storageClientPtr is not nil
func ListObjectsInBucket(
	storageClientPtr *storage.Client,
	bucketName string,
) (
	bucketList []string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tBucketPtr        *storage.BucketHandle
		tObjectAttributes *storage.ObjectAttrs
		tObjectIterator   *storage.ObjectIterator
	)

	if storageClientPtr == nil || bucketName == ctv.VAL_EMPTY {
		errorInfo.Error = errs.ErrRequiredArgumentMissing
	} else {
		tBucketPtr, errorInfo = getBucket(storageClientPtr, bucketName)

		// Create a list object for the bucket.
		tObjectIterator = tBucketPtr.Objects(ctx, nil)

		for {
			tObjectAttributes, errorInfo.Error = tObjectIterator.Next()
			if errorInfo.Error == iterator.Done {
				errorInfo.Error = nil
				break
			}
			if errorInfo.Error != nil {
				break
			}
			bucketList = append(bucketList, tObjectAttributes.Name)
		}
	}

	return
}

// ReadBucketObject - returns the contains of the bucket's named file.
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: bucketPtr is not nil
func ReadBucketObject(
	storageClientPtr *storage.Client,
	bucketName string,
	fileName string,
) (
	contents []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tBucketPtr *storage.BucketHandle
		tReader    *storage.Reader
	)

	if storageClientPtr == nil || bucketName == ctv.VAL_EMPTY || fileName == ctv.VAL_EMPTY {
		errorInfo.Error = errs.ErrRequiredArgumentMissing
	} else {
		if tBucketPtr, errorInfo = getBucket(storageClientPtr, bucketName); errorInfo.Error == nil {
			// Create an object for the specified file.
			if tReader, errorInfo.Error = tBucketPtr.Object(fileName).NewReader(context.Background()); errorInfo.Error == nil {
				defer func(tReader *storage.Reader) {
					_ = tReader.Close()
				}(tReader)
				// Read the contents of the reader into a byte slice.
				contents, errorInfo.Error = ioutil.ReadAll(tReader)
			}
		}
	}

	return
}
