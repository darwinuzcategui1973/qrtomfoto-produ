package dtos

type ConformarEmpresas struct {
	Rif          string `json:"rif" validate:"min=3,max=10,containsany=JjVvGgEePpCc"`
	Nombre       string `json:"nombre" validate:"required,min=2,max=100"`
	Conformacion string `json:"conformacion" validate:"required,min=8"`
	Usuario      int64  `json:"usuario" validate:"required"`
	Viewers      int64  `json:"viewers" validate:"required"`
}


/*
v : Natural de la República Bolivariana de Venezuela.
J : Persona Jurídica.
E : Extranjero con residencia en Venezuela.
P : Agente registrado con Pasaporte.
G : Ente Gubernamental.
c // consejo comunal
*/
