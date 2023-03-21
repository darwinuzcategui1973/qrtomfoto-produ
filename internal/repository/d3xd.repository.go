package repository

import (
	"context"
	"log"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/entity"
)

// from BdEmpleado.empleados e  ;
// aqui voy un query prueba
const (
	qryListaproductos = `
		select pcode,
	preferencia,
	pdescribe,
	pdepartamento,
	pventa1,
	pventa2,
	pventa3,
	pventa4,
	pventa5,
	pempaque,
	pmedida,
	pmedempaque,
	pexiste,
	pintercode,
	Pdatevoid,
	pmoneydif,
	ptasac
		from d3xd_dbase.productos  ORDER BY pexiste desc;`

	qryBuscarListaproductos = `
	select pcode,
	preferencia,
	pdescribe,
	pdepartamento,
	pventa1,
	pventa2,
	pventa3,
	pventa4,
	pventa5,
	pempaque,
	pmedida,
	pmedempaque,
	pexiste,
	pintercode,
	Pdatevoid,
	pmoneydif,
	ptasac
				from d3xd_dbase.productos
				WHERE pdescribe LIKE CONCAT('%',?,'%') ORDER BY pdescribe;`

	qryBuscarCodigoBarraproductos = `
	select pcode,
	preferencia,
	pdescribe,
	pdepartamento,
	pventa1,
	pventa2,
	pventa3,
	pventa4,
	pventa5,
	pempaque,
	pmedida,
	pmedempaque,
	pexiste,
	pintercode,
	Pdatevoid,
	pmoneydif,
	ptasac from d3xd_dbase.productos WHERE pcode = ?;`

	qryBuscarCodigoBarraproductos1 = `
	select pcode,
	preferencia,
	pdescribe,
	pdepartamento,
	pventa1,
	pventa2,
	pventa3,
	pventa4,
	pventa5,
	pempaque,
	pmedida,
	pmedempaque,
	pexiste,
	pintercode,
	Pdatevoid,
	pmoneydif,
	ptasac from d3xd_dbase.productos WHERE pintercode = ?;`

	qryModificarCodigoBarraproductos = `
	UPDATE d3xd_dbase.productos 
	SET pintercode =? 
	WHERE pcode = ?;`
)

/*
"UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5"
*/

/*
UPDATE nombre_tabla
SET columna1 = valor1, columna2 = valor2
WHERE columna3 = valor3
*/

func (r *repo) GetListaDeEnD3xdProductos(ctx context.Context) ([]entity.Produtosd3xd, error) {
	productos := []entity.Produtosd3xd{}
	//err := r.db.SelectContext(ctx, &users, qryLista)
	err := r.db.SelectContext(ctx, &productos, qryListaproductos)

	if err != nil {
		return nil, err
	}

	return productos, nil

}

func (r *repo) GetBuscarListaDeEnD3xdProductos(ctx context.Context, buscar string) ([]entity.Produtosd3xd, error) {
	productos := []entity.Produtosd3xd{}
	println("*********-- GetBuscarListaDeEnD3xdProductos --****")
	println(buscar)

	err := r.db.SelectContext(ctx, &productos, qryBuscarListaproductos, buscar)
	//err.:= r.db.Se
	if err != nil {
		//println(r.db.GetContext(ctx, &productos, qryBuscarListaproductos, buscar))
		return nil, err
	}

	return productos, nil

}

func (r *repo) GetBuscarCodigoDeBarraEnD3xdProductos(ctx context.Context, buscar string) ([]entity.Produtosd3xd, error) {
	productos := []entity.Produtosd3xd{}
	println("*********GetBuscarCodigoDeBarraEnD3xdProductos****")
	println(buscar)
	//println("*********ante*******************")
	//println(len(productos))
	//busUserID: userID,
	//buscar:= buscar

	err := r.db.SelectContext(ctx, &productos, qryBuscarCodigoBarraproductos, buscar)
	//println("*********despues*******************")
	valor := 0
	valor = len(productos)
	//println(valor)
	if valor == 0 {
		println("realizo  otra busqueda algo pintercode ")
		r.db.SelectContext(ctx, &productos, qryBuscarCodigoBarraproductos1, buscar)
	}

	//err.:= r.db.Se
	if err != nil {
		//println(r.db.GetContext(ctx, &productos, qryBuscarListaproductos, buscar))
		return nil, err
	}

	/*
		if valor = 0 {
			println("aqui es cero ")
		}
	*/

	return productos, nil

}

func (r *repo) ModificarCodigoDeBarraEnD3xdProductos(ctx context.Context, codigoProductoAModificar, codigoDeBarra string) error {
	data := entity.CodigoBarra{
		CodigoProductoAModificar: codigoProductoAModificar,
		CodigoDeBarra:            codigoDeBarra,
	}

	log.Println("ModificarCodigoDeBarraEnD3xdProductos aqui va grbar todo ok")
	log.Println(codigoDeBarra)
	//consulta, err := r.db.NamedExecContext(ctx, qryModificarCodigoBarraproductos, data)
	//consulta, err := r.db.NamedExecContext(ctx, qryModificarCodigoBarraproductos, codigoDeBarra, codigoProductoAModificar)
	_, err := r.db.Query(qryModificarCodigoBarraproductos, data.CodigoDeBarra, codigoProductoAModificar)
	//log.Println(res)
	return err
}

func (r *repo) GetProdD3XD_repositorio(ctx context.Context, id string) (*entity.CodigoBarra, error) {
	p := &entity.CodigoBarra{}

	err := r.db.GetContext(ctx, p, qryBuscarCodigoBarraproductos, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
