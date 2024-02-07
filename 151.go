package main

import (
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	i, j, res := 0, len(s)-1, ""

	for j >= 0 {
		for j >= 0 && s[j] == ' ' { j-- }

		i = j
		for i >= 0 && s[i] != ' ' { i-- }

		word := s[i+1 : j+1]

		if len(res) == 0 {
			res = word
		} else {
			res = res + " " + word
		}

		j = i
	}

	return res
}

// func main() {
// 	s1 := "the sky is blue"
// 	s2 := "  hello world  "
// 	s3 := "a good   example"

// 	println(reverseWords(s1)) // Output: "blue is sky the"
// 	println(reverseWords(s2)) // Output: "world hello"
// 	println(reverseWords(s3)) // Output: "example good a"
// }
