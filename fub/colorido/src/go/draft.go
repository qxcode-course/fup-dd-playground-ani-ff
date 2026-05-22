package main

import "fmt"

func main() {
	var n int
	var p string

	fmt.Scan(&n, &p)

	fmt.Print("[ ")

	for i := 0; i <= 10; i++ {
		if i == n {
			continue
		}

		if i == 10 {
			fmt.Print("ceu ")
		} else {
			fmt.Print(i, p, " ")
		}

		if p == "d" {
			p = "e"
		} else {
			p = "d"
		}
	}

	fmt.Println("]")
}
