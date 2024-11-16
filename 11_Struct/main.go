package main

import "fmt"

func main(){
	fmt.Println("Structs in GoLang")
	// no inheritance in golang; No super or parent

	Ayush := User{"Ayush","ayush@gmail.com",true,21}


	fmt.Println(Ayush)

	fmt.Printf("hitesh details are: %+v\n",Ayush)

	fmt.Printf("Name is %v and Email is %v.", Ayush.Name, Ayush.Email)



}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

