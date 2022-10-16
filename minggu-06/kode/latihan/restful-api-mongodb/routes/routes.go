package routes

import (
	user "mongodb-restfulapi/controllers"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
		}
	)
	
	router.Run(":8081")
}