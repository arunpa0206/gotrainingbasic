package main
import "fmt"
func main() {
  fmt.Println("1")
  Counter()
}
func Counter() {
  i := 1
  for {
    fmt.Printf("%3d", i)
    i++
    if i == 10 {
      break
    }
  }
}
