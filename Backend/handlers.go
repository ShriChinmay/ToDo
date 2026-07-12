package main

import(
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"strconv"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := TaskList{
		Message: "Request Successful",
		Tasks:   db,
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(resp)
	if err != nil {
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	newTask := Task{}
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		returnError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if len(newTask.Taskname) == 0 {
		returnError(w, "Name of task is required", http.StatusBadRequest)
		return
	}
	db[nextId] = newTask
	nextId++
	resp := msgResp{
		Message: "Task added successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err = enc.Encode(resp)
	if err != nil {
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func modifyTodo(w http.ResponseWriter, r *http.Request) {
	body := Task{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		returnError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	idstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		returnError(w, "A valid ID is required", http.StatusBadRequest)
		return
	}
	_, found := db[id]
	if !found {
		returnError(w, "Task not found", http.StatusNotFound)
		return
	}
	modified := db[id]
	if len(body.Taskname) > 0 {
		modified.Taskname = body.Taskname
	}
	modified.Completed = body.Completed
	db[id] = modified
	resp := msgResp{
		Message: "Modified successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err = enc.Encode(resp)
	if err != nil {
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		returnError(w, "A valid ID is required", http.StatusBadRequest)
		return
	}
	_, found := db[id]
	if !found {
		returnError(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := msgResp{
		Message: "Deleted successfully",
	}
	delete(db, id)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err = enc.Encode(resp)
	if err != nil {
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	idstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		returnError(w, "Valid id required", http.StatusBadRequest)
		return
	}
	todo, found := db[id]
	if !found {
		returnError(w, "ID not found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(todo)
	if err != nil {
		returnError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}