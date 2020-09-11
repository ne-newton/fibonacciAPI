package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//endpoint for fibonacci.current(), outputs result as string
//buffered channel is engaged to block other API functions and then released at end
func current(c *gin.Context) {
	buf <- true
	c.String(
		http.StatusOK,
		"current -> %v",
		f.current(),
		)
	<- buf
}

//endpoint for fibonacci.next(), outputs result as string
//buffered channel is engaged to block other API functions and then released at end
func next(c *gin.Context) {
	buf <- true
	c.String(
		http.StatusOK,
		"next -> %v",
		f.next(),
		)
	<- buf
}

//endpoint for fibonacci.previous(), outputs result as string
//buffered channel is engaged to block other API functions and then released at end
func previous(c *gin.Context) {
	buf <- true
	if p, err := f.previous(); err != nil {
		c.String(
			http.StatusOK,
			"error: %s",
			fmt.Sprint(err),
			)
	} else {
		c.String(
			http.StatusOK,
			"previous -> %v",
			p,
			)
	}
	<- buf
}