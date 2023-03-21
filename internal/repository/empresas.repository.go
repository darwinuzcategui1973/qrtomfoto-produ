package repository

import (
	context "context"
	"log"
	"time"

	entity "github.com/darwinuzcategui1973/qrtomfoto-produ/internal/entity"
)

const (
	qryInsertEmpresa = `
		insert into EMPRESA (
			rif, 
			nombre,
			conformacion,
			usuario,
			viewers,
			MACAddress,
			sistema,
			version,
			status,
			fechaInstalacion,
			fechaVencimiento
			
			) 
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	qryGetEmpresaByRif = `
				select
					id,
					rif,
					nombre
				from  EMPRESA
				where rif = ?;`
)

// GetEmpresaByRif implements Repository
func (r *repo) GetEmpresaByRif(ctx context.Context, rif string) (*entity.Empresa, error) {
	u := &entity.Empresa{}
	err := r.db.GetContext(ctx, u, qryGetEmpresaByRif, rif)
	if err != nil {
		return nil, err
	}

	//return u, nil
	return u, err
}

// viewers,
// MACAddress
// SaveEmpresa implements Repository
func (r *repo) SaveEmpresa(ctx context.Context, rif string, nombre string, conform string, usuario int64, viewers int64, mac string, sistema string, version string, status string, fecha time.Time) error {
	var datetime = time.Now()
	//dt := datetime.Format(time.RFC3339)
	//dt1:= dt.
	_, err := r.db.ExecContext(ctx, qryInsertEmpresa, rif, nombre, conform, usuario, viewers, mac, sistema, version, status, datetime, fecha)
	log.Println(err)
	return err
}
