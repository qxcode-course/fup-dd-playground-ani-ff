package main
import "fmt"
func main() {
    fome := true
    dinheiro := true
    if !dinheiro {
        fmt.Println("Vou ficar com fome, estou liso")
    } else if !fome {
        fmt.Println("Não estou com fome, vou economizar")
    } else {
        fmt.Println("Vou almoçar")
    }


    // idade := 6789
    // if idade > 0 && idade <= 12 {
    //     fmt.Println("Criança")
    // } else if idade < 18 {
    //     fmt.Println("Adolescente")
    // } else if idade < 65 {
    //     fmt.Println("Adulto")
    // } else if idade >= 65 && idade < 100 {
    //     fmt.Println("Idoso")
    // } else {
    //     fmt.Println("Não existe")
    // }



    // idade := 6789
    // switch{
    // case idade < 13:
    //     fmt.Println("criança")
    // case "seg", "ter", "qua":
    //     fmt.Println("aula")
    // }


    // if idade > 0 && idade <= 12 {
    //     fmt.Println("Criança")
    // } else if idade < 18 {
    //     fmt.Println("Adolescente")
    // } else if idade < 65 {
    //     fmt.Println("Adulto")
    // } else if idade >= 65 && idade < 100 {
    //     fmt.Println("Idoso")
    // } else {
    //     fmt.Println("Não existe")
    // }












    // PEDRA PAPEL TESOURA --------------------------
    // jog1 := "pedra"
    // jog2 := "papel"
    // fmt.Scan()
    // if jog1 == jog2 {
    //     fmt.Println("Empate")
    // } else if (jog1 == "pedra" && jog2 == "tesoura") || 
    //     (jog1 == "papel" && jog2 == "pedra") ||
    //     (jog1 == "tesoura" && jog2 == "papel") {
    //         fmt.Println("jog1")
    //     } else {
    //         fmt.Println("jog2")
    //     }

    // JOQUE EM PÔ --------------------------------
    // jog1 := 0
    // jog2 := 1
    // jog3 := 0
    // if jog1 == jog2 && jog2 == jog3 {
    //     fmt.Println("empate")
    // } else if jog2 == jog3 {
    //     fmt.Println("jog1")
    // } else if jog1 == jog3 {
    //     fmt.Println("jog2")
    // } else {
    //     fmt.Println("jog3")
    // }

    
}
