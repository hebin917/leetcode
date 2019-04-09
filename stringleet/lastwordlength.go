package main

import "fmt"

func main() {
	s := " "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	length := 0
	wflag := false
	tail := len(s) -1

	for ; tail >= 0; tail-- {
		if s[tail] != ' ' && !wflag {
			wflag = true
			length++
			continue
		}else if s[tail]==' ' && wflag {
			break
		}
		if s[tail] != ' ' && wflag {
			length++
			continue
		}
	}
	return length
}
