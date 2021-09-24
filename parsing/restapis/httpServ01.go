package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",HelloHandler)
	fmt.Println("Server started at port 8089")
	log.Fatal(http.ListenAndServe(":8089",nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello, there")
}
