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

	ctv "github.com/sty-holdings/sharedServices/v2024/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
	hlp "github.com/sty-holdings/sharedServices/v2024/helpers"
)

// BuildTLSTemporaryFiles - creates temporary files for TLS information.
// The function checks if the TLSCABundle, TLSCert, and TLSPrivateKey in tlsInfo are provided. If any of these values are empty,
// the function returns an error indicating the missing.
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, returned from WriteFile
//	Verifications: None
func BuildTLSTemporaryFiles(
	tempDirectory string,
	tlsInfo TLSInfo,
) (
	errorInfo errs.ErrorInfo,
) {

	if tlsInfo.TLSCABundle == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.LBL_MISSING_PARAMETER, ctv.FN_TLS_CA_BUNDLE))
		return
	} else {
		if errorInfo = hlp.WriteFile(fmt.Sprintf("%v/%v", tempDirectory, TLS_CA_BUNDLE_FILENAME), []byte(tlsInfo.TLSCABundle), 0744); errorInfo.Error != nil {
			return
		}
	}
	if tlsInfo.TLSCert == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.LBL_MISSING_PARAMETER, ctv.FN_TLS_CERTIFICATE))
		return
	} else {
		if errorInfo = hlp.WriteFile(fmt.Sprintf("%v/%v", tempDirectory, TLS_CERT_FILENAME), []byte(tlsInfo.TLSCert), 0744); errorInfo.Error != nil {
			return
		}
	}
	if tlsInfo.TLSPrivateKey == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", ctv.LBL_MISSING_PARAMETER, ctv.FN_TLS_PRIVATE_KEY))
		return
	} else {
		if errorInfo = hlp.WriteFile(fmt.Sprintf("%v/%v", tempDirectory, TLS_PRIVATE_KEY_FILENAME), []byte(tlsInfo.TLSPrivateKey), 0744); errorInfo.Error != nil {
			return
		}
	}

	return
}

// Decrypt - decrypts an encrypted base64 string using AES-GCM mode of operation.
//
// Customer Messages: None
// Errors:
// - ErrInvalidBase64: If the encryptedMessage could not be decoded from base64.
// - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// - ErrAEADInitialization: If the AEAD object could not be created.
// - ErrDecryption: If the ciphertext could not be decrypted.
// Verifications: None
func Decrypt(
	uId string,
	keyB64 string,
	encryptedMessageB64 string,
) (
	decryptedMessage string,
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

	if encryptedMessageB64 == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%s%s", ctv.LBL_MESSAGE, ctv.TXT_IS_MISSING))
		return
	}
	if keyB64 == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%s%s", ctv.LBL_SECRET_KEY, ctv.TXT_IS_MISSING))
		return
	}
	if uId == ctv.VAL_EMPTY {
		errorInfo = errs.NewErrorInfo(errs.ErrRequiredArgumentMissing, fmt.Sprintf("%s%s", ctv.LBL_USERNAME, ctv.TXT_IS_MISSING))
		return
	}

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(keyB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, uId))
		return
	}

	if tCiphertext, errorInfo.Error = base64.StdEncoding.DecodeString(encryptedMessageB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, uId))
		return
	}
	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, uId))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, uId))
		return
	}

	tNonceSize = tAESGCM.NonceSize()
	tNonce, tCiphertext = tCiphertext[:tNonceSize], tCiphertext[tNonceSize:]

	if tPlaintext, errorInfo.Error = tAESGCM.Open(nil, tNonce, tCiphertext, nil); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, uId))
		return
	}

	decryptedMessage = string(tPlaintext)

	return
}

// DecryptToByte - will call decrypt and convert the return string to []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func DecryptToByte(
	uId string,
	keyB64 string,
	encryptedMessageB64 string,
) (
	decryptedMessage []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tDecryptedMessage string
	)

	if tDecryptedMessage, errorInfo = Decrypt(uId, keyB64, encryptedMessageB64); errorInfo.Error == nil {
		decryptedMessage = []byte(tDecryptedMessage)
	}

	return
}

// Encrypt - encrypts a message using AES-GCM mode of operation.
// The function encrypts message and initializes the AES cipher
// with the provided key.
//
// Returns:
// - encryptedMessage: base64 encoded string
//
// Customer Messages: None
// Errors:
// - ErrInvalidBase64: If the encryptedMessage could not be decoded from base64.
// - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// - ErrAEADInitialization: If the AEAD object could not be created.
// - ErrDecryption: If the ciphertext could not be decrypted.
// Verifications: None
func Encrypt(
	uId string,
	keyB64 string,
	message string,
) (
	encryptedMessageB64 string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tAESGCM     cipher.AEAD
		tBlock      cipher.Block
		tCiphertext []byte
		tDecodedKey []byte
		tNonce      []byte
	)

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(keyB64); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_USERNAME, uId))
		return
	}

	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_USERNAME, uId))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.LBL_USERNAME, uId))
		return
	}

	tNonce = make([]byte, tAESGCM.NonceSize())
	//goland:noinspection GoUnhandledErrorResult
	io.ReadFull(rand.Reader, tNonce) // Use a cryptographically secure RNG

	tCiphertext = tAESGCM.Seal(tNonce, tNonce, []byte(message), nil)
	encryptedMessageB64 = base64.StdEncoding.EncodeToString(tCiphertext)

	return
}

// EncryptToByte - will call encrypt and convert the return base 64 string to []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func EncryptToByte(
	uId string,
	keyB64 string,
	message string,
) (
	encryptedMessageB64 []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tEncryptedMessageB64 string
	)

	if tEncryptedMessageB64, errorInfo = Encrypt(uId, keyB64, message); errorInfo.Error == nil {
		encryptedMessageB64 = []byte(tEncryptedMessageB64)
	}

	return
}

// EncryptFromByteToByte - will call encrypt the []byte and return a base 64 []byte
//
//	Customer Messages: None
//	Errors: returned by Decrypt
//	Verifications: None
func EncryptFromByteToByte(
	uId string,
	keyB64 string,
	message []byte,
) (
	encryptedMessageB64 []byte,
	errorInfo errs.ErrorInfo,
) {

	encryptedMessageB64, errorInfo = EncryptToByte(uId, keyB64, string(message))

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
// - SymmetricKey: The generated symmetric key base64 encoded
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
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%s%s", ctv.LBL_SYMMETRIC_KEY, ctv.TXT_SERVICE_FAILED))
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
