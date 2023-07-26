package utils

import (
	"bytes"
	"encoding/gob"
)

func EncodeForRedisStore(data interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeFromRedisStore(data string, res interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader([]byte(data)))

	err := dec.Decode(res)

	if err != nil {
		return err
	}

	return nil
}
