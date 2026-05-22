package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Print("[ ")
	for ; a > b; a-- {
		fmt.Print(a, " ")
	}
	fmt.Println("]")
}
