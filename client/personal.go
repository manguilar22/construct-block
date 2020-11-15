package network

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/manguilar22/construct-block/utils/oop"
)

// Unlock Accounts on Host (IPC)
func UnlockAllAccounts(ks *keystore.KeyStore, password string) {
	hostAccounts := ks.Accounts()
	for e := range hostAccounts {
		ks.Unlock(hostAccounts[e], password)
	}
}

// Create Account on Host (IPC)
func CreateAccount(ks *keystore.KeyStore, password string) (accounts.Account) {
	createAccount, _ := ks.NewAccount(password)
	return createAccount
}


// Signing from Go
func SigningTransaction(ks *keystore.KeyStore, password string) {
	// Create a new account to sign transactions with
	signer, _ := ks.NewAccount(password)
	block := oop.EthBlock{}.Hash
	txHash := common.HexToHash(block)

	// Sign a transaction with a single authorization
	signature, _ := ks.SignHashWithPassphrase(signer, password, txHash.Bytes())

	// Sign a transaction with multiple manually cancelled authorizations
	_ = ks.Unlock(signer, password)
	signature, _ = ks.SignHash(signer, txHash.Bytes())
	_ = ks.Lock(signer.Address)

	/*
	// Sign a transaction with multiple automatically cancelled authorizations
	_ = ks.TimedUnlock(signer, password, time.Second)
	signature, _ = ks.SignHash(signer, txHash.Bytes())
	*/

	fmt.Println(signature)
}