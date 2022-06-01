package main

import "fmt"

type Vertex5 struct {
	X int
	Y int
}

func main20() {
	v := Vertex5{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
