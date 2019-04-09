package main

import (
	"strings"
	"bytes"
	"fmt"
	"math"
)

func main() {

	//A := "a"
	//B := "aa"
	//fmt.Println(repeatedStringMatch(A, B))

	A:="abc"
	B:="cabcabca"

	fmt.Println(repeatedStringMatch(A, B))

	//A = "aaabbb"
	//B = "ab"
	//fmt.Println(repeatedStringMatch(A, B))
	//
	//A = "aaac"
	//B = "aac"
	//fmt.Println(repeatedStringMatch(A, B))

}

func repeatedStringMatch1(A string, B string) int {
	if len(A) == 0 || len(B) == 0 {
		return -1
	}
	length := 0
	step := len(B)
	buf := bytes.Buffer{}
	orig := A
	if strings.Index(A, B) != -1 {
		return length + 1
	}
	for i := 0; strings.Index(orig, B) == -1; i++ {
		if i > step+1 {
			break
		}
		buf.WriteString(A)
		orig = buf.String()
		length ++
	}

	//if len !=-1 {
	//	return len+1
	//}
	return length
}

func repeatedStringMatch(A string, B string) int {
	q := int(math.Ceil(float64(len(B) - 1)/float64(len(A)+1)))
	for i := range (make([]int, 2)) {
		s := strings.Repeat(A, q+i)
		if strings.Index(s, B) != -1 {
			return q + i
		}
	}

	return -1
}



