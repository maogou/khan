package help

import "os"

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mErr := os.MkdirAll(path, os.ModePerm)
		if mErr != nil {
			return mErr
		}
	}

	return nil
}
