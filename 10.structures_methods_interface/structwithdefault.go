package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (obj *Employee) Default() {
	fmt.Println(obj.Name)
	fmt.Println(obj.Age)

}

func main() {

	emp1 := Employee{Name: ""}
	emp1.Default()

	fmt.Println(emp1)

	emp2 := Employee{Age: 0}
	emp2.Default()

	fmt.Println(emp2)
}
