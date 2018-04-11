package decrypt

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestOracle(t *testing.T) {
	plaintext, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal("Can't read file.")
	}

	for i := 0; i < 10; i++ {
		cipher, choice := OracleEncrypt(plaintext)
		guess := Oracle(cipher)
		log.Println(choice, guess)
		if choice != guess {
			t.Errorf("Choice: %d, Guess: %d", choice, guess)
		}
	}
}

func TestOracleDecryptECB(t *testing.T) {
	plaintext := OracleDecryptECB()
	log.Println(string(plaintext[:]))
}
