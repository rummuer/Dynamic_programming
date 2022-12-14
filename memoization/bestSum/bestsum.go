// What is the best way to do it - Optimisation problem
package main

import "fmt"

func main() {
	m := make(map[int][]int)
	s := make([]int, 0, 4)
	s = append(s, 5, 3, 4, 7)
	fmt.Println(bestSum(7, s, m))
}

func bestSum(targetSum int, nums []int, mem map[int][]int) []int {

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
	best := []int{}
	fmt.Println(len(best))
	for _, num := range nums {
		remainder := targetSum - num
		result := bestSum(remainder, nums, mem)
		if result != nil {
			result = append(result, num)

			if (len(best) == 0) || (len(result) < len(best)) {
				best = result
				//fmt.Println(best, result)
			}
		}
	}
	mem[targetSum] = best
	return best
}

// Brute force - O(n^m * m) and O(m * m) the extra m in time complexity comes from the fact that copying elements in line 28 takes m time.
// Dynamic O(n *m * m ) and O(m*m). The extra m in space comes from storing the elements for all the subtress

// canSum --> Decision problem
// howSum -> Combinatoric problem
// bestSum -> Optimization problem
