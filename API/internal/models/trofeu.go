package models

import (
	"log"
	"net/http"
)

type Trofeu struct {
	CodTrof int    `json:"codTrof"`
	CodJog  int    `json:"codJog"`
	CodTorn int    `json:"codTorn"`
	Torneio string `json:"torneio"`
	Posicao int    `json:"posicao"`
}

type Trofeus []Trofeu

// Valida uma instância de Trofeu
func (t Trofeu) IsValid() (bool, string) {
	if t.CodTrof < 0 || t.CodTrof > 99999999 {
		return false, "Código de troféu inválido."
	}

	if t.CodJog <= 0 || t.CodJog > 99999999 {
		return false, "Código de jogador inválido."
	}

	if t.CodTorn <= 0 || t.CodTorn > 99999999 {
		return false, "Código de torneio inválido."
	}

	if t.Posicao <= 0 || t.Posicao > 3 {
		return false, "Posição do torneio deve ser de 1º, 2º ou 3º lugar."
	}

	return true, ""
}

// Retorna os troféus de um jogador específico
func (t *Trofeus) GetTrofeusJogador(codJog int) (int, string) {
	selectTrof := `
		SELECT Torneios.titulo, Trofeus.posicao
		FROM Trofeus
		INNER JOIN Torneios
		ON Trofeus.cod_torn = Torneios.cod_torn
		WHERE cod_jog = ?;`

	rows, err := E.DB.Query(selectTrof, codJog)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao retornar prêmios do jogador. Informe o erro a um" +
				"mantenedor do projeto ou abra uma issue em nosso github."
	}

	for rows.Next() {
		var trof Trofeu
		err := rows.Scan(&trof.Torneio, &trof.Posicao)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber prêmios do jogador."
		}

		*t = append(*t, trof)
	}

	if err := rows.Err(); err != nil {
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}
