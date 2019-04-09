package main

import (
	"strings"
	"fmt"
	"unicode"
)

func main() {
	word :="Leetcode"
	fmt.Println(unicode.IsUpper(rune(word[0])))
	fmt.Println(detectCapitalUse("FlaG"))

}

func detectCapitalUse(word string) bool {
	if strings.Compare(word, strings.ToUpper(word)) == 0 || strings.Compare(word, strings.ToLower(word)) == 0 {
		return true
	}else if unicode.IsUpper(rune(word[0])) {
		for i,w :=range word {
			if i==0 {
				continue
			}
			if unicode.IsUpper(rune(w)) {
				return false
			}

		}
		return true
	}
	return false
}
