package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("welcome to files in golang")
	content := "this needs to be in file - learncode.online"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil{
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("length is: ",length)
	defer file.Close()
	readFile("./mylcofile.txt")
}


func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data inside the file is \n",string(databyte))
}


func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}
