package repository

import (
	"context"
	"dentistry-clinic/internal/domain/dentista"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DentistaRepository struct {
	collection *mongo.Collection
}

func NovoDentistaRepository(db *mongo.Database) *DentistaRepository {
	return &DentistaRepository{collection: db.Collection("dentistas")}
}

func (r *DentistaRepository) Criar(ctx context.Context, d *dentista.Dentista) error {
	_, err := r.collection.InsertOne(ctx, d)
	return err
}

func (r *DentistaRepository) BuscarTodos(ctx context.Context) ([]dentista.Dentista, error) {
	var dentistas []dentista.Dentista

	cursor, err := r.collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var dentista dentista.Dentista
		if err := cursor.Decode(&dentista); err != nil {
			return nil, err
		}
		fmt.Printf("Dentista encontrado: %+v\n", dentista)
		dentistas = append(dentistas, dentista)
	}
	return dentistas, nil
}
