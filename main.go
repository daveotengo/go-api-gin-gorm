package main

import (
	controllers "example/go-api-gin-gorm/Controllers"
	"example/go-api-gin-gorm/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func YourCustomMiddleware(c *gin.Context) {
	// Configure Gin to trust specific proxies
	c.Request.Header.Set("X-Forwarded-For", "trusted-proxy-ip1, trusted-proxy-ip2")

	// Continue processing the request
	c.Next()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PATCH("/posts/:id", controllers.PostsUpdateById)
	r.DELETE("/posts/:id", controllers.PostsDeleteById)

	// r.Set("TrustProxy", true)
	// r.Use(YourCustomMiddleware)
	r.Run()

}
