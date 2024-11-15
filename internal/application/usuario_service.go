package application

import (
	"context"
	"dentistry-clinic/internal/domain/usuario"
	"dentistry-clinic/internal/infrastructure/repository"
	"dentistry-clinic/internal/infrastructure/security"
	"errors"
)

type UsuarioService struct {
	repo *repository.UsuarioRepository
}

func NovoUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) RegistrarUsuario(ctx context.Context, nome, email, senha, role string) (*usuario.Usuario, error) {
	novoUsuario := usuario.NovoUsuario(nome, email, senha, role)
	if err := s.repo.Criar(ctx, novoUsuario); err != nil {
		return nil, err
	}
	return novoUsuario, nil
}

func (s *UsuarioService) Login(ctx context.Context, email, senha string) (string, error) {
	user, err := s.repo.BuscarPorEmail(ctx, email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}
	if !user.VerificarSenha(senha) {
		return "", errors.New("senha incorreta")
	}
	return security.GenerateJWT(user.ID, string(user.Role))
}
