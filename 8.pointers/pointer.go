package main
import "fmt"
func main() {
  var len int = 10
  var p *int
  p = &len

  fmt.Println("Value stored in len = ", len)
  fmt.Println("Address of len = ", &len)
  fmt.Println("Value stored in variable p = ", p)
}
