package main
import "fmt"
func main() {
    
    company := map[int]string{
        22:"Tech",
        33:"Techco",
        44:"Techcovery",
    }
    for key, value:= range company {
        
       fmt.Println(key, value)
    }

}
