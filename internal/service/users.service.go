package service

import (
	"context"
	"errors"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/encryption"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("role not found")
)

func (s *serv) RegisterUser(ctx context.Context, name, password string) error {

	u, _ := s.repo.GetUserByName(ctx, name)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	return s.repo.SaveUser(ctx, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, name, password string) (*models.User, error) {
	u, err := s.repo.GetUserByName(ctx, name)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:   u.ID,
		Name: u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {
	//darwin
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userID, roleID)
}

func (s *serv) GetUsuariosListas(ctx context.Context) ([]models.User, error) {
	//uu, err := s.repo.GetListaUsuario(ctx)
	uu, err := s.repo.GetLista(ctx)
	if err != nil {
		return nil, err
	}

	usuarios := []models.User{}

	for _, u := range uu {
		usuarios = append(usuarios, models.User{
			ID:   u.ID,
			Name: u.Name,
		})

	}

	return usuarios, nil
}
