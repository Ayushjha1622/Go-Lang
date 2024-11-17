package main

import "fmt"


func main(){
	day := "Saturday"

	switch day {
	case "Monday","Tuesday","Wednesday","Thursday","Friday":
		fmt.Println("Weekday")
	case "Saturday","Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
		
	}
}