package models

import "log"

type Promoção struct {
	ID            int    `json:"ID"`
	IDProduto     int    `json:"IDProduto"`
	DiaDaSemana   int    `json:"DiaDaSemana"`
	HorarioInicio string `json:"HorarioInicio"`
	HorarioFim    string `json:"HorarioFim"`
}

func (h *BaseHandler) GetPromoções() []Promoção {
	rows, err := h.db.Query("SELECT Id, IdProduto, DiaDaSemana, HorarioInicio, HorarioFim FROM Promoção")
	if err != nil {
		log.Println(err)
	}
	var P []Promoção

	for rows.Next() {
		temp := Promoção{}
		err := rows.Scan(&temp.ID, &temp.IDProduto, &temp.DiaDaSemana, &temp.HorarioInicio, &temp.HorarioFim)
		if err != nil {
			log.Println(err)
		}
		P = append(P, temp)
	}

	return P
}
