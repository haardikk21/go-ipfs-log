package identity

// Provider is an interface for identities
type Provider interface {
	GetID(options string) (string, error)
	GetType() string

	SignIdentity(data []byte, mimeType string) ([]byte, error)
	VerifyIdentity(identity Identity) bool
}
