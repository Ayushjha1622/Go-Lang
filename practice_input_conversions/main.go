package main

import "fmt"

func main(){

	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)


	fmt.Print("enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n",name,age)
}