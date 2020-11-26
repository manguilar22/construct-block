package main

import (
	"context"
	"fmt"
	"github.com/manguilar22/construct-block/client"
	"log"
	"os"
)

/*
var version = os.Getenv("VERSION")
var hostname = os.Getenv("HOSTNAME")
var ipcFile = os.Getenv("IPC_FILE")
var publicKey = os.Getenv("PUBLIC_KEY")
var hostEthDataDir = os.Getenv("DATADIR")
var hostKeyStoreFilePath = os.Getenv("KEYSTORE")
var privateKey = os.Getenv("PRIVATE_KEY")
var keyStorePassword = os.Getenv("PASSWORD_FILE")
*/

func main(){
	//deployHTTPNetwork(hostname)
	var hostKeyStoreFilePath = os.Getenv("KEYSTORE")
	var ipcFile = os.Getenv("IPC_FILE")
	anotherClient, _, _,_ := network.ConnectEthIPC(ipcFile, hostKeyStoreFilePath)
	//client, _ := ethclient.Dial(ipcFile)
	header, err := anotherClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.Int64())
}

func deployIPC(wallet string, ipcFile string) {
	/*
	// Start Ethereum Node (No cmd.Wait() because we want to background process)
	cmd := exec.Command("./scripts/startU18.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatalf("main: Failed running *.sh file because %+v.\n",err)
	}
	fmt.Printf("Ethereum node process PID: %d\n", cmd.Process.Pid)
	*/

	// Private Keystore and Host IPC
	client, _, _, _ := network.ConnectEthIPC(ipcFile, wallet)
	fmt.Println(client)
	b, _ := client.BlockByNumber(context.Background(), nil)
	fmt.Println(b.Number().Int64())
	//fmt.Println(ks.Accounts())

}





