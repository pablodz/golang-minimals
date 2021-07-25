// Imports
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {
	// http packages
	fmt.Println("Hello, world.")

	// Any strcuture
	newMux:=http.ServeMux()
	
	newMux.HandleFunc("/randomFloat",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w,rand.Float64())
	}

	http.ListenAndServe(":8000", mux)

}
