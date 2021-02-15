package models

import "log"

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
	rows, err := h.db.Query("SELECT Id, IdRestaurante, Foto, Nome, Preço, Categoria, Promoção, DescriçãoPromoção, PreçoPromoção FROM Produtos")
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
