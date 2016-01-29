package main

import (
		"fmt"; 
		)

func main() {
	var v int = 97

	Testfunc := func(v int) (half int, even bool) {
			 half = v / 2
			 even = (v % 2) == 0
			 return
			 }
			 
	h, e := Testfunc(v)
	fmt.Printf("Function expression: Arg = %d, Half = %d, Even? = %t\n", v, h, e)
	v++
	h, e = Testfunc(v)
	fmt.Printf("Function expression: Arg = %d, Half = %d, Even? = %t\n\n", v, h, e)
 }
 