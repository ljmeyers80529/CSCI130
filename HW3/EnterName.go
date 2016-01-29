package main

import (
		"fmt"; 
		)

func main() {
	var s string

	fmt.Print("Enter name> ")
	fmt.Scanln(&s)
	fmt.Printf("Hello %s\n\n", s)
}

