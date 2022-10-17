package main

import (
	"fmt"
)

func main() {
	var p, l, t, b int;

	fmt.Scan(&p, &l, &t, &b)

	v := (p * l * t)
	fmt.Println((v / 1000000) < 1 && b / 1000 < 30)
	
}