package ch11

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func F7_sort() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data)
	Sort(a)
	// fmt.Println(a)
	logrus.Info(a)
}

type Person struct {
	name string
	age  int
}
type Any interface{}

func F9_case_type() {
	var val interface{}
	val = 5
	logrus.Infof("val is: %v\n", val)

	per := new(Person) // {"Rob", 55}
	per.name = "Rob"
	per.age = 55
	val = per

	switch t := val.(type) {
	case int:
		logrus.Infof("Type int %T\n", t)
	case string:
		logrus.Infof("Type string %T\n", t)
	case bool:
		logrus.Infof("Type bool %T\n", t)
	case *Person:
		logrus.Infof("Type *Person %T\n", t)
	default:
		logrus.Infof("Unexpected type %T\n", t)
	}
}

type Obj interface{}

func F9_interface_array() {
	a := make([]interface{}, 0, 4)
	logrus.Infof("type a: %T\n", a)
	a = append(a, 1)
	logrus.Info(a)

	b := [...]interface{}{0, 1, 2, 3}
	fmt.Println("----------")
	logrus.Infof("type b: %T\n", b)
	func(x []interface{}) {
		x[0] = 10
	}(b[:])
	logrus.Info(b)
}

func F12_overload() {
	mysum := func(a ...interface{}) int {
		logrus.Info("len: ", len(a))
		logrus.Info(a)
		all := 0
		for _, v := range a {
			all += v.(int) // reflect.ValueOf(v).Interface().(int)
		}
		// for _, v := range a {
		// 	all += v
		// }
		return all
	}
	logrus.Info(mysum(1, 2, 3))
}
