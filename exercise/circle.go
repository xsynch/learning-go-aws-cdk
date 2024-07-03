package main

import "fmt"

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func (c Circle) circumference() {
	fmt.Println(2 * c.Radius * PI )
}

func (c Circle) area(){
	fmt.Println(PI * (c.Radius * c.Radius) )
}

func main(){
	myCircle := Circle{
		Radius: 2.5,
	}

	myCircle.circumference()
	myCircle.area()
}