package repository

import (
	"context"
	"dentistry-clinic/internal/domain/consulta"
	"dentistry-clinic/internal/domain/pagamento"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConsultaRepository struct {
	colletion *mongo.Collection
	db        *mongo.Database
}

func NovoConsultaRepository(db *mongo.Database) *ConsultaRepository {
	return &ConsultaRepository{colletion: db.Collection("consultas")}
}

func (r *ConsultaRepository) Criar(ctx context.Context, c *consulta.Consulta) error {
	result, err := r.colletion.InsertOne(ctx, c)

	consultaID := result.InsertedID.(string)

	repoValor := NovoValorconsultaRepository(r.db)
	repoPagamento := NovoPagamentoRepository(r.db)

	go func(consultaID string) {
		valor, err := repoValor.ValorDaConsulta(ctx, consultaID)
		if err != nil {
			fmt.Printf("Erro ao buscar valor para consulta %s: %v\n", consultaID, err)
			return
		}
		pagamento := &pagamento.Pagamento{
			ConsultaId: consultaID,
			Valor:      valor,
			Status:     "Pendente",
		}
		err = repoPagamento.SalvarPagamento(ctx, pagamento)
		if err != nil {
			fmt.Printf("Erro ao criar pagamento para consulta %s: %v\n", consultaID, err)
		}
	}(consultaID)
	return err
}

func (r *ConsultaRepository) BuscarTodas(ctx context.Context) ([]*consulta.Consulta, error) {
	var consultas []*consulta.Consulta

	cursor, err := r.colletion.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var consulta consulta.Consulta
		err := cursor.Decode(&consulta)
		if err != nil {
			return nil, err
		}
		consultas = append(consultas, &consulta)
	}

	return consultas, nil
}

func (r *ConsultaRepository) BuscarPorID(ctx context.Context, id string) (*consulta.Consulta, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var consulta consulta.Consulta
	err = r.colletion.FindOne(ctx, bson.M{"_id": objID}).Decode(&consulta)
	return &consulta, err
}

func (r *ConsultaRepository) BuscarPorPaciente(ctx context.Context, id string) (*consulta.Consulta, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var consulta consulta.Consulta
	filter := bson.M{"paciente_id": objID}
	err = r.colletion.FindOne(ctx, filter).Decode(&consulta)
	return &consulta, err
}
