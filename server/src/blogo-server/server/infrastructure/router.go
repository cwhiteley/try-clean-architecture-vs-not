package infrastructure

import (
	"blogo-server/server/handler"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

// Router - router api server
var Router *gin.Engine

func init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3030"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	router.POST("/api/v1/posts", func(c *gin.Context) { handler.CreatePost(c) })
	router.GET("/api/v1/posts", func(c *gin.Context) { handler.GetPost(c) })
	router.GET("/api/v1/posts/:id", func(c *gin.Context) { handler.ListPosts(c) })
	router.DELETE("/api/v1/posts/:id", func(c *gin.Context) { handler.DeletePost(c) })

	Router = router
}
