package main

import (
	"fmt"
	"strconv"
)

func main(){
	edad := "treinta"
	//convirtiendo a cadena
	edad2 := 22
	edad_str := strconv.Itoa(edad2)

	//convirtiendo a entero
	edad3 := "22"
	//la conversion retorna 2 valores
	edad_int,_ := strconv.Atoi(edad3)

	fmt.Println("Mi edad es " +edad)
	fmt.Println("Mi edad es "+edad_str)
	fmt.Println(edad2 + edad_int)
}