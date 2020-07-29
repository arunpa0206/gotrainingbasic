package main
import (
    "container/list"
    "fmt"
)
func main() {
  fruits := list.New()
  fruits.PushBack("apple")
  fruits.PushBack("banana")
  fruits.PushFront("mango")
  for fruitlist := fruits.Front(); fruitlist!= nil; fruitlist=fruitlist.Next() {
    fmt.Println(fruitlist.Value)
  }
}
