package models

type Trofeu struct {
	CodTrof int `json:"codTrof"`
	CodJog  int `json:"codJog"`
	CodTorn int `json:"codTorn"`
	Posicao int `json:"posicao"`
}

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
