package main

import "fmt"

func main() {
	var a int;

	fmt.Scan(&a)

	fmt.Print(((a % 4 == 0) || (a % 400 == 0)) && !(a % 100 == 0))
	
}