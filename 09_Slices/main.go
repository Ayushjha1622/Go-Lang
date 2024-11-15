package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome on slices")

	var fruitList = []string{"apple","banana","peach"}
	fmt.Printf("type of fruitList here %T\n",fruitList)

	fruitList = append(fruitList, "mango","melon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int,4)
	highScore[0] =234
	highScore[1] =357
	highScore[2] =278
	highScore[3] =298
	// highScore[4] =777

	highScore = append(highScore, 555, 666, 321)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))


	// how to remove a value from slices based on index

	var courses = []string{"reactjs","javascript","swift","ruby"}
	fmt.Println(courses)
	var index int =2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)


}