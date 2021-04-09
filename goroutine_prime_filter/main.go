package main

import "fmt"

func Generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int { //这里的in是只读管道
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	q := Generate()
	for i := 1; i <= 100; i++ {
		prime := <-q
		fmt.Printf("%d\n", prime)
		q = PrimeFilter(q, prime)
	}
}
