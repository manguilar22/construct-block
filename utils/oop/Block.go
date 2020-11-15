package oop

type EthBlock struct {
	BlockNumber       uint64 `json:"block_number"`
	Time              uint64 `json:"block_time"`
	MiningDifficulty  uint64 `json:"mining_difficulty"`
	Hash              string `json:"hash"`
	BlockTransactions int    `json:"block_transactions"`
	TransactionCount uint   `json:"transaction_count"`
}