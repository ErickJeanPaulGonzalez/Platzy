package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gorilla/mux")

func postHandler(w http.ResponseWriter, r *http.Request ){
	fmt.Println("received POST resquest")

	defer r.Body.Close()

	
	body, err := io.ReadAll(r.Body)
	if err != nil{
		fmt.Println("error reading the request")
		
		return
	}
	fmt.Printf(string(body))
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")
	fmt.Println("server listen on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil{
		fmt.Println(err.Error())
	}
}