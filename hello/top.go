package main

import (
  "fmt"
  "world/flowcontrol"
)

func main() {
  flowcontrol.F0()
  flowcontrol.F1()
  fmt.Println(flowcontrol.Sqrt(2), flowcontrol.Sqrt(-4))
  fmt.Println(flowcontrol.Pow(3, 2, 10), flowcontrol.Pow(3, 3, 20))
  // flowcontrol.F8_customize_sqrt(169)
  flowcontrol.My_customize_extraction(27)
  // hehello.test_struct_method()
  // fmt.Println("hello world!")
  // read(write())
  // test_slice_array()
  // test_map()
  // test_struct_method()
}
