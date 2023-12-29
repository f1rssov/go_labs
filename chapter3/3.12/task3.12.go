package main

import (
	"fmt"
	"unicode"
)

func IsAnagram(s1, s2 string) bool {
	var runeS1, runeS2 = []rune(s1), []rune(s2)
	var lenS1, lenS2 = len(runeS1), len(runeS2)

	if lenS1 != lenS2 {
		return false
	} else {
		for i := 0; i < lenS1; i++ {
			if unicode.ToLower(runeS1[i]) != unicode.ToLower(runeS2[lenS2-i-1]) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Print(IsAnagram("12345", "54312"))
}
