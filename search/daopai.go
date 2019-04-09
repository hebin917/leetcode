package main

import (
	"fmt"
	"strings"
)

func daopai(str []string) map[string][]string {
	keymap := make(map[string][]string)
	for i := 0; i < len(str); i++ {
		for k, v := range strings.Split(str[i], " ") {
			key_index := fmt.Sprintf("%d_%d", i, k)
			fmt.Printf("%T \n", v)
			if values,ok := keymap[v]; ok {
				values = append(values, key_index)
				keymap[v] = values
			} else {
				values := []string{}
				values = append(values, key_index)
				keymap[v] = values
			}
		}
	}
	return keymap
}

func main() {
	str := []string{"How are you", "are you ok", "i'm fine", "thank you"}
	strmap := daopai(str)
	for k, v := range strmap {
		fmt.Printf("k:  %v,  v: %v", k, v)
	}
}
