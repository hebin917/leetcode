package main

import (
	"bytes"
	"fmt"
	)

func main() {
	fmt.Println(toLowerCase("HffelLLG"))

}

func toLowerCase(str string) string {
	ret := bytes.Buffer{}
	for _, v := range str {

		if v >= 'A' && v <= 'Z' {
			ret.WriteByte(byte(v - 'A' + 'a'))
			continue
		}
		ret.WriteByte(byte(v))
	}
	return ret.String()
}
