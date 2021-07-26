package main

import "net/http"

func main() {
	//Serving assets and files such as JS, CSS and Images
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
