package main

import "fmt"

/*
In number theory, the integer square root (isqrt) of a positive integer n is the positive integer m which is the greatest integer less than or equal to the square root of n,
根据牛顿定律：
	x2=n
	x=n/x
	x/2 =n/x*2
	x-x/2=n/(x*2)
	x = x/2 + n/(x*2)
    x= 1/2(x + n/x)
 */

func isqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}


/**
利用 mid*mid <x <(mid+1) *(mid+1)
 */
func mysqrt(x int) int{
	if x==0 {
		return 0
	}
	low,high:=1,x

	for low <=high {
		mid := low + (high-low)/2
		if mid > x/mid {
			high=mid -1
		}else {
			if mid+1 >x/(mid +1){
				return mid
			}
			low=mid +1
		}
	}
	return -1
}

func main() {
	fmt.Println(isqrt(27))
	fmt.Print(mysqrt(27))
}
