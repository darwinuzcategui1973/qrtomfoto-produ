package models

type Produtosd3xd struct {
	//ID            int64   `json:"pcode"`
	ID            string  `json:"pcode"`
	Preferencia   string  `json:"preferencia"`
	Pdescribe     string  `json:"pdescribe"`
	Pdepartamento string  `json:"pdepartamento"`
	Pventa1       float64 `json:"pventa1"`
	Pventa2       float32 `json:"pventa2"`
	Pventa3       float32 `json:"pventa3"`
	Pventa4       float32 `json:"pventa4"`
	Pventa5       float32 `json:"pventa5"`
	Pempaque      string  `json:"pempaque"`
	Pmedida       string  `json:"pmedida"`
	Pmedempaque   string  `json:"pmedempaque"`
	Pexiste       int64   `json:"pexiste"`
	Pintercode    string  `json:"pintercode"`
	Pdatevoid     string  `json:"Pdatevoid"`
	Pmoneydif     int64   `json:"pmoneydif"`
	Ptasac        float32 `json:"ptasac"`
}
