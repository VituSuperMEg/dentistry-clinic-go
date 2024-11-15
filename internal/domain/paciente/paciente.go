package paciente

import "time"

type Paciente struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	Nome         string    `json:"nome"`
	CPF          string    `json:"cpf"`
	Telefone     string    `json:"telefone"`
	CriadoEm     time.Time `json:"criado_em"`
	AtualizadoEm time.Time `json:"atualizado_em"`
}

func NovoPaciente(nome, cpf, telefone string) *Paciente {
	paciente := &Paciente{
		Nome:     nome,
		CPF:      cpf,
		Telefone: telefone,
		CriadoEm: time.Now(),
	}

	return paciente
}
