package main
import "fmt"
func main() {
  
  var a int = 1
   for a < 10 {
     if a == 5 {
       a = a + 1;
       continue;
     }
     fmt.Printf("value of a: %d\n", a);
     a++;
   }
 }

 
