// Can you do it? yes/no - Decision problem

package main

import "fmt"

func main() {
	m := make(map[int]bool)
	s := make([]int, 0, 2)
	s = append(s, 7, 14)
	fmt.Println(canSum(300, s, m))
}

func canSum(targetSum int, nums []int, m map[int]bool) bool {
	if targetSum < 0 {
		return false
	}
	if targetSum == 0 {
		return true
	}
	if _, ok := m[targetSum]; ok {
		return m[targetSum]
	}
	for _, num := range nums {
		r := targetSum - num
		m[targetSum] = canSum(r, nums, m)
		if m[targetSum] == true {
			return true
		}
	}

	return false

}

// Brute force recurrsion O(n^m)  and O(m) where n is array length and m is target sum
// Dynamic O(m*n) and O(m)
