package decrypt

import (
	"github.com/lttviet/cryptopals/stringutil"
	"github.com/lttviet/cryptopals/xor"
)

type Text struct {
	cipher, plaintext string
	score             int
}

func (t *Text) Score() int {
	return t.score
}

// Given a hex string, find the best ascii text based on scores
func DecryptSingleByteXOR(s string) Text {
	var texts []Text
	t := Text{cipher: s}

	arr := stringutil.DecodeHexStr(s)
	for i := 0; i < 256; i++ {
		t.plaintext = string(xor.SingleByteXOR(arr, byte(i)))
		t.score = stringutil.Score(t.plaintext)

		texts = append(texts, t)
	}
	return max(texts)
}

// Returns Text with the highest score
func max(arr []Text) Text {
	var maxT Text
	for _, t := range arr {
		if t.score > maxT.score {
			maxT = t
		}
	}
	return maxT
}
