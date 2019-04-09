package main

import (
	"fmt"
	"strings"
	"bytes"
)

func main() {
	fmt.Println(strings.Repeat("a", 10))
	s := "I speak Goat Latin"
	fmt.Printf("%T\n", s[0])

	fmt.Println(toGoatLatin(s))
}

func toGoatLatin(S string) string {
	comm := "aeiouAEIOU"
	goat := make(map[int]bool)
	for _, c := range comm {
		goat[int(c)] = true
	}

	t := func(s string) string {
		_, ok := goat[int(s[0])]
		tmp := s
		if !ok {
			tmp = s[1:] + s[:1]
		}
		return tmp + "ma"
	}

	buf := bytes.Buffer{}
	for i, s := range strings.Split(S, " ") {
		tmp := fmt.Sprintf("%s%s ",t(s),strings.Repeat("a",i+1))
		buf.WriteString(tmp)
	}
	return strings.TrimRight(buf.String()," ")
}
