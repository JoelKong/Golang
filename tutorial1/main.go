package main

import "fmt"

func main() {
	var intNum int
	var intArr [3]int32
	var intArr2 = [...]int32{1,2,3}
	fmt.Println(intNum)
	printMe("ff")
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder
}