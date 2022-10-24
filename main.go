package main

import "flag"

var isDecrypt bool
var isTriple bool
var inputFile string
var outputFile string
var key string

func main() {
	flag.BoolVar(&isTriple, "triple", false, "sets encryption method to TripleDES")
	flag.BoolVar(&isDecrypt, "decrypt", false, "sets mode to decrypt")
	flag.StringVar(&inputFile, "i", "", "sets input file")
	flag.StringVar(&outputFile, "o", "", "sets output file")
	flag.StringVar(&key, "key", "", "sets key (used only for `des` mode)")
	flag.Parse()

	
}
