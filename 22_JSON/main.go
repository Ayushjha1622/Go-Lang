package main

import "fmt"

type course struct{
	Name 		string
	Price 		int
	Platform 	string
	Password 	string
	Tags 		[]string
}

func main(){
	fmt.Println("welcome to JSON video")
}


func EncodeJson(){

	lcoCouses := []courses{
		{"GO", 100, "Udemy", 123, "coursera", 499,"abc123",[]string{"web-dev","js"}},
	}
}