package main

import (
	"fmt"
	"sort"
)

func main(){
	// var fruitList = []string{"apple","banana","mango","melon","guava"}
	// fmt.Println(fruitList)

	// fruitList = append(fruitList, "cherry", "lichi")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	marks := make([]int,4)
	marks[0] = 90
	marks[1] = 80
	marks[2] = 70
	marks[3] = 60
	fmt.Println(marks)


	sort.Ints(marks)
	fmt.Println(marks)
	fmt.Println(sort.IntsAreSorted(marks))


	
	var courses = []string{"reactjs","javascript","swift","ruby"}
	fmt.Println(courses)
	var index int =2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)






}