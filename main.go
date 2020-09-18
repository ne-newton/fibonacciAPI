package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

// fibonacci variable that is used to keep track of current and previous number
var f fibonacci

// database instance for writing fibonacci numbers for restart
var DB *gorm.DB

// buffered channel created, sends only dummy values to tie up access to fibonacci struct
var buf = make(chan bool, 1)

func main() {
	// setup connection to database, check for any previous fibonacci numbers
	checkDB()

	// create server
	s := gin.Default()

	//setup endpoints
	s.GET("/current", current)
	s.GET("/next", next)
	s.GET("/previous", previous)

	// start server, using endless for graceful restarts
	if err := endless.ListenAndServe("localhost:8080", s); err != nil {
		log.Fatal(err)
	}
}
