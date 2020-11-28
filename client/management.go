package network

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// GetAllAccounts Get All Accounts when connected to (RPC)
func GetAllAccounts(ks *keystore.KeyStore) (result string) {
	accs := ks.Accounts()
	result = "|Account Number|Address|Hash|URL|\n|:---|:---:|:---:|---:|"
	for idx, e := range accs {
		result += fmt.Sprintf("|%d|%s|%s|%s|\n", idx, e.Address.Hex(),e.Address.Hash().String(), e.URL.String())
	}
	return
}

// GetAllBlocks Get All Blocks (HTTP) (time complexity: O(n))
func GetAllBlocks(client *ethclient.Client) (blockData []*types.Block) {
	totalBlocks, _ := client.BlockByNumber(context.Background(), nil)
	var length = totalBlocks.Number().Int64()
	blockData = make([]*types.Block, length-length)

	for i := 0; i < int(length); i++ {
		var iterator = big.NewInt(int64(i))
		blockN, err := client.BlockByNumber(context.Background(),iterator)
		if err != nil {
			log.Fatalf("management: Problem retrieving block number %d of %d\n", iterator, length)
		}

		count, _ := client.TransactionCount(context.Background(), blockN.Hash())


		blockData = append(blockData, blockN)
		log.Printf("management: Block number %d of %d with a transaction count of %d\n",iterator, length, count)
	}
	return
}

// SearchTxBlocks Get All TX Blocks
func SearchTxBlocks(client *ethclient.Client) (blockData []*types.Block) {
	totalBlocks, _ := client.BlockByNumber(context.Background(), nil)
	var length = totalBlocks.Number().Int64()
	blockData = make([]*types.Block, length-length)

	for i := 0; i < int(length); i++ {
		var iterator = big.NewInt(int64(i))
		blockN, err := client.BlockByNumber(context.Background(),iterator)
		if err != nil {
			log.Fatalf("management: Problem retrieving block number %d of %d\n", iterator, length)
		}

		count, countErr := client.TransactionCount(context.Background(), blockN.Hash())
		if countErr != nil {
			log.Fatalf("management: error retrieving transaction count because %+v\n", countErr)
		}

		if count > 0 {
			blockData = append(blockData, blockN)
			log.Printf("management: Block number %d of %d with a transaction count of %d\n",iterator, length, count)
		}
	}
	log.Println("management: SearchTxBlocks Complete")
	return
}

// GetReceipt Get receipt from transaction struct
func GetReceipt(client *ethclient.Client, transaction *types.Transaction) (types.Message, *types.Receipt) {
	chainID, _ := client.NetworkID(context.Background())
	msg, _ := transaction.AsMessage(types.NewEIP155Signer(chainID))
	receipt, _  := client.TransactionReceipt(context.Background(), transaction.Hash())
	log.Printf("management: receipt status %d\n", receipt.Status)
	return msg, receipt
}
