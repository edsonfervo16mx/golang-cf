package main

import(
	"net/http"
	"encoding/json"
)

type Curso struct{
	Title string `json:"titulo"`
	NumeroVideos int `json:"numero_videos"`
}

type Cursos []Curso

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		cursos := Cursos{
			{"Curso de Go", 30},
			{"Curso de Ruby", 20},
			{"Curso de NodeJS", 50},
			{"Curso de PHP", 40},
			{"Curso de Java", 50},
		}
		json.NewEncoder(w).Encode(cursos)
	})

	http.ListenAndServe(":8000",nil)
}