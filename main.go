package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/manguilar22/construct-block/client"
	"github.com/manguilar22/construct-block/client/oop"
	"github.com/manguilar22/construct-block/server"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var version = os.Getenv("VERSION")
var hostname = os.Getenv("HOSTNAME")
var ipcFile = os.Getenv("IPC_FILE")
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
	client, ks, _, _ := network.ConnectEthIPC(ipcFile, hostKeyStoreFilePath)
	defer client.Close()

	// Pre-Load Data (HOST)
	allAccountData := server.LoadAllData(hostKeyStoreFilePath)
	allAccountDataJson := server.MakeEthereumAccountList(hostKeyStoreFilePath)
	// Pre-Load Data (CLIENT)
	// Get All Blocks
	allBlocks := network.GetAllBlocks(client)
	// Get all blocks with tx
	allTxBlocks := network.SearchTxBlocks(client)

	// Template Data
	blockDataHTML := network.MakeBlocks(allBlocks)
	txDataHTML := network.MakeTransactions(allTxBlocks)
	blockchainData := oop.Blockchain{Hostname:hostname,BlockData: blockDataHTML,TransactionData: txDataHTML}

	// Go-Server
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		tpl.Execute(writer, blockchainData)
	})
	router.HandleFunc("/blockchain", func(writer http.ResponseWriter, request *http.Request) {
		jsonB, _ := json.Marshal(blockchainData)
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonB)
	})

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

	router.HandleFunc("/eth/tx/receipt", func(writer http.ResponseWriter, request *http.Request) {
		
	})
	
	router.HandleFunc("/accounts", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(allAccountDataJson)
	})
	router.HandleFunc("/accounts/balance", func(writer http.ResponseWriter, request *http.Request) {
		allBalances := make(map[string]string)
		for _, accountHex := range ks.Accounts() {
			balance, _ := client.BalanceAt(context.Background(), accountHex.Address, nil)
			allBalances[accountHex.Address.Hex()] = balance.String()
		}
		jsonB, _ := json.Marshal(allBalances)
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonB)
	})
	router.HandleFunc("/accounts/{account}", func(writer http.ResponseWriter, request *http.Request) {
		r := mux.Vars(request)
		value := r["account"]
		jsonBytes, _ := json.Marshal(allAccountData[value])

		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonBytes)
	})
	router.HandleFunc("/accounts/id/{num}", func(writer http.ResponseWriter, request *http.Request) {
		r := mux.Vars(request)
		value := r["num"]
		accountHex := make(map[int]string)
		counter := 0

		for k,_ := range allAccountData {
			accountHex[counter] = k
			counter += 1
		}

		num, _ := strconv.Atoi(value)
		accountNum := accountHex[num]

		fmt.Printf("main: num=%d, value=%s, accountNum=%+v\n",num,value, accountNum)

		jsonBytes, _ := json.Marshal(allAccountData[accountNum])
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonBytes)
	})

	srv := &http.Server{
		Handler: router,
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}
