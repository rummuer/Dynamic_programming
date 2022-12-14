// Write a function 'canConstruct(target, wordBank)' that accepts a target string and array of strings
// The function should return a boolean indicating whether or not the 'target'
//can be construcyed by concatenating elements of the 'wordbank' array
// You may reuse elements of 'wordbank' as many times as needed.

//canConstruct(abcdef, [ab,abc,cd,def,abcd]) --> true
//canConstruct(skateboard,[bo,rd,ate,t,ska,sk,boar]) --> false
//canConstruct('',[dog,mouse,cat]) --> true

//Do not take strings in the middle. Always take prefixes only

package main

import "fmt"

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                              //true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                               //false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                             //true
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) //false

}

func canConstruct(targetString string, words []string) bool {
	table := make([]bool, len(targetString)+1)
	table[0] = true
	for i := 0; i < len(targetString); i++ {
		if table[i] == true {
			for j := 0; j < len(words); j++ {
				if i+len(words[j]) < len(table) {
					if targetString[i:i+len(words[j])] == words[j] {
						table[i+len(words[j])] = true
					}
				}
			}
		}

	}

	return table[len(targetString)]
}

// complexity
//Time - O(m^2 * n)
//Space - O(m)
