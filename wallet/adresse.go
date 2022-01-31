// adresse
package BlockchainDocumentWallet

import (
	"blockchain/encoding"
	"bytes"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:4]
}

func HashPubKey(pubkey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubkey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func MakeAdress(version []byte, pubkey []byte) string {
	pubKeyHash := HashPubKey(pubkey)

	payload := append(version, pubKeyHash...)
	checksum := checksum(payload)

	fullPayload := append(payload, checksum...)
	address := BlockchainDocumentEncoding.EncodeBase58(fullPayload)

	return address
}

func ValidateAddress(address string) bool {
	fullPayload := BlockchainDocumentEncoding.DecodeBase58(address)
	actualChecksum := fullPayload[len(fullPayload)-4:]
	payload := fullPayload[:len(fullPayload)-4]
	targetChecksum := checksum(payload)

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}
