package stringutil

import (
	"encoding/hex"
	"log"
)

func DecodeHexStr(str string) []byte {
	byteArr, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return byteArr
}
