package decrypt

import (
	"crypto/aes"
	"github.com/lttviet/cryptopals/stringutil"
	"github.com/lttviet/cryptopals/xor"
	"log"
)

const (
	MINKEYSIZE = 2
	MAXKEYSIZE = 40
)

type Block []byte

type Text struct {
	cipher, plaintext string
	score             int
	key               byte
}

func (t *Text) Score() int {
	return t.score
}

// Given a byte array, find the best ascii text based on scores
func DecryptSingleByteXOR(arr []byte) Text {
	var texts []Text
	var t Text

	for i := 0; i < 256; i++ {
		t.key = byte(i)
		t.plaintext = string(xor.SingleByteXOR(arr, t.key))
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

func DecryptRepeatingKeyXOR(cipher []byte) []byte {
	keySize := findKeySize(cipher)
	blks := breakCiphertext(cipher, keySize)
	blks = transpose(blks)

	var key []byte
	for _, b := range blks {
		t := DecryptSingleByteXOR(b)
		key = append(key, t.key)
	}

	return xor.RepeatingKeyXOR(cipher, key)
}

func findKeySize(cipher []byte) int {
	keySize, minScore := 0, 9999.0
	for i := MINKEYSIZE; i <= MAXKEYSIZE; i++ {
		s := score(cipher, i)
		if s < minScore {
			keySize, minScore = i, s
		}
	}
	return keySize
}

// Returns their average Hammington Distance
func score(cipher Block, keySize int) float64 {
	// First 4 blocks of keySize-length of ciphertext
	blks := breakCiphertext(cipher, keySize)[:4]

	var s, len float64
	for c := range generateCombination(blks) {
		s += float64(hammingtonDistance(c[0], c[1]))
		len++
	}
	s /= len              // average
	s /= float64(keySize) // normalise
	return s
}

// 2-combination
func generateCombination(blks []Block) chan []Block {
	ch := make(chan []Block)

	//defer close(ch)
	go func() {
		for i := 0; i < len(blks); i++ {
			for j := i + 1; j < len(blks); j++ {
				ch <- []Block{blks[i], blks[j]}
			}
		}
		close(ch)
	}()
	return ch
}

// bit differece
func hammingtonDistance(arr1, arr2 Block) int {
	n := 0
	for i, _ := range arr1 {
		for b := arr1[i] ^ arr2[i]; b != 0; b &= b - 1 {
			n++
		}
	}
	return n
}

// Breaks ciphertext into blocks of KeySize length
func breakCiphertext(cipher Block, keySize int) []Block {
	var blks []Block
	for i := 0; i < len(cipher); i += keySize {
		blks = append(blks, cipher[i:i+keySize])
	}
	return blks
}

func transpose(blks []Block) []Block {
	result := make([]Block, len(blks[0]))
	for _, blk := range blks {
		for i, b := range blk {
			result[i] = append(result[i], b)
		}
	}
	return result
}

func DecryptAES128ECB(data, key []byte) []byte {
	if len(key) != 16 {
		log.Fatal("Key must be 16 bytes")
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
