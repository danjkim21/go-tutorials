package main

import (
  "fmt"
  "errors"
)

func main() {
  fmt.Println("Hello world!")

  var intNum int = 10
  name := "Danny boy"
  var1, var2, var3 := 1, 2, 3
  const pi = 3.14
  const printText string = "Goodbye world!"

  fmt.Println(intNum, name, var1, var2, var3, pi)

  var divisionResult, remainder, err = intDivision(10, 3)

  if err != nil {
    fmt.Println(err.Error())
  } else if remainder == 0 {
    fmt.Printf("division: %v", divisionResult)
  } else {
    fmt.Printf("division: %v remainder: %v", divisionResult, remainder)
  }

  printMe(printText)
}

func printMe( printValue string) {
  fmt.Println(printValue)
}

func intDivision (numerator int, denominator int) (int, int, error) {
  var err error

  if denominator == 0 {
    err = errors.New("cannot divide by zero")

    return 0, 0, err
  }

  fmt.Printf("dividing %v by %v", numerator, denominator)

  var division int = numerator / denominator
  var remainder int = numerator % denominator

  return division, remainder, err
}
