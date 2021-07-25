// Imports
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	// http packages
	fmt.Println("Hello, world.")

	// Any strcuture
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(1000))
	})

	http.ListenAndServe(":8000", newMux)

}
