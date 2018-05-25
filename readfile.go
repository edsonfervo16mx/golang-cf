package main

import(
	"io/ioutil"
	"fmt"
)

func main(){
	file_data,err := ioutil.ReadFile("./hola.txt")

	if err != nil{
		fmt.Println("Error")
	}

	fmt.Println(string(file_data))
}