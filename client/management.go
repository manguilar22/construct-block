package network

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manguilar22/construct-block/utils"
	"github.com/manguilar22/construct-block/utils/oop"
	"log"
	"math/big"
	"time"
)

// TransactionsCount (HTTP)
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
		utils.BlockDataTextBanner(blockN)
		fmt.Printf("Transaction Count: %d\n", count)

	}
}

// Accounts (RPC)
func GetAllAccounts(ks *keystore.KeyStore) (result string) {
	accs := ks.Accounts()
	result = "|Account Number|Address|Hash|URL|\n|:---|:---:|:---:|---:|"
	for idx, e := range accs {
		result += fmt.Sprintf("|%d|%s|%s|%s|\n", idx, e.Address.Hex(),e.Address.Hash().String(), e.URL.String())
	}
	return
}

// TransactionsCount (HTTP)
func GetAllBlocks(client *ethclient.Client) (blockData []oop.EthBlock) {

	totalBlocks, _ := client.BlockByNumber(context.Background(), nil)
	var length = totalBlocks.Number().Int64()
	blockData = make([]oop.EthBlock, length-length)


	for i := 0; i < int(length); i++ {
		time.Sleep(50000)								// TODO: REMOVE
		var iterator = big.NewInt(int64(i))
		blockN, err := client.BlockByNumber(context.Background(),iterator)
		if err != nil {
			log.Fatalf("management: Problem retrieving block number %d of %d\n", iterator, length)
		}

		count, _ := client.TransactionCount(context.Background(), blockN.Hash())

		block := oop.EthBlock{
				BlockNumber:       blockN.Number().Uint64(),
				Time:              blockN.Time(),
				MiningDifficulty:  blockN.Difficulty().Uint64(),
				Hash:              blockN.Hash().Hex() ,
				BlockTransactions: len(blockN.Transactions()),
				TransactionCount: count,
		}

		blockData = append(blockData, block)
		fmt.Printf("management: Block number %d of %d with a transaction count of %d\n",iterator, length, count)
	}
	return
}

// TransactionsCount (HTTP)
func GetAllBlocksN(client *ethclient.Client, N int) (blockData []oop.EthBlock) {
	length := N
	blockData = make([]oop.EthBlock, 0)

	for i := 0; i < length; i++ {
		var iterator = big.NewInt(int64(i))
		blockN, err := client.BlockByNumber(context.Background(),iterator)
		if err != nil {
			log.Fatalf("management: Problem retrieving block number %d of %d\n", iterator, length)
		}

		count, _ := client.TransactionCount(context.Background(), blockN.Hash())

		block := oop.EthBlock{
			BlockNumber:       blockN.Number().Uint64(),
			Time:              blockN.Time(),
			MiningDifficulty:  blockN.Difficulty().Uint64(),
			Hash:              blockN.Hash().Hex() ,
			BlockTransactions: len(blockN.Transactions()),
			TransactionCount: count,
		}

		blockData = append(blockData, block)
		fmt.Printf("management: Block number %d of %d with a transaction count of %d\n",iterator, length, count)
	}
	return
}

// TransactionsCount (HTTP)
func GetTransaction(client *ethclient.Client, currentBlock oop.EthBlock) {
	blockNumber := big.NewInt(int64(currentBlock.BlockNumber))
	block, _ := client.BlockByNumber(context.Background(), blockNumber)

	for _, tx := range block.Transactions() {
		chainID, _ := client.NetworkID(context.Background())
		transaction := oop.Transactions{
			NetworkID:  chainID.Uint64(),
			BlockHash:  tx.Hash().Hex(),
			BlockValue: tx.Value().String(),
			Gas:        tx.Gas(),
			GasPrice:   tx.GasPrice().Uint64(),
			Nonce:      tx.Nonce(),
			Data:       tx.Data(),
			ToHex:      tx.To().String(),
		}
		fmt.Println(transaction)
	}
}

func GetReceipt(client ethclient.Client, transaction types.Transaction) (types.Message, *types.Receipt) {
	chainID, _ := client.NetworkID(context.Background())
	msg, _ := transaction.AsMessage(types.NewEIP155Signer(chainID))
	receipt, _  := client.TransactionReceipt(context.Background(), transaction.Hash())
	return msg, receipt
}
