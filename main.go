package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const output = "<html>" +
	"<head><title>Hello, world!</title></head>" +
	"<body><h3>Hello, World!</h3><p>An example go app.</p></body>" +
	"</html>\n\n"

func HelloWorldGet(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte(output))
}

func main() {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err.Error())
	}

	// Add APIs and start server
	router.GET("/", HelloWorldGet)

	// Start serving the application
	configuredPort := os.Getenv("PORT")
	if configuredPort == "" {
		configuredPort = "8080"
	}
	port, err := strconv.Atoi(configuredPort)
	if err != nil {
		panic(fmt.Sprintf("Unable to get port: %v\n", err))
	}

	log.Println("Starting on port", port)

	// Start serving the application
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err.Error())
	}
}
