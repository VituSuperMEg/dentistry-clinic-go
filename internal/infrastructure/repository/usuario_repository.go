package repository

import (
	"context"
	"dentistry-clinic/internal/domain/usuario"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
	collection *mongo.Collection
}

func NovoUsuarioreRepository(db *mongo.Database) *UsuarioRepository {
	return &UsuarioRepository{collection: db.Collection("usuarios")}
}

func (r *UsuarioRepository) Criar(ctx context.Context, u *usuario.Usuario) error {
	_, err := r.collection.InsertOne(ctx, u)
	return err
}

func (r *UsuarioRepository) BuscarPorEmail(ctx context.Context, email string) (*usuario.Usuario, error) {
	var usuario usuario.Usuario
	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&usuario)
	return &usuario, err
}
