package main

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A, &B)

	fmt.Print("[ ")

	for i, j := A, B; i <= B; i, j = i+1, j-1 {
		fmt.Print(i, " ", j)

		if i != B {
			fmt.Print(" ")
		}
	}

	fmt.Println(" ]")
}
