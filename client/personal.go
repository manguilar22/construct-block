package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/manguilar22/construct-block/client/oop"
	"log"
	"math/big"
)

// UnlockAllAccounts Unlock Accounts on Host Machine. (RPC)
func UnlockAllAccounts(ks *keystore.KeyStore, password string) {
	hostAccounts := ks.Accounts()
	for e := range hostAccounts {
		ks.Unlock(hostAccounts[e], password)
		fmt.Printf("Unlocking Account: %s\n", hostAccounts[e].Address.String())
	}
}

// CreateAccount Create Account on Host (RPC)
func CreateAccount(ks *keystore.KeyStore, password string) (accounts.Account) {
	createAccount, _ := ks.NewAccount(password)
	return createAccount
}

// SigningTransaction Signing from Go (RPC)
func SigningTransaction(ks *keystore.KeyStore, signingAddress string, password string) {
	// Create a new account to sign transactions with
	//signer, _ := ks.NewAccount(password)
	signer := accounts.Account{Address: common.HexToAddress(signingAddress)}
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

// GetAddressBalance get balance in term of wei (RPC & HTTP)
func GetAddressBalance(client *ethclient.Client, address string) (string, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "0", err
	}

	return balance.String(), nil
}

func SendEther(client *ethclient.Client, privateKeyMAC string, to string, wei int64) {
	privateKey, err := crypto.HexToECDSA(privateKeyMAC)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("personal: %s sending a transaction", fromAddress.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("transaction error when creating Nonce: %+v\n", err)
	}

	value := big.NewInt(wei) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice := big.NewInt(1000000000)//, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, _ := client.NetworkID(context.Background())

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("%+v = %d + (%d * %d)",err, value, gasLimit, gasPrice)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func SendEtherRaw(client *ethclient.Client, privateKeyMAC string, to string, wei int64) {
	privateKey, err := crypto.HexToECDSA(privateKeyMAC)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(wei)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funds: value + (gas * gasPrice) = %d + (%d * %d)", value, gasLimit, gasPrice)
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	rawTxBytes, err = hex.DecodeString(rawTxHex)

	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}