package license

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"
)

func (l *License) KeyGen(out string) error {
	if len(out) == 0 {
		out = time.Now().Format("20060102150405")
	}

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return err
	}

	publicKeyHex := base64.StdEncoding.EncodeToString(publicKey)
	privateKeyHex := base64.StdEncoding.EncodeToString(privateKey)

	if err = os.WriteFile(out+".pri", []byte(privateKeyHex), 0644); err != nil {
		return err
	}

	if err = os.WriteFile(out+".pub", []byte(publicKeyHex), 0644); err != nil {
		return err
	}

	return nil
}
