package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Parseform() err: Iv", err)
		return
	}
	fmt.Fprintf(w, "Post request succesfull ")
	name := r.FormValue("name")
	address := r.FormValue("adress")
	fmt.Fprintf(w, "name = ", name)
	fmt.Fprintf(w, "address = ", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello"{
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w, "Method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello!")
 
}

func main(){
	fileServer := http.FileServer(http.Dir("./static")) 
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err !=nil {
		log.Fatal(err)
	}
}