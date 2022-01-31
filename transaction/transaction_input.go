// transaction_input
package BlockchainDocumentTransaction

import (
	"fmt"
)

type TxInput struct {
	Id      []byte
	Index   int64
	Value   int64
	Address string
}

type TxInputSigned struct {
	Id        []byte
	Index     int64
	Value     int64
	Signature []byte
	PubKey    []byte
}

func main() {
	fmt.Println("Hello World!")
}
