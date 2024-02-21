package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Perimeter () float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) Area () float64 {
	return r.width * r.height
}

func main () {
	rec1 := Rectangle{ width: 5, height: 10 }
	fmt.Println(rec1.Perimeter())
	fmt.Println(rec1.Area())
}
