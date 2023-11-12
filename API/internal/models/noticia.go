package models

import "net/http"

type Noticia struct {
	CodNotc        int    `json:"codNotc"`
	CodAutor       string `json:"codAutor"`
	Titulo         string `json:"titulo"`
	Noticia        string `json:"noticia"`
	DataPublicacao string `json:"dataPublicacao"`
}

type Feed []Noticia

// Valida uma instância de Noticia
func (n Noticia) IsValid() (bool, string) {
	if n.CodNotc < 0 || n.CodNotc > 99999999 {
		return false, "Código de notícia inválido."
	}

	if len(n.CodAutor) < 16 || len(n.CodAutor) > 36 {
		return false, "Código de autor inválido."
	}

	if len(n.Titulo) < 4 || len(n.Titulo) > 255 {
		return false, "Título de notícia deve conter de 4 a 255 caracteres."
	}

	if n.Noticia == "" {
		return false, "Corpo da notícia não pode estar vazio."
	}

	if len(n.DataPublicacao) != 19 {
		return false, "Data de publicação deve estar no modelo YYYY-MM-DD HH:MM:SS." +
			"Ex: 2023-01-01 12:00:00"
	}

	return true, ""
}

// Retorna um feed com as 6 últimas notícias
func (f Feed) GetFeed() (int, string) {
	//selectFeed := "SELECT * FROM Noticias ORDER BY data_publicacao LIMIT 6;"

	return http.StatusOK, ""
}
