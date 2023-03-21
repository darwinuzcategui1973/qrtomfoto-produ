package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/users")
	products := e.Group("/products")
	empresas := e.Group("/conformacion")
	d3xd := e.Group("/d3xd")
	// lo nuevo de qrTomFoto
	qrtomfotos := e.Group("/qrtomfotos")

	//users
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)
    //qrtomfotos
	qrtomfotos.POST("/incluir", a.IncluirQrTomFoto)
	//qrtomfotos.GET("/listar", a.GetListasQrTomFotos)


	// realizandos los get
	users.GET("/usuarios", a.GetUsarios)
	// products
	products.POST("", a.AddProduct)
	//conformacion
	empresas.POST("regitar", a.EmpresaConforma)
	//d3xd
	d3xd.GET("/productos", a.GetProductosD3xD)
	d3xd.GET("/buscar", a.GetBuscarProductosD3xDDescripcion)
	d3xd.GET("/barcode", a.GetBuscarProductosD3xDCodigoBarra)
	d3xd.POST("/barcode", a.ModificarBarcode)

}
