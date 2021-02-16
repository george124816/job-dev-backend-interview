package models

import (
	"errors"
	"log"
)

type Produto struct {
	ID                int    `json:"Id"`
	IDRestaurante     int    `json:"IdRestaurante"`
	Foto              []byte `json:"Foto"`
	Nome              string `json:"Nome"`
	Preço             int    `json:"Preço"`
	Categoria         string `json:"Categoria"`
	Promoção          bool   `json:"Promoção"`
	DescriçãoPromoção string `json:"DescriçãoPromoção"`
	PreçoPromoção     int    `json:"PreçoPromoção"`
}

func (h *BaseHandler) GetProdutos() []Produto {
	rows, err := h.db.Query("SELECT Id, IdRestaurante, Foto, Nome, Preço, Categoria, Promoção, DescriçãoPromoção, PreçoPromoção FROM Produto")
	if err != nil {
		log.Println(err)
	}
	var P []Produto

	for rows.Next() {
		temp := Produto{}
		err := rows.Scan(&temp.ID, &temp.IDRestaurante, &temp.Foto, &temp.Nome, &temp.Preço, &temp.Categoria, &temp.Promoção, &temp.DescriçãoPromoção, &temp.PreçoPromoção)
		if err != nil {
			log.Println(err)
		}
		P = append(P, temp)
	}

	return P
}

func (h *BaseHandler) GetProdutosByRestaurante(ID int) ([]Produto, error) {
	var Produtos []Produto
	stmt, err := h.db.Prepare("SELECT Id, IdRestaurante, Foto, Nome, Preço, Categoria, Promoção, DescriçãoPromoção, PreçoPromoção FROM Produto WHERE IdRestaurante = (?)")
	if err != nil {
		log.Println(err)
		return Produtos, err
	}
	rows, err := stmt.Query(ID)
	if err != nil {
		log.Println(err)
		return Produtos, err
	}

	for rows.Next() {
		var temp Produto
		err := rows.Scan(&temp.ID, &temp.IDRestaurante, &temp.Foto, &temp.Nome, &temp.Preço, &temp.Categoria, &temp.Promoção, &temp.DescriçãoPromoção, &temp.PreçoPromoção)
		if err != nil {
			log.Println(err)
		}
		Produtos = append(Produtos, temp)
	}

	if len(Produtos) == 0 {
		return Produtos, errors.New("O restaurante não contém produtos.")
	}

	return Produtos, nil
}

func (h *BaseHandler) InserirProduto(P Produto) (Produto, error) {
	stmt, err := h.db.Prepare("INSERT INTO Produto (IdRestaurante, Foto, Nome, Preço, Categoria, Promoção, DescriçãoPromoção,PreçoPromoção) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return P, err
	}

	res, err := stmt.Exec(P.IDRestaurante, P.Foto, P.Nome, P.Preço, P.Categoria, P.Promoção, P.DescriçãoPromoção, P.PreçoPromoção)
	if err != nil {
		log.Println(err)
		return P, err
	}
	RA, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return P, errors.New("Não foi possivel inserir o produto")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	if RA != 0 {
		P.ID = int(id)
	}
	return P, nil
}

func (h *BaseHandler) AlterarProduto(ID int, ProdutoInput Produto) (Produto, error) {
	var ProdutoOutput Produto

	stmt, err := h.db.Prepare("UPDATE Produto SET IdRestaurante = (?), Foto = (?), Nome = (?), Preço = (?), Categoria = (?), Promoção = (?), DescriçãoPromoção = (?), PreçoPromoção = (?) WHERE Id = (?)")
	if err != nil {
		log.Println(err)
		return Produto{}, err
	}

	res, err := stmt.Exec(ProdutoInput.IDRestaurante, ProdutoInput.Foto, ProdutoInput.Nome, ProdutoInput.Preço, ProdutoInput.Categoria, ProdutoInput.Promoção, ProdutoInput.DescriçãoPromoção, ProdutoInput.PreçoPromoção, ID)
	if err != nil {
		log.Println(err)
		return Produto{}, err
	}

	RA, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return Produto{}, err
	}
	if RA == 1 {
		ProdutoOutput = ProdutoInput
		ProdutoOutput.ID = ID
	} else {
		log.Println("Não foi possivel atualizar")
		return Produto{}, err
	}
	return ProdutoOutput, nil
}

func (h *BaseHandler) ExcluirProduto(ID int) error {
	stmt, err := h.db.Prepare("DELETE FROM Produto WHERE Id = (?)")
	if err != nil {
		log.Println(err)
		return err
	}
	res, err := stmt.Exec(ID)
	if err != nil {
		log.Println(err)
		return err
	}

	RA, _ := res.RowsAffected()
	if RA != 1 {
		log.Println("Erro ao deletar um Produto")
		return errors.New("Erro ao deletar um produto")
	} else {
		return nil
	}
}
