package infrastructure

import (
	"blogo-server/api/interfaces/controllers"

	gin "github.com/gin-gonic/gin"
)

// Router - router api server
var Router *gin.Engine

func init() {
	router := gin.Default()

	postController := controllers.NewPostController(NewSQLHandler())

	router.GET("/api/v1/posts", func(c *gin.Context) { postController.Index(c) })
	router.GET("/api/v1/posts/:id", func(c *gin.Context) { postController.Show(c) })

	Router = router
}
