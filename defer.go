package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	readFile()
}

func readFile() bool{
	archivo , err := os.Open("./hola.txt")
	//defer siempre se va a ejecutar al finalizar la funcion actual
	defer func(){
		archivo.Close()
		fmt.Println("Ejecucion del defer")
	}()

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
	return true
}