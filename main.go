package main

import (
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
	err = router.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
