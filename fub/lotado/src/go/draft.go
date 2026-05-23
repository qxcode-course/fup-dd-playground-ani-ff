package main

import "fmt"

func main() {
	var c, m, p int
	fmt.Scan(&c)

	for {
		fmt.Scan(&m)
		p += m

		if p == 0 {
			fmt.Println("vazio")
		} else if p < c {
			fmt.Println("ainda cabe")
		} else if p < 2*c {
			fmt.Println("lotado")
		} else {
			fmt.Println("hora de partir")
			break
		}
	}
}
