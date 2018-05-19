package main

import (
	"fmt"
)

//definiendo una interface
type User interface{
	Permisos() int
	Nombre() string
}

//definiendo una estructura para admin
type Admin struct{
	nombre string 
}

//definiendo metodos para al estructura admin

func (this Admin) Permisos() int{
	return 5
}

func (this Admin) Nombre() string{
	return this.nombre
}

//definiendo una estructura para editor
type Editor struct{
	nombre string 
}

//definiendo metodos para al estructura editor

func (this Editor) Permisos() int{
	return 3
}

func (this Editor) Nombre() string{
	return this.nombre
}

//creando una funcion para autenticar usuarios
func auth(user User) string{
	nivel := user.Permisos()
	if nivel > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}else if nivel == 3{
		return user.Nombre() + " tiene permisos de editor"
	}else{
		return ""
	}
}

func main(){
	usuarios := []User{Admin{"Edson"},Editor{"Fulano"}}
	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}
}
