/**
Write a function 'canSum(targetSum,numbers)` that takes in a targetSum and an array of numbers as arguments

The function should return a boolean indicating whether or not it is possible to generate the
targetSum using numbers from the array

You may use an element of the array as many times as needed

You may assume that all input numbers are nonnegative

*/

package main

import "fmt"

func main() {

	fmt.Println(canSum(7, []int{2, 3}))       //true
	fmt.Println(canSum(7, []int{5, 3, 4, 7})) //true
	fmt.Println(canSum(7, []int{3, 4, 5}))    //true
	fmt.Println(canSum(7, []int{2, 4}))       //false
	fmt.Println(canSum(8, []int{2, 3, 5}))    //true
	fmt.Println(canSum(300, []int{7, 14}))    //false

}
func canSum(targetNum int, numbers []int) bool {
	table := make([]bool, targetNum+1)
	table[0] = true
	for i := 0; i < len(table); i++ {
		if table[i] == true {
			for j := 0; j < len(numbers); j++ {
				if i+numbers[j] < len(table) {
					table[i+numbers[j]] = true
				}
			}
		}
	}
	return table[targetNum]
}
