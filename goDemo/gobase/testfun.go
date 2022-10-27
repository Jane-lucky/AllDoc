package main

import "fmt"

func day3() {
	slice := make([]int, 5)
	slice = append(slice, 1, 2, 3)

	s := make([]int, 0)
	s = append(s, 1, 2, 3)

	fmt.Println(slice) //[0 0 0 0 0 1 2 3]
	fmt.Println(s)     //[1 2 3]
}

func day2E() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func day2R() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
