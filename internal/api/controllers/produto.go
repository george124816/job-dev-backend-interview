package controllers

import (
	"bytes"
	"io"
	"log"
	"strconv"

	"github.com/george124816/job-dev-backend-interview/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetProdutos(c *gin.Context) {
	produtos := handle.GetProdutos()
	c.JSON(200, produtos)
}

func GetProdutosByRestaurante(c *gin.Context) {

	var Produtos []models.Produto

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Insira um inteiro"})
		return
	}
	Produtos, err = handle.GetProdutosByRestaurante(ID)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, Produtos)
	}
}

func CriarProduto(c *gin.Context) {

	var P models.Produto

	foto, _, err := c.Request.FormFile("Foto")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Foto inválida"})
		return
	}

	file := bytes.NewBuffer(nil)
	if _, err := io.Copy(file, foto); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Foto inválida"})
		return
	}

	P.IDRestaurante, _ = strconv.Atoi(c.PostForm("IdRestaurante"))
	P.Foto = file.Bytes()
	P.Nome = c.PostForm("Nome")
	P.Preço, _ = strconv.Atoi(c.PostForm("Preço"))
	P.Categoria = c.PostForm("Categoria")
	P.Promoção, _ = strconv.ParseBool(c.PostForm("Promoção"))
	P.DescriçãoPromoção = c.PostForm("DescriçãoPromoção")
	P.PreçoPromoção, _ = strconv.Atoi(c.PostForm("PreçoPromoção"))

	log.Println(P)
	c.JSON(200, "Ok")

	P, err = handle.InserirProduto(P)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, P)
	}
}

func AlterarProduto(c *gin.Context) {

	var P models.Produto

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Insira um inteiro"})
		return
	}

	foto, _, err := c.Request.FormFile("Foto")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Foto inválida"})
		return
	}

	file := bytes.NewBuffer(nil)
	if _, err := io.Copy(file, foto); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Foto inválida"})
		return
	}

	P.IDRestaurante, _ = strconv.Atoi(c.PostForm("IdRestaurante"))
	P.Foto = file.Bytes()
	P.Nome = c.PostForm("Nome")
	P.Preço, _ = strconv.Atoi(c.PostForm("Preço"))
	P.Categoria = c.PostForm("Categoria")
	P.Promoção, _ = strconv.ParseBool(c.PostForm("Promoção"))
	P.DescriçãoPromoção = c.PostForm("DescriçãoPromoção")
	P.PreçoPromoção, _ = strconv.Atoi(c.PostForm("PreçoPromoção"))

	P, err = handle.AlterarProduto(ID, P)
	if err != nil {
		log.Println(err)
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, P)
	}
}

func DeletarProduto(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Insira um inteiro"})
		return
	}

	err = handle.ExcluirProduto(ID)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, "Deleted")
	}
}
