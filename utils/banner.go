package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// BlockDataTextBanner show all the properties of a individual block struct
func BlockDataTextBanner(block *types.Block) {
	fmt.Printf("Block Number: %d\n",block.Number().Uint64())
	fmt.Printf("Block Time: %d\n",block.Time())
	fmt.Printf("Mining Difficulty: %d\n", block.Difficulty().Uint64())
	fmt.Printf("Hash: %s\n",block.Hash().Hex())
	fmt.Printf("Block TransactionsCount: %d\n",len(block.Transactions()))
}

// PrintTransactions view all transactions in the blockchain (Text)
func PrintTransactions(data []*types.Block) {
	for e := range data {
		for errorCode, tx := range data[e].Transactions() {
			if errorCode != 1 {
				// No Error occurs.
				fmt.Printf("Hash: %s\n",tx.Hash().Hex())
				fmt.Printf("Value: %s\n",tx.Value().String())
				fmt.Printf("Gas: %s\n",tx.Gas())
				fmt.Printf("Gas Price: %s\n",tx.GasPrice().Uint64())
				fmt.Printf("Nonce: %s\n",tx.Nonce())
				fmt.Printf("BlockData: %+v\n",tx.Data())
				//fmt.Println(tx.To().Hex())
			}
		}
	}
}

// ViewAllBlocks view all blocks in the connected blockchain. (Text)
func ViewAllBlocks(client *ethclient.Client) {

	totalBlocks, _ := client.BlockByNumber(context.Background(), nil)

	var length = totalBlocks.Number().Int64()

	for i := 0; i < int(length); i++ {
		var iterator = big.NewInt(int64(i))
		blockN, err := client.BlockByNumber(context.Background(),iterator)
		if err != nil {
			log.Fatalf("management: Problem retrieving block number %d of %d\n", iterator, length)
		}

		count, _ := client.TransactionCount(context.Background(), blockN.Hash())
		BlockDataTextBanner(blockN)
		fmt.Printf("Transaction Count: %d\n", count)

	}
}

