package main
import "fmt"
func main() {
    var wifiint, loginint, adminint int
    fmt.Scan(&wifiint, &loginint, &adminint)

    wifi := wifiint == 1
    login := loginint == 1
    admin := adminint == 1


    if !wifi {
        fmt.Println("you must connect to wifi")
    } else if !login {
        fmt.Println("you need to login first")
    } else if !admin {
        fmt.Println("you must to login as admin")
    } else {
        fmt.Println("done")
    }
}
