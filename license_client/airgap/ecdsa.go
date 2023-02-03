package airgap

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
)

var curve = elliptic.P256()

func KeyFromBytes(key []byte) *ecdsa.PrivateKey {
	n := new(big.Int)
	n.SetBytes(key)

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = n
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(n.Bytes())

	return priv
}

func GenerateSignature(key *ecdsa.PrivateKey, signingString string) (string, error) {
	hash := sha1.Sum([]byte(signingString))
	r, s, err := ecdsa.Sign(rand.Reader, key, hash[:])
	if err != nil {
		return "", fmt.Errorf("error generating activation code: %w", err)
	}
	return encode(r, s), nil
}

func encode(r, s *big.Int) string {
	signature := hexToAscii(numberToString(r)) + hexToAscii(numberToString(s))
	return base64.StdEncoding.EncodeToString([]byte(signature))
}

func numberToString(num *big.Int) string {
	return fmt.Sprintf("%064x", num)
}

func hexToAscii(s string) string {
	decoded, _ := hex.DecodeString(s)
	return string(decoded)
}
