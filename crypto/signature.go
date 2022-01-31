// signature
package BlockchainDocument

import (
	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/edwards25519"
)

var curve = edwards25519.SuiteEd25519{}
var sha256 = curve.Hash()

type Signature struct {
	R kyber.Point
	S kyber.Scalar
}

func (S *Signature) MarshalBinary() ([]byte, error) {
	r, err := S.R.MarshalBinary()
	if err != nil {
		return nil, err
	}
	s, err := S.S.MarshalBinary()
	if err != nil {
		return nil, err
	}

	binary := append(r, s...)

	return binary, nil
}

func (S *Signature) UnmarshalBinary(data []byte) {
	S.R.UnmarshalBinary(data[:32])
	S.S.UnmarshalBinary(data[32:])
}

func Hash(s string) []byte {
	sha256.Reset()
	sha256.Write([]byte(s))

	return sha256.Sum(nil)
}

func SignSchnorrRSwithoutP(m string, x kyber.Scalar) *Signature {
	if m == "" {
		panic("Erreur: Message a signé vide.")
	}
	if x.Equal(curve.Scalar().Zero()) {
		panic("Erreur: Clé privé à zero.")
	}

	// obtenir le point générateur
	G := curve.Point().Base()

	// obtenir k de mainère aléatoire
	k := curve.Scalar().Pick(curve.RandomStream())

	// R = k * G (likewise, r = g^k)
	R := curve.Point().Mul(k, G)

	// Hash(m || r)
	e := curve.Scalar().SetBytes(Hash(m + R.String()))

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, x))

	return &Signature{R: R, S: s}
}

func SignSchnorrRSwithP(m string, x kyber.Scalar) *Signature {
	if m == "" {
		panic("Erreur: Message a signé vide.")
	}
	if x.Equal(curve.Scalar().Zero()) {
		panic("Erreur: Clé privé à zero.")
	}

	// obtenir le point générateur
	G := curve.Point().Base()

	// obtenir k de mainère aléatoire
	k := curve.Scalar().Pick(curve.RandomStream())

	//Clé public P
	P := curve.Point().Mul(x, G)

	// R = k * G (likewise, r = g^k)
	R := curve.Point().Mul(k, G)

	// Hash(m || r)
	e := curve.Scalar().SetBytes(Hash(m + P.String() + R.String()))

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, x))

	return &Signature{R: R, S: s}
}

func VerifySchnorrRSwithoutP(m string, S Signature, Y kyber.Point) bool {
	if m == "" {
		panic("Erreur: Message a signé vide.")
	}
	if S.R == nil || S.S == nil || S.R.Equal(curve.Point()) || S.S.Equal(curve.Scalar()) {
		panic("Erreur: Signature malformé.")
	}
	if Y.Equal(curve.Point().Null()) {
		panic("Erreur: la clé public est null.")
	}

	// générateur
	G := curve.Point().Base()

	// e = Hash(m || R)
	e := curve.Scalar().SetBytes(Hash(m + S.R.String()))

	// s * G = r - e * y
	sGv := curve.Point().Sub(S.R, curve.Point().Mul(e, Y))

	// 's * G'
	sG := curve.Point().Mul(S.S, g)

	// Equality check; ensure signature and public key outputs to s * G.
	return sG.Equal(sGv)
}

func VerifySchnorrRSwithP(m string, S Signature, Y kyber.Point) bool {
	if m == "" {
		panic("Erreur: Message a signé vide.")
	}
	if S.R == nil || S.S == nil || S.R.Equal(curve.Point()) || S.S.Equal(curve.Scalar()) {
		panic("Erreur: Signature malformé.")
	}
	if Y.Equal(curve.Point().Null()) {
		panic("Erreur: la clé public est null.")
	}

	// générateur
	G := curve.Point().Base()

	// e = Hash(m || R)
	e := curve.Scalar().SetBytes(Hash(m + Y.String() + S.R.String()))

	// s * G = r - e * y
	sGv := curve.Point().Sub(S.R, curve.Point().Mul(e, Y))

	// 's * G'
	sG := curve.Point().Mul(S.S, g)

	// Equality check; ensure signature and public key outputs to s * G.
	return sG.Equal(sGv)
}
