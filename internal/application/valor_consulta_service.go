package application

import (
	"context"
	valorconsulta "dentistry-clinic/internal/domain/valor_consulta"
	"dentistry-clinic/internal/infrastructure/repository"
)

type ValorConsultaService struct {
	repo *repository.ValorconsultaRepository
}

func NovoValorConsultaService(repo *repository.ValorconsultaRepository) *ValorConsultaService {
	return &ValorConsultaService{repo: repo}
}

func (s *ValorConsultaService) Criar(ctx context.Context, valor float64) (*valorconsulta.Valorconsulta, error) {
	valorConsultas := valorconsulta.NovaConsulta(valor)
	if err := s.repo.Criar(ctx, valorConsultas); err != nil {
		return nil, err
	}
	return valorConsultas, nil
}

func (s *ValorConsultaService) ListarConsultas(ctx context.Context) ([]*valorconsulta.Valorconsulta, error) {
	return s.repo.Listar(ctx)
}
