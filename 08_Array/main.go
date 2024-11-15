package main

import(
	"fmt"
)
func main(){
	fmt.Println("welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "cherry"
	fruitList[3] = "date"

	fmt.Println("fruit List is: ", fruitList)
	fmt.Println("fruit List is: ", len(fruitList))


	var vegList = [4]string{"potato","beans","tomato","gobhi"}
	fmt.Println(vegList)
}