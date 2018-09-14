package main

import "fmt"

func main(){
	var va int
	var vb int

	fmt.Println("Ingresar el valor A")
	fmt.Scanf("%d\n",&va)
	fmt.Println("Ingresar el valor B")
	fmt.Scanf("%d\n",&vb)

	if va == vb {
		fmt.Println("Son iguales")
	}else{
		if va > vb {
			fmt.Println("A es mayor")
		}else{
			fmt.Println("B es mayor")
		}
	}
}