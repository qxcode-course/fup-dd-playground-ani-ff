package main

import "fmt"

func main() {
	var a, b, p int
	fmt.Scan(&a, &b)

	if a < b {
		p = 1
	} else {
		p = -1
	}

	fmt.Print("[ ")
	for ; a != b; a += p {
		fmt.Print(a, " ")
	}
	fmt.Println("]")
}
