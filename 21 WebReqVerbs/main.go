package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main() {
	fmt.Println("Welcome to web verb video - LCO")
	//PerformGetRequest()

	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformPostRequest(){
	const myURL = "http://localhost:3000/post"
	// fake json payload

	requestBody := strings.NewReader(`
	{
	"coursename":"Let's 
	}
	`)
}
