package entity

type Qrtomfoto struct {
	ID        int64  `db:"id"`
	Rif       string `db:"rif"`
	Ubicacion string `db:"ubicacion"`
	Nombre    string `db:"nombre"`
	Archivo    []int64 `db:"archivo"`
	FechaInicioEnvio string `db:"fechaInicioEnvio"`
}

/*
/*
Archivo   blob `db:"archivo"`

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