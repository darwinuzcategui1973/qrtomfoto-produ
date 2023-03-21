package repository

import (
	//"log"
	"context"
	"encoding/json"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/entity"
)

const (
	qryInsertQrTomFoto = `
		insert into QRTOMFOTO (rif, ubicacion, nombre, archivo)
		values ( ?, ?, ?, ?);`

		qryGetQrTomFotoByRif = `
		select
			id,
			rif
		from QRTOMFOTO
		where rif = ?;`
	

	//from BdEmpleado.empleados e  ;
	// aqui voy un query prueba
	qryListasQrtomfotos = `
		select id, rif, ubicacion, fechaInicioEnvio  from QRTOMFOTO;`
)


func (r *repo) SaveQrTomFoto(ctx context.Context,rif, ubicacion, nombre string, archivo []int64) error {
	//comilla:="'"
	//archivo1 := strings.Join(archivo)
	//archivo1 := strconv.Itoa(archivo)
	archivo1, _ := json.Marshal(archivo)
	//var archivo1 = strings(archivo)+comilla
	_, err := r.db.ExecContext(ctx, qryInsertQrTomFoto, rif, ubicacion, nombre, archivo1)
	
	//log.Println("***************desde repositorio*******")
//	log.Println("aqui veo si sw queda*************************")
//	log.Println(err)
//	log.Println("aqui genero el error************************")
	return err
}

func (r *repo) GetQrTomByRif(ctx context.Context, rif string) (*entity.Qrtomfoto, error) {
	q := &entity.Qrtomfoto{}
	err := r.db.GetContext(ctx, q, qryGetQrTomFotoByRif, rif)
	if err != nil {
		return nil, err
	}

	return q, err
}

func (r *repo) GetListaQrTomFoto(ctx context.Context) ([]entity.Qrtomfoto, error) {
	qrtomfotos := []entity.Qrtomfoto{}
	err := r.db.SelectContext(ctx, &qrtomfotos, qryListasQrtomfotos)
	if err != nil {
		return nil, err
	}

	return qrtomfotos, nil

}

