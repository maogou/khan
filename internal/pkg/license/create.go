package license

import (
	"os"
)

func (l *License) Create(priPath, out string) error {
	privateKey, err := ReadPrivateKey(priPath)

	if err != nil {
		return err
	}

	encoded, err := l.Encode(privateKey)

	if err != nil {
		return err
	}

	if err = os.WriteFile(out, encoded, 0644); err != nil {
		return err
	}

	return nil
}
