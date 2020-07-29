package main
import "fmt"
func CallByValue(x int) {
    x = 70
}

func main() {
    var x int = 10
    fmt.Printf("Before Function Call, value of x is = %d", x)
    CallByValue(x)
    fmt.Printf("\nAfter Function Call, value of x is = %d", x)
}
