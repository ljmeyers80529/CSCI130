package main

import (
		"fmt"; 
		)

func main() {
	for i := 0; i < 100; i++ {
		f := (i % 3) == 0
		b := (i % 5) == 0
		if (f && b) {
			fmt.Println("FizzBuzz ")
		} else {
			if (f && !b) {
				fmt.Println("Fizz ")
			} else {
				if (!f && b) {
					fmt.Println("Buzz ")
				} else {
					fmt.Println(i)
				}
			}
		}
	}
}

