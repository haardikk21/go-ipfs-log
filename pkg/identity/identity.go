package identity

import (
	"encoding/json"
	"errors"
)

// Identity represents an identity
type Identity struct {
	ID         string
	PublicKey  []byte
	Signatures *Signatures
	Type       string
	Provider   Provider
}

// Signatures is the collection of signatures used by the identity
type Signatures struct {
	ID        []byte `json:"id"`
	PublicKey []byte `json:"publicKey"`
}

// NewIdentity returns an instance of Identity
func NewIdentity(id string, publicKey []byte, signatures *Signatures, identityType string, provider Provider) (*Identity, error) {
	if id == "" {
		return nil, errors.New("missing field for identity: ID")
	}

	if publicKey == nil {
		return nil, errors.New("missing field for identity: Public Key")
	}

	if signatures == nil {
		return nil, errors.New("missing field for identity: Signatures")
	}

	if identityType == "" {
		return nil, errors.New("missing field for identity: Type")
	}

	if provider == nil {
		return nil, errors.New("missing field for identity: Provider")
	}

	identity := Identity{
		ID:         id,
		PublicKey:  publicKey,
		Signatures: signatures,
		Type:       identityType,
		Provider:   provider,
	}

	return &identity, nil
}

func (i *Identity) MarshalJSON() ([]byte, error) {
	var tmpIdentity struct {
		ID         string      `json:"id"`
		PublicKey  []byte      `json:"publicKey"`
		Signatures *Signatures `json:"signatures"`
		Type       string      `json:"type"`
	}

	tmpIdentity.ID = i.ID
	tmpIdentity.PublicKey = i.PublicKey
	tmpIdentity.Signatures = i.Signatures
	tmpIdentity.Type = i.Type

	return json.Marshal(tmpIdentity)
}
