package service

import (
	"context"
	"errors"
	"log"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
)

var (
	ErrCodigoAlreadyExists = errors.New("Codigo de Barra Existes")
	ErrCodigoNoExists      = errors.New("Codigo de Producto no existe En la base datos")
	//ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	//ErrRoleNotFound       = errors.New("role not found")
)

func (s *serv) GetProductosd3xdListas(ctx context.Context) ([]models.Produtosd3xd, error) {
	pp, err := s.repo.GetListaDeEnD3xdProductos(ctx)
	if err != nil {
		return nil, err
	}

	productos := []models.Produtosd3xd{}

	for _, p := range pp {
		productos = append(productos, models.Produtosd3xd{
			ID:            p.ID,
			Preferencia:   p.Preferencia,
			Pdescribe:     p.Pdescribe,
			Pdepartamento: p.Pdepartamento,
			Pventa1:       float64(p.Pventa1),
			Pventa2:       p.Pventa2,
			Pventa3:       p.Pventa3,
			Pventa4:       p.Pventa4,
			Pventa5:       p.Pventa5,
			Pempaque:      p.Pempaque,
			Pmedida:       p.Pmedida,
			Pmedempaque:   p.Pmedempaque,
			Pexiste:       p.Pexiste,
			Pintercode:    p.Pintercode,
			Pdatevoid:     p.Pdatevoid,
			Pmoneydif:     p.Pmoneydif,
			Ptasac:        p.Ptasac,
		})

	}

	return productos, nil
}

func (s *serv) GetBuscarProductosd3xdListas(ctx context.Context, buscar string) ([]models.Produtosd3xd, error) {
	pp, err := s.repo.GetBuscarListaDeEnD3xdProductos(ctx, buscar)
	println(buscar)
	if err != nil {
		return nil, err
	}

	productos := []models.Produtosd3xd{}

	for _, p := range pp {
		productos = append(productos, models.Produtosd3xd{
			ID:            p.ID,
			Preferencia:   p.Preferencia,
			Pdescribe:     p.Pdescribe,
			Pdepartamento: p.Pdepartamento,
			Pventa1:       float64(p.Pventa1),
			Pventa2:       p.Pventa2,
			Pventa3:       p.Pventa3,
			Pventa4:       p.Pventa4,
			Pventa5:       p.Pventa5,
			Pempaque:      p.Pempaque,
			Pmedida:       p.Pmedida,
			Pmedempaque:   p.Pmedempaque,
			Pexiste:       p.Pexiste,
			Pintercode:    p.Pintercode,
			Pdatevoid:     p.Pdatevoid,
			Pmoneydif:     p.Pmoneydif,
			Ptasac:        p.Ptasac,
		})

	}

	return productos, nil
}

func (s *serv) GetBuscarCodigodeBarraProductosd3xd(ctx context.Context, buscar string) ([]models.Produtosd3xd, error) {
	pp, err := s.repo.GetBuscarCodigoDeBarraEnD3xdProductos(ctx, buscar)
	println(buscar)
	if err != nil {
		return nil, err
	}
	//determinar el la longitud del codigo de barr
	if len(buscar) < 4 {
		return nil, err
	}

	productos := []models.Produtosd3xd{}

	for _, p := range pp {
		productos = append(productos, models.Produtosd3xd{
			ID:            p.ID,
			Preferencia:   p.Preferencia,
			Pdescribe:     p.Pdescribe,
			Pdepartamento: p.Pdepartamento,
			Pventa1:       float64(p.Pventa1),
			Pventa2:       p.Pventa2,
			Pventa3:       p.Pventa3,
			Pventa4:       p.Pventa4,
			Pventa5:       p.Pventa5,
			Pempaque:      p.Pempaque,
			Pmedida:       p.Pmedida,
			Pmedempaque:   p.Pmedempaque,
			Pexiste:       p.Pexiste,
			Pintercode:    p.Pintercode,
			Pdatevoid:     p.Pdatevoid,
			Pmoneydif:     p.Pmoneydif,
			Ptasac:        p.Ptasac,
		})

	}

	return productos, nil
}

// modificar
func (s *serv) ModificarBarcode(ctx context.Context, codigoProductoAModificar, codigoDeBarra string) error {
	//darwin
	//producto, err := s.repo.GetProdD3XD_repositorio(crx, codigoProductoAModificar)
	log.Println("------------ ModificarBarcode---------------------")
	log.Println(codigoProductoAModificar)
	log.Println(codigoDeBarra)
	log.Println("---------------------------------")
	// GetBuscarCodigoDeBarraEnD3xdProductos
	// lo primero es servicio validad que el codigo de producto exista y codigo de barra
	_codigoNoExiste, err := s.repo.GetBuscarCodigoDeBarraEnD3xdProductos(ctx, codigoProductoAModificar)
	//esta bien existe en la base de datos el codigo a modificar
	if err != nil {
		log.Println("aquiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
		return err
	}
	println(len(_codigoNoExiste))
	if len(_codigoNoExiste) == 0 {
		println("el codidgo no existe  en la base de datos ")
		println(_codigoNoExiste)
		println(len(_codigoNoExiste))
		return ErrCodigoNoExists

	}

	_invalidoCodigoBarra, err := s.repo.GetBuscarCodigoDeBarraEnD3xdProductos(ctx, codigoDeBarra)

	if err != nil {
		return err
	}

	if len(_invalidoCodigoBarra) >= 1 {
		//println("el codidgo existe  en la base de datos ")
		//println(_invalidoCodigo)
		//println(len(_invalidoCodigo))
		return ErrCodigoAlreadyExists

	}

	return s.repo.ModificarCodigoDeBarraEnD3xdProductos(ctx, codigoProductoAModificar, codigoDeBarra)
	//return nil //s.repo.ModificarCodigoDeBarraEnD3xdProductos(ctx, codigoProductoAModificar, codigoDeBarra)
}

func (s *serv) GetProduD3xD(ctx context.Context, id string) (*models.Produtosd3xd, error) {
	p, err := s.repo.GetProdD3XD_repositorio(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Produtosd3xd{
		ID:         p.CodigoProductoAModificar,
		Pintercode: p.CodigoDeBarra,
	}

	return product, nil
}
