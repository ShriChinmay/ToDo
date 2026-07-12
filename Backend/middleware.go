package main

import(
    "net/http"
    "fmt"
    "time"
)

func Logging (next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        fmt.Printf("Request started\n %s, %s,", r.Method, r.URL.Path)
        now:=time.Now()
        next.ServeHTTP(w, r)
        t:=time.Since(now)
        fmt.Printf("Time taken: %v\n", t)
        fmt.Println("Request finished")
        
    })
}