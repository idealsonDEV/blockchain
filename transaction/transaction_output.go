// transaction_output
package BlockchainDocumentTransaction

import (
	"fmt"
)

type TxOutput struct {
	Index   int64
	Value   int64
	Address string
}

func NewTxOutput(ind int64, vout int64, to string) *TxOutput {
	tx := TxOutput{Index: ind, Value: vout, Address, to}
	return &tx
}
