package xor

import (
	"log"
)

// xor 2 equal-length buffer, returns a byte array
func FixedXOR(arr1, arr2 []byte) []byte {
	if len(arr1) != len(arr2) {
		log.Fatal("Byte arrays don't have the same length.")
	}

	var result []byte
	for i, _ := range arr1 {
		result = append(result, arr1[i]^arr2[i])
	}
	return result
}
