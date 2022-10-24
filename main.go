package main

import (
	"crypto/cipher"
	"crypto/des"
	"flag"
	"fmt"
	"log"
)

var isDecrypt bool
var isTriple bool
var inputFile string
var outputFile string
var key string

func AllignSize(block []byte) {
	rem := len(block) % 8
	if rem == 0 {
		return
	}

	block = append(block, make([]byte, rem)...)
}

func AsyncCrypt(cipher cipher.Block, block []byte, isDecrypt bool) []byte {
	AllignSize(block)
	pages := len(block) / des.BlockSize
	for i := 0; i < pages; i++ {
		if isDecrypt {
			go cipher.Decrypt(block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()], block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()])
		} else {
			go cipher.Encrypt(block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()], block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()])
		}
	}
	return block
}

func main() {
	var err error

	flag.BoolVar(&isTriple, "triple", false, "sets encryption method to TripleDES")
	flag.BoolVar(&isDecrypt, "decrypt", false, "sets mode to decrypt")
	flag.StringVar(&inputFile, "i", "", "sets input file")
	flag.StringVar(&outputFile, "o", "", "sets output file")
	flag.StringVar(&key, "key", "", "sets key (used only for `des` mode)")
	flag.Parse()

	var cipher cipher.Block
	if isTriple {
		cipher, err = des.NewTripleDESCipher([]byte(key))
	} else {
		cipher, err = des.NewCipher([]byte(key))
	}
	if err != nil {
		log.Fatalln(err)
	}
	result := make([]byte, 8)
	cipher.Encrypt([]byte("pudgeboo"), result)
	fmt.Println(result)
}
