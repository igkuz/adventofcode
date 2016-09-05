package main

import (
  "fmt"
  "crypto/md5"
  "strconv"
)

const secretKey = "ckczppom"

func main() {
  found := false
  i := 1
  for !found {
    sum := md5.Sum([]byte(secretKey + strconv.Itoa(i)))
    strSum := fmt.Sprintf("%x", sum)
    fmt.Println("Key ", secretKey + strconv.Itoa(i), " Sum: ", strSum)
    fmt.Printf("1st 6 symbols: %s\n", strSum[:6])
    if strSum[:6] == "000000" {
      found = true
      fmt.Println("Final I: ", i)
    } else {
      i += 1
    }
  }
}
