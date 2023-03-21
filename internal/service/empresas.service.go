package service

import (
	"context"
	"errors"
	"time"

	//"log"
	//"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/encryption"
)

var (
	ErrEmpresaAlreadyExists = errors.New("Empresa already exists")
	ErrInvalidConformacion  = errors.New("invalid conformacion")
)

// ConformarEmpresa implements Service
func (s *serv) ConformarEmpresa(ctx context.Context, rif, nombre, conformacion string, usuario, viewers int64, mac, sistema, version, status string, fecha time.Time) error {
	u, _ := s.repo.GetEmpresaByRif(ctx, rif)
	if u != nil {
		return ErrEmpresaAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(conformacion))
	if err != nil {
		return err
	}

	conf := encryption.ToBase64(bb)

	return s.repo.SaveEmpresa(ctx, rif, nombre, conf, usuario, viewers, mac, sistema, version, status, fecha)
}
