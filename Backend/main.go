package main

import(
	"net/http"
	"encoding/json"
)

var db=make(map[int] Task)
var nextID=1

type Task struct{
	ID int `json:"id"`
	Task string `json:"task"`
	Completed bool `json:"completed"`
}

type TaskList struct{
	Message string `json:"message"`
	Tasks map[int]Task `json:"tasks"`
}

type errResp struct{
	Message string `json:"message"`
}

func returnError(w http.ResponseWriter, errMsg string, code int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc:=json.NewEncoder(w)
	enc.SetIndent("", "    ")
	resp:=errResp{
		Message: errMsg,
	}
	err:=enc.Encode(resp)
	if (err!=nil){
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request){
	if (r.Method==http.MethodGet){
		w.Header().Set("Content-Type", "application/json")
		resp:=TaskList{
			Message:"Request Successful",
			Tasks: db,
		}
		encoder:=json.NewEncoder(w)
		encoder.SetIndent("", "    ")
		err:=encoder.Encode(resp)
		if (err!=nil){
			returnError(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
}

func main(){
	http.HandleFunc("/todos", todoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}