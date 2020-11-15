package network

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ConnectEthHttp(hostName string) (client *ethclient.Client, err error){

	var url string = fmt.Sprintf("http://%s",hostName)

	client, err = ethclient.Dial(url)
	defer client.Close()

	if err != nil {
		log.Fatalln("Error Connecting Remotely")
	}

	return
}

func ConnectEthIPC(ipcFilePath string, walletKeystoreDirPath string) (client *ethclient.Client,ks *keystore.KeyStore, am *accounts.Manager, err error) {
	ks = keystore.NewKeyStore(walletKeystoreDirPath,keystore.StandardScryptN,keystore.StandardScryptP)
	am = accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)

	client,err = ethclient.Dial(ipcFilePath)
	defer client.Close()

	if err != nil {
		log.Fatalln("Backend: Not Connected to Client")
	}
	return
}
