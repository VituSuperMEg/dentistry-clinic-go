package consulta

import (
	"dentistry-clinic/internal/domain/dentista"
	"dentistry-clinic/internal/domain/paciente"
	"time"
)

type StatusConsulta string
type TipoAtendimento string

const (
	TipoAtendimentoConsulta TipoAtendimento = "consulta"
	TipoAtendimentoExame    TipoAtendimento = "exame"
	TipoAtendimentoMaterial TipoAtendimento = "material"
	TipoAtendimentoOutros   TipoAtendimento = "outros"
)

const (
	StatusConsultaAgendada            StatusConsulta = "agendada"
	StatusConsultaRealizada           StatusConsulta = "realizada"
	StatusConsultaCancelada           StatusConsulta = "cancelada"
	StatusConsultaReagendada          StatusConsulta = "reagendada"
	StatusConsultaRetirada            StatusConsulta = "retirada"
	StatusConsultaAguardando          StatusConsulta = "aguardando"
	StatusConsultaAprovada            StatusConsulta = "aprovada"
	StatusConsultaReprovada           StatusConsulta = "reprovada"
	StatusConsultaExpirada            StatusConsulta = "expirada"
	StatusConsultaAguardandoPagamento StatusConsulta = "aguardando_pagamento"
	StatusConsultaPaga                StatusConsulta = "paga"
)

type Consulta struct {
	ID              string             `json:"id" bson:"_id,omitempty"`
	DestitaID       string             `json:"destita_id"`
	Destita         *dentista.Dentista `json:"destita,omitempty"`
	PacienteID      string             `json:"paciente_id"`
	Paciente        *paciente.Paciente `json:"paciente,omitempty"`
	DataAgendamento time.Time          `json:"data_agendamento"`
	DataRealizacao  *time.Time         `json:"data_realizacao"`
	DataRetorno     *time.Time         `json:"data_retorno,omitempty" bson:"data_retorno,omitempty"`
	TipoAtendimento TipoAtendimento    `json:"tipo_atendimento,string"`
	Observacao      string             `json:"observacao"`
	StatusConsulta  StatusConsulta     `json:"status_consulta"`
	ValorConsultaID string             `json:"valor_consulta_id'`
	CriadoEm        time.Time          `json:"criado_em"`
}

func NovaConsulta(destitaId, pacienteId, observacao, valor_consultaId string, status StatusConsulta, dataRetorno, data_realizacao *time.Time, data_agendamento time.Time, tipo TipoAtendimento) *Consulta {
	consulta := &Consulta{
		DestitaID:       destitaId,
		PacienteID:      pacienteId,
		Observacao:      observacao,
		ValorConsultaID: valor_consultaId,
		StatusConsulta:  status,
		CriadoEm:        time.Now(),
		DataRetorno:     dataRetorno,
		DataAgendamento: data_agendamento,
		DataRealizacao:  data_realizacao,
		TipoAtendimento: tipo,
	}
	return consulta
}
