package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, response *http.Request) {
	if response.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("405 - Method not allowed"))

	}
}

func main() {
	http.HandleFunc("/city", mainLogic)
	http.ListenAndServe(":8000", nil)
}

// curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'
