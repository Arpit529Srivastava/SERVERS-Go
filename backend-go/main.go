package main

import (
	"fmt"
	"log"
	"net/http"
)
func formhandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
	}
	fmt.Fprintf(w,"POST request successful")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w,"name is= %s\n", name)
	fmt.Fprintf(w,"address= %s\n", address)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "love you 3000 khushi!!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./frontend"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Println("Server Started:")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
