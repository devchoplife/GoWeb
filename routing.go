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
}
