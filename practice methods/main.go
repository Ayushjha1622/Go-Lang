package main

import (
	"fmt"
)

type employee struct {
	Name     string
	Age      int
	Position string
}

// Method to get the position of the employee
func (e employee) GetPosition() {
	fmt.Println("Position is:", e.Position)
}

// Method to get the information of the employee
func (e employee) GetInfo() {
	fmt.Println("Name is:", e.Name)
}

func main() {
	// Correctly initializing the employee struct
	Ayush := employee{Name: "Ayush", Age: 21, Position: "Head"}
	
	// Calling the methods using parentheses
	Ayush.GetInfo()
	Ayush.GetPosition()
}
