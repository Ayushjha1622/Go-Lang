package main


import(
	"fmt"
	
)

func main()  {
	marks := 45
	var Marks = &marks
	fmt.Println(*Marks)
	fmt.Println(Marks)

}