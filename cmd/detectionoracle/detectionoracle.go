package main

import (
	"github.com/lttviet/cryptopals/decrypt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need a filepath.")
	}

	plaintext, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	decrypt.Oracle(plaintext)
}
