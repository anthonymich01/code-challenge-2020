package models

type Contract struct {
	Address     string `json:"address"`
	BlockHeight int64  `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   string `json:"block_time"`
	ERC20       bool   `json:"erc20"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Decimals    string `json:"decimals"`
}
