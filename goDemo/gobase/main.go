package main

import "fmt"

func hello(num []int) []int {
	num[0] = 18
	return num
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a B = c
	var b A = c
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())
}
