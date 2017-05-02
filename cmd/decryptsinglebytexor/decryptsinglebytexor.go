package main

import (
	"fmt"
	"github.com/lttviet/cryptopals/decrypt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires a hex strings.")
	}

	result := decrypt.DecryptSingleByteXOR(os.Args[1])
	fmt.Println(result)
}
