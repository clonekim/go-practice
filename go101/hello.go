package main

import (
	"fmt"
	"math"
	"strconv"
)

type Circle struct {
	x float32
	y float32
	z float32
	a float64
}

func getArea(c *Circle) {
	area := float64(math.Pi * c.x * c.y * c.z)
	fmt.Println("Area is " + strconv.FormatFloat(area, 'f', -1, 32))
	c.a = area
}

func main() {

	c := Circle{1, 4, 5, 0}
	fmt.Println(c)

	getArea(&c)
	fmt.Println(c)
}
