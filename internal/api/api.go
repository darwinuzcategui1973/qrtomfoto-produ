package api

import (
	//"context"
	//"os"

	"os"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{
		serv:          serv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string, conformado bool) error {
	//func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)

	//const conformado = true
	//if ()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.POST},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))

	// AllowOrigins:     []string{"http://127.0.0.1:5500"},

	confirmado := conformado
	if confirmado != true {
		print("Sistema no esta conformado ** COMUNICARSE AL 0414.921.32.35**  ")
		os.Exit(-1)

	}

	return e.Start(address)
}
