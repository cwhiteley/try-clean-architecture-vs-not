package infrastructure

import (
	"blogo-server/api/interfaces/controllers"

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

	postController := controllers.NewPostController(NewSQLHandler())

	router.POST("/api/v1/posts", func(c *gin.Context) { postController.Create(c) })
	router.GET("/api/v1/posts", func(c *gin.Context) { postController.Index(c) })
	router.GET("/api/v1/posts/:id", func(c *gin.Context) { postController.Show(c) })
	router.DELETE("/api/v1/posts/:id", func(c *gin.Context) { postController.Delete(c) })

	Router = router
}
