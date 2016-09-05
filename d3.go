package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
)

type Coords struct {
  X int
  Y int
}

func (c *Coords) Up() {
  c.Y += 1
}

func (c *Coords) Down() {
  c.Y -= 1
}

func (c *Coords) Right() {
  c.X += 1
}

func (c *Coords) Left() {
  c.X -= 1
}

func (c *Coords) ParseCommand(command string) {
  switch command {
  case "<":
    c.Left()
  case ">":
    c.Right()
  case "^":
    c.Up()
  case "v":
    c.Down()
  default:
    fmt.Println("Not recognized command: ", command)
  }
}

func (c *Coords) Serialize() string {
  return strconv.Itoa(c.X) + strconv.Itoa(c.Y)
}

func main() {
  data, err := ioutil.ReadFile("d3.txt")
  if err != nil {
    fmt.Println("Error reading d3.txt: ", err)
  }

  position := &Coords{0, 0}
  visitedHouses := make(map[string]int)
  visitedHouses[position.Serialize()] = 1

  for _, v := range data {
    //if v != 10 {
      position.ParseCommand(string(v))
      //fmt.Println(visitedHouses[position.Serialize()])
      visitedHouses[position.Serialize()] += 1
    //}
  }
  fmt.Println("Visited houses d3_1: ", visitedHouses)
  fmt.Println("Houses Len d3_1: ", len(visitedHouses))

  deliverWithRobot(data)
}

func deliverWithRobot(data []byte) {
  robotPosition := &Coords{0, 0}
  santaPosition := &Coords{0, 0}
  visitedHouses := make(map[string]int)
  visitedHouses[santaPosition.Serialize()] = 2

  i := 1
  for _, v := range data {
    if (i % 2) != 0 {
      santaPosition.ParseCommand(string(v))
      visitedHouses[santaPosition.Serialize()] += 1
    } else if (i % 2) == 0 {
      robotPosition.ParseCommand(string(v))
      visitedHouses[robotPosition.Serialize()] += 1
    }
    i += 1
  }

  fmt.Println("Visited houses d3_2: ", visitedHouses)
  fmt.Println("Total steps: ", i)
  fmt.Println("Houses Len d3_2: ", len(visitedHouses))
  fmt.Println("Santa Position: ", santaPosition, " Robot Position: ", robotPosition)
}
