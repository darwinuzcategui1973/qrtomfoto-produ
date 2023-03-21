package models

type Empresa struct {
	ID           int64  `json:"id"`
	Rif string `json:"rif"`
	Status       string `json:"status"`
	Nombre       string `json:"nombre"`
	Conformacion string `json:"conformacion"`
	FechaInsta   string `json:"fechaInstalacion"`
	FechaVenci   string `json:"fechaVencimiento"`
	Sistema      string `json:"sistema"`
	Usuario      int64  `json:"usuario"`
	Viewers      int64  `json:"viewers"`
	MacAddr      string `json:"macddress"`
	FechaIni     string `json:"fechaInicioSistema"`
	Version      string `json:"version"`
}
