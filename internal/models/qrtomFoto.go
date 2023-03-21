package models

type QrtomFoto struct {
	ID    int64  `json:"id"`
	Rif  string `json:"rif"`
	Ubicacion  string `json:"ubicacion"`
	Nombre  string `json:"nombre"`
	Archivo  []int64 `json:"archivo"`
	FechaInicioEnvio  string `json:"fechaInicioEnvio"`
}


/*

Rif       string `db:"rif"`
	Ubicacion string `db:"ubicacion"`
	Nombre    string `db:"nombre"`
	Archivo   blob `db:"archivo"`
	FechaInicioEnvio string `db:"fechaInicioEnvio"`
*/