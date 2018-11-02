package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Microsecond * 1)
	}
}

func main() {
	go count()
	time.Sleep(time.Microsecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Microsecond * 5)
}
