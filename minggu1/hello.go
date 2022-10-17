package main

import "fmt"

func main() {
	var bilangan1 float64
	var bilangan2 float64
	var hasil float64

	fmt.Scan(&bilangan1, &bilangan2)

	hasil = (bilangan1 + bilangan2) / 2.0

	fmt.Println(hasil)
	
}