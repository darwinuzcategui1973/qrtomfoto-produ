package entity

type Produtosd3xd struct {
	//ID            int64   `db:"pcode"`
	ID            string  `db:"pcode"`
	Preferencia   string  `db:"preferencia"`
	Pdescribe     string  `db:"pdescribe"`
	Pdepartamento string  `db:"pdepartamento"`
	Pventa1       float32 `db:"pventa1"`
	Pventa2       float32 `db:"pventa2"`
	Pventa3       float32 `db:"pventa3"`
	Pventa4       float32 `db:"pventa4"`
	Pventa5       float32 `db:"pventa5"`
	Pempaque      string  `db:"pempaque"`
	Pmedida       string  `db:"pmedida"`
	Pmedempaque   string  `db:"pmedempaque"`
	Pexiste       int64   `db:"pexiste"`
	Pintercode    string  `db:"pintercode"`
	Pdatevoid     string  `db:"Pdatevoid"`
	Pmoneydif     int64   `db:"pmoneydif"`
	Ptasac        float32 `db:"ptasac"`
}
