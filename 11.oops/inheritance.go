package main

import (
    "fmt"
)

// creating an interface
type Vehicle interface {
    Move()
    Honk()
}

type Car struct {
    Color string
    Brand string
}

// Car has method Move(), automatically stated implement Vehicle interface
func (c *Car) Move() {
    fmt.Print("Car is moving")
}

func(c *Car) Honk() {
    fmt.Print("Teeeet... Teet")
}

type Audi struct {
    EngineType string
    Car
}

// passing an interface
func honking(v Vehicle) { 
    fmt.Print(v)
    v.Honk() 

}


func main() {
    var car = Car{Color:"Green"}
    var audi = Audi{}
    audi.Car = car
    audi.Brand = "Audi"
    audi.EngineType = "EC"

    honking(&audi)
}