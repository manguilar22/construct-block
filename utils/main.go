package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/manguilar22/construct-block/utils/oop"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
)

func WeiToEther(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance,big.NewFloat(math.Pow10(18)))
	fmt.Printf("Ether: %v\n",ethValue)
	return ethValue
}

func ImportJsonBlockData(filepath string) (data []oop.EthBlock){
	jsonB, _ := ioutil.ReadFile(filepath)
	err := json.Unmarshal(jsonB, &data)
	if err != nil {
		log.Fatalln("utils: Error importing Blockchain data.")
	}
	return
}

func BlockDataTextBanner(block *types.Block) {
	fmt.Printf("Block Number: %d\n",block.Number().Uint64())
	fmt.Printf("Block Time: %d\n",block.Time())
	fmt.Printf("Mining Difficulty: %d\n", block.Difficulty().Uint64())
	fmt.Printf("Hash: %s\n",block.Hash().Hex())
	fmt.Printf("Block TransactionsCount: %d\n",len(block.Transactions()))
}

func SaveDataJson2(data []oop.EthBlock, fileName string) {

	file, _ := os.Create(fileName)
	defer file.Close()

	for e := range data {
		jsonB, _ := json.Marshal(data[e])
		saveLine := string(jsonB) + ","
		file.Write([]byte(saveLine))
	}
	fmt.Println("Utils: Creating File Complete (SaveDataJson2)")
}

func SaveDataJson1(data []oop.EthBlock, fileName string) {

	var result string = ""

	for e := range data {
		jsonS, _ := json.Marshal(data[e])

		result += string(jsonS)+","

	}

	err := ioutil.WriteFile(fileName, []byte(result), 0644)

	if err != nil {
		log.Fatalln("utils: Failed Writing to File")
	}
	fmt.Println("Utils: Creating File Complete (SaveDataJson1)")

}

