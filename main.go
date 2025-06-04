package main

import "fmt"

func main() {
	d := NewDequeue(3, 3)

	go squares(c)

type OptionalInt struct {
	Value int
	IsSet bool
}

	fmt.Println("main stop")
}

func squares(c chan int) {
	for i := 0; i <= 2; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}
