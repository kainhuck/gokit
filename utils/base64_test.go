package utils

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	x := "123456asdfgh!@#$%^&*"
	xx := EncodeBase64(x)
	fmt.Println(xx)
	xxx, err := DecodeBase64(xx)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(xxx)

}
