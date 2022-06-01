package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m1 = map[string]Vertex1{
	"Bell Labs": Vertex1{
		40.68433, -74.39967,
	},
	"Google": Vertex1{
		37.42202, -122.08408,
	},
}

func main5() {
	fmt.Println(m1)
}
