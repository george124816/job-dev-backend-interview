package router

import (
	"github.com/george124816/job-dev-backend-interview/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	gin.ForceConsoleColor()
	router := gin.Default()

	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", controllers.Hello)
		v1.GET("/restaurantes", controllers.GetRestaurantes)
		v1.GET("/produtos", controllers.GetProdutos)
		v1.GET("/promoções", controllers.GetPromoções)
		v1.POST("/restaurante", controllers.InsertRestaurante)
		v1.GET("/restaurante/:id", controllers.GetRestaurante)
		v1.PUT("/restaurante/:id", controllers.AlterarRestaurante)
		v1.DELETE("/restaurante/:id", controllers.ExcluirRestaurante)
		v1.GET("/restaurante/:id/produtos", controllers.GetProdutosByRestaurante)
	}

	return router
}
