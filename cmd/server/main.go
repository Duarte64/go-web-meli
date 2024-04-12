package main

import (
	"os"

	"github.com/Duarte64/go-web-meli/cmd/server/handler"
	"github.com/Duarte64/go-web-meli/cmd/server/middleware/guards"
	"github.com/Duarte64/go-web-meli/docs"
	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/Duarte64/go-web-meli/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error ao carregar o arquivo .env")
	}

	db := store.New(store.FileType, "./users.json")

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, Gabriel!",
		})
	})

	routeUsers := router.Group("/users")
	routeUsers.Use(guards.TokenAuthMiddleware())
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
