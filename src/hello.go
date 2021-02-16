package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

	rand.Seed(time.Now().UnixNano())
	var guess int
	var answer int = rand.Intn(10) + 1

	fmt.Print("Your guess?")

	// Scanで入力を受け取る
	fmt.Scanf("%v", &guess)

	if answer == guess {
		fmt.Println("Bingo!!")
	} else if answer > guess {
		fmt.Println("answer is more big")
	} else {
		fmt.Println("answer is more small")
	}

	fmt.Printf("Your guess is %v\n", guess)

	// var foo string = "foo"

	// %vで変数の埋め込み
	// fmt.Printf("%v", foo)

	fmt.Println("hoge")
	fmt.Println("foo")
	fmt.Println("a", 100, true)

	var price int
	fmt.Println("Price >")
	fmt.Scan(&price)
	fmt.Printf("%dYen\n", price)
}
