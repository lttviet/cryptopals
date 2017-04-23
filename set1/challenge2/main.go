package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Requires 2 hex strings.")
	}

	arr1, arr2 := decodeHexStr(os.Args[1]), decodeHexStr(os.Args[2])
	arr := xor(arr1, arr2)

	str := hex.EncodeToString(arr)
	fmt.Println(str)
}

func xor(arr1, arr2 []byte) []byte {
	var result []byte
	for i, _ := range arr1 {
		result = append(result, arr1[i]^arr2[i])
	}

	return result
}

func decodeHexStr(str string) []byte {
	byteArr, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	return byteArr
}
