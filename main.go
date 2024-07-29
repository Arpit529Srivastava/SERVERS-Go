package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Started:")
	if err:= http.ListenAndServe(":8080", err); err!=nil{
		log.Fatal(err)
	}
}