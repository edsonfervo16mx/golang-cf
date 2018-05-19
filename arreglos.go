package main

import "fmt"

func main() {
	//declaracion de arreglos
	var arreglo [10]int
	fmt.Println(arreglo)

	//declaracion de arreglos con valores
	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)

	//saber la longitud del arreglo
	fmt.Println(len(arreglo2))

	//declaracion de arreglos incompletos
	arreglo3 := [5]int{1, 2, 3}
	for i := 0; i < len(arreglo3); i++ {
		fmt.Println(arreglo3[i])
	}

	//declarando arreglos multidimensionales
	var matriz [2][2]int
	fmt.Println(matriz)
}
