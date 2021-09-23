package flowcontrol

import (
  "fmt"
  "math"
)

func F0() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}

func F1() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

func Sqrt(x float64) string {
  if x < 0 {
    return Sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func Pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func F8_customize_sqrt(x float64) float64 {
  z := 1.0
  for i := 0; i < 10; i++ {
    fmt.Printf("i:%v z:%v\n", i, z)
    z -= (z*z - x) / (2 * z)
  }
  return z
}

func My_customize_extraction(x float64) float64 {
  // https://blog.csdn.net/taoqick/article/details/104034602
  // 泰勒展开
  now := 1.0
  for i := 0; i < 10; i++ {
    pre := now
    now = 2*pre/3 + x/(3*pre*pre)
    fmt.Printf("i:%v z:%v\n", i, now)
  }
  return now
}
