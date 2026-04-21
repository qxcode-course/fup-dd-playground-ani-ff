package main
import "fmt"

// funções impuras 
func mostrar_vetor() {
    arr := []int {9, 8, 4, 5, 6, 2, 3}
    fmt.Print("[ ")
    for i, valor := range arr {
        if i != 0 {
            fmt.Print(", ")
        }
        fmt.Printf("%v", valor)
    }
     fmt.Print("]\n")

}

func main() {

    // int
    // float64
    // bool
    // rune
    // string

    // [] (lista) de int

    // var lista []int = []int{1, 4, 2, 5, 7} // lista de estática
    // fmt.Println(lista)

    // var nomes []string = []string{"uva", "ovo", "eva"}
    // fmt.Println(nomes[2])

    // var qtd int
    // fmt.Scan(&qtd)
    // var idades []int = make([]int, qtd) // make - cria 10 elementos vazio // consigo digitar a quantidade
    // fmt.Println(idades)

    // PERCORRER A LISTA
    // arr := []int {9, 8, 4, 5, 6, 2, 3}
    // fmt.Print("[ ") // so pra deixar bonito
    // for i := 0; i < len(arr) - 1; i++ { // i := 2 (de onde começa) e len(arr) - 2 (onde termina)
    //     fmt.Printf("%d ", arr[i])
    // }
    // fmt.Print("]\n") // so pra deixar bonito


    // arr := []int {9, 8, 4, 5, 6, 2, 3}
    // fmt.Print("[ ")
    // for i, valor := range arr {
    //     if i != 0 {
    //         fmt.Print(", ")
    //     }
    //     fmt.Printf("%v", valor)
    // }
    //  fmt.Print("]\n")


    
}
