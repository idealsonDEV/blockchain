// wallet
package BlockchainDocumentWallet

import (
	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/edwards25519"
)

var curve = edwards25519.SuiteEd25519{}

type Wallet struct {
	PrivateKey kyber.Scalar
	PublicKey  kyber.Point
}

func NewWalletRandom() *Wallet {
	priv := curve.Scalar().Pick(curve.RandomStream())
	pubk := curve.Point().Mul(priv, curve.Point().Base())
	return &Wallet{PrivateKey: priv, PublicKey: pubk}
}

func NewWalletFromSeed(seed []byte) *Wallet {
	priv := curve.Scalar().SetBytes(seed)
	pubk := curve.Point().Mul(priv, curve.Point().Base())
	return &Wallet{PrivateKey: priv, PublicKey: pubk}
}

func (w *Wallet) GetAdressVer1() string {
	version := []byte{byte(0x86)}
	pubkbin, err := w.PublicKey.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return MakeAdress(version, pubkbin)
}
