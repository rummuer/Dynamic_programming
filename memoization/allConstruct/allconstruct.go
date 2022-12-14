/**
Write a function "allConstruct(target, wordBank)" that accepts a target string and an array of strings.

The function should return a 2D array containg all of the ways that the "target" can be constructed by
concatenating elements of the "wordBank" array. Each element of the 2D array should represent one combination
that constructs the target

You may reuse elements of "wordBank" as many times as needed
*/

/**
allConstruct(hello,[cat,dog,mouse]) ---> []
allConstruct("",[cat,dog,mouse]) -->[[]]
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}
	s1 := []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}
	s2 := []string{"e", "ee", "eee", "eeee", "eeee", "eeeee", "eeeeee"}
	memo := make(map[string][][]string)
	fmt.Println(allConstruct("abcdef", s, memo))
	fmt.Println(allConstruct("skateboard", s1, memo))
	fmt.Println(allConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", s2, memo))
}
func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		m := make([][]string, 1)
		fmt.Println("m=", m)
		return m
	}
	result := make([][]string, 0, 10)
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.TrimPrefix(target, word)
			suffixWays := allConstruct(suffix, wordBank, memo)
			targetWays := make([][]string, 0, 10)
			for _, arr := range suffixWays {
				arr = append([]string{word}, arr...)
				targetWays = append(targetWays, arr)
			}
			for _, arr := range targetWays {
				result = append(result, arr)
			}

		}
	}
	memo[target] = result
	return result
}

// Complexity is  Sames as canConstruct
// The brute force and DP complexities are same
// Time --> O(n^m) and space --> O(m)
// The time complexity is not changed due to the fact that the worst is to find all the possibilitites
