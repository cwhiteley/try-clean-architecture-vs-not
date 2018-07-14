package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"blogo-server/api/interfaces/controllers"
)

// Router - router api server
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.GET("/posts", func(c *gin.Context) { postController.Index(c) })
	router.GET("/posts/:id", func(c *gin.Context) { postController.Show(c) })

	Router = router
}
