package main

import "fmt"


func main(){
	fmt.Println("If else in golang ")
	LoginCount := 23 
	var result string

	if LoginCount <10 {
		result = "regular user"
	} else if LoginCount > 10{
		result = "watch out"
	} else{
		result = "exactly 10 login count"
	}	
	fmt.Println(result)



	if num := 3; num <10{
		fmt.Println("less than 10")
	} else{
		fmt.Println("num is NOT less than 10")
	}
}