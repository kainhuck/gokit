package serialize

import (
	"fmt"
	"github.com/kainhuck/gokit/serialize/json"
	"testing"
)

func TestSerialize(t *testing.T) {
	raw := "12.34567"
	bts, _ := json.EncodeBytes(raw)
	fmt.Println(bts)

	b, er := json.DecodeBytesToInterface(bts)
	fmt.Println(b)
	fmt.Println(er)

	var r uint64 = 13

	bb, _ := json.EncodeBytes(r)
	fmt.Println(json.DecodeBytesToUint64(bb))
}
