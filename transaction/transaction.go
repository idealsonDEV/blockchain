// transaction
package BlockchainDocumentTransaction

import (
	"crypto/sha256"
	"encoding/json"
)

const coinbase = 50

type Transaction struct {
	Version  int64
	ID       []byte
	Tx_Input []TxInputSigned
	Tx_Oupt  []TxOutput
	Fee      int64
	LockTime int64
}

func (tx *Transaction) TxJson() string {
	txb, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return string(txb)
}

func (tx *Transaction) TxHash() []byte {
	txjson := tx.TxJson()
	txhash := sha256.Sum256(txjson)
	return txhash
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{}
	txout := NewTXOutput(0, coinbase, to)
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{*txout}}
	tx.ID = tx.Hash()

	return &tx
}
