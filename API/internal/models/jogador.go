package models

import (
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Jogador struct {
	CodJog         int         `json:"codJog"`
	Nome           string      `json:"nome"`
	Apelido        null.String `json:"apelido"`
	TituloAFEB     null.String `json:"tituloAFEB"`
	Info           null.String `json:"info"`
	EloRapid       null.Int    `json:"eloRapid"`
	EloBlitz       null.Int    `json:"eloBlitz"`
	DataNascimento string      `json:"dataNascimento"`
}

type RankingJogadores []Jogador

// Valida uma instância de Jogador
func (j Jogador) IsValid() (bool, string) {
	if j.CodJog < 0 || j.CodJog > 99999999 {
		return false, "Código de jogador inválido."
	}

	if len(j.Nome) < 4 || len(j.Nome) > 70 {
		return false, "Nome de jogador deve conter de 4 a 70 caracteres."
	}

	if j.Apelido.Valid {
		if len(j.Apelido.String) > 20 {
			return false, "Apelido deve conter até 20 caracteres."
		}
	}

	if j.TituloAFEB.Valid {
		if j.TituloAFEB.String != "MNB" && j.TituloAFEB.String != "CMB" &&
			j.TituloAFEB.String != "MB" && j.TituloAFEB.String != "MIB" &&
			j.TituloAFEB.String != "GMB" {
			return false, "Título de mestre inválido."
		}
	}

	if j.Info.Valid {
		if len(j.Info.String) > 255 {
			return false,
				"Informações de jogador não podem ultrapassar 255 caracteres."
		}
	}

	if j.EloRapid.Valid {
		if j.EloRapid.Int64 < 0 || j.EloRapid.Int64 > 9999 {
			return false, "Rating rápido deve estar entre 0 e 9999."
		}
	}

	if j.EloBlitz.Valid {
		if j.EloBlitz.Int64 < 0 || j.EloBlitz.Int64 > 9999 {
			return false, "Rating blitz deve estar entre 0 e 9999."
		}
	}

	if len(j.DataNascimento) != 10 {
		return false, "Data de nascimento deve estar no padrão: YYYY-MM-DD."
	}

	return true, ""
}

// Retorna o Top 10 jogadores da AFEB
func (r *RankingJogadores) GetTop10AFEB() (int, string) {
	selectRanking := "SELECT * FROM Jogadores ORDER BY elo_rapido LIMIT 10;"

	rows, err := E.DB.Query(selectRanking)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao receber ranking do banco de dados."
	}

	for rows.Next() {
		var j Jogador
		err := rows.Scan(&j.CodJog, &j.Nome, &j.Apelido, &j.TituloAFEB, &j.Info,
			&j.EloRapid, &j.EloBlitz, &j.DataNascimento)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				"Erro ao receber jogadores do banco de dados."
		}

		*r = append(*r, j)
	}

	if err := rows.Err(); err != nil {
		return http.StatusInternalServerError,
			"Erro ao conectar ao banco de dados."
	}

	defer rows.Close()

	return http.StatusOK, ""
}

// Encontra dados de um jogador
func (j *Jogador) GetJogador(codJog string) (int, string) {
	selectJog := "SELECT * FROM Jogadores WHERE cod_jog = ?;"
	row := E.DB.QueryRow(selectJog, codJog)

	err := row.Scan(&j.CodJog, &j.Nome, &j.Apelido, &j.TituloAFEB, &j.Info,
		&j.EloRapid, &j.EloBlitz, &j.DataNascimento)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, "Jogador não encontrado."
	}

	return http.StatusOK, ""
}

// Salva um jogador novo no banco de dados
func (j Jogador) RegistrarJogador() (int, string) {
	insert := `
		INSERT INTO Jogadores (nome, apelido, titulo_AFEB, info, elo_rapido,
		elo_blitz, data_nascimento) VALUES(?, ?, ?, ?, ?, ?, ?);`

	_, err := E.DB.Exec(insert, j.Nome, j.Apelido, j.TituloAFEB, j.Info,
		j.EloRapid, j.EloBlitz, j.DataNascimento)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao registrar usuário. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusCreated, ""
}

// Atualiza os dados de um jogador no sistema
func (j Jogador) EditarJogador() (int, string) {
	update := `
		UPDATE Jogadores SET nome = ?, apelido = ?, titulo_AFEB = ?, info = ?,
		elo_rapido = ?, elo_blitz = ?, data_nascimento = ? WHERE cod_jog = ?;`

	_, err := E.DB.Exec(update, j.Nome, j.Apelido, j.TituloAFEB, j.Info,
		j.EloRapid, j.EloBlitz, j.DataNascimento, j.CodJog)
	if err != nil {
		return http.StatusInternalServerError,
			"Erro ao editar jogador. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Exclui um jogador do sistema
func (j Jogador) ExcluirJogador() (int, string) {
	delete := "DELETE FROM Jogadores WHERE cod_jog = ?;"
	_, err := E.DB.Exec(delete, j.CodJog)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			"Erro ao excluir jogador. Informe o erro a um mantenedor do " +
				"projeto ou abra uma issue em nosso github."
	}

	return http.StatusOK, ""
}

// Verifica se um jogador existe a partir de seu código de jogador da AFEB
func JogadorExiste(codJog int) bool {
	selectJog := "SELECT nome FROM Jogadores WHERE cod_jog = ?;"
	row := E.DB.QueryRow(selectJog, codJog)

	var nome string
	if err := row.Scan(&nome); err != nil {
		log.Println(err)
		return false
	}

	return true
}
