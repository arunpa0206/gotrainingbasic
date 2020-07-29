package main

import "fmt"

func main() {

var arr1 [2][2] int
arr1[0][0] = 100
arr1[0][1] = 200
arr1[1][0] = 300
arr1[1][1] = 400

fmt.Println("Elements of array 2")
for p:= 0; p<2; p++{
for q:= 0; q<2; q++{
fmt.Println(arr1[p][q])

}
}
}
