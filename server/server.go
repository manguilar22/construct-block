package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var keystorePath = os.Getenv("KEYSTORE")

// loadAllData Loads all data from host system keystore
func LoadAllData(keystoreDirectoryPath string) (map[string]map[string]interface{}){
	ethereumDatadirAccounts, _ := ioutil.ReadDir(keystoreDirectoryPath)
	parsedData := make(map[string]map[string]interface{}) // Eth Address : Account Info
	for _ , e := range ethereumDatadirAccounts {
		parse := strings.Split(e.Name(),"--")
		f, _ := ioutil.ReadFile(keystoreDirectoryPath+"/"+e.Name())
		var anyJson map[string]interface{}
		err := json.Unmarshal(f, &anyJson)
		if err != nil {log.Printf("main: Problem reading ethereum account data from json. %+v\n", err)}

		parsedData[parse[2]] = anyJson
	}
	return parsedData
}

// MakeEthereumAccountList List All Accounts on the blockchain.
func MakeEthereumAccountList(keyStoreDirectoryPath string) (res []byte){
	allData := LoadAllData(keyStoreDirectoryPath)
	data, _ := json.Marshal(allData)
	res = data
	return
}
