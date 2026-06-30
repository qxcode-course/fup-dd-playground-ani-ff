package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	matriz := make([][]int, n)

	for i := 0; i < n; i++ {
		matriz[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	maiorColuna := 0
	maiorValor := -1

	for col := 0; col < n; col++ {
		soma := 0

		for lin := 0; lin < n; lin++ {
			valor := matriz[lin][col]
			soma += valor * valor
		}

		if soma > maiorValor {
			maiorValor = soma
			maiorColuna = col
		}
	}

	fmt.Println(maiorColuna)
}