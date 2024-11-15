package despesa

import "time"

type TipoDespesa string

const (
	Material TipoDespesa = "material"
	Luz      TipoDespesa = "luz"
	Agua     TipoDespesa = "agua"
	Outros   TipoDespesa = "outros"
)

type Despesa struct {
	ID         string      `json:"id,omitempty"`
	Tipo       TipoDespesa `json:"tipo"`
	Valor      float64     `json:"valor"`
	Data       string      `json:"data"`
	Observacao string      `json:"observacao"`
	CriadoEm   time.Time   `json:"criado_em"`
}
