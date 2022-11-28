package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//const initPath = `/square`

func main() {

	fmt.Println("Hello world")

	router := gin.Default()

	router.GET("/ping",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"res": "pong",
			})
		})
	log.Fatal(
		router.Run(),
	)

}
