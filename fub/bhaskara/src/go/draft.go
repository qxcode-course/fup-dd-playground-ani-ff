package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C float64
	fmt.Scan(&A, &B, &C)

	if A == 0 {
		fmt.Println("nao ha raiz real")
		return
	}

	delta := B*B - 4*A*C

	if delta > 0 {
		x1 := (-B + math.Sqrt(delta)) / (2 * A)
		x2 := (-B - math.Sqrt(delta)) / (2 * A)

		if math.Abs(x1) < 1e-9 {
			x1 = 0
		}
		if math.Abs(x2) < 1e-9 {
			x2 = 0
		}

		fmt.Printf("%.2f\n", x1)
		fmt.Printf("%.2f\n", x2)

	} else if delta == 0 {
		x := -B / (2 * A)

		if math.Abs(x) < 1e-9 {
			x = 0
		}

		fmt.Printf("%.2f\n", x)

	} else {
		fmt.Println("nao ha raiz real")
	}
	
}
