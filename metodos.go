package main

import "fmt"

type User struct{
	edad int
	nombre,apellido string
}

//metodos

func (this User) nombre_completo() string{
	return this.nombre + " "+ this.apellido
}

func (this *User) set_name(n string){
	this.nombre = n
}

func main(){
	usuario := new(User)
	usuario.nombre ="Edson"
	usuario.apellido ="Ventura"
	usuario.edad =26

	usuario.set_name("Luis")

	fmt.Println(usuario.nombre)
	fmt.Println(usuario.apellido)
	fmt.Println(usuario.edad)

	fmt.Println(usuario.nombre_completo())
}