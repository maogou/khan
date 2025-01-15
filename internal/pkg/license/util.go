package license

import (
	"crypto/ed25519"
	"encoding/base64"
	"os"
)

func ReadPublicKey(path string) (ed25519.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return DecodePublicKey(data)
}

func ReadPublicKeyFromStr(str string) (ed25519.PublicKey, error) {
	return DecodePublicKey([]byte(str))
}

func ReadPrivateKey(path string) (ed25519.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return DecodePrivateKey(data)
}

func DecodePublicKey(data []byte) (ed25519.PublicKey, error) {
	decoded, err := decode(data)
	if err != nil {
		return nil, err
	}
	return ed25519.PublicKey(decoded), nil
}

func DecodePrivateKey(data []byte) (ed25519.PrivateKey, error) {
	decoded, err := decode(data)
	if err != nil {
		return nil, err
	}
	return ed25519.PrivateKey(decoded), nil
}

func decode(b []byte) ([]byte, error) {
	enc := base64.StdEncoding
	buf := make([]byte, enc.DecodedLen(len(b)))
	n, err := enc.Decode(buf, b)
	return buf[:n], err
}
