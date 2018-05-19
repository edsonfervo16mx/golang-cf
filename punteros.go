package main

import "fmt"

func main(){
	var x,y *int
	entero :=5
	x = &entero
	y = &entero

	*x = 6

	//direcciones de memoria
	fmt.Println(x)
	fmt.Println(y)

	//valores almacenados en memoria
	fmt.Println(*x)
	fmt.Println(*y)
}