package main

func main() {
  assert(1, 1)
  assert(3, sum(1, 2))
  assert(4, sum(1, 3))
  assert(4.5, sum(1.5, 3))
  assert(4.5, sum(1,1))
}
/*
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
*/