package main

import "fmt"

func main() {
	var a int;

	fmt.Scan(&a)

	mine := a / 100
	oboli := (a - (mine * 100))*6
	attic := mine / 60
	fmt.Print(attic, mine - (attic * 60), oboli)
	
}