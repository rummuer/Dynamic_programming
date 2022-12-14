package main

import "fmt"

func main() {
	m := make(map[int]int)
	fmt.Println(fib(40, m))
}

func fib(f int, m map[int]int) int {
	if _, ok := m[f]; ok {
		return m[f]
	}
	if f <= 2 {
		return 1
	}
	m[f] = fib(f-1, m) + fib(f-2, m)
	return m[f]
}

// Time and space complexities

// Brute force - recursion - O(2^n) and O(n)
// Dynamic Mem - O(n) and O(n)
