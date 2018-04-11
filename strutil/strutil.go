package strutil

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"unicode"
)

func DecodeHexStr(str string) []byte {
	byteArr, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return byteArr
}

func HexToBase64(hexstr string) string {
	byteArr := DecodeHexStr(hexstr)
	return base64.StdEncoding.EncodeToString(byteArr)
}

func DecodeBase64Str(str string) []byte {
	byteArr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return byteArr
}

// Scores an ascii string based on english letter frequency
func Score(str string) int {
	letterFreq := map[string]int{
		"a": 8, "b": 1, "c": 3, "d": 4, "e": 13,
		"f": 2, "g": 2, "h": 6, "i": 7, "j": 1,
		"k": 1, "l": 4, "m": 2, "n": 7, "o": 8, "p": 2,
		"q": 1, "r": 6, "s": 6, "t": 9, "u": 3, "v": 1,
		"w": 2, "x": 1, "y": 2, "z": 1, " ": 14,
	}

	var score int
	for _, char := range str {
		char = unicode.ToLower(char)

		if _, ok := letterFreq[string(char)]; ok {
			score += letterFreq[string(char)]
		} else if unicode.IsNumber(char) {
			score++
		} else if unicode.IsPunct(char) {
			continue
		} else if unicode.IsSpace(char) {
			score--
		} else {
			score -= 20
		}
	}
	return score
}

// appending to a block a number of bytes of padding
func PKCS7Padding(block []byte, blocksize int) []byte {
	if len(block) > blocksize {
		log.Fatal("Size of block is greater than given blocksize.")
	}

	pad := blocksize - len(block)
	for i := 0; i < pad; i++ {
		block = append(block, byte(pad))
	}
	return block
}
