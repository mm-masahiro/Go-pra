package main

import "fmt"

func main() {
	var greet string = "Hello, World"

	const greet_morning string = "Good Mornig"

	const (
		a = 1 + 2
		b
		c
	)

	const (
		d = iota
		e
	)

	const (
		f = 1 << iota
		g
		h
	)

	fmt.Println(greet)
	fmt.Println(greet_morning)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g, h)
}
