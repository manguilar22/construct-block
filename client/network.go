package network

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// ConnectEthHTTP Connect to an ethereum node with HTTP
func ConnectEthHTTP(hostName string) (client *ethclient.Client, err error){

	var url  = fmt.Sprintf("http://%s",hostName)

	client, err = ethclient.Dial(url)

	if err != nil {
		log.Fatalln("Error Connecting Remotely")
	}
	return
}

// ConnectEthWS Connect to an ethereum node with WS
func ConnectEthWS(hostName string) (client *ethclient.Client, err error){

	var url  = fmt.Sprintf("ws://%s",hostName)

	client, err = ethclient.Dial(url)

	if err != nil {
		log.Fatalln("Error Connecting Remotely")
	}
	return
}

// ConnectEthIPC Connect to an ethereum node with IPC
func ConnectEthIPC(ipcFilePath string, walletKeystoreDirPath string) (client *ethclient.Client,ks *keystore.KeyStore, am *accounts.Manager, err error) {
	ks = keystore.NewKeyStore(walletKeystoreDirPath,keystore.StandardScryptN,keystore.StandardScryptP)
	am = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)

	client,err = ethclient.Dial(ipcFilePath)

	if err != nil {
		log.Fatalln("network: client not connected to ethereum network")
	}
	return
}
