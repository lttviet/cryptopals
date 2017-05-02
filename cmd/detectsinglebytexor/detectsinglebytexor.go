package main

import (
	"bufio"
	"fmt"
	"github.com/lttviet/cryptopals/decrypt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Requires a filepath and a minimum score limit.")
	}

	minScore, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := decrypt.DecryptSingleByteXOR(scanner.Text())

		if t.Score() >= minScore {
			fmt.Printf("%+v\n", t)
		}
	}
}
