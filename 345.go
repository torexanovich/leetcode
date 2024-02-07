package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75

// func main() {
// 	fmt.Println(reverseVowels("itE")) //expected: leotcede
// }

/*
2pointers i j
for loop
check for vowel
hold the pointer on the vowel's index till the other one is met
then swap
continue
*/

func reverseVowels(s string) string {
	v, i, j, b := "aeiouAEIOU", 0, len(s)-1, []byte(s)

	for i < j {
		if !strings.Contains(v, string(s[i])) {
			i++
			continue
		}
		if !strings.Contains(v, string(s[j])) {
			j--
			continue
		}
		
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}
