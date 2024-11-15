package usuario

import "golang.org/x/crypto/bcrypt"

type Role string

const (
	Admin        Role = "admin"
	Dentista     Role = "dentista"
	Recepcioista Role = "recepcionista"
)

type Usuario struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
	Role  Role   `json:"role"`
}

func NovoUsuario(nome, email, senha, role string) *Usuario {
	hash, _ := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return &Usuario{
		ID:    "",
		Nome:  nome,
		Email: email,
		Senha: string(hash),
		Role:  Role(role),
	}
}

func (u *Usuario) VerificarSenha(senha string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(senha)) == nil
}
