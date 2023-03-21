package service

import (
	"context"
	"time"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/repository"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, name, password string) error
	// incluir qrtomfoto
	IncluirQrTomFoto(ctx context.Context, rif, ubicacion, nombre string, archivo []int64 ) error
	//GetQrTomFotosListas(ctx context.Context) ([]models.QrtomFoto, error)

	LoginUser(ctx context.Context, name, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
	AddProdcut(ctx context.Context, product models.Product, userName string) error
	// listas usuarios
	GetUsuariosListas(ctx context.Context) ([]models.User, error)
	// Empresa conformacion
	ConformarEmpresa(ctx context.Context, rif, nombre, conformacion string, usuario, viewers int64, mac string, sistema string, version string, status string, fecha time.Time) error
	// listas productosd3xd
	GetProductosd3xdListas(ctx context.Context) ([]models.Produtosd3xd, error)

	GetBuscarProductosd3xdListas(ctx context.Context, buscar string) ([]models.Produtosd3xd, error)
	GetBuscarCodigodeBarraProductosd3xd(ctx context.Context, buscar string) ([]models.Produtosd3xd, error)

	ModificarBarcode(ctx context.Context, codigoProductoAModificar, codigoDeBarra string) error
}

/*
Rif     string `json:"rif" validate:"required,rif"`
	Nombre  string `json:"nombre" validate:"required"`
	Conform string `json:"conformacion" validate:"required,min=8"`
	Usuario int64 `json:"usuario" validate:"required"`
	Viewers int64 `json:"viewers" validate:"required"`
*/

type serv struct {
	repo repository.Repository
}

/*
// GetBuscarProductosd3xdListas implements Service

	func (*serv) GetBuscarProductosd3xdListas(ctx context.Context, buscar string) ([]models.Produtosd3xd, error) {
		panic("unimplemented")
	}
*/
func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
