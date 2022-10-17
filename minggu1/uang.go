package main

import "fmt"

func main() {
	var a int;

	fmt.Scan(&a)

	sepuluh := a / 10000
	lima := (a - (sepuluh * 10000)) / 5000
	seribu := (a - (sepuluh * 10000) - (lima * 5000)) / 1000
	fmt.Print(sepuluh, lima, seribu)
	
}