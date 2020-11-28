package server

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/manguilar22/construct-block/client"
	"github.com/manguilar22/construct-block/client/oop"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GetEthereumStatusIndex(hostname string, client *ethclient.Client,tpl *template.Template) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")
		// Get All Blocks
		allBlocks := network.GetAllBlocks(client)
		// Get all blocks with tx
		allTxBlocks := network.SearchTxBlocks(client)

		// Template Data
		blockDataHTML := network.MakeBlocks(allBlocks)
		txDataHTML := network.MakeTransactions(allTxBlocks)
		blockchainData := oop.Blockchain{Hostname:hostname,BlockData: blockDataHTML,TransactionData: txDataHTML}
		tpl.Execute(writer, blockchainData)
	}
}

func GetBlockchainData(client *ethclient.Client, hostname string) http.HandlerFunc {
	// Get All Blocks
	allBlocks := network.GetAllBlocks(client)
	// Get all blocks with tx
	allTxBlocks := network.SearchTxBlocks(client)
	// Template Data
	blockDataHTML := network.MakeBlocks(allBlocks)
	txDataHTML := network.MakeTransactions(allTxBlocks)
	blockchainData := oop.Blockchain{Hostname:hostname,BlockData: blockDataHTML,TransactionData: txDataHTML}

	return func(writer http.ResponseWriter, request *http.Request) {
		jsonB, _ := json.Marshal(blockchainData)
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonB)
	}
}

func GetAllAccounts(hostKeyStoreFilePath string) http.HandlerFunc {
	allAccountDataJson := MakeEthereumAccountList(hostKeyStoreFilePath)

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(allAccountDataJson)
	}
}

func GetAllAccountBalance(client *ethclient.Client, ks *keystore.KeyStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		allBalances := make(map[string]string)
		for _, accountHex := range ks.Accounts() {
			balance, _ := client.BalanceAt(context.Background(), accountHex.Address, nil)
			allBalances[accountHex.Address.Hex()] = balance.String()
			log.Printf("Getting Balance of account %s and has %d wei", accountHex.Address.String(), balance)
		}
		jsonB, _ := json.Marshal(allBalances)
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonB)
	}
}

func FindByAccount(hostKeyStoreFilePath string) http.HandlerFunc {
	allAccountData := LoadAllData(hostKeyStoreFilePath)

	return  func(writer http.ResponseWriter, request *http.Request) {
		r := mux.Vars(request)
		value := r["account"]
		jsonBytes, _ := json.Marshal(allAccountData[value])

		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonBytes)
	}
}

func GetBalanceByAccount(client *ethclient.Client) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		r := mux.Vars(request)
		value := r["account"]
		//jsonBytes, _ := json.Marshal(allAccountData[value])
		currentBalance , _ := client.BalanceAt(context.Background(),common.HexToAddress("0x"+value),nil)

		writer.Header().Set("Content-Type","application/json")
		writer.Write([]byte(currentBalance.String()))
	}
}

func GetAccountById(hostKeyStoreFilePath string) http.HandlerFunc {
	allAccountData := LoadAllData(hostKeyStoreFilePath)

	return func(writer http.ResponseWriter, request *http.Request) {
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

		log.Printf("main: num=%d, value=%s, accountNum=%+v\n",num,value, accountNum)

		jsonBytes, _ := json.Marshal(allAccountData[accountNum])
		writer.Header().Set("Content-Type","application/json")
		writer.Write(jsonBytes)
	}
}

