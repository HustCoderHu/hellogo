package ch11

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func F1_interface() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func F2_poly() {
	r := Rectangle{5, 3}
	q := &Square{5}

	fmt.Println("Looping through shapes for area ...")
	shapes := []Shaper{r, q}

	for _, shape := range shapes {
		fmt.Println("shape details:", shape)
		fmt.Println("Area of this shape is:", shape.Area())
	}
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

func F3_valuable() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func F4_type_interface() {
	var areaIntf Shaper
	sq1 := Square{5}
	areaIntf = &sq1

	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("the type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("the type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
