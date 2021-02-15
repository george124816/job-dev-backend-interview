package models

import (
	"errors"
	"log"
)

type Restaurante struct {
	ID            int    `json:"ID"`
	Foto          []byte `json:"Foto"`
	Nome          string `json:"Nome"`
	Endereço      string `json:"Endereço"`
	Funcionamento string `json:"Funcionamento"`
}

func (h *BaseHandler) GetRestaurantes() ([]Restaurante, error) {
	rows, err := h.db.Query("SELECT Id, Foto, Nome, Endereço, Funcionamento FROM Restaurante")
	if err != nil {
		log.Println(err)
	}
	var R []Restaurante

	for rows.Next() {
		temp := Restaurante{}
		err := rows.Scan(&temp.ID, &temp.Foto, &temp.Nome, &temp.Endereço, &temp.Funcionamento)
		if err != nil {
			log.Println(err)
		}
		R = append(R, temp)
	}
	if len(R) == 0 {
		return R, errors.New("Não há restaurante cadastrado")
	}

	return R, nil
}

func (h *BaseHandler) GetRestaurante(ID int) (Restaurante, error) {

	var R Restaurante

	stmt, err := h.db.Prepare("SELECT Id, Foto, Nome, Endereço, Funcionamento FROM Restaurante Where ID = (?)")
	if err != nil {
		log.Println(err)
		return R, err
	}
	rows, err := stmt.Query(ID)
	if err != nil {
		log.Println(err)
		return R, err
	}
	for rows.Next() {
		err := rows.Scan(&R.ID, &R.Foto, &R.Nome, &R.Endereço, &R.Funcionamento)
		if err != nil {
			log.Println(err)
			return R, err
		}
	}
	if R.ID == 0 {
		return R, errors.New("Restaurante não encontrado")
	}
	return R, nil
}

func (h *BaseHandler) CadastrarRestaurante(R Restaurante) (Restaurante, error) {
	stmt, err := h.db.Prepare("INSERT INTO Restaurante (Foto, Nome, Endereço, Funcionamento) VALUES (?,?,?,?)")
	if err != nil {
		log.Println(err)
		return Restaurante{}, errors.New("Erro ao preparar o statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(R.Foto, R.Nome, R.Endereço, R.Funcionamento)
	if err != nil {
		log.Println(err)
		return Restaurante{}, errors.New("Erro ao inserir o restaurante")

	}

	Id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return Restaurante{}, errors.New("Não foi possivel obter o id após inserir o restaurante.")
	}

	R.ID = int(Id)

	return R, nil
}
