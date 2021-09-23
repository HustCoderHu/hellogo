package main

import "fmt"

func write() (name string, age int) {
  return "Alice", 10
}

func read(name string, age int) {
  fmt.Println(name, "already", age, "years old")
}

func test_slice_array() {
  var str string
  str = "Alice welcome everyone"
  fmt.Println(str)

  arr0 := [4]int{1, 2, 3: 10}
  fmt.Println(arr0)

  slice0 := make([]int, 5)
  for i := 0; i < 5; i++ {
    slice0 = append(slice0, i)
  }
  slice0 = append(slice0, 6)
  fmt.Println(slice0)
}

func test_map() {
  // var strmap map[int]string
  strmap := make(map[int]string)
  // strmap := map[int]string
  for i := 0; i < 5; i++ {
    strmap[i] = "a"
  }
  fmt.Println(strmap)
}

type user struct {
  name string
  age  int
}

func (u *user) read() string {
  return fmt.Sprintf("%s is already %d years old", u.name, u.age)
}

func test_struct_method() {
  u := &user{
    name: "Bob",
    age:  1,
  }
  fmt.Println(u.name, "is already", u.age, "years old")
  fmt.Println(u.read())
}
