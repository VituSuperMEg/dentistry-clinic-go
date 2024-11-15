package application

import (
	"context"
	"dentistry-clinic/internal/domain/dentista"
	"dentistry-clinic/internal/infrastructure/repository"
)

type DentistaService struct {
	repo *repository.DentistaRepository
}

func NovoDentistaService(repo *repository.DentistaRepository) *DentistaService {
	return &DentistaService{repo: repo}
}

func (s *DentistaService) Criar(ctx context.Context, nome, cro string) (*dentista.Dentista, error) {
	novoDentista := dentista.NovoDentista(nome, cro)
	if err := s.repo.Criar(ctx, novoDentista); err != nil {
		return nil, err
	}
	return novoDentista, nil
}

func (s *DentistaService) Lista(ctx context.Context) ([]dentista.Dentista, error) {
	return s.repo.BuscarTodos(ctx)
}
