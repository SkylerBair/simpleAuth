package main

import (
	"fmt"
	"log"
	"net/http"
)

var users map[string]string = map[string]string{
	"Skyler": "active", "Dylan": "inactive",
}

func main() {
	http.HandleFunc("/users/", getUsers)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", users)))
}
