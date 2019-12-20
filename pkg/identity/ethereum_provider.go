package identity

import (
	"errors"
)

type EthereumProvider struct {
	Wallet *Wallet
	Type   string
}

func NewEthereumProvider(wallet *Wallet) *EthereumProvider {
	return &EthereumProvider{
		Wallet: wallet,
		Type:   "ethereum",
	}
}

func (p *EthereumProvider) GetID() (string, error) {
	account, err := p.Wallet.GetAccount(0)
	if err != nil {
		return "", err
	}

	return account.Address.String(), nil
}

func (p *EthereumProvider) GetType() string {
	return p.Type
}

func (p *EthereumProvider) SignIdentity(data []byte, mimeType string) ([]byte, error) {
	if p.Wallet == nil {
		return nil, errors.New("wallet is required")
	}

	account, err := p.Wallet.GetAccount(0)
	if err != nil {
		return nil, err
	}

	return p.Wallet.SignData(account, mimeType, data)
}

func (p *EthereumProvider) VerifyIdentity(identity Identity) bool {
	return false
}
