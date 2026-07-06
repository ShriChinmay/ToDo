package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

var db=make(map[int] Task)
var nextId=1

type Task struct{
	Taskname string `json:"task"`
	Completed bool `json:"completed"`
}

type TaskList struct{
	Message string `json:"message"`
	Tasks map[int]Task `json:"tasks"`
}

type msgResp struct{
	Message string `json:"message"`
}

func returnError(w http.ResponseWriter, errMsg string, code int){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc:=json.NewEncoder(w)
	enc.SetIndent("", "    ")
	resp:=msgResp{
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
	if (r.Method==http.MethodPost){
		newTask:=Task{}
		err:=json.NewDecoder(r.Body).Decode(&newTask)
		if (err!=nil){
			returnError(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if (len(newTask.Taskname)==0){
			returnError(w, "Name of task is required", http.StatusBadRequest)
			return
		}
		db[nextId]=newTask
		nextId++
		resp:=msgResp{
			Message: "Task added successfully",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		enc:=json.NewEncoder(w)
		enc.SetIndent("", "    ")
		err= enc.Encode(resp)
		if (err!=nil){
			returnError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	if (r.Method==http.MethodPut || r.Method==http.MethodPatch){
		body:=Task{}
		err:=json.NewDecoder(r.Body).Decode(&body)
		if (err!=nil){
			returnError(w, "Bad Request", http.StatusBadRequest)
			return
		}
		idstr:=r.URL.Query().Get("id")
		id, err:=strconv.Atoi(idstr)
		if (err!=nil){
			returnError(w, "A valid ID is required", http.StatusBadRequest)
			return
		}
		_, found:=db[id]
		if (!found){
			returnError(w, "Task not found", http.StatusNotFound)
			return
		}
		modified:=db[id]
		if (len(body.Taskname)>0){
			modified.Taskname=body.Taskname
		}
		modified.Completed=body.Completed
		db[id]=modified
		resp:=msgResp{
			Message: "Modified successfully",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc:=json.NewEncoder(w)
		enc.SetIndent("", "    ")
		err= enc.Encode(resp)
		if (err!=nil){
			returnError(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	if (r.Method==http.MethodDelete){
		idstr:=r.URL.Query().Get("id")
		id, err:=strconv.Atoi(idstr)
		if (err!=nil){
			returnError(w, "A valid ID is required", http.StatusBadRequest)
			return
		}
		_, found:=db[id]
		if (!found){
			returnError(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp:=msgResp{
			Message: "Deleted successfully",
		}
		delete(db, id)
		enc:=json.NewEncoder(w)
		enc.SetIndent("", "    ")
		err= enc.Encode(resp)
		if (err!=nil){
			returnError(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	returnError(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main(){
	R:=chi.NewRouter()
	R.HandleFunc("/todos", todoHandler)
	err := http.ListenAndServe(":8080", R)
	if err != nil {
		panic(err)
	}
}