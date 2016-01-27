
// larry meyers
// CSCI 130
// 26-Jan-2016

package main

import "fmt"

func main() {
	var var3 uint8
	
	var1 := `Type test`
	var2 := 100.0
	var3 = 45

	fmt.Printf("var1 is of type %T with a value --> ", var1)
	fmt.Println(var1)

	fmt.Printf("var2 is of type %T with a value --> ", var2)
	fmt.Println(var2)
	
	fmt.Printf("var3 is of type %T with a value --> ", var3)
	fmt.Println(var3)
}
