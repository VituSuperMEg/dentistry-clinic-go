package application

import (
	"context"
	"dentistry-clinic/internal/domain/paciente"
	"dentistry-clinic/internal/infrastructure/repository"
)

type PacienteService struct {
	repo *repository.PacienteRepository
}

func NovoPacienteService(repo *repository.PacienteRepository) *PacienteService {
	return &PacienteService{repo: repo}
}

func (s *PacienteService) RegistrarPaciente(ctx context.Context, nome, cpf, telefone string) (*paciente.Paciente, error) {
	novoPaciente := paciente.NovoPaciente(nome, cpf, telefone)
	if err := s.repo.Criar(ctx, novoPaciente); err != nil {
		return nil, err
	}
	return novoPaciente, nil
}

func (s *PacienteService) ListarPacientes(ctx context.Context) ([]paciente.Paciente, error) {
	return s.repo.BuscarTodos(ctx)
}

func (s *PacienteService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
