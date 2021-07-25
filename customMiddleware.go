package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing before")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing after")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler")
	w.Write([]byte("OK"))
}

func main() {
	// main
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}
