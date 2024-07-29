package main

import (
	"fmt"
	"log"
	"net/http"
)
func hellohandler(w http.ResponseWriter, r *http.Request ){
	if r.URL.Path!="/hello"{
		http.Error(w,"status not found", http.StatusNotFound)
		return
	}
	
}
func main() {
	fmt.Println("Server Started:")
	if err:= http.ListenAndServe(":8080", err); err!=nil{
		log.Fatal(err)
	}
}