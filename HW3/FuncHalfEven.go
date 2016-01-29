package main

import (
		"fmt"; 
		)

func main() {
	var v int = 97

	h,e := TestFunction(v)
	fmt.Printf("Normal function: Arg = %d, Half = %d, Even? = %t\n", v, h, e)
	v++
	h,e = TestFunction(v)
	fmt.Printf("Normal function: Arg = %d, Half = %d, Even? = %t\n\n", v, h, e)
 }
 
 func TestFunction(v int) (half int, even bool) {
	half = v / 2
	even = (v % 2) == 0
	return
}

