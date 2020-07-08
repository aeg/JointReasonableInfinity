package main

import "fmt"

func assert(a float64, b float64) {
  if a == b {
    fmt.Print("Success!!: ")
  } else {
    fmt.Print("Fail!!: ")
  }
  fmt.Printf("%v == %v",a , b)
  fmt.Println()
}

func assertStr(a string, b string) {
  var result bool = a == b
  fmt.Print(a + " == " + b + " : ")
  fmt.Println(result)
}