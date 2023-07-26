package utils

import (
	"io/ioutil"
)

func LoadPrivateKeyFromDisk() ([]byte, error) {
	env := ParseEnv()

	keyData, err := ioutil.ReadFile(env.KeyPath)

	if err != nil {
		return nil, err
	}

	return keyData, nil
}
