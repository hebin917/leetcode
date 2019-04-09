package main

import (
	"bytes"
	)

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	uniqueMorseRepresentations(words)
	//fmt.Println(int('a'))
}

func uniqueMorseRepresentations(words []string) int {
	moser := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	rest := make(map[string]string)
	buf := bytes.Buffer{}
	for _, word := range words {
		for _,v :=range word {
			index := int(v)-int('a')
			buf.WriteString(moser[index])
		}
		//fmt.Println(buf.String())
		rest[buf.String()] = word
		buf =bytes.Buffer{}
	}
	return len(rest)
}

