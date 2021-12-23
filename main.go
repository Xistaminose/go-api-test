package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"

	routes "./routes"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}