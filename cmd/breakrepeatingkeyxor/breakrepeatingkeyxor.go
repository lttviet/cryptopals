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

	b64, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, len(b64))
	end, err2 := base64.StdEncoding.Decode(data, b64)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(string(decrypt.DecryptRepeatingKeyXOR(data[:end])))
}
