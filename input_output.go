package main

import (
	"fmt"
)

func main(){
	edad := 22
	nombre := "Edson"
	estado := true
	estatura := 1.72
	descripcion := "Para las cadenas"
	//solicitando valores al terminal
	var tu_edad int
	
	fmt.Printf("Mi edad es %d\n", edad)
	fmt.Printf("Mi nombre es %v\n", nombre)
	fmt.Printf("Mi estado es %t\n", estado)
	fmt.Printf("Mi estatura es %f\n", estatura)
	fmt.Printf("Mi descripcion es %s\n", descripcion)

	//solicitando valores al terminal
	fmt.Println("Coloca tu edad")
	fmt.Scanf("%d\n",&tu_edad)
	fmt.Printf("Tu edad es %d\n", tu_edad)
}