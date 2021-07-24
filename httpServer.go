package main

import (
	"fmt"
	"net/http"
)

//A basic HTTp server is capable of the following
/*Process dynamic requests: Process incoming requests from
users who browse the website, log into their accounts or post images.

Serve static assets: Serve JavaScript, CSS and images to browsers
to create a dynamic experience for the user.

Accept connections: The HTTP Server must listen on a specific port to
be able to accept connections from the internet.*/

//PROCESS DYNAMIC REQUESTS
/*
	For the dynamic aspect, the http.Request contains all information
	about the request and it’s parameters. You can read GET
	parameters with r.URL.Query().Get("token") or POST
	parameters (fields from an HTML form) with r.FormValue("email").
*/

//SERVING STATIC ASSETS
//Examples of statc assets - JS, CSS, Images
//We use the inbuilt http.FileServer(http.Dir("static/"))
//It needs to know where to serve files from

/*
	Once our file server is in place, we just need to point a url
	path at it, just like we did with the dynamic requests. One
	thing to note: In order to serve files correctly, we need to
	strip away a part of the url path. Usually this is the name
	of the directory our files live in.
*/

//http.Handle("/static/", http.StripPrefix("/static/", fs))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
