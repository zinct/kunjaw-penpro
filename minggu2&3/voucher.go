package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string;

	fmt.Scan(&a)

	d1, _ := strconv.ParseInt(string(a[0]),10, 0)
	d3, _ := strconv.ParseInt(string(a[2]),10, 0)
	d4, _ := strconv.ParseInt(string(a[3]),10, 0)

	diskon := (d1 % 2 == 0)
	cashback := (d1 + d3) % d4 == 0

	fmt.Print(diskon, cashback, (diskon || cashback) && !(diskon && cashback))
	
}