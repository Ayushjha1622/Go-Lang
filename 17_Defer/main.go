package main

import "fmt"

func main(){
	defer fmt.Println("world")
	fmt.Println("hello")
	MyDefer()
}

func MyDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}