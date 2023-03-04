package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hey!")
	})
	fmt.Println("Server is running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}