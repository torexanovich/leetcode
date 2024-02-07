package main

import (
    "strings"
)

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    
    end := len(s) - 1
    var start int
    var reversed string
    
    for end >= 0 {
        // Move end pointer to the end of the current word
        for end >= 0 && s[end] == ' ' {
            end--
        }
        
        // Move start pointer to the beginning of the current word
        start = end
        for start >= 0 && s[start] != ' ' {
            start--
        }
        
        // Extract and append word to the result string
        word := s[start+1 : end+1]
        if len(reversed) == 0 {
            reversed = word
        } else {
            reversed = reversed + " " + word
        }
        
        // Move end pointer to the position just before the current word
        end = start
    }
    
    return reversed
}

func main() {
    // Test cases
    s1 := "the sky is blue"
    s2 := "  hello world  "
    s3 := "a good   example"
    
    // Output the reversed strings
    println(reverseWords(s1)) // Output: "blue is sky the"
    println(reverseWords(s2)) // Output: "world hello"
    println(reverseWords(s3)) // Output: "example good a"
}
