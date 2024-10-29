package main

import (
	"fmt"
	"sync"
)

func main() {
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	n := 5

	event := func(i int) int {
		return i
	}

	// Заполняем каналы в отдельной горутине
	go func() {
		defer close(in1)
		defer close(in2)
		for j := 0; j < n; j++ {
			in1 <- j
			in2 <- j
		}
		
	}()

	merge2Channels(event, in1, in2, out, n)

	 //Чтение результатов из канала
	for val := range out {
		fmt.Print(val, " ")
	}
	
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var wg sync.WaitGroup

	val1 := make([]int, n)
	val2 := make([]int, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			x1, ok1 := <-in1
			if !ok1 {
				return
			}
			val1[i] = fn(x1)
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			x2, ok2 := <-in2
			if !ok2 {
				return
			}
			val2[i] = fn(x2)
		}(i)
	}

	
	go func() {
		wg.Wait()
		defer close(out)
		for i := 0; i < n; i++ {
			//fmt.Println(val1[i], "   ", val2[i])
			out <- val1[i] + val2[i]
		}
		fmt.Println(val1, "   ", val2)
	}()
}