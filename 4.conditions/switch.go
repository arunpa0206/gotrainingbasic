package main
import "fmt"
func main() {
  var number = 20
  switch number {
  case 10:
    fmt.Println("10")
  case 20:
    fmt.Println("20")
  case 30:
    fmt.Println("30")
  default:
    fmt.Println("Not in 10,20,30")
}
}
