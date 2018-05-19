package main

import "fmt"

func main() {
	a := 8
	b := 8

	if a > b {
		fmt.Printf("%d es mayor que %d\n",a,b)
	}else if b > a {
		fmt.Printf("%d es mayor que %d\n",b,a)
	}else{
		fmt.Println("Son iguales")
	}
}