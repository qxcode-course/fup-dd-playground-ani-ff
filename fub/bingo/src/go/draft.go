package main
import "fmt"
func main() {
	cartela := [4][4]int{
		{1, 9, 27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66},
	}
	numeros := make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Scan(&numeros[i])
	}
	contador := 0
	for _, num := range numeros {
		encontrou := false

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if cartela[i][j] == num {
					contador++
					encontrou = true
					break
				}
			}
			if encontrou {
				break
			}
		}
	}

	fmt.Println(contador)
}