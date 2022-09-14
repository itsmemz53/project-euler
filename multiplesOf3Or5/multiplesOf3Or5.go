//Multiples of 3 or 5
//Problem 1
//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
	fmt.Println(findSumOfMultiplesOf3Or5(3, 5, 1000))
}

func findSumOfMultiplesOf3Or5(x int, y int, n int) int {
	return sumMultiples(x, n) + sumMultiples(y, n) - sumMultiples(x*y, n)
}

func sumMultiples(value int, n int) int {
	sum := (n - 1) / value
	return sum * (sum + 1) / 2 * value
}
