package sharedServices

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt/v5"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
	hlp "github.com/sty-holdings/sharedServices/v2025/helpers"
	vldts "github.com/sty-holdings/sharedServices/v2025/validators"
)

// Decrypt - decrypts a Base64 string using AES-GCM mode of operation and returns a string.
//
// Customer Messages: None
// Errors:
// - ErrInvalidBase64: If the encryptedMessage could not be decoded from Base64.
// - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// - ErrAEADInitialization: If the AEAD object could not be created.
// - ErrDecryption: If the ciphertext could not be decrypted.
// Verifications: None
func Decrypt(
	internalUserID string,
	keyB64 string,
	valueB64 string,
) (
	decryptedValue string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tAESGCM     cipher.AEAD
		tBlock      cipher.Block
		tCiphertext []byte
		tDecodedKey []byte
		tNonce      []byte
		tNonceSize  int
		tPlaintext  []byte
	)

	if errorInfo = checkEncryptDecryptParameters(internalUserID, keyB64, valueB64); errorInfo.Error != nil {
		return
	}

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(keyB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, errs.BuildLabelValue(ctv.LBL_SERVICE_JWT, ctv.LBL_KEY_B64, ctv.TXT_DECODE_FAILED))
		return
	}
	if tCiphertext, errorInfo.Error = base64.StdEncoding.DecodeString(valueB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_VALUE_B64, ctv.TXT_DECODE_FAILED))
	}

	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_KEY_DECODED, internalUserID))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DECIPHER_BLOCK, internalUserID))
		return
	}

	tNonceSize = tAESGCM.NonceSize()
	tNonce, tCiphertext = tCiphertext[:tNonceSize], tCiphertext[tNonceSize:]

	if tPlaintext, errorInfo.Error = tAESGCM.Open(nil, tNonce, tCiphertext, nil); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_INTERNAL_USER_ID, internalUserID))
		return
	}

	decryptedValue = string(tPlaintext)

	return
}

// DecryptToByte - will call decrypt using a Base64 string and returns a []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func DecryptToByte(
	internalUserID string,
	keyB64 string,
	valueB64 string,
) (
	decryptedValue []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tDecryptedValue string
	)

	if tDecryptedValue, errorInfo = Decrypt(internalUserID, keyB64, valueB64); errorInfo.Error == nil {
		decryptedValue = []byte(tDecryptedValue)
	}

	return
}

// DecryptByteToByte - will call decrypt using a Base64 []byte and returns a []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func DecryptByteToByte(
	internalUserID string,
	keyB64 string,
	valueB64 []byte,
) (
	decryptedValue []byte,
	errorInfo errs.ErrorInfo,
) {

	decryptedValue, errorInfo = DecryptToByte(internalUserID, keyB64, string(valueB64))

	return
}

// DecryptByteToString - will call decrypt using a Base64 []byte and return a string
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func DecryptByteToString(
	internalUserID string,
	keyB64 string,
	valueB64 []byte,
) (
	decryptedMessage string,
	errorInfo errs.ErrorInfo,
) {

	decryptedMessage, errorInfo = Decrypt(internalUserID, keyB64, string(valueB64))

	return
}

// Encrypt - encrypts a value using AES-GCM mode of operation and returns a Base64 string
// The function encrypts message and initializes the AES cipher
// with the provided key.
//
// Returns:
// - encryptedMessage: Base64 encoded string
//
// Customer Messages: None
// Errors:
// - ErrInvalidBase64: If the encryptedMessage could not be decoded from Base64.
// - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// - ErrAEADInitialization: If the AEAD object could not be created.
// - ErrDecryption: If the ciphertext could not be decrypted.
// Verifications: None
func Encrypt(
	internalUserID string,
	keyB64 string,
	value string,
) (
	encryptedValueB64 string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tAESGCM     cipher.AEAD
		tBlock      cipher.Block
		tCiphertext []byte
		tDecodedKey []byte
		tNonce      []byte
	)

	if errorInfo = checkEncryptDecryptParameters(internalUserID, keyB64, value); errorInfo.Error != nil {
		return
	}

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(keyB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_KEY_B64, internalUserID))
		return
	}

	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_KEY_DECODED, internalUserID))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_DECIPHER_BLOCK, internalUserID))
		return
	}

	tNonce = make([]byte, tAESGCM.NonceSize())
	//goland:noinspection GoUnhandledErrorResult
	io.ReadFull(rand.Reader, tNonce) // Use a cryptographically secure RNG

	tCiphertext = tAESGCM.Seal(tNonce, tNonce, []byte(value), nil)
	encryptedValueB64 = base64.StdEncoding.EncodeToString(tCiphertext)

	return
}

// EncryptToByte - will call encrypt using a string and returns a Base64 []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func EncryptToByte(
	internalUserID string,
	keyB64 string,
	value string,
) (
	encryptedValueB64 []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tEncryptedMessageB64 string
	)

	if tEncryptedMessageB64, errorInfo = Encrypt(internalUserID, keyB64, value); errorInfo.Error == nil {
		encryptedValueB64 = []byte(tEncryptedMessageB64)
	}

	return
}

// EncryptByteToByte - will call encrypt using a []byte and returns a Base64 []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func EncryptByteToByte(
	internalUserID string,
	keyB64 string,
	value []byte,
) (
	encryptedValueB64 []byte,
	errorInfo errs.ErrorInfo,
) {

	encryptedValueB64, errorInfo = EncryptToByte(internalUserID, keyB64, string(value))

	return
}

// EncryptByteToString - will call encrypt the []byte and return a Base64 string
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func EncryptByteToString(
	internalUserID string,
	keyB64 string,
	encryptedValueB64 []byte,
) (
	decryptedMessage string,
	errorInfo errs.ErrorInfo,
) {

	decryptedMessage, errorInfo = Encrypt(internalUserID, keyB64, string(encryptedValueB64))

	return
}

// GenerateJWT
// Create a new token object, specifying signing method and the claims
// you would like it to contain.
// func GenerateJWT(privateKey, requesterId, period string, duration int64) (jwtServices string, errorInfo errs.ErrorInfo) {
//
// 	var (
// 		tDuration      time.Duration
// 		tPrivateKey    *rsa.PrivateKey
// 		tRawPrivateKey []byte
// 	)
//
// 	if privateKey == ctv.VAL_EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v'", ctv.FN_PRIVATE_KEY, ctv.VAL_EMPTY))
// 		log.Println(errorInfo.Error)
// 	} else {
// 		if requesterId == ctv.VAL_EMPTY || period == ctv.VAL_EMPTY || duration < 1 {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v' %v: '%v'", ctv.FN_REQUESTER_ID, requesterId, ctv.FN_PERIOD, period, ctv.FN_DURATION, duration))
// 			log.Println(errorInfo.Error)
// 		} else {
// 			if cv.IsPeriodValid(period) && duration > 0 {
// 				tRawPrivateKey = []byte(privateKey)
// 				if tPrivateKey, errorInfo = ParsePrivateKey(tRawPrivateKey); errorInfo.Error == nil {
// 					switch strings.ToUpper(period) {
// 					case "M":
// 						tDuration = time.Minute * time.Duration(duration)
// 					case "H":
// 						tDuration = time.Hour * time.Duration(duration)
// 					case "D":
// 						tDuration = time.Hour * time.Duration(duration*24)
// 					default:
// 						tDuration = time.Hour * time.Duration(duration)
// 					}
// 					jwtServices, errorInfo.Error = jwt2.NewWithClaims(jwt2.SigningMethodRS512, jwt2.MapClaims{
// 						"requesterId": requesterId,
// 						"Issuer":      ctv.CERT_ISSUER,
// 						"Subject":     requesterId,
// 						"ExpiresAt":   time.Now().Add(tDuration).String(),
// 						"NotBefore":   time.Now(),
// 					}).SignedString(tPrivateKey)
// 				}
// 			}
// 		}
// 	}
//
// 	return
// }

// GenerateRSAKey - generates an RSA key pair with a specified number of bits.
// The function returns the private and public keys. If an error occurs during key generation,
// the error is returned in the errorInfo parameter.
//
// Parameters:
// - rsaBits: The number of bits for the RSA key.
//
// Returns:
// - privateKey: The generated RSA private key.
// - publicKey: The corresponding RSA public key.
// - errorInfo: An error information object containing details about any errors that occurred.
//
// Customer Messages: None
// Errors: Any error during key generation.
// Verifications: None
func GenerateRSAKey(rsaBits int) (
	privateKey crypto.PrivateKey,
	publicKey crypto.PublicKey,
	errorInfo errs.ErrorInfo,
) {

	var (
		_PrivateKey *rsa.PrivateKey
	)

	if _PrivateKey, errorInfo.Error = rsa.GenerateKey(rand.Reader, rsaBits); errorInfo.Error != nil {
		log.Println(errorInfo.Error)
	}

	if errorInfo.Error == nil {
		// The public key is a part of the *rsa.PrivateKey struct
		publicKey = _PrivateKey.Public()
		privateKey = _PrivateKey
	}

	return
}

// GenerateSymmetricKey - generates a Symmetric key with 32 bytes.
//
// Returns:
// - SymmetricKey: The generated symmetric key Base64 encoded
//
// Customer Messages: None
// Errors: Any error during key generation.
// Verifications: None
func GenerateSymmetricKey() (
	SymmetricKey string,
) {

	var (
		errorInfo errs.ErrorInfo
		seed      int64
		tKey      []byte
	)

	seed = time.Now().UnixNano() + int64(runtime.NumCPU())

	tKey = make([]byte, 32)
	binary.LittleEndian.PutUint64(tKey, uint64(seed))
	if _, errorInfo.Error = rand.Reader.Read(tKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%s%s", ctv.LBL_JWT_SYMMETRIC_KEY, ctv.TXT_SERVICE_FAILED))
		errs.PrintErrorInfo(errorInfo)
	}

	SymmetricKey = base64.StdEncoding.EncodeToString(tKey)

	return
}

// ParsePrivateKey - parses the provided raw private key in PEM format.
// If the parsing is successful, it returns the parsed private key.
// Otherwise, it returns an error indicating the failure.
//
// Customer Messages: None
// Errors: None
// Verifications: None
func ParsePrivateKey(tRawPrivateKey []byte) (
	privateKey *rsa.PrivateKey,
	errorInfo errs.ErrorInfo,
) {

	if privateKey, errorInfo.Error = jwt.ParseRSAPrivateKeyFromPEM(tRawPrivateKey); errorInfo.Error != nil {
		errorInfo.Error = errors.New("Unable to parse the private key referred to in the configuration file.")
		log.Println(errorInfo.Error)
	}

	return
}

// RemoveTLSTemporaryFiles - removes the temporary CA Bundle, Certificate, and Private Key files.
//
//	Customer Messages: None
//	Errors: Return from RemoveFile
//	Verifications: None
func RemoveTLSTemporaryFiles(
	tempDirectory string,
) (errorInfo errs.ErrorInfo) {

	if errorInfo = hlp.RemoveFile(fmt.Sprintf("%v/tls-ca-bundle.crt", tempDirectory)); errorInfo.Error == nil {
		if errorInfo = hlp.RemoveFile(fmt.Sprintf("%v/tls-ca-cert.crt", tempDirectory)); errorInfo.Error == nil {
			if errorInfo = hlp.RemoveFile(fmt.Sprintf("%v/tls-private.key", tempDirectory)); errorInfo.Error == nil {
			}
		}
	}

	return
}

// Private Functions

// CheckEncryptDecryptParameters - Checks the required parameter for Encrypt and Decrypt functions.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func checkEncryptDecryptParameters(internalUserID string, keyB64 string, value string) (errorInfo errs.ErrorInfo) {

	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_JWT, internalUserID, ctv.FN_INTERNAL_USER_ID); errorInfo.Error != nil {
		return
	}
	if errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_JWT, keyB64, ctv.FN_KEY_B64); errorInfo.Error != nil {
		return
	}
	errorInfo = vldts.CheckValueNotEmpty(ctv.LBL_SERVICE_JWT, value, ctv.FN_VALUE_B64)

	return
}
