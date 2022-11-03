package signature

const (
	PubKeyKey          = "publicKey"
	PvtKeyKey          = "privateKey"
	PubKeyNamespaceKey = "publicKeyNamespace"
	PubKeyNameKey      = "publicKeyName"

	keyBitSize = 2048

	// Headers.
	DigestHeader    = "Digest"
	SignatureHeader = "Signature"

	// Signature String Construction.
	headerFieldDelimiter = ": "
	headersDelimiter     = "\n"

	// Signature Parameters.
	createdParameter               = "created"
	pubKeySecretNameParameter      = "pubKeySecretName"
	pubKeySecretNamespaceParameter = "pubKeySecretNamespace"
	signatureParameter             = "Signature"
	parameterKVSeparater           = "="
	parameterValueDelimiter        = "\""
	parameterSeparater             = ","
)

// TODO include expires header.
var defaultHeaders = []string{createdParameter} //nolint:gochecknoglobals
