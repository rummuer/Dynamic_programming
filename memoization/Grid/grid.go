package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int)
	fmt.Println(grid(18, 18, m))
}

func grid(m int, n int, m1 map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if _, ok := m1[key]; ok {
		return m1[key]
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	m1[key] = grid(m-1, n, m1) + grid(m, n-1, m1)
	return m1[key]
}

// Brute force recursion - O(2^(n+m)) O(n+m)

// Dynamic - O(n+m) O(n+m)
