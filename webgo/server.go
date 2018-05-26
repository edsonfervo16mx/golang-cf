package main

import(
	"net/http"
)
func main(){

	//localhost:8000
	/*
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"index.html")
	})
	/**/

	//localhost:8000/teste.html
	/*
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,r.URL.Path[1:])
	})
	/**/

	//localhost:8000/public/hola.html
	/**/
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/",http.StripPrefix("/public/",fileServer))
	/**/
	http.ListenAndServe(":8000",nil)
}