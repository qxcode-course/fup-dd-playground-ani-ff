package main
import "fmt"
func main() {
	var salario float64
	var aumento float64
	fmt.Scan(&salario)
	if salario <= 1000.00 {
		aumento = 0.20
	} else if salario <= 1500.00 {
		aumento = 0.15
	} else if salario <= 2000.00 {
		aumento = 0.10
	} else {
		aumento = 0.05
	}
	novoSalario := salario + salario*aumento
	fmt.Printf("%.2f\n", novoSalario)
}
