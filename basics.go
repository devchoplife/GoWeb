package main

import (
	"fmt"
	"net/http"
)

func main() {
	//The net/http package from the standard library contains all functionalities
	//about the HTTP protocol. This includes (among many other things) an HTTP client and an HTTP server

	//Create a handler which receives all incoming HTTP connections from browsers, HTTP clients or API requests.
	//This function receives 2 parameters
	//An http.ResponseWriter which is where you write your text/html response to.
	//An http.Request which contains all information about this HTTP request including things like the URL or header fields.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
	})

	//Listen for HTTP connections

	//THe request handler alone cannot accept any HTTP connections, An HTTP server has to listen on a port to pass
	//connection on to the request handler
	//Port 80 is in most cases the default port for web traffic

	http.ListenAndServe(":80", nil)
}
