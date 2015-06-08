package main

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello DevOps! This is your welcome task!")
}

func main() {
	http.HandleFunc("/", welcomeHandler)

	fmt.Println("Started, serving on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
