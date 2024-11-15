package dentista

import "time"

type Dentista struct {
	ID       string    `json:"id" bson:"_id,omitempty"`
	Nome     string    `json:"nome"`
	CRO      string    `json:"cro"`
	CriadoEm time.Time `json:"criado_em"`
}

func NovoDentista(nome, cro string) *Dentista {
	return &Dentista{
		Nome:     nome,
		CRO:      cro,
		CriadoEm: time.Now(),
	}
}
