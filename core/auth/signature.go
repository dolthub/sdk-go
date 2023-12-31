package auth

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAwEt7ZQWSzeBYNBwyi3KW
+XU/6I+ju90R5ZpTvSE4VWL8KSgVJ6bWKIhaKTL2hbUpIQQgS5ZKfa4SEEMdIm5k
xe0u2F64JFAVQunx/1O/UXsD7ADVt93Q/hxA9Npa16tKepZoyhi79sxpaxyy/WTd
sZKuTLApX7bWX6/XPhaPjNyiWeeK2Naka44B+F+PwN/ey3rZUra3pBltShwy2dmK
IxJVmprf5ttMYBB+ouqPin3VgDw5Jq1FZwpLBiOe+ogR+sHu7QYWwLq6AgC4e3Jq
gryEvqfJr/XkvOnAdcAIZm6tK6DdnLWCy6Onrk7t0VQK5nnGF1CKE+jcAByWopcI
VCwfYrs1UCSY1YibKXNJoyvEhgLyCC+KAgXKf8omcg1Q18XKu5XSrvrCzXLnRnaF
NpFDjcg5AJJab78hX7qGPC+e8PjuBYwh2vtx5mFj7/c+T59JM/vXwwvW9DsnDztD
WEDFhzGanU71NgwrZbRoNSyTW0UjbZsyJVSIX6233Ng0y5L9mDe8p8P+u1B0sXAA
ozEKL/yG4Qu5r4LIw4iv6JVPT1xbWlH4Vc4KaN3toaf4+G0EkFy6ncvBncWifLAR
SoWpQ5YsygDdQNVYdmMsoQ76UTuNxo3eZ0sQJjbQVBqlkcXrAdtUCpuojPRsl/Xs
9zVOJrzZPG0I98E5quNbRkMCAwEAAQ==
-----END PUBLIC KEY-----`

func VerifySignatureV2(signature, json string) error {
	key := getKey()

	h := sha256.New()
	h.Write([]byte(json))
	hash := h.Sum(nil)

	signatureDecoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hash[:], signatureDecoded)
}

func getKey() *rsa.PublicKey {
	decoded, _ := pem.Decode([]byte(publicKey))
	publicKey, _ := x509.ParsePKIXPublicKey(decoded.Bytes)
	return publicKey.(*rsa.PublicKey)
}
