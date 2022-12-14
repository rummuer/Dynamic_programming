/* write a function "countConstruct(target,wordBank)" that accepts a target string and an array of strings
The function should return the number of ways that the "target" can be constructed by concatenating
elements of the "wordBank" array
You may reuse elements of "wordBank" as many times as needed */

package main

import "fmt"

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                                                       //1
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))                                        //0
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeee", "eeeee", "eeeeee"})) //0
	fmt.Println(canConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                                                       //2
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))                                      //4

}
func canConstruct(targetString string, words []string) int {
	table := make([]int, len(targetString)+1)
	table[0] = 1
	for i := 0; i < len(table); i++ {
		if table[i] != 0 {
			for j := 0; j < len(words); j++ {
				if i+len(words[j]) < len(table) {
					if targetString[i:i+len(words[j])] == words[j] {
						table[i+len(words[j])] += table[i]
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
