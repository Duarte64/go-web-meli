package main

import (
	"github.com/Duarte64/go-web-meli/cmd/server/handler"
	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, Gabriel!",
		})
	})

	routeUsers := router.Group("/users")
	{
		routeUsers.GET("", u.GetAll())
		routeUsers.GET("/:id", u.GetById())
		routeUsers.DELETE("/:id", u.Delete())
		routeUsers.PATCH("/:id", u.Patch())
		routeUsers.POST("", u.Store())
		routeUsers.PUT("/:id", u.Update())
	}

	router.Run(":8080")
}
