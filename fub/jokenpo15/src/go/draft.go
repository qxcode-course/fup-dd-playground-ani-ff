package main

import "fmt"

func main() {
	x := 0 // criou antes para n zerar o 0 sempre

	// for { // loop infinitob
	//     fmt.Println(x)
	//     x += 1
	// }

	// for { // so sai quando o número ficar maior que 10
	//     x += 1
	//     fmt.Println(x)
	//     if x > 10 {
	//         break
	//     }
	// }

	for { // so os números ímpares
		x += 1
		if x%2 == 0 {
			continue // volta pro começo
		}
		fmt.Println(x)
		if x > 10 {
			break
		}
	}

	// for x := 0; x < 10; x += 1 { // o que faz antes de começar o laço, o que faz pra saber se entra no laço, e o que faz no final do laço
	//     fmt.Println(x)
	// }
}
