package main

import (
	"encoding/json"
	"net/http"
)

func returnError(w http.ResponseWriter, errMsg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	resp := msgResp{
		Message: errMsg,
	}
	err := enc.Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
