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

	b64Cipher, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	hexCipher, err2 := base64.StdEncoding.DecodeString(string(b64Cipher))
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decrypt.DecryptRepeatingKeyXOR(hexCipher)))
}
