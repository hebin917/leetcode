package main

import "fmt"

func guessNumber(n int) int {
	lo, hi := 0, n

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

func guess(n int) int {
	if n == 6 {
		return 0
	} else if n > 6 {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(guessNumber(6))
}
