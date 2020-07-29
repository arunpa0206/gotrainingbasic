package main
import "fmt"
func main() {
  person := struct {
    firstName, lastName string
    salary int
    fullTime bool
   }{
     firstName: "Monica",
     lastName: "Geller",
     salary:    1200,
  }
    fmt.Println(person)
}
