package json

import (
	"encoding/json"
)


func EncodeBytes(key interface{}) ([]byte, error) {
	return json.Marshal(key)
}

func EncodeString(key interface{}) (string, error) {
	bts, err := json.Marshal(key)
	return string(bts), err
}