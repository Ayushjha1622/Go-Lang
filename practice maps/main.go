package main

import "fmt"

func main(){
	fmt.Println("this is maps practice")

		MyMaps := make(map[int]int)
		MyMaps[1] = 151
		MyMaps[2] = 152
		MyMaps[3] = 153
		MyMaps[4] = 154
		MyMaps[5] = 155
		MyMaps[6] = 156
		MyMaps[7] = 157

		fmt.Println(MyMaps)

		delete(MyMaps, 3)
		fmt.Println(MyMaps)

}