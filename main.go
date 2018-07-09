package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func serveError(w http.ResponseWriter, r *http.Request) {
	fmt.Println("There's no way I'll work!")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["id"]
	fileName := "files/" + pageId + ".html"
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(Port, nil)
}
