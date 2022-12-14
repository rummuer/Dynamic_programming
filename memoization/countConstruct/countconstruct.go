// write a function "countConstruct(target,wordBank)" that accepts a target string and an array of strings
// The function should return the number of ways that the "target" can be constructed by concatenating
//elements of the "wordBank" array
// You may reuse elements of "wordBank" as many times as needed

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"ab", "abc", "cd", "def", "abcd"}
	s1 := []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}
	s2 := []string{"e", "ee", "eee", "eeee", "eeee", "eeeee", "eeeeee"}
	memo := make(map[string]int)
	fmt.Println(canConstruct("abcdef", s, memo))
	fmt.Println(canConstruct("skateboard", s1, memo))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", s2, memo))
}
func canConstruct(target string, wordBank []string, memo map[string]int) int {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return 1
	}
	totalCount := 0
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.TrimPrefix(target, word)
			numWays := canConstruct(suffix, wordBank, memo)
			totalCount += numWays
		}
	}
	memo[target] = totalCount
	return totalCount
}

// Complexity is  Sames as canConstruct
