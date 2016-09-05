package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  data, err := ioutil.ReadFile("d1.txt")

  if err != nil {
    fmt.Println("Error while reading from file: ", err)
  }

  floor := 0
  positionFounded := false
  position := -1

  for k, v := range data {
    el := string(v)
    if el == "(" {
      floor += 1
    } else if el == ")" {
      floor -= 1
    }

    if floor == -1 && !positionFounded {
      position = k + 1
      positionFounded = true
    }
  }

  fmt.Println("Resulting floor: ", floor)
  fmt.Println("Position 1st time meet basement: ", position)
}
