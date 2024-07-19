package main

import (
	"fmt"
	"hello/balancingcharge"
	"hello/balancingcharge/egat"
)

func main() {

	egat := &egat.EgatFactory{
		Load: 1,
		Plan: 2,
		Rate: 3,
		Type: "consumer",
	}

	var egatfactory balancingcharge.Factory = egat

	fmt.Println("BalancingCharge:", egatfactory.GetBalancingCharge())

	// square := shapes.Square{
	// 	Side: 5,
	// }

	// triangle := geometry.Triangle{Base: 3, Height: 4}
	// circle := circle.Circle{Foo: 2, Bar: 2}

	// Polymorphism - assign implementing structs to interface type
	// var shape1 shapes.Shape = &square
	// var shape2 shapes.Shaper = &triangle
	// var shape3 shapes.Shaper = &circle

	// fmt.Println("Square Area:", shape1.Area())
	// fmt.Println("Triangle Area:", shape2.Area())
	// fmt.Println("Circle Area:", shape3.Area())

	// baz := baz.Baz{}

	// var foo1 foo.Bar = &baz

	// fmt.Println("foo:", foo1.Area())

}
