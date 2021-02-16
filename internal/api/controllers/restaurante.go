package controllers

import (
	"bytes"
	"io"
	"log"
	"strconv"

	"github.com/george124816/job-dev-backend-interview/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetRestaurantes(c *gin.Context) {
	restaurantes, err := handle.GetRestaurantes()
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, restaurantes)
	}
}

func InsertRestaurante(c *gin.Context) {

	var R models.Restaurante

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

	R.Foto = file.Bytes()
	R.Nome = c.PostForm("Nome")
	R.Endereço = c.PostForm("Endereço")
	R.Funcionamento = c.PostForm("Funcionamento")

	if R.Foto != nil && R.Nome != "" && R.Endereço != "" && R.Funcionamento != "" {
		restaurante, err := handle.CadastrarRestaurante(R)
		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, restaurante)
		return
	}
	c.JSON(400, gin.H{"error": "Campos vazios"})
}

func GetRestaurante(c *gin.Context) {

	if c.Param("id") != "" {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Insira um inteiro"})
			return
		}
		R, err := handle.GetRestaurante(ID)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, R)
		}
	} else {
		c.JSON(400, gin.H{"error": "Parametro vazio"})
	}
}

func AlterarRestaurante(c *gin.Context) {

	var R models.Restaurante

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Insira um inteiro"})
		return
	}

	foto, _, err := c.Request.FormFile("Foto")
	if err != nil {
		log.Println(err)
		return
	}

	file := bytes.NewBuffer(nil)
	if _, err := io.Copy(file, foto); err != nil {
		log.Println(err)
		return
	}

	R.Foto = file.Bytes()
	R.Nome = c.PostForm("Nome")
	R.Endereço = c.PostForm("Endereço")
	R.Funcionamento = c.PostForm("Funcionamento")

	if ID != 0 && R.Foto != nil && R.Nome != "" && R.Endereço != "" && R.Funcionamento != "" {
		R, err = handle.AlterarRestaurante(ID, R)
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			c.JSON(200, R)
		}
	} else {
		c.JSON(400, gin.H{"error": "Parametro vazio"})
	}
}

func ExcluirRestaurante(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Insira um inteiro"})
		return
	}
	err = handle.ExcluirRestaurante(ID)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, "Restaurante excluido com sucesso!")
	}
}
