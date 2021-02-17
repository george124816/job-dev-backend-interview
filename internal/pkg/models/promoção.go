package models

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

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

func (h *BaseHandler) CadastrarPromoção(P Promoção) error {
	stmt, err := h.db.Prepare("INSERT INTO Promoção (IdProduto, DiaDaSemana, HorarioInicio, HorarioFim) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(P.IDProduto, P.DiaDaSemana, P.HorarioInicio, P.HorarioFim)
	if err != nil {
		return err
	}

	RA, _ := res.RowsAffected()
	if RA != 1 {
		return errors.New("Não foi possivel inserir a promoção")
	}

	return nil
}

func ValidaPromoção(P Promoção) error {
	match, _ := regexp.MatchString("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$", P.HorarioInicio)
	if !match {
		log.Println("Digite um horario de inicio válido")
		return errors.New("Horario de inicio inválido")
	}

	match, _ = regexp.MatchString("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$", P.HorarioFim)
	if !match {
		log.Println("Digite um horario de Fim válido")
		return errors.New("Horario de fim inválido")
	}

	hora1, _ := strconv.Atoi(P.HorarioInicio[:2])
	minuto1, _ := strconv.Atoi(P.HorarioInicio[3:])
	hora2, _ := strconv.Atoi(P.HorarioFim[:2])
	minuto2, _ := strconv.Atoi(P.HorarioFim[3:])

	minuto1 += hora1 * 60
	minuto2 += hora2 * 60

	if minuto1 > minuto2 || minuto2-minuto1 < 15 {
		return errors.New("Diferença do horario de inicio e fim deve ser maior que 15 minutos")
	}
	if minuto2-minuto1 < 15 {
		return errors.New("")
	}

	if P.DiaDaSemana < 1 && P.DiaDaSemana > 7 {
		return errors.New("Dia da semana inválido")
	}

	return nil
}
