package repository

import (
	"context"
	"time"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context,  name, password string) error
	GetUserByName(ctx context.Context, name string) (*entity.User, error)

	// qrtomfoto
	SaveQrTomFoto(ctx context.Context,  rif, ubicacion, nombre string, archivo []int64) error
	GetQrTomByRif(ctx context.Context, rif string) (*entity.Qrtomfoto, error)
	GetListaQrTomFoto(ctx context.Context) ([]entity.Qrtomfoto, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRoles(ctx context.Context, userID int64) ([]entity.UserRole, error)
	// interface pare listas usuarios
	GetListaUsuario(ctx context.Context) ([]entity.User, error)
	GetLista(ctx context.Context) ([]entity.User, error)

	SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error
	GetProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id int64) (*entity.Product, error)
	// interface de conformacion de empresa
	SaveEmpresa(ctx context.Context, rif, nombre, conform string, usuario, viewers int64, mac string, sistema string, version string, status string, fecha time.Time) error
	GetEmpresaByRif(ctx context.Context, rif string) (*entity.Empresa, error)
	// interface pare listas productos d3x3d
	GetListaDeEnD3xdProductos(ctx context.Context) ([]entity.Produtosd3xd, error)

	GetBuscarListaDeEnD3xdProductos(ctx context.Context, buscar string) ([]entity.Produtosd3xd, error)
	//GetLista(ctx context.Context) ([]entity.User, error)
	GetBuscarCodigoDeBarraEnD3xdProductos(ctx context.Context, buscar string) ([]entity.Produtosd3xd, error)

	GetProdD3XD_repositorio(ctx context.Context, id string) (*entity.CodigoBarra, error)

	ModificarCodigoDeBarraEnD3xdProductos(ctx context.Context, codigoProductoAModificar, codigoDeBarra string) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
