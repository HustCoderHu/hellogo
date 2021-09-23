package flowcontrol

import (
  "fmt"
  "math"
  "runtime"
  "time"
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

func F9_switch() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
  }
}

func F10_switch_evaluation_order() {
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  fmt.Println("today:", today)
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }
}

func F11_switch_with_no_condition() {
  t := time.Now()
  fmt.Println("today:", t)
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}

// https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
// defer 关键字
// 返回值 = xxx
// 调用defer函数
// 空的return
func test_defer3() (r int) {
  defer func(r int) {
    r = r + 5
  }(r)
  return 1
}

func F12_defer() {
  fmt.Println(test_defer3())
}
