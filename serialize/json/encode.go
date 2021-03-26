package json

import (
	"encoding/json"
)


func EncodeBytes(key interface{}) ([]byte, error) {
	return json.Marshal(key)
}