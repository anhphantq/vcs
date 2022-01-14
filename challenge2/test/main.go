package main

import "fmt"

type test interface {
	Check() int
}

type hic struct {
	wibu int
}

func (a hic) Check() int {
	return 1
}

func main() {
	var a test = hic{1}
	b, ok := a.(hic)
	fmt.Print(a, b, ok)
}
