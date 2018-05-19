package main

import "fmt"

func main(){
	/*
		el slice es un arreglo extendido que no requiere declarar su tamaño inicial, sin embargo permite establecer su tamaño y capacidad maxima, ya que nos ayuda a que en caso de desbordar los valores del arreglo lo reconstruye y permite extender su capacidad y no mostrar error en el codigo.
	*/
	
	slice := make([]int,3,5)
	slice = append(slice,2)
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}