package main

import (
	"fmt"
)

func main() {
	var intSlice = new([7]int)[0:4]
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
}
