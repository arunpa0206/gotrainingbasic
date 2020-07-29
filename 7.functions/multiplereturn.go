package main
import "fmt"
func  nums() (int, int) {
    return 3, 7
}
func main() {
    a, b := nums()
    fmt.Println(a)
    fmt.Println(b)
    _, c := nums()
    fmt.Println(c)
}
