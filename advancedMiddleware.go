package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Here we define a new type Middleware which makes it eventually easier to
chain multiple middlewares together

Form more information check https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
*/

/*
CREATING A NEW MIDDLEWARE

func createNewMiddleware() Middleware {

    // Create a new Middleware
    middleware := func(next http.HandlerFunc) http.HandlerFunc {

        // Define the http.HandlerFunc which is called by the server eventually
        handler := func(w http.ResponseWriter, r *http.Request) {

            // ... do middleware things

            // Call the next middleware/handler in chain
            next(w, r)
        }

        // Return newly created handler
        return handler
    }

    // Return newly created middleware
    return middleware
}
*/

type Middleware func(http.HandlerFunc) http.HandlerFunc

//logging logs all requests with its path and the time it took to process
func logging() Middleware {
	//Create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		//Define the http handlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			//Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			//Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

//Method ensures that URL can only be requested with a specifi method else,
//returns a 400 Bad Request
func Method(m string) Middleware {
	//Create a bnew Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		//Define the http.handlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			//Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			//Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

//This applies middlewares to an http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), logging()))
	http.ListenAndServe(":8080", nil)
}
