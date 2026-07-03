package main

import(
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request){

}

func main(){
	http.HandleFunc("/tasks", taskHandler)
}