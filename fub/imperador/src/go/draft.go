package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	matriz := make([][]string, n)

	leaoLinha, leaoColuna := -1, -1

	// Leitura da matriz
	for i := 0; i < n; i++ {
		matriz[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matriz[i][j])

			if matriz[i][j] == "L" {
				leaoLinha = i
				leaoColuna = j
			}
		}
	}

	gladiadores := 0
	condenados := 0

	// Contagem de pontos
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			// Se existe leão, anula linha e coluna dele
			if (leaoLinha != -1 && i == leaoLinha) || (leaoColuna != -1 && j == leaoColuna) {
				continue
			}

			if matriz[i][j] == "G" {
				gladiadores += 2
			} else if matriz[i][j] == "C" {
				if i+j == n-1 { // diagonal secundária
					condenados += 2
				} else {
					condenados++
				}
			}
		}
	}

	if gladiadores > condenados {
		fmt.Println("Gladiadores")
	} else if condenados > gladiadores {
		fmt.Println("Condenados a morte")
	} else {
		fmt.Println("Ninguem")
	}
}