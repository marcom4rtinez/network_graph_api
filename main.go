package main

import (
	"network_graph_api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // allow requests from any origin
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"} // allow all HTTP methods
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(config))

	router.GET("/", handlers.Ping)
	router.GET("/getNodes", handlers.GetAllNodeNames)
	router.GET("/getLinks", handlers.GetAllLinkConnections)

	router.Run()
}
