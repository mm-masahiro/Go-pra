package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

//ここがエントリーポイント
func main() {
	// fmt.Printf("Good Morning")
	var num int
	num = 1
	fmt.Println(num)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var hoge int
	hoge = 1
	if hoge == 0 {
		 fmt.Println("OK")
	}
	// trueじゃなかったらここが走る
	fmt.Println("No")

	fmt.Println(pow(3, 2, 10))
}
