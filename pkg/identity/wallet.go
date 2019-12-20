package identity

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	*hdwallet.Wallet
}

// NewWallet creates new HD wallet
func NewWallet(mnemonic string) (*Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return &Wallet{wallet}, nil
}

// GetAccount returns Ethereum account
func (w *Wallet) GetAccount(keyIndex int) (accounts.Account, error) {
	pathStr := fmt.Sprintf("m/44'/60'/0'/0/%d", keyIndex)
	path := hdwallet.MustParseDerivationPath(pathStr)
	return w.Derive(path, true)
}
