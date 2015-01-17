package main

import (
	"fmt"
	"net/http"
)

func serveHelloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

func main(){
	router := initRouter()
	http.Handle("/", router)
	http.ListenAndServe(":1123", nil)
}
