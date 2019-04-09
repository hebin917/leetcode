package main


func numSpecialEquivGroups(A []string) int {
	total  := make(map[int32[]]int)
	for _, spec := range A {
		count := make([]int32, 52)
		for i, v := range spec {
			count[(v-'a') + int32(26*(i%2))]++
		}

	}
}

func main() {

}
