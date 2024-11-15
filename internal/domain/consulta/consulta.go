package consulta

import (
	"dentistry-clinic/internal/domain/paciente"
	"time"
)

type consulta struct {
	ID         string             `json:"id" bson:"_id,omitempty"`
	PacienteID string             `json:"paciente_id"`
	Paciente   *paciente.Paciente `json:"paciente,omitempty"`
	Data       time.Time          `json:"data"`
	Observacao string             `json:"observacao"`
	CriadoEm   time.Time          `json:"criado_em"`
}
