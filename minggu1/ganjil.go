package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64;

	fmt.Scan(&a)

	fmt.Print(int(math.Abs(a))%2 > 0 && int(a) < 0)
	
}