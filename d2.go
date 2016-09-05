package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "sort"
  "strconv"
)

func main() {
  file, err := os.Open("d2.txt")
  if err != nil {
    fmt.Println("Error opening d2.txt: ", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  total := 0.

  for scanner.Scan() {
    dimensions := strings.Split(scanner.Text(), "x")
    fmt.Println("Input line: ", dimensions)
    var dims []float64
    dims = make([]float64, 3, 3)
    for k, v := range dimensions {
      dims[k], _ = strconv.ParseFloat(v, 64)
    }
    sort.Float64s(dims)
    perimeter := 2*dims[0] + 2*dims[1]
    volume := dims[0]*dims[1]*dims[2]
    //l := strconv.ParseFloat(dimensions[0], 64) //ParseFloat because math.Min accepts float params
    //w := strconv.ParseFloat(dimensions[1], 64)
    //h := strconv.ParseFloat(dimensions[2], 64)
    //min1 := math.Min(l, w)
    total += volume + perimeter
  }
  fmt.Printf("Total ribbon: %f", total)
}
