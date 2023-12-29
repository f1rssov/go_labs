package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	flag_sha384 := flag.Bool("sha384", false, "use sha384")
	flag_sha512 := flag.Bool("sha512", false, "use sha512")
	flag.Parse()

	inputData := []byte(flag.Arg(0))

	var result []byte

	switch {
	case *flag_sha384:
		hash := sha512.Sum384(inputData)
		result = hash[:]
	case *flag_sha512:
		hash := sha512.Sum512(inputData)
		result = hash[:]
	default:
		hash := sha256.Sum256(inputData)
		result = hash[:]
	}

	fmt.Printf("%x\n", result)

}
