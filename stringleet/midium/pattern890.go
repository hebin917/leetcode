package midium

import "fmt"

func main() {
	words := []string{"abc","deq","mee","aqq","dkd","ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words,pattern))
}

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	for _, s := range words {

		if len(s) == len(pattern) {
			word_to_pattern_dict := make(map[rune]rune)
			pattern_to_word_dict := make(map[rune]rune)
			flag := true
			for j, c := range pattern {
				_, wok := word_to_pattern_dict[c]
				_, pok := pattern_to_word_dict[c]
				if !wok && !pok {
					pattern_to_word_dict[c] = rune(s[j])
					word_to_pattern_dict[rune(s[j])] = c

					//word_to_pattern_dict[rune(s[j])] = c
					//pattern_to_word_dict[c] = rune(s[j])
				} else if wok && pok {
					continue
				} else {
					flag = false
					break
				}
			}
			if flag {
				res=append(res, s)
			}
		}
	}
	return res

}

/*
func zip(lists ...[]string) func() []string {
	zip := make([]string, len(lists))
	i := 0
	return func() []string {
		for j := range lists {
			if i >= len(lists[j]) {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}*/
