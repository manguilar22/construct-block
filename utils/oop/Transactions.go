package oop

type Transactions struct {
	NetworkID uint64 `json:"network_id"`
	BlockHash string `json:"block_hash"`
	BlockValue string `json:"block_value"`
	Gas uint64 `json:"gas"`
	GasPrice uint64 `json:"gas_price"`
	Nonce uint64 `json:"nonce"`
	Data []byte `json:"data"`
	ToHex string `json:"to_hex"`
}
