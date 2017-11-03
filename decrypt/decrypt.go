package decrypt

import (
	"crypto/aes"
	"github.com/lttviet/cryptopals/strutil"
	"github.com/lttviet/cryptopals/xor"
	"log"
	"math/rand"
	"time"
)

const (
	MINKEYSIZE = 2
	MAXKEYSIZE = 40
	AES128SIZE = 16
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
		t.score = strutil.Score(t.plaintext)

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

func EncryptAES128ECB(plaintext, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// ciphertext length is multiple of key size
	ciphertextLen := len(key)
	for ; len(plaintext) > ciphertextLen; ciphertextLen += len(key) {
	}

	plaintext = strutil.PKCS7Padding(plaintext, ciphertextLen) // pad plaintext
	ciphertext := make([]byte, ciphertextLen)

	for bs, be := 0, AES128SIZE; bs < len(plaintext); bs, be = bs+AES128SIZE, be+AES128SIZE {
		cipher.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}

	return ciphertext
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

	for bs, be := 0, AES128SIZE; bs < len(data); bs, be = bs+AES128SIZE, be+AES128SIZE {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

// check if the raw byte is encrypted with ECB
func DetectAES128ECB(raw []byte) bool {
	seen := make(map[[16]byte]bool, len(raw)/16)
	for i := 0; i < len(raw); i += 16 {
		var key [16]byte
		copy(key[:], raw[i:i+16])

		if seen[key] {
			return true
		} else {
			seen[key] = true
		}
	}
	return false
}

func EncryptAES128CBC(plaintext, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// ciphertext length is multiple of key size
	ciphertextLen := len(key)
	for ; len(plaintext) > ciphertextLen; ciphertextLen += len(key) {
	}

	plaintext = strutil.PKCS7Padding(plaintext, ciphertextLen) // pad plaintext
	ciphertext := make([]byte, ciphertextLen)

	for i := 0; i < ciphertextLen; i += 16 {
		var src []byte

		if i == 0 {
			// first block of plaintext
			src = xor.FixedXOR(plaintext[i:i+16], iv)
		} else {
			src = xor.FixedXOR(plaintext[i:i+16], ciphertext[i-16:i])
		}
		cipher.Encrypt(ciphertext[i:i+16], src)
	}

	return ciphertext
}

func DecryptAES128CBC(ciphertext, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	plaintext := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext); i += 16 {
		cipher.Decrypt(plaintext[i:i+16], ciphertext[i:i+16])

		var decrypted []byte

		if i == 0 {
			decrypted = xor.FixedXOR(plaintext[i:i+16], iv)
		} else {
			decrypted = xor.FixedXOR(plaintext[i:i+16], ciphertext[i-16:i])
		}
		copy(plaintext[i:i+16], decrypted)
	}

	return plaintext
}

// Generates random byte array of given len
func randomBytes(len int) []byte {
	arr := make([]byte, len)
	_, err := rand.Read(arr)
	if err != nil {
		log.Fatal(err)
	}

	return arr
}

// Generates random aes key, 16 random bytes
func randomAESKey() []byte {
	return randomBytes(AES128SIZE)
}

func OracleEncrypt(plaintext []byte) ([]byte, int) {
	rand.Seed(time.Now().UnixNano())

	before := rand.Intn(6) + 5
	plaintext = append(randomBytes(before), plaintext...)

	after := rand.Intn(6) + 5
	plaintext = append(plaintext, randomBytes(after)...)

	var cipher []byte
	key := randomAESKey()
	choice := rand.Intn(2)

	if choice == 0 {
		iv := randomBytes(AES128SIZE)
		cipher = EncryptAES128CBC(plaintext, key, iv)
	} else {
		cipher = EncryptAES128ECB(plaintext, key)
	}

	return cipher, choice
}

// Returns 1 if ecb, 0 if cbc
func Oracle(cipher []byte) int {
	if DetectAES128ECB(cipher) {
		return 1
	} else {
		return 0
	}
}
