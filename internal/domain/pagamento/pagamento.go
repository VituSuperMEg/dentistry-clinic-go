package pagamento

type Banco string
type FormaPagamento string
type Status string

const (
	CartaoCredito FormaPagamento = "cartao_credito"
	CartaoDebito  FormaPagamento = "cartao_debito"
	Dinheiro      FormaPagamento = "dinheiro"
	Pix           FormaPagamento = "pix"
)

const (
	AguardandoPagamento Status = "Pendente"
	PagamentoConfirmado Status = "Confirmado"
	PagamentoRejeitado  Status = "Rejeitado"
)

type Pagamento struct {
	ID         string          `json:"id,omitempty"`
	ConsultaId string          `json:"consulta_id,omitempty"`
	Valor      float64         `json:"valor"`
	Forma      *FormaPagamento `json:"forma_pagamento"`
	Status     Status          `json:"status"`
}
