package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("list of all languages:",languages)
	fmt.Println("list of all languages: ",languages["JS"])


	delete(languages, "RB")
	fmt.Println("list of all languages:",languages)


	// loops 

	for key, value := range languages{
		fmt.Println("for key %v, value is %v\n,",key,value)
	}

}