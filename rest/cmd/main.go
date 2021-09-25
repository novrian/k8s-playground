package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Printf("entered on %s", "/")
		c.JSON(http.StatusOK, Response{
			Path:    "/",
			Message: "Just output message on root",
		})
	})
	log.Fatal(r.Run(":8080"))
}

type Response struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}
