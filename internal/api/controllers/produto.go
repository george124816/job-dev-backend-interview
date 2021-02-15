package controllers

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, "Ok")
}

func GetProdutos(c *gin.Context) {
	produtos := handle.GetProdutos()
	c.JSON(200, produtos)
}
