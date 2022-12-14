/**
Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal is to travel to the
bottom-right corner. You may only move down or right

In how many ways can you travel to the goal on a grid with dimension m * n?

Write a function 'gridTraveller(m,n)` that calculates this.

*/

package main

import "fmt"

func main() {
	fmt.Println(grid(3, 3))   // 6
	fmt.Println(grid(1, 1))   // 1
	fmt.Println(grid(2, 3))   // 3
	fmt.Println(grid(3, 2))   // 3
	fmt.Println(grid(18, 18)) // 2333606220

}

func grid(m int, n int) int {
	table := make([][]int, (m + 1))
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, n+1)
	}
	table[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j+1 <= n {
				table[i][j+1] += table[i][j]
			}
			if i+1 <= m {
				table[i+1][j] += table[i][j]
			}
		}
	}
	return table[m][n]
}

// Complexity
// Time - O(mn)
// Space - O(mn)
