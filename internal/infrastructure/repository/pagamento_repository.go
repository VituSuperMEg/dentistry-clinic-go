package repository

import (
	"context"
	"dentistry-clinic/internal/domain/pagamento"

	"go.mongodb.org/mongo-driver/mongo"
)

type PagamentoRepository struct {
	colletion *mongo.Collection
}

func NovoPagamentoRepository(db *mongo.Database) *PagamentoRepository {
	return &PagamentoRepository{colletion: db.Collection("pagamentos")}
}

func (r *PagamentoRepository) SalvarPagamento(ctx context.Context, p *pagamento.Pagamento) error {
	_, err := r.colletion.InsertOne(ctx, p)
	return err
}
