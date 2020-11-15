package oop

type Blockchain struct {
	Data []EthBlock `json:"data"`
}

func (b *Blockchain) GetTransactionBlocks() (data []EthBlock) {
	var blockData []EthBlock = b.Data

	for e := range blockData {
		if (blockData[e].TransactionCount > 0) {
			data = append(data, blockData[e])
		}
	}
	return
}


func (b *Blockchain) FilterBlocks(pred func(EthBlock) bool) (data []EthBlock) {
	var blockData []EthBlock = b.Data

	for e := range blockData {
		if pred(blockData[e]) {
			data = append(data, blockData[e])
		}
	}
	return
}