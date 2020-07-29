package main
import "fmt"
func main() {

  array := [7]string{"This", "is", "the", "tutorial","of", "Go", "language"}

  fmt.Println("Array:", array)

  slice := array[1:6]

  fmt.Println("Slice:", slice)
  fmt.Printf("Length of the slice: %d", len(slice))
  fmt.Printf("Capacity of the slice: %d", cap(slice))
}
