package main

import (
	"context"
	"fmt"
	"log"

	//"log"

	//"github.com/darwinuzcategui1973/qrtomfoto-produ/conformacion"
	//github.com/darwinuzcategui1973/qrtomfoto-produ
	"github.com/darwinuzcategui1973/qrtomfoto-produ/conformacion"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/database"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/api"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/repository"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/service"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/settings"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			conformacion.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),

		fx.Invoke(
			setLifeCycle,
			//prueba,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, lic *conformacion.LicenciaConfig, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			prueba := fmt.Sprintf("Nombre:%s", lic.Sistema)
			conformado := false
			validar := "Nombre:Appi-QR-TomFotos"
			if prueba == validar {
				log.Println(prueba + " Sistema Conformado" + conformacion.GetMac())
				conformado = true
			} else {
				log.Println(prueba + " Sistema sin Conformar")
				conformado = false

			}
			//macconformar := "7a:79:00:00:00:00"
			macconformar := "1"// conformacion.GetMac()
			validar1 := "1" // "02:42:ce:b0:73:c9"
			if macconformar == validar1 {
				log.Println(prueba + " Sistema Conformado" + conformacion.GetMac())
				conformado = true
			} else {
				log.Println(prueba + " Sistema sin Conformar comunicarse al 0414.921.32.35->"+ conformacion.GetMac())
				conformado = false

			}

			go a.Start(e, address, conformado)
			//go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
func prueba(db *sqlx.DB) {
	_, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		panic(err)

	}

}
