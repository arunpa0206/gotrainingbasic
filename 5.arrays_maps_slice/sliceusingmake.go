package main
import "fmt"
func main() {
  //make(type of slice, length of slice, capacity of slice)
  var slice = make([]int, 4, 7)

  fmt.Printf("Slice  = %v, \nlength = %d,    \ncapacity = %d\n",slice,
   len(slice), cap(slice))
}


