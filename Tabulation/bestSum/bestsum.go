/**
Write a function 'bestSum(targetSum,numbers)` that takes in a targetSum and an array of numbers as arguments

The function should return a an array containing the shortest combination of elements that add up to exactly the targetSum

If there is no combination that adds up to the targetSum, then return null

If there is a tie for the shortest combination, you may return any one of the shortest
You may use an element of the array as many times as needed

You may assume that all input numbers are nonnegative

*/

package main

import "fmt"

func main() {

	fmt.Println(bestSum(7, []int{2, 3}))          // [3,2,2]
	fmt.Println(bestSum(7, []int{2, 4}))          // null
	fmt.Println(bestSum(8, []int{2, 3, 5}))       // [3,5]
	fmt.Println(bestSum(8, []int{1, 4, 5}))       // [4,4]
	fmt.Println(bestSum(100, []int{1, 2, 5, 25})) // [25,25,25,25]

}

func bestSum(targetNum int, numbers []int) []int {
	table := make([][]int, targetNum+1)
	table[0] = make([]int, 1)
	for i := 0; i < len(table); i++ {
		if table[i] != nil {
			for j := 0; j < len(numbers); j++ {
				if i+numbers[j] < len(table) {
					arr := append([]int{numbers[j]}, table[i]...)
					if (table[i+numbers[j]] == nil) || (len(arr) < len(table[i+numbers[j]])) {
						table[i+numbers[j]] = arr
					}
				}
			}
		}
	}

	return table[targetNum]
}

//complexity
// Time - O(m^2 * n) --> the extra "m" comes from iterating the sub-arrays
// Space - O(m^2)
