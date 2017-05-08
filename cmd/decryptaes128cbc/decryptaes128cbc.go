package main

import (
	"encoding/base64"
	"fmt"
	"github.com/lttviet/cryptopals/decrypt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need a filepath.")
	}

	key := []byte("YELLOW SUBMARINE")
	iv := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

	b64, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	raw := make([]byte, len(b64))
	n, err2 := base64.StdEncoding.Decode(raw, b64)
	if err2 != nil {
		log.Fatal(err2)
	}

	decrypted := decrypt.DecryptAES128CBC(raw[:n], key, iv)
	fmt.Println(string(decrypted))
}
