package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

var f fibonacci

//buffered channel created, sends only dummy values to tie up access to fibonacci struct
var buf = make(chan bool, 1)

func main() {
	//create server
	s := gin.Default()

	//setup endpoints
	s.GET("/current", current)
	s.GET("/next", next)
	s.GET("/previous", previous)

	//start server, using endless for graceful restarts
	if err := endless.ListenAndServe("localhost:8080", s); err != nil {
		log.Fatal(err)
	}
}
