package controllers

import (
	"log"
	"strconv"

	"github.com/george124816/job-dev-backend-interview/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetPromoções(c *gin.Context) {
	promoções := handle.GetPromoções()
	c.JSON(200, promoções)
}

func AdicionarPromoção(c *gin.Context) {
	var P models.Promoção
	var err error

	P.DiaDaSemana, err = strconv.Atoi(c.PostForm("DiaDaSemana"))
	if err != nil {
		log.Println("Dia da semana deve ser um inteiro de 1-7")
		c.JSON(400, gin.H{"error": "Dia da semana deve ser um inteiro de 1-7"})
		return
	}
	P.IDProduto, err = strconv.Atoi(c.PostForm("IdProduto"))
	if err != nil {
		log.Println("ID do Produto deve ser um inteiro")
		c.JSON(400, gin.H{"Error": "Id do produto deve ser um inteiro"})
		return
	}
	P.HorarioInicio = c.PostForm("HorarioInicio")
	P.HorarioFim = c.PostForm("HorarioFim")

	err = (models.ValidaPromoção(P))
	if err != nil {
		log.Println(err)
		c.JSON(400, err.Error())
		return
	}

	err = handle.CadastrarPromoção(P)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, "Promoção inserida com sucesso.")
	}
}
