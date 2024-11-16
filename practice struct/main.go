package main

import "fmt"

type Person struct{
	name string
	age int
	city string
}


func main(){
	person1 := Person{
		name: "ayush",
		age: 21,
		city: "kanpur",
	}

	fmt.Println(person1)
	fmt.Println("Name:",person1.name)
	fmt.Println("age:",person1.age)
	fmt.Println("city:",person1.city)



	person1.age = 31
	fmt.Println("age after updating:",person1.age)

}