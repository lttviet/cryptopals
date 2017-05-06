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
	if len(os.Args) < 3 {
		log.Fatal("Required 16 bytes KEY and a Base64-encoded textfile.")
	}

	key := []byte(os.Args[1])

	b64, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, len(b64))
	end, err2 := base64.StdEncoding.Decode(data, b64)
	if err2 != nil {
		log.Fatal(err2)
	}

	decrypted := decrypt.DecryptAES128ECB(data[:end], key)
	fmt.Println(string(decrypted))
}
