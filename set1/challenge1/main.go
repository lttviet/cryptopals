package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires a hex string.")
	}

	hexstr := os.Args[1]
	byteArr, err := hex.DecodeString(hexstr)
	if err != nil {
		log.Fatal(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(byteArr)
	fmt.Printf("ascii:%s\nbase64:%s\n", byteArr, base64Str)
}
