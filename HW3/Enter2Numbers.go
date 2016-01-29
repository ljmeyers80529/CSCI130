package main

import (
		"fmt"; 
		)

func main() {
	var n1, n2 int
	
	fmt.Print("Enter 2 numbers, a small number then a large number> ")
	fmt.Scanf("%d %d", &n1, &n2)
	fmt.Printf("Remainder of  %d/%d = %d", n2, n1, n2%n1)
}

