package main

import (
	"fmt"
	)

func main() {

	fmt.Println(3%2)
	a := "1011"
	b := "1011"
	fmt.Println(addBinary(a,b))
}


func addBinary(a string, b string) string {
	res:=""
	alen := len(a) - 1
	blen := len(b) - 1
	var c uint8 = 0

	for alen >= 0 || blen >= 0 || c == 1 {
		if alen >= 0 {
			c += a[alen] - '0'
		} else {
			c += 0
		}

		if blen>=0 {
			c += b[blen]-'0'
		}else {
			c+=0
		}
		alen--
		blen--
		res = fmt.Sprintf("%d%s",c%2,res)
		c=c/2
	}
	return res
}


