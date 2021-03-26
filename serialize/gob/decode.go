package gob

import (
	"bytes"
	"encoding/gob"
)

func DecodeBytesToString(bts []byte) (value string, err error) {
	buf := bytes.NewReader(bts)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&value)

	return value, err
}
