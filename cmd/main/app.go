package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
)

//const initPath = `/square`

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func HelloHandler(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}

	var obj any
	err = json.Unmarshal(jsonData, &obj)
	if err != nil {
		panic(err)
	}

	ctx.String(http.StatusOK, "hello from server")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"req": obj,
	//})
}

func main() {

	fmt.Println("Hello world")

	router := gin.Default()

	router.GET("/", IndexHandler)
	router.POST("/hello", HelloHandler)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Fatal(router.RunListener(listener))
	//log.Fatal(
	//	router.Run(),
	//)

}
