package main

import "fmt"

func main() {
	var a float64;

	fmt.Scan(&a)

	fmt.Print((a*a + 1/a)*(a*a + 1/a))
	
}