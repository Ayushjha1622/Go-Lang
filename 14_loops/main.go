package main

import "fmt"

func main(){
	fmt.Println("welcome to loops in golang")
	days := []string{"sunday","tuesday","wednesday","friday","saturday"}
	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }


	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Printf("index is and value is %v\n,", day)

	// }

	rougeValue := 1



	for rougeValue < 10 {

		if rougeValue ==2 {
			goto youtube
			
		}

		if rougeValue ==5 {
			break
		}
		fmt.Println("value is: ", rougeValue)
		rougeValue++;
		
	}
	youtube:
		fmt.Println("jumping at youtube")


	
}