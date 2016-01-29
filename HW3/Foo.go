package main

import (
		"fmt"; 
		)

func main() {
	aslice := []int {6,2,16,11,67,88,3,70,31}
	sum1 := foo()
	sum2 := foo(1,3)
	sum3 := foo(1,6,11)
	sum4 := foo(aslice...)
	fmt.Printf("Sums of numbers : %d, %d, %d, %d", sum1, sum2, sum3, sum4)
}

func foo(n ...int) int {
	var cl = 0;
	for _, v := range n {
		cl += v
	}
	return cl
}