package main
import "fmt"
  type emp struct {
    name   string
    empid  int
    salary int
}
func main() {
      emp1 := new(emp)
    emp1.name = "XYZ"
    emp1.empid = 1555
    emp1.salary = 25000
    fmt.Println(emp1)
 }
