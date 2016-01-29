package main

import (
		"fmt"; 
		)

func main() {

	// output will be true.
	// true && false => false
	// false && true => false
	// false && false => false, but given the not (!) this becomes true.
	// any false within an AND equation will be false. 
	// any true with in an OR equation will return a true.
	// thus, given fasle OR false OR true gives a true.
	fmt.Println((true && false) || (false && true) || !(false && false))
}
