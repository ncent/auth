package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"log"
	"math/big"
)

type ECDSASignature struct {
	R, S *big.Int
}

func hash(b []byte) []byte {
	h := sha256.New()
	// hash the body bytes
	h.Write(b)
	// compute the SHA256 hash
	return h.Sum(nil)
}

func IsValidSignature(publicKey string, signature string, secret string) (bool, error) {

	pubBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Printf("Invalid public key HEX %+v", err)
		return false, err
	}
	log.Printf("pubBytes: %+v", pubBytes)

	x, y := elliptic.Unmarshal(elliptic.P256(), pubBytes)
	log.Printf("Pub x,y: %+v %+v", x, y)

	pub := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	log.Printf("Pub: %+v", pub)

	// first decode the signature to extract the DER-encoded byte string
	der, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	// unmarshal the R and S components of the ASN.1-encoded signature into our
	// signature data structure
	sig := &ECDSASignature{}
	_, err = asn1.Unmarshal(der, sig)
	if err != nil {
		return false, err
	}
	// compute the SHA256 hash of our message
	h := hash([]byte(secret))
	// validate the signature!
	valid := ecdsa.Verify(
		&pub,
		h,
		sig.R,
		sig.S,
	)
	return valid, nil
}
