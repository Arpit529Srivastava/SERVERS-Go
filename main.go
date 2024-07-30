package main

import (
	"fmt"
	"log"
	"net/http"
)
func formhandler(w http.ResponseWriter, r *http. Request){
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w,"post request successful \n")
	UserNamee:= r.FormValue("UserName")
	Password:= r.FormValue("Password")
	fmt.Fprintf(w, "Username =%s\n", UserNamee)
	fmt.Fprintf(w, "Password =%s\n", Password)


}
func hellohandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"working with wrong method", http.StatusNotFound)
	}
	fmt.Fprintf(w,"WELCOME TO MY SERVER!!")
}
func main(){
	fileServer:= http.FileServer(http.Dir("./frontend"))
	//handling the root route
	http.Handle("/",fileServer)
	http.HandleFunc("/signin", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("server start connecting at port 8080")
	if err:= http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal(err)
	}


}