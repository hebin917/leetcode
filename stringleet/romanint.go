package main

import "fmt"

func main() {
	s:= "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	roman := make(map[uint8]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	tmp :=roman[s[0]]
	switch len(s) {
	case 0:
		return 0
	case 1:
		return roman[s[0]]

	}

	for i:=0;i<len(s)-1;i++ {
		if roman[s[i]] <roman[s[i+1]]  {
			tmp += roman[s[i+1]] - roman[s[i]]
		}else {
			tmp +=roman[s[i]]
		}
	}

	return tmp


}
