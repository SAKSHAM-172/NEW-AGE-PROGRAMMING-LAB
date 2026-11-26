package main

import (
	"fmt"
	"time"
)

type Result struct {
	Task  string
	Value []int
}

func squares(n int, ch chan<- Result) {
	fmt.Println("[squares] started")
	time.Sleep(300 * time.Millisecond)
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*i)
	}
	fmt.Println("[squares] completed")
	ch <- Result{"Squares", out}
}

func cubes(n int, ch chan<- Result) {
	fmt.Println("[cubes] started")
	time.Sleep(100 * time.Millisecond)
	out := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		out = append(out, i*i*i)
	}
	fmt.Println("[cubes] completed")
	ch <- Result{"Cubes", out}
}

func fibonacci(n int, ch chan<- Result) {
	fmt.Println("[fibonacci] started")
	time.Sleep(200 * time.Millisecond)
	out := make([]int, 0, n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		out = append(out, a)
		a, b = b, a+b
	}
	fmt.Println("[fibonacci] completed")
	ch <- Result{"Fibonacci", out}
}

func main() {
	const n = 10
	ch := make(chan Result)

	fmt.Println("main: launching goroutines")
	go squares(n, ch)
	go cubes(n, ch)
	go fibonacci(n, ch)

	for i := 0; i < 3; i++ {
		r := <-ch
		fmt.Printf("main received %-10s: %v\n", r.Task, r.Value)
	}

	fmt.Println("main: all tasks finished")
}
