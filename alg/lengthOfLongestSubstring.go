package alg

import "fmt"

func LenthOfLongestSubstring() {
	var s string = "dfsffdsafdzsx"
	fmt.Println(lenthOfLongestSubstring(s))
}

func lenthOfLongestSubstring(s string) (length int) {
	if s == "" {
		return
	}

	length = 0
	for i := 0; i < len(s); i++ {
		aa := map[byte]bool{}
		aa[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := aa[s[j]]; ok {
				break
			}
			aa[s[j]] = true
			sMax := j - i + 1
			if sMax > length {
				length = sMax
			}
		}
	}
	return
}
