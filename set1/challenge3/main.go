package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"unicode"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires a hex strings.")
	}

	var results []string
	var scores []int

	arr := decodeHexStr(os.Args[1])

	for _, char := range alphabet {
		charByteArr := []byte(string(char))
		result := xor(arr, charByteArr)

		results = append(results, result)
		scores = append(scores, score(result))
	}

	maxIndex, _ := max(scores)
	fmt.Println(results[maxIndex])
}

func xor(arr1, arr2 []byte) string {
	var result []byte
	for i, _ := range arr1 {
		// if len(arr2) < len(arr1), modulus to repeat arr2 byte
		result = append(result, arr1[i]^arr2[i%len(arr2)])
	}

	return string(result)
}

func decodeHexStr(str string) []byte {
	byteArr, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}

	return byteArr
}

func score(str string) int {
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

func max(arr []int) (int, int) {
	// return index and value of a max value in array
	var maxIndex, max int
	for i, val := range arr {
		if val > max {
			maxIndex, max = i, val
		}
	}
	return maxIndex, max
}
