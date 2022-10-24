package main

import (
	"crypto/cipher"
	"crypto/des"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var isDecrypt bool
var isTriple bool
var inputFile string
var outputFile string
var key string

func AllignSize(block []byte) ([]byte, int) {
	rem := len(block) % des.BlockSize
	if rem == 0 {
		return block, rem
	}

	block = append(block, make([]byte, rem)...)
	return block, rem
}

func AsyncCrypt(cipher cipher.Block, block []byte, isDecrypt bool) []byte {
	block, rem := AllignSize(block)
	pages := len(block) / des.BlockSize
	wg := sync.WaitGroup{}
	for i := 0; i < pages-1; i++ {
		wg.Add(1)
		if isDecrypt {
			go func(i int) {
				cipher.Decrypt(block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()], block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()])
				wg.Done()
			}(i)

		} else {
			go func(i int) {
				cipher.Encrypt(block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()], block[i*cipher.BlockSize():(i+1)*cipher.BlockSize()])
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	block = block[:len(block)-rem]
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
	fmt.Println("key:", key)

	if len(inputFile) == 0 {
		log.Fatalln("no input file was given")
	}

	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalln("error opening file:", err)
	}

	result := AsyncCrypt(cipher, input, isDecrypt)

	if len(outputFile) == 0 {
		fmt.Println(`no output file was given, writing to "output.txt"`)
		outputFile = "output.txt"
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln("error opening file:", err)
	}
	defer file.Close()
	file.Write(result)
}
