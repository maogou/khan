package license

import "smallBot/internal/pkg/license/licenseutil"

func (l *License) Verify(pub, file string) (*License, error) {
	publicKey, err := licenseutil.ReadPublicKey(pub)

	if err != nil {
		return nil, err
	}

	pl, err := DecodeFile(file, publicKey)
	if err != nil {
		return nil, err
	}

	return pl, nil
}
