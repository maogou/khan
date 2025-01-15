package license

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"
)

func (l *License) KeyGen(path, name string) error {
	if len(name) == 0 {
		name = time.Now().Format("20060102150405")
	}

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return err
	}

	publicKeyHex := base64.StdEncoding.EncodeToString(publicKey)
	privateKeyHex := base64.StdEncoding.EncodeToString(privateKey)

	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if err = os.WriteFile(path+"/"+name+".pri", []byte(privateKeyHex), 0644); err != nil {
		return err
	}

	if err = os.WriteFile(path+"/"+name+".pub", []byte(publicKeyHex), 0644); err != nil {
		return err
	}

	return nil
}
