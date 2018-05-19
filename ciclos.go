package main

import "fmt"

func main(){
	//alternativa 1
	for i:=1;i<10;i++{
		fmt.Println("Hola Mundo")
	}

	//alternativa 2
	j:=1
	for j<10{
		fmt.Println("Otra vez hola")
		j++
	}

	//manejo del break
	l:=1
	for {
		fmt.Println(l)
		l++
		if l == 10{
			break
		}
	}

}