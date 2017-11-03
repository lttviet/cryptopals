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

	cipher, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	guess := decrypt.Oracle(cipher)
	if guess == 0 {
		log.Println("Oracle thinks ecb")
	} else {
		log.Println("Oracle thinks cbc")
	}
}
