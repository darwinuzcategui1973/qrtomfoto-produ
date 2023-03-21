package api

import (
	"log"
	"net/http"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/conformacion"

	"github.com/darwinuzcategui1973/qrtomfoto-produ/encryption"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/api/dtos"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/models"
	"github.com/darwinuzcategui1973/qrtomfoto-produ/internal/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

type listaParams struct {
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
	Buscar string `query:"buscar"`
}

type listaPrueba struct {
	Conformacion string `json:"conformacion" validate:"required,min=8"`
}

var (
	lista = []string{"0 usuario 1", "1 usuario 2", "2 usuario 3", "3 usuario 4"}
)

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx,  params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se Creo Sastifatoriamente el usuario"})
}

// Incluir QrTomFoto

func (a *API) IncluirQrTomFoto(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.IncluirQrtomfoto{}
	
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	//IncluirQrTomFoto(ctx context.Context,rif, ubicacion, nombre, archivo string) error {
	err = a.serv.IncluirQrTomFoto(ctx,  params.Rif, params.Ubicacion, params.Nombre, params.Archivo)
	//log.Println("***************desde handle*******")
	//log.Println("pasa el archivo************")
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Rif ya esistes!"})
		}



		//return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se Creo Sastifatoriamente el Documento en BD"})
}
// las lista qrtomfoto
/*
func (a *API) GetListasQrTomFotos(c echo.Context) error {

	ctx := c.Request().Context()
	params := listaParams{}
	err := c.Bind(&params)

	qrTomFotosLista, err := a.serv.GetQrTomFotosListas(ctx)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query0 parametros"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if params.Offset > len(qrTomFotosLista) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query1 parametros"})

	}

	if params.Limit < 0 || params.Limit > len(qrTomFotosLista) {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query2 parametros"})

	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(qrTomFotosLista)
	}


	return c.JSON(http.StatusOK, qrTomFotosLista[from:to])

}
*/

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}
	//params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Name, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)

	// visualizar los header
	for key, values := range c.Request().Header {
		log.Println("****************key**********")
		log.Println(key)
		log.Println("**************************")
		for _, value := range values {

			log.Println("----------valor--------")
			log.Println(value)
			log.Println("----------valor--------")
		}
	}

	/*
			for key, values := range context.Request().Header {
		            fmt.Println(key)
		            for _,value := range values {
		                fmt.Println(value)
		            }
		        }
	*/
	log.Println("******* login*********")

	return c.JSON(http.StatusOK, map[string]string{"success": "true", "token": token, "nombre": params.Name})
}

func (a *API) AddProduct(c echo.Context) error {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Unauthorized"})
	}

	name := claims["name"].(string)

	ctx := c.Request().Context()
	params := dtos.AddProduct{}

	err = c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	p := models.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	err = a.serv.AddProdcut(ctx, p, name)
	if err != nil {
		log.Println(err)

		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Invalid permissions"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	// visualizar los header
	for key, values := range c.Request().Header {
		log.Println("*ooo***************key**********")
		log.Println(key)
		log.Println("****oo**********************")
		for _, value := range values {

			log.Println("--ii--------valor--------")
			log.Println(value)
			log.Println("----ii------valor--------")
		}
	}

	return c.JSON(http.StatusOK, nil)
}

// GET DE USUARIOS

func (a *API) GetUsarios(c echo.Context) error {

	ctx := c.Request().Context()
	params := listaParams{}
	err := c.Bind(&params)

	usuarioLista, err := a.serv.GetUsuariosListas(ctx)
	//datos, err := Get
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query0 parametros"})
	}
	//println(len(usuarioLista))

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if params.Offset > len(usuarioLista) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query1 parametros"})

	}

	if params.Limit < 0 || params.Limit > len(usuarioLista) {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query2 parametros"})

	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(usuarioLista)
	}

	var darwin string
	var darwin2 conformacion.LicenciaConfig
	//darwin = conformacion.GetMac()

	darwin = conformacion.GetSistema()
	darwin2 = *conformacion.GetDatos()

	log.Print(darwin2.Sistema)
	log.Print(darwin)

	return c.JSON(http.StatusOK, usuarioLista[from:to])

}

func (a *API) EmpresaConforma(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.ConformarEmpresas{}
	mac := conformacion.GetMac()

	sistema := conformacion.GetSistema()
	datos := *conformacion.GetDatos()
	fecha := conformacion.GetFecha()
	//fecha :

	log.Println(fecha)

	log.Println(datos)

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	//err = a.serv.ConformarEmpresa(ctx, params.Rif, params.Nombre, params.Conform, params.Usuario, params.Viewers)
	//err = a.serv.ConformarEmpresa(ctx, params.Rif, params.Nombre, params.Conformacion, mac, params.Usuario, params.Viewers)
	err = a.serv.ConformarEmpresa(ctx, params.Rif, params.Nombre, params.Conformacion, params.Usuario, params.Viewers, mac, sistema, datos.Version, datos.Status, fecha)
	//log.Println(_)
	if err != nil {
		if err == service.ErrEmpresaAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "Empresa already exists"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusCreated, responseMessage{Message: "Se Creo Sastifatoriamente el usuario"})
}

func (a *API) GetProductosD3xD(c echo.Context) error {

	ctx := c.Request().Context()
	params := listaParams{}
	err := c.Bind(&params)

	println("**** Llamado el metodo ********************")
	ProductosDeD3xDLista, err := a.serv.GetProductosd3xdListas(ctx)
	//datos, err := Get
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query0 parametros"})
	}
	println(len(ProductosDeD3xDLista))

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if params.Offset > len(ProductosDeD3xDLista) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query1 parametros"})

	}

	if params.Limit < 0 || params.Limit > len(ProductosDeD3xDLista) {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query2 parametros"})

	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(ProductosDeD3xDLista)
	}

	return c.JSON(http.StatusOK, ProductosDeD3xDLista[from:to])

}

func (a *API) GetBuscarProductosD3xDDescripcion(c echo.Context) error {

	ctx := c.Request().Context()
	var buscar string
	//buscar := ""
	params := listaParams{}
	err := c.Bind(&params)

	println("**** Llamado el metodo buscar ********************")

	if params.Buscar != "" {
		buscar = params.Buscar
	}
	println("************************")
	println(buscar)
	println("************************")

	ProductosDeD3xDLista, err := a.serv.GetBuscarProductosd3xdListas(ctx, buscar)
	//datos, err := Get
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query0 parametros"})
	}
	println(len(ProductosDeD3xDLista))

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if params.Offset > len(ProductosDeD3xDLista) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query1 parametros"})

	}

	if params.Limit < 0 || params.Limit > len(ProductosDeD3xDLista) {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query2 parametros"})

	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(ProductosDeD3xDLista)
	}

	return c.JSON(http.StatusOK, ProductosDeD3xDLista[from:to])

}

func (a *API) GetBuscarProductosD3xDCodigoBarra(c echo.Context) error {

	ctx := c.Request().Context()
	var buscar string
	//buscar := ""
	params := listaParams{}
	err := c.Bind(&params)

	println("**** Llamado el metodo buscar  CodigoBarra********************")

	if params.Buscar != "" {
		buscar = params.Buscar
	}
	println("************************")
	println(buscar)
	println("************************")

	ProductosDeD3xDLista, err := a.serv.GetBuscarCodigodeBarraProductosd3xd(ctx, buscar)
	//datos, err := Get
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if len(buscar) < 4 {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "no se Puede Ejecutar Codigo de barra no valido"})
	}

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query0 parametros"})
	}
	println(len(ProductosDeD3xDLista))

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	if params.Offset > len(ProductosDeD3xDLista) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query1 parametros"})

	}

	if params.Limit < 0 || params.Limit > len(ProductosDeD3xDLista) {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido query2 parametros"})

	}

	var from, to int
	if params.Offset > 0 {
		from = params.Offset
	}
	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(ProductosDeD3xDLista)
	}

	return c.JSON(http.StatusOK, ProductosDeD3xDLista[from:to])

}

// AQUI ES LA LOGICA DE MODIFICAR
func (a *API) ModificarBarcode(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.ModificarBarcode{}
	log.Println(ctx)
	log.Println(params)

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalido el Json del body request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	p := models.Produtosd3xd{
		ID:         params.Id,
		Pintercode: params.CodBarra,
	}
	log.Println("ModificarBarcode")
	log.Println(params.Id)
	log.Println(p.ID)
	log.Println(params.CodBarra)
	log.Println(p.Pintercode)
	log.Println(p.Pdescribe)
	err = a.serv.ModificarBarcode(ctx, params.Id, params.CodBarra)

	if err != nil {
		log.Println("---------------pimer if---------------")

		log.Println(err)

		if err == service.ErrCodigoAlreadyExists {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Codigo de Barra Existe en la Base de Datos"})
		}
		if err == service.ErrCodigoNoExists {
			return c.JSON(http.StatusForbidden, responseMessage{Message: "Codigo de Producto  NO Existe en la Base de Datos"})
		}

		log.Println("-------------aquir if---------------")

		//return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: err.Error()})
	}

	// visualizar los header
	/*
		for key, values := range c.Request().Header {
			log.Println("*ooo***************key**********")
			log.Println(key)
			log.Println("****oo**********************")
			for _, value := range values {

				log.Println("--ii--------valor--------")
				log.Println(value)
				log.Println("----ii------valor--------")
			}
		}
	*/
	//response.JSON(w, r, http.StatusOK, response.Map{"token": token})

	type respuesta struct {
		Respuesta bool
		Mensaje   string `json:"mensaje"`
		Registro  int
	}

	r := respuesta{Respuesta: true, Mensaje: "Se grabo Sastifactoriamnet el codigo de barra", Registro: 1}
	//	r2 := respuesta{Respuesta: true}
	return c.JSON(http.StatusOK, respuesta(r))

}
