package main

import (
	"github.com/golang-collections/collections/stack"
	"fmt"
)

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	st := stack.Stack{}
	for _, w := range s {
		switch w {
		case '[':
			st.Push(']')
		case '{':
			st.Push('}')
		case '(':
			st.Push(')')
		default:
			if st.Len() == 0 || st.Pop() != w {
				return false
			}
		}

	}
	return true
}
