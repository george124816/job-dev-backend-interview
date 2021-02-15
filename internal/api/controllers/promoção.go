package controllers

import "github.com/gin-gonic/gin"

func GetPromoções(c *gin.Context) {
	promoções := handle.GetPromoções()
	c.JSON(200, promoções)
}
