package main
import "fmt"
func main() {

    fruits := [3]string{"Apple", "Mango", "Banana"}

    for index,element := range fruits{
      
      fmt.Println(index)
      fmt.Println(element)
    }
  }
