package models

type Usuario struct {
	CodUsu   string `json:"codUsu"`
	Username string `json:"username"`
	Senha    string `json:"senha"`
	Adm      bool   `json:"adm"`
}

// Valida uma instância de Usuario
func (u Usuario) IsValid() (bool, string) {
	if u.CodUsu == "" || len(u.CodUsu) > 16 {
		return false, "Código de usuário inválido."
	}

	if len(u.Username) < 5 || len(u.Username) > 30 {
		return false, "Username deve conter de 5 a 30 caracteres."
	}

	if len(u.Senha) < 8 || len(u.Senha) > 32 {
		return false, "Senha de usuário deve conter entre 8 e 32 caracteres"
	}

	return true, ""
}
