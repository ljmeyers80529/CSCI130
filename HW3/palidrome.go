package main

import (
		"fmt";
		)

func main() {
	var num1, num2, lp int = 0,0,0
	var rslt int 
	
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			rslt = i * j
			if isPalidrome(rslt) && (rslt > lp) {
				lp = rslt
				num1 = i
				num2 = j
			}
		}
	}
	
	fmt.Printf("Largest palidrome from 2 3-digit numbers = %d.\nValues found are i = %d and j = %d", lp, num1, num2)
}

func isPalidrome(pal int) bool {
	var vs string = fmt.Sprint(pal)
	var slen int = len(vs)
	var size int = slen / 2
	var secHalf int = size + (slen % 2)
	var first string = vs[0:size]
	var second string = vs[secHalf:slen]
	var reverse string = ""
	var rev []byte
	
 	for i:=len(second) - 1; i >= 0; i-- {
		rev = append(rev, second[i])
	}
	reverse = string(rev)
	return first == reverse
}


// calculate the largest palidrome result from 2 3-digit numbers, numbers between 100-999.
// problem is to detemine of a given number is a palidrome or not.
// have to take into account if the number has an even number of digitis or an odd number.
// Solution is to convert the number to a string and split the string in half.
// need to take the second half of the split string and reverse the order of the digits then compare the two halfs.
// if the number has an odd number of digits, the central digit is simply ignored.
// To get the largest palidrome, just simply loop through all the numbers, multiply, check if palidrome, and save the largest one and the multipliers.