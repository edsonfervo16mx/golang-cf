package main

import "fmt"

type User struct{
	edad int
	nombre,apellido string
}

func main(){
	usuario := new(User)
	usuario.nombre ="Edson"
	usuario.apellido ="Ventura"
	usuario.edad =26

	fmt.Println(usuario.nombre)
	fmt.Println(usuario.apellido)
	fmt.Println(usuario.edad)
}