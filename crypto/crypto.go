package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {

	privkey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		log.Fatal(err)
	}

	return privkey, &privkey.PublicKey
}

func EncryptWithPublicKey(secretMessage string, pub *rsa.PublicKey) []byte {

	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, []byte(secretMessage), nil)

	if err != nil {
		log.Fatal(err)
	}

	return ciphertext
}

func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)

	if err != nil {
		log.Fatal(err)
	}

	return plaintext
}
