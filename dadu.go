package main

import "fmt"

func main() {
	var a, b, c int;

	fmt.Scan(&a, &b, &c)

	fmt.Print((a % 2 == 0) && (b % 2 == 0) && (c % 2 == 0))
	
}