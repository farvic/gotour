package main

import "fmt"

type Vertex0 struct {
	Lat, Long float64
}

var m0 = map[string]Vertex0{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main4() {
	fmt.Println(m0)
}
