package decrypt

import (
	"github.com/lttviet/cryptopals/stringutil"
	"github.com/lttviet/cryptopals/xor"
	"unicode"
)

// Given a hex string, find the best ascii text based on scores
func DecryptSingleByteXOR(s string) string {
	arr := stringutil.DecodeHexStr(s)

	var results []string
	var scores []int
	for i := 0; i < 256; i++ {
		singleByte := byte(i)
		result := string(xor.SingleByteXOR(arr, singleByte))

		results = append(results, result)
		scores = append(scores, scoring(result))
	}

	maxIndex, _ := max(scores)
	return results[maxIndex]
}

// Scores an ascii string based on english letter frequency
func scoring(str string) int {
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

// Returns index and value of a max value in array
func max(arr []int) (int, int) {
	var maxIndex, max int
	for i, val := range arr {
		if val > max {
			maxIndex, max = i, val
		}
	}
	return maxIndex, max
}
