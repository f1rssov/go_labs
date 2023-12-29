package main

import (
	"crypto/sha256"
	"fmt"
)

func conuter_byte(x, y *[32]byte) int {
	count := 0

	for i := 0; i < len(x); i++ {
		diffBits := x[i] ^ y[i]

		for diffBits > 0 {
			if diffBits&1 == 1 {
				count++
			}
			diffBits >>= 1
		}
	}

	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", c1, c2)

	CountDiffbits := conuter_byte(&c1, &c2)
	fmt.Print(CountDiffbits)
}
