/**
Write a function 'howSum(targetSum,numbers)` that takes in a targetSum and an array of numbers as arguments

The function should return a an array containing any combination of elements that add up to exactly the targetSum

If there is no combination that adds up to the targetSum, then return null

If there are multiple combinations possible, you may return any single one
You may use an element of the array as many times as needed

You may assume that all input numbers are nonnegative

*/

package main

import "fmt"

func main() {

	fmt.Println(howSum(7, []int{2, 3}))       // [3,2,2]
	fmt.Println(howSum(7, []int{5, 3, 4, 7})) // [4,3]
	fmt.Println(howSum(7, []int{2, 4}))       // null
	fmt.Println(howSum(8, []int{2, 3, 5}))    // [2,2,2,2]
	fmt.Println(howSum(300, []int{7, 14}))    // null

}

func howSum(targetNum int, numbers []int) []int {
	table := make([][]int, targetNum+1)
	table[0] = make([]int, 1)
	for i := 0; i < len(table); i++ {
		if table[i] != nil {
			for j := 0; j < len(numbers); j++ {
				if i+numbers[j] < len(table) {
					table[i+numbers[j]] = append([]int{numbers[j]}, table[i]...)
				}
			}
		}
	}

	return table[targetNum]
}

//complexity
// Time - O(m^2 * n) --> the extra "m" comes from iterating the sub-arrays
// Space - O(m^2)
