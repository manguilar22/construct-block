package main

import (
	"fmt"
	"github.com/manguilar22/construct-block/client"
	"github.com/manguilar22/construct-block/utils/oop"
	"os"
)

var version = os.Getenv("VERSION")
var hostname = os.Getenv("HOSTNAME")
var ipcFile = os.Getenv("IPCFILE")
var publicKey = os.Getenv("PUBLIC_KEY")
var hostKeyStoreFilePath = os.Getenv("KEYSTORE")
var privateKey = os.Getenv("PRIVATE_KEY")
var keyStorePassword = os.Getenv("PASSWORD")

func main(){
	//deployHTTPNetwork(hostname)
	deployIPC(hostKeyStoreFilePath, ipcFile)
}

func deployIPC(wallet string, ipcFile string) {
	// Private Keystore and Host IPC
	_, ks, _, _ := network.ConnectEthIPC(ipcFile, wallet)
	network.UnlockAllAccounts(ks, keyStorePassword)
	str := network.GetAllAccounts(ks)
	fmt.Println(str)

	network.SigningTransaction(ks, keyStorePassword)
}

func deployHTTPNetwork(hostName string) {
	client, _ := network.ConnectEthHttp(hostName)

	data := network.GetAllBlocks(client)
	fmt.Println(".........Loading")
	fmt.Println(len(data))

	// Blockchain Struct
	b1 := oop.Blockchain{Data: data}
	transactionBlocks := b1.GetTransactionBlocks()

	fmt.Println(transactionBlocks)

	//client, _ := network.ConnectEthHttp(hostName)

	fmt.Println(data)
}




