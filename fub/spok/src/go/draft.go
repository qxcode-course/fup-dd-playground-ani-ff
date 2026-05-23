package main

import "fmt"

func inv(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func main() {
	var n int
	fmt.Scan(&n)

	if n == inv(n) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
