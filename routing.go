package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Routing using gorilla/mux

//Creating a new router
/*
The router is the main router for your web application and
will later be passed as parameter to the server. It will receive
all HTTP connections and pass it on to the request handlers
you will register on it. You can create a new router like so:
*/

//r := mux.NewRouter()

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You have requested the book: %s on page %s", title, page)
	})

	http.ListenAndServe(":80", r)

	//ROUTING METHODS
	//Restrict the request handler to specific methods
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//DOMAINS AND SUBDOMAINS
	//Restrict the request handler to specific hostnames and domains
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	//SCHEMES
	//Restrict the request handler to specific schemes
	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	//PATH PREFIXES AND SUBROUTERS
	//Restrict the request handler to specific path prefixes
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)
}
