// factorial problem solved with recursion
package main

func main() {
	println(factorial(5))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
