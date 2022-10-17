package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b = 3, 12;

	fmt.Scan(&a, &b)

	r, _ := new(big.Float).SetPrec(3).SetString("3")

	fmt.Print((1.0/3.0)*3.14*(a*a*b))
	
}