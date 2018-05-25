package main

import(
	"fmt"
	"time"
	"strings"
)

func main(){
	//mi_nombre_lentooo("Edson")
	go mi_nombre_lentooo("Edson")
	fmt.Println("Aburrido")

	//para obligar al programa a terminar
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lentooo(name string){
	letras := strings.Split(name,"")
	for _,letra := range(letras){
		time.Sleep(1000* time.Millisecond)
		fmt.Println(letra)
	}
}