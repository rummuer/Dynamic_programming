// Write a function 'canConstruct(target, wordBank)' that accepts a target string and array of strings
// The function should return a boolena indicating whether or not the 'target'
//can be construcyed by concatenating elements of the 'wordbank' array
// You may reuse elements of 'wordbank' as many times as needed.

//canConstruct(abcdef, [ab,abc,cd,def,abcd]) --> true
//canConstruct(skateboard,[bo,rd,ate,t,ska,sk,boar]) --> false
//canConstruct('',[dog,mouse,cat]) --> true

//Do not take strings in the middle. Always take prefixes only
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"ab", "abc", "cd", "def", "abcd"}
	s1 := []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}
	memo := make(map[string]bool)
	fmt.Println(canConstruct("abcdef", s, memo))
	fmt.Println(canConstruct("skateboard", s1, memo))
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.TrimPrefix(target, word)
			if canConstruct(suffix, wordBank, memo) == true {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

// Complexity
// Brute force -- Time --- O(n^m * m) --> the extra "m" comes from trimming the prefix
// Space -- O(m * m)  --> The extra "m" comes from storing new string on every stack call

// memoized
//Time --> O(n* m^2)
//Space --> O(m^2)
