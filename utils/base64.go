package utils

import (
	"encoding/base64"
)

func EncodeBase64(str string) string{
	bts := []byte(str)
	return base64.URLEncoding.EncodeToString(bts)
}

func DecodeBase64(str string)(string, error) {
	decoded, err := base64.URLEncoding.DecodeString(str)
	return string(decoded), err
}