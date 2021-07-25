package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, resp *restful.Response) {
	// W response
	x0 := time.Now()
	io.WriteString(resp, fmt.Sprintf("%s", x0))
	fmt.Println(x0)
}

func main() {
	// Create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// ADD ENDPOINT
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

// http://localhost:8000/ping
