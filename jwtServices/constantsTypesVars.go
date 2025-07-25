package sharedServices

import (
	"crypto"
)

//goland:noinspection ALL
const (
	CERT_ED25519             = "ED25519"
	CERT_RSA                 = "RSA"
	CERT_RS265               = "RS256"
	CERT_ECDSACURVE          = "ECDSACURVE"
	CERT_ECDSACURVE_P224     = "P224"
	CERT_ECDSACURVE_P256     = "P256"
	CERT_ECDSACURVE_P384     = "P384"
	CERT_ECDSACURVE_P521     = "P521"
	CERT_PRIVATE_KEY         = "RSA PRIVATE KEY"
	CERT_PUBLIC_KEY          = "PUBLIC KEY"
	CERTIFICATE              = "CERTIFICATE"
	TLS_CA_BUNDLE_FILENAME   = "tls-ca-bundle.crt"
	TLS_CERT_FILENAME        = "tls-cert.crt"
	TLS_PRIVATE_KEY_FILENAME = "tls-private.key"
	// Parameters
	PARAMETER_TLS_CERT            = "tls-certificate"
	PARAMETER_TLS_CERT_FQN        = "tls-certificate-fqn"
	PARAMETER_TLS_PRIVATE_KEY     = "tls-private-key"
	PARAMETER_TLS_PRIVATE_KEY_FQM = "tls-private-key-fqn"
	PARAMETER_TLS_CA_BUNDLE       = "tls-ca-bundle"
	PARAMETER_TLS_CA_BUNDLE_FQN   = "tls-ca-bundle-fqn"
	// Toekn
	TOKEN_TYPE_ID      = "id"
	TOKEN_TYPE_ACCESS  = "access"
	TOKEN_TYPE_REFRESH = "refresh"
)

//goland:noinspection ALL
const (
	JWT_PAYLOAD_SUBJECT_FN      = "SUBJECT"
	JWT_PAYLOAD_CLAIMS_FN       = "CLAIMS"
	JWT_PAYLOAD_AUDIENCE_FN     = "AUDIENCE"
	JWT_PAYLOAD_REQUESTOR_ID_FN = "REQUESTOR_ID"
	JWT_PAYLOAD_EXPIRES_FN      = "EXPIRES"
	JWT_PAYLOAD_ISSUER_FN       = "ISSUER"
	JWT_PAYLOAD_ISSUED_AT_FN    = "ISSUED_AT"
)

type GenerateCertificate struct {
	CertFileName       string
	Certificate        []byte
	Host               string
	PublicKey          crypto.PublicKey
	PrivateKey         crypto.PrivateKey
	PrivateKeyFileName string
	RSABits            int
	SelfCA             bool
	ValidFor           string
}

// TLSInfo files
type TLSInfo struct {
	TLSCert          string `json:"tls_certificate" yaml:"tls_certificate"`
	TLSCertFQN       string `json:"tls_certificate_fqn" yaml:"tls_certificate_fqn"`
	TLSPrivateKey    string `json:"tls_private_key" yaml:"tls_private_key"`
	TLSPrivateKeyFQN string `json:"tls_private_key_fqn" yaml:"tls_private_key_fqn"`
	TLSCABundle      string `json:"tls_ca_bundle" yaml:"tls_ca_bundle"`
	TLSCABundleFQN   string `json:"tls_ca_bundle_fqn" yaml:"tls_ca_bundle_fqn"`
}
