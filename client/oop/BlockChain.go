package oop

type Blockchain struct {
	Hostname string `json:"hostname"`
	BlockData []EthBlock `json:"block_data"`
	TransactionData []Transaction `json:"transaction_data"`
}

// TODO: Delete?
func (b *Blockchain) GetTransactionBlocks() (data []EthBlock) {
	var blockData = b.BlockData

	for e := range blockData {
		if (blockData[e].BlockTransactions > 0) {
			data = append(data, blockData[e])
		}
	}
	return
}

// TODO: Delete?
func (b *Blockchain) FilterBlocks(pred func(EthBlock) bool) (data []EthBlock) {
	var blockData []EthBlock = b.BlockData

	for e := range blockData {
		if pred(blockData[e]) {
			data = append(data, blockData[e])
		}
	}
	return
}