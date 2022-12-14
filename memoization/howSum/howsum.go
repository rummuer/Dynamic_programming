// How will you do it --> combinatoric problem
package main

import "fmt"

func main() {
	m := make(map[int][]int)
	s := make([]int, 0, 4)
	s = append(s, 7, 14)
	fmt.Println(howSum(300, s, m))
}

func howSum(targetSum int, nums []int, mem map[int][]int) []int {

	if _, ok := mem[targetSum]; ok {
		return mem[targetSum]
	}
	if targetSum < 0 {
		return nil
	}
	if targetSum == 0 {
		emptySlice := []int{}
		return emptySlice
	}
	for _, num := range nums {
		remainder := targetSum - num
		result := howSum(remainder, nums, mem)
		if result != nil {
			result = append(result, num)
			mem[targetSum] = result
			return mem[targetSum]
		}
	}
	mem[targetSum] = nil
	return nil
}

// Brute force - O(n^m * m) and O(m) the extra m in time complexity comes from the fact that copying elements in line 28 takes m time.
// Dynamic O(n *m * m ) and O(m*m). The extra m in space comes from storing the elements for all the subtress
