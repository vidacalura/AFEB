package models

import "gopkg.in/guregu/null.v3"

type Modo string

const (
	Online     Modo = "online"
	Presencial Modo = "presencial"
)

type Torneio struct {
	CodTorn       int         `json:"codTorn"`
	Titulo        string      `json:"titulo"`
	Descricao     string      `json:"descricao"`
	Comentarios   null.String `json:"comentarios"`
	DataInicio    string      `json:"dataInicio"`
	DataFim       null.String `json:"dataFim"`
	Modo          Modo        `json:"modo"`
	Participantes int         `json:"participantes"`
	PlacarFinal   null.String `json:"placarFinal"`
}

// Valida uma instância de Torneio
func (t Torneio) IsValid() (bool, string) {
	if t.CodTorn < 0 || t.CodTorn > 99999999 {
		return false, "Código de torneio inválido."
	}

	if len(t.Titulo) < 7 || len(t.Titulo) > 100 {
		return false, "Título de torneio deve conter de 7 a 100 caracteres."
	}

	if len(t.Descricao) < 10 || len(t.Descricao) > 255 {
		return false, "Descrição do torneio deve conter de 10 a 255 caracteres."
	}

	if len(t.DataInicio) != 10 {
		return false, "Data de início do torneio deve estar no modelo YYYY-MM-DD."
	}

	if t.DataFim.Valid {
		if len(t.DataFim.String) != 10 {
			return false, "Data de encerramento do torneio deve estar no modelo YYYY-MM-DD."
		}
	}

	if t.Modo != Online && t.Modo != Presencial {
		return false, "Modo de torneio deve ser: online ou presencial."
	}

	if t.Participantes <= 1 || t.Participantes > 999 {
		return false, "Torneio não pode ter menos de 2 participantes ou mais que 999."
	}

	return true, ""
}
