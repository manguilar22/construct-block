package utils

import (
	"encoding/json"
	"fmt"
	"github.com/manguilar22/construct-block/client/oop"
	"io/ioutil"
	"log"
	"os"
)

// ImportBlockchainJsonData import blockchain data from a json file.
func ImportBlockchainJsonData(filepath string) (data []oop.EthBlock){
	jsonB, _ := ioutil.ReadFile(filepath)
	err := json.Unmarshal(jsonB, &data)
	if err != nil {
		log.Fatalln("utils: Error importing Blockchain data.")
	}
	return
}

// SaveDataToJson save block data to json file.
func SaveDataToJson(data []oop.EthBlock, fileName string) {
	file, _ := os.Create(fileName)
	defer file.Close()

	jsonB, _ := json.Marshal(data)
	saveLine := string(jsonB)
	file.Write([]byte(saveLine))

	fmt.Println("Utils: Creating File Complete (SaveDataToJson)")
}

// SaveDataJsonPermission save block data to json file with OS file permission.
func SaveDataJsonPermission(data []oop.EthBlock, fileName string, osFileMode os.FileMode) {

	var result string = ""

	for e := range data {
		jsonS, _ := json.Marshal(data[e])

		result += string(jsonS)+","

	}

	err := ioutil.WriteFile(fileName, []byte(result), osFileMode)

	if err != nil {
		log.Fatalln("utils: Failed Writing to File")
	}
	fmt.Println("Utils: Creating File Complete (SaveDataJsonPermission)")

}

