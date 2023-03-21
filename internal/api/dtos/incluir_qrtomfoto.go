package dtos

type IncluirQrtomfoto struct {
	Rif       string `json:"rif" validate:"required"`
	Ubicacion string `json:"ubicacion" validate:"required,min=8"`
	Nombre    string `json:"nombre" validate:"required"`
	Archivo   []int64 `json:"Archivo" validate:"required"`
}

/*

CREATE TABLE IF NOT EXISTS QRTOMFOTO (
    id int not null auto_increment,
    rif varchar(255) not null,
    ubicacion  varchar (255) not null DEFAULT "noHayUbicacion",
    nombre varchar(255) not null,
    archivo blob not null,
    fechaInicioEnvio TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
);


*/