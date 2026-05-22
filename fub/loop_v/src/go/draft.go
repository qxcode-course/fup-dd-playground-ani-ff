package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Print("[ ")

	for {
		if a == b {
			break
		}

		if a%2 == 0 {
			a++
			continue
		}

		fmt.Print(a, " ")
		a++
	}

	fmt.Println("]")
}