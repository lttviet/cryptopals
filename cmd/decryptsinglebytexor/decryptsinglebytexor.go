package main

import (
	"fmt"
	"github.com/lttviet/cryptopals/decrypt"
	"github.com/lttviet/cryptopals/stringutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires a hex strings.")
	}

	arr := stringutil.DecodeHexStr(os.Args[1])
	result := decrypt.DecryptSingleByteXOR(arr)
	fmt.Printf("%+v\n", result)
}
