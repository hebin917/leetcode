package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeCircle("LL"))
	fmt.Println(judgeCircle("UDLRL"))
	fmt.Println(judgeCircle("LR"))
	fmt.Println(judgeCircle("DL"))
}

func judgeCircle0(moves string) bool {
	res := make(map[int32]int)
	for _, v := range moves {
		res[v] += 1
	}

	if res['L'] != res['R'] {
		return false
	}
	if res['U'] != res['D'] {
		return false
	}
	return true

}

func judgeCircle(moves string) bool {
	var x, y int
	for _, v := range moves {
		switch v {
		case 'L':
			x += 1
		case 'R':
			x -= 1
		case 'D':
			y -= 1
		case 'U':
			y += 1
		}

	}

	return x == 0 && y == 0

}
