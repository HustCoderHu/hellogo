package methods

import (
	"fmt"
	"image"
	"io"
	"math"
	"os"
	"strings"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		fmt.Printf("v: %v\n", v)
		return 0.0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func F1_methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	p := Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p.Abs())
}

type Abser interface {
	Abs() float64
}

func F9_interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	var t *Vertex
	a = t
	fmt.Printf("a.abs: %v", a.Abs())
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func F15_type_assertions() {
	var i interface{} = "hello"

	s := i.(string)
	//fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func F21_readers() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	for i, v := range b {
		if ('a' <= v && v <= 'm') || ('A' <= v && v <= 'M') {
			b[i] = v + 13
		} else if ('n' <= v && v <= 'z') || ('N' <= v && v <= 'Z') {
			b[i] = v - 13
		}
	}
	return
}

func F23_rot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func F24_Images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
