// main
package main

import (
	wallet "blockchain/wallet"
	"fmt"
)

func main() {
	w := wallet.NewWalletRandom()
	fmt.Println(w.PrivateKey)
	fmt.Println(w.PublicKey)
	fmt.Println(w.GetAdressVer1())
	fmt.Println(wallet.ValidateAddress(w.GetAdressVer1()))
}
