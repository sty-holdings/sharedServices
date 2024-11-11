package sharedServices

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"

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

// Decrypt - decrypts an encrypted string using AES-GCM mode of operation.
// The function decodes the encryptedMessage from base64 and initializes the AES cipher
// with the provided key. It then creates an AEAD object using the AES cipher and extracts
// the nonce and ciphertext from the decoded message. Finally, it decrypts the ciphertext
// using the AEAD object and returns the decrypted message.
//
// Customer Messages: None
// Errors:
// - ErrInvalidBase64: If the encryptedMessage could not be decoded from base64.
// - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// - ErrAEADInitialization: If the AEAD object could not be created.
// - ErrDecryption: If the ciphertext could not be decrypted.
// Verifications: None
func Decrypt(
	username string,
	key string,
	encryptedMessage string,
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

	if len(encryptedMessage) == ctv.VAL_ZERO {
		decryptedMessage = encryptedMessage
		return
	}

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(key); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, username))
		return
	}

	if tCiphertext, errorInfo.Error = base64.StdEncoding.DecodeString(encryptedMessage); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, username))
		return
	}
	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, username))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, username))
		return
	}

	tNonceSize = tAESGCM.NonceSize()
	tNonce, tCiphertext = tCiphertext[:tNonceSize], tCiphertext[tNonceSize:]

	if tPlaintext, errorInfo.Error = tAESGCM.Open(nil, tNonce, tCiphertext, nil); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_USERNAME, username))
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
	clientId string,
	key string,
	encryptedMessage string,
) (
	decryptedMessage []byte,
	errorInfo errs.ErrorInfo,
) {

	var (
		tDecryptedMessage string
	)

	if tDecryptedMessage, errorInfo = Decrypt(clientId, key, encryptedMessage); errorInfo.Error == nil {
		decryptedMessage = []byte(tDecryptedMessage)
	}

	return
}

// Encrypt - encrypts a message using AES-GCM mode of operation.
// // The function decodes the encryptedMessage from base64 and initializes the AES cipher
// // with the provided key. It then creates an AEAD object using the AES cipher and extracts
// // the nonce and ciphertext from the decoded message. Finally, it decrypts the ciphertext
// // using the AEAD object and returns the decrypted message.
// //
// // Customer Messages: None
// // Errors:
// // - ErrInvalidBase64: If the encryptedMessage could not be decoded from base64.
// // - ErrCipherInitialization: If the AES cipher could not be initialized with the provided key.
// // - ErrAEADInitialization: If the AEAD object could not be created.
// // - ErrDecryption: If the ciphertext could not be decrypted.
// // Verifications: None
func Encrypt(
	clientId string,
	key string,
	message string,
) (
	encryptedMessage string,
	errorInfo errs.ErrorInfo,
) {

	var (
		tAESGCM     cipher.AEAD
		tBlock      cipher.Block
		tCiphertext []byte
		tDecodedKey []byte
		tNonce      []byte
	)

	if tDecodedKey, errorInfo.Error = base64.StdEncoding.DecodeString(key); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_CLIENT_ID, clientId))
		return
	}

	if tBlock, errorInfo.Error = aes.NewCipher(tDecodedKey); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_CLIENT_ID, clientId))
		return
	}

	if tAESGCM, errorInfo.Error = cipher.NewGCM(tBlock); errorInfo.Error != nil {
		errorInfo = errs.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.FN_CLIENT_ID, clientId))
		return
	}

	tNonce = make([]byte, tAESGCM.NonceSize())
	io.ReadFull(rand.Reader, tNonce) // Use a cryptographically secure RNG

	tCiphertext = tAESGCM.Seal(tNonce, tNonce, []byte(message), nil)
	encryptedMessage = base64.StdEncoding.EncodeToString(tCiphertext)

	return
}

// GenerateJWT
// Create a new token object, specifying signing method and the claims
// you would like it to contain.
// func GenerateJWT(privateKey, requestorId, period string, duration int64) (jwtServices string, errorInfo errs.ErrorInfo) {
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
// 		if requestorId == ctv.VAL_EMPTY || period == ctv.VAL_EMPTY || duration < 1 {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v' %v: '%v'", ctv.FN_REQUESTOR_ID, requestorId, ctv.FN_PERIOD, period, ctv.FN_DURATION, duration))
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
// 						"requestorId": requestorId,
// 						"Issuer":      ctv.CERT_ISSUER,
// 						"Subject":     requestorId,
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
