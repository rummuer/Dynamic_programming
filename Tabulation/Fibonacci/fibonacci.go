/**
Write a function 'fib(n)' that takes in a number as an argument
The function should return the n-th number of the Fibonacci sequence

The 0th number of the sequence is 0
The 1st number of the sequence is 1

To generate the next number of the sequence, we sum the previous two
*/

package main

import "fmt"

func main() {
	fmt.Println(fib(6))  // 8
	fmt.Println(fib(7))  // 13
	fmt.Println(fib(50)) // 12586269025
}

func fib(n int) int {
	table := make([]int, n+1)
	table[1] = 1
	for i := 0; i <= len(table); i++ {
		if i+1 < len(table) {
			table[i+1] += table[i]
		}
		if i+2 < len(table) {
			table[i+2] += table[i]
		}
	}
	return table[n]
}

// Complexities

// Time --> O(n) length of the array
// Space --> O(n) length of the array
