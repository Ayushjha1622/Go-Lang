package main

import (
	"os"
	"fmt"
)

func main() {
	// CREATING A FILE
					file, err := os.Create("file.txt")
					if err != nil{
						fmt.Println("error: ",err)
						return
					}

					defer file.Close()
					fmt.Println("file created successfully")

	// OPEN A FILE

					// file, err := os.Open("file.txt")
					// if err != nil{
					// 	fmt.Println("error: ",err)
					// 	return
					// }

					// defer file.Close()
					// fmt.Println("file opened successfully")

	// WRITING TO A FILE
					// file, err := os.Create("file.txt")
					// if err != nil{
					// 	fmt.Println("error: ",err)
					// 	return
					// }

					// defer file.Close()
					

					// _, err = file.WriteString("hello, Go")
					// if err != nil {
					// 	fmt.Println("error:",err)
					// 	return
					// }

					// fmt.Println("data written successfully")


	// APPENDING NEW CONTENT
					// file, err := os.OpenFile("file.txt",os.O_APPEND|os.O_WRONLY,0644)
					// if err != nil {
					// 	fmt.Println("error:",err)
					// 	return
					// }		
					// defer file.Close()
				
					// _, err = file.WriteString("appending new content")
					// if err != nil {
					// 	fmt.Println("error:",err)
					// 	return
					// }
					// fmt.Println("content appended successfully")


	// DELETING A FILE

					// err := os.Remove("file.txt")
					// if err != nil {
					// 	fmt.Println("error:", err)
					// 	return
					// 	}
					// 	fmt.Println("file deleted successfully")

	


}