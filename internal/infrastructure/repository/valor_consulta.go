package repository

import (
	"context"
	valorconsulta "dentistry-clinic/internal/domain/valor_consulta"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ValorconsultaRepository struct {
	collection *mongo.Collection
}

func NovoValorconsultaRepository(db *mongo.Database) *ValorconsultaRepository {
	return &ValorconsultaRepository{
		collection: db.Collection("valor_consulta"),
	}
}

func (r *ValorconsultaRepository) Criar(ctx context.Context, v *valorconsulta.Valorconsulta) error {
	_, err := r.collection.InsertOne(ctx, v)
	return err
}

func (r *ValorconsultaRepository) Listar(ctx context.Context) ([]*valorconsulta.Valorconsulta, error) {
	var valorConsultas []*valorconsulta.Valorconsulta

	cursor, err := r.collection.Find(ctx, bson.D{{}}, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var v valorconsulta.Valorconsulta
		err := cursor.Decode(&v)
		if err != nil {
			return nil, err
		}
		valorConsultas = append(valorConsultas, &v)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return valorConsultas, nil
}

func (r *ValorconsultaRepository) ValorDaConsulta(ctx context.Context, id string) (float64, error) {
	var result valorconsulta.Valorconsulta

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, errors.New("valor da consulta não encontrado")
		}
		return 0, err
	}

	return result.Valor, nil
}
