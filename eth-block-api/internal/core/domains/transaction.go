package domains

type Transaction struct {
	Hex           string `json:"hex"`
	Value         uint64 `json:"value"`
	Gas           uint64 `json:"gas"`
	GasPrice      uint64 `json:"gasPrice"`
	Nonce         uint64 `json:"nonce"`
	Reciever      string `json:"reciever"`
	Sender        string `json:"sender"`
}
