package repository

import (
	"context"
	"dentistry-clinic/internal/domain/paciente"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PacienteRepository struct {
	colletion *mongo.Collection
}

func NovoPacienteRepository(db *mongo.Database) *PacienteRepository {
	return &PacienteRepository{colletion: db.Collection("pacientes")}
}

func (r *PacienteRepository) Criar(ctx context.Context, p *paciente.Paciente) error {
	_, err := r.colletion.InsertOne(ctx, p)
	return err
}

func (r *PacienteRepository) BuscarPorId(ctx context.Context, id string) (*paciente.Paciente, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var paciente paciente.Paciente
	err = r.colletion.FindOne(ctx, bson.M{"_id": objID}).Decode(&paciente)
	return &paciente, err
}

func (r *PacienteRepository) BuscarTodos(ctx context.Context) ([]paciente.Paciente, error) {
	var pacientes []paciente.Paciente

	cursor, err := r.colletion.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var paciente paciente.Paciente
		if err := cursor.Decode(&paciente); err != nil {
			return nil, err
		}
		pacientes = append(pacientes, paciente)
	}

	return pacientes, nil
}
