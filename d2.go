package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"
)

func main() {
  // высчитывать необходимые величины
  // складывать в накопитель

  file, err := os.Open("d2.txt")
  if err != nil {
    fmt.Println("Error opening d2.txt: ", err)
  }
  defer file.Close()

  totalWrap := 0.0
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    fmt.Println("Input string: ", scanner.Text())

    dimensions := strings.Split(scanner.Text(), "x")
    fmt.Println("Dimensions: ", dimensions)

    l, _ := strconv.ParseFloat(dimensions[0], 64)
    w, _ := strconv.ParseFloat(dimensions[1], 64)
    h, _ := strconv.ParseFloat(dimensions[2], 64)

    v1 := l*w
    v2 := w*h
    v3 := h*l
    minv := math.Min(v1, math.Min(v2, v3))

    fmt.Println("Box surface: ", 2*(v1+v2+v3), "Small wrap: ", minv)

    totalWrap += (2 * (v1 + v2 + v3) + minv)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading d2.txt: ", err)
  }

  fmt.Printf("Total wrap: %.f", totalWrap)
}
