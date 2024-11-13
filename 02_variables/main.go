package main

import "fmt"

const LoginToken string = "ghjhjkhjhc"   //public


func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)



	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.555565644455465465
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// IMPLICIT type
	var website = "learncodeonline.in"
	fmt.Println(website)


	// NO VAR STYLE

	numberOfUser := 30000
	fmt.Println(numberOfUser)


	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}