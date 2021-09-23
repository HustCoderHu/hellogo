package moretypes

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

func F6_array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func F8_slice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func F11_slice_len_cap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func F13_make_slice() {
	a := make([]int, 5)
	printSlice1("a", a)

	b := make([]int, 0, 5)
	printSlice1("b", b)

	c := b[:2]
	printSlice1("c", c)

	d := c[2:5]
	printSlice1("d", d)
}

func F14_complex_slice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func F15_slice_append() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func F16_range() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Printf("pow type:%T\n", pow)
	printSlice(pow)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

type Vertex struct {
	Lat, Long float64
}

func F19_map() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

func F22_mutating_maps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	sArray := strings.Fields(s)
	countMap := make(map[string]int)
	for _, str := range sArray {
		countMap[str] = countMap[str] + 1
	}
	return countMap
}

func F23_map_exercise() {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func F24_wrapper() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
