package main

import (
		"fmt"; 
		)

func main() {
	large := LargestNumber( 56,2,16,11,67,88,3,70,31)
	fmt.Printf("Largest number is : %d", large)
}

func LargestNumber(n ...int) int {
	var cl = 0;
	
	if len(n) == 0 {
		return 0
	}
	for _, v := range n {
		if v > cl {
			cl = v;
		}
	}
	return cl
}