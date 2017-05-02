package xor

func SingleByteXOR(arr []byte, b byte) []byte {
	var result []byte
	for i, _ := range arr {
		result = append(result, arr[i]^b)
	}
	return result
}
