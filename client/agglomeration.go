package network

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/manguilar22/construct-block/client/oop"
)

// MakeBlocks Turn block data to struct
func MakeBlocks(blockData []*types.Block) (collection []oop.EthBlock) {
	for _, block := range blockData {
		if len(block.Transactions()) > 0 {
			ethBlock := oop.EthBlock{
				BlockNumber:       block.Number().Uint64(),
				Time:              block.Time(),
				MiningDifficulty:  block.Difficulty().Uint64(),
				Hash:              block.Hash().Hex(),
				BlockTransactions: len(block.Transactions()),
			}
			collection = append(collection, ethBlock)
		}
	}
	return
}

// MakeTransactions Turn blocks to transaction struct
func MakeTransactions(transactionBlocks []*types.Block) (collection []oop.Transaction) {
	for _, e := range transactionBlocks {
		for _, tx := range e.Transactions() {
			save := oop.Transaction{
				NetworkID: tx.ChainId().Uint64(),
				BlockHash: tx.Hash().Hex(),
				BlockValue: tx.Value().String(),
				Gas:  tx.Gas(),
				GasPrice: tx.GasPrice().Uint64(),
				Nonce : tx.Nonce(),
				Data : tx.Data(),
			}
			collection = append(collection, save)
		}
	}
	return
}
