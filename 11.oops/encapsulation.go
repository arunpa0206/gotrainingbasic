package main
import "fmt"
type Student struct{
  name string
  mark int

}
func(s *Student) getname() {
  return s.name
}
func(s *Student) Getmark() {
  return s.mark
}

func main() {
    s := &Student{
      name: "vijay",
      mark : 56,
    }
    fmt.Print(s.getname())
    fmt.Print(s.Getmark())
    
}
