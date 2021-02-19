package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	const x int = 1

	// if x == 1 {
	// 	println("x=1")
	// } else if x == 2 {
	// 	println("x=2")
	// } else {
	// 	println("X not equal 1 & 2")
	// }

	for i := 0; i < 5; i = i + 1 {
		if i%2 == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}

	t := time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(6)

	switch n + 1 {
	case 6:
		fmt.Println("大吉")
	case 5, 4:
		fmt.Println("中吉")
	case 3, 2:
		fmt.Println("小吉")
	default:
		fmt.Println("凶")
	}

}
