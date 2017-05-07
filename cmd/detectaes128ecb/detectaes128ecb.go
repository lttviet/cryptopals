package main

import (
	"bufio"
	"fmt"
	"github.com/lttviet/cryptopals/decrypt"
	"github.com/lttviet/cryptopals/strutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need a filename")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexStr := scanner.Text()
		raw := strutil.DecodeHexStr(hexStr)
		if decrypt.DetectAES128ECB(raw) {
			fmt.Println(hexStr)
		}
	}
}
