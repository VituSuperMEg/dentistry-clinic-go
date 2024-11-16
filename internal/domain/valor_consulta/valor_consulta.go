package valorconsulta

type Valorconsulta struct {
	ID    string  `json:"id" bson:"_id,omitempty"`
	Valor float64 `json:"valor"`
}

func NovaConsulta(valor float64) *Valorconsulta {
	return &Valorconsulta{Valor: valor}
}
