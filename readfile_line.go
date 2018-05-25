package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	archivo , err := os.Open("./hola.txt")

	if err != nil{
		fmt.Println("Error")
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan(){
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}

	archivo.Close()
}