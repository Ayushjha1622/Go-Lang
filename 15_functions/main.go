package main

import "fmt"

// func greeter()  {
// 	fmt.Println("namaste")

	
// }

// func greeterTwo()  {
// 	fmt.Println("another method")
// }



func main(){
	fmt.Println("welcome to functions")
	// greeter()
	// greeterTwo()

	// result := adder(3,5)
	// fmt.Println("result is: ", result)

	proRes := proAdder(2,5,8,7)
	fmt.Println("pro result is: ", proRes)
}

// func adder(valOne int, valTwo int) int{
// 	return valOne + valTwo
// }


func proAdder(values ...int) int  {
	total := 0

	for _, val := range values{
		total += val
	}

	return total
	
}