package main

import (
	"time"
	"fmt"

)


func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	newsDate := time.Date(2020,time.December,12,34,23,0,0,time.UTC)
	fmt.Println(newsDate)

	
}