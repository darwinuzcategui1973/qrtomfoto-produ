package entity

type Empresa struct {
	ID           int64  `db:"id"`
	Rif          string `db:"rif"`
	Status       string `db:"status"`
	Nombre       string `db:"nombre"`
	Conformacion string `db:"conformacion"`
	FechaInsta   string `db:"fechaInstalacion"`
	FechaVenci   string `db:"fechaVencimiento"`
	Sistema      string `db:"sistema"`
	Usuario      int64  `db:"usuario"`
	Viewers      int64  `db:"viewers"`
	MacAddr      string `db:"macddress"`
	FechaIni     string `db:"fechaInicioSistema"`
	Version      string `db:"version"`
}
