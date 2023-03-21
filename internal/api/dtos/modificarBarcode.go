package dtos

type ModificarBarcode struct {
	Id       string `json:"id" validate:"required"`
	CodBarra string `json:"codBarra" validate:"required"`
}
