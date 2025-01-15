package license

import (
	"crypto/ed25519"
	"os"
)

func Parse(pub, file string) (*License, error) {
	var (
		publicKey ed25519.PublicKey
		err       error
	)
	if _, err = os.Stat(pub); err != nil {
		publicKey, err = ReadPublicKeyFromStr(pub)
	} else {
		publicKey, err = ReadPublicKey(pub)
	}

	if err != nil {
		return nil, err
	}

	pl, err := DecodeFile(file, publicKey)
	if err != nil {
		return nil, err
	}

	return pl, nil
}
