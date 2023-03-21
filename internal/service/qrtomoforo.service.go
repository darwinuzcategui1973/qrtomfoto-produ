package service

import (
	"context"
	"errors"
	//"log"

//	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
)

var (
	ErrQrFotoExisteRif  = errors.New("QrTomdaFoto Rif ya existes!")
	
)

func (s *serv) IncluirQrTomFoto(ctx context.Context,rif string, ubicacion string, nombre string, archivo []int64) error {

	//q, _ := s.repo.GetQrTomByRif(ctx, rif)
	//if q != nil {
	//	return ErrQrFotoExisteRif
	//}
	
	//log.Println("***************desde servicio*******")
	//log.Println("pasa bien en el servicio***********************")

	return s.repo.SaveQrTomFoto(ctx, rif, ubicacion, nombre, archivo )
}

/*
func (s *serv) GetQrTomFotosListas(ctx context.Context) ([]models.QrtomFoto, error) {
	qq, err := s.repo.GetQrTomFotosListas(ctx)
	if err != nil {
		return nil, err
	}

	qrtomfotos := []models.QrtomFoto{}

	for _, q := range qq {
		qrtomfotos = append(qrtomfotos, models.QrtomFoto{
			ID:   q.ID,
			Rif:  q.Rif,
			Ubicacion: q.Ubicacion,
			Nombre: q.Nombre,
			//Archivo: q.Archivo,
			FechaInicioEnvio: q.FechaInicioEnvio,
		})

	}

	return qrtomfotos, nil
	

}

*/