package json

import (
	"encoding/json"
)

func DecodeBytesToInterface(bts []byte) (value interface{}, err error)  {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToString(bts []byte) (value string, err error){
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToBool(bts []byte) (value bool, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToUint8(bts []byte) (value uint8, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToUint16(bts []byte) (value uint16, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToUint32(bts []byte) (value uint32, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToUint64(bts []byte) (value uint64, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToUint(bts []byte) (value uint, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToInt8(bts []byte) (value int8, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToInt16(bts []byte) (value int16, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToInt32(bts []byte) (value int32, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToInt64(bts []byte) (value int64, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToInt(bts []byte) (value int, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToFloat32(bts []byte) (value float32, err error) {
	err = json.Unmarshal(bts, &value)
	return
}

func DecodeBytesToFloat64(bts []byte) (value float64, err error) {
	err = json.Unmarshal(bts, &value)
	return
}