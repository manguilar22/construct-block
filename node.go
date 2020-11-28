package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/manguilar22/construct-block/client"
	"github.com/manguilar22/construct-block/server"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var version = os.Getenv("VERSION")
var hostname = os.Getenv("HOSTNAME")
var publicKey = os.Getenv("PUBLIC_KEY")
var hostEthDataDir = os.Getenv("DATADIR")
var hostKeyStoreFilePath = os.Getenv("KEYSTORE")
var privateKey = os.Getenv("PRIVATE_KEY")
var keyStorePassword = os.Getenv("PASSWORD_FILE")

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./public/*.gohtml"))
}

func main(){

	// Private Keystore and Host IPC
	client, ks, _, _ := network.ConnectEthIPC(hostname, hostKeyStoreFilePath)
	defer client.Close()

	// Get All Blocks
	allBlocks := network.GetAllBlocks(client)
	// Get all blocks with tx
	allTxBlocks := network.SearchTxBlocks(client)

	// Server Routes
	router := mux.NewRouter()

	router.HandleFunc("/", server.GetEthereumStatusIndex(hostname, client, tpl))
	router.HandleFunc("/blockchain", server.GetBlockchainData(client,hostname))

	router.HandleFunc("/accounts", server.GetAllAccounts(hostKeyStoreFilePath))
	router.HandleFunc("/accounts/balance",server.GetAllAccountBalance(client,ks))
	router.HandleFunc("/accounts/id/{num}", server.GetAccountById(hostKeyStoreFilePath))
	router.HandleFunc("/accounts/{account}", server.FindByAccount(hostKeyStoreFilePath))
	router.HandleFunc("/accounts/{account}/balance",server.GetBalanceByAccount(client))

	router.HandleFunc("/eth", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, ks.Accounts())
	})
	router.HandleFunc("/eth/blocks", func(writer http.ResponseWriter, request *http.Request) {
		data := allBlocks
		fmt.Fprintln(writer, data)
	})
	router.HandleFunc("/eth/blocks/tx", func(writer http.ResponseWriter, request *http.Request) {
		data := allTxBlocks
		fmt.Fprintln(writer, data)
	})

	srv := &http.Server{
		Handler: router,
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}
