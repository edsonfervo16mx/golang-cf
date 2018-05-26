package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool{
	//cambio del nombre del archivo intencionalmente para crear un error / panic
	archivo , err := os.Open("./holaaaaaaaaaaaaaaa.txt")
	//defer siempre se va a ejecutar al finalizar la funcion actual
	defer func(){
		archivo.Close()
		fmt.Println("Ejecucion del defer")
		//evitar el panic con esta funcion recover
		recover()
	}()

	if err != nil{
		//fmt.Println("Error")
		//ejecucion del panic
		panic(err)
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