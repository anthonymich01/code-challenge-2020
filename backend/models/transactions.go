package models

type Transaction struct {
	Txid        string `json:"txid"`
	BlockHeight int64  `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   string `json:"block_time"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       int64  `json:"value"`
	GasProvided int64  `json:"gas_provided"`
	GasUsed     int64  `json:"gas_used"`
	GasPrice    int64  `json:"gas_price"`
	Status      string `json:"status"`
}
