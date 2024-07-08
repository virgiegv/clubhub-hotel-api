package dto

type UpdateCompanyDTO struct {
	Name      string      `json:"name"`
	TaxNumber string      `json:"tax_number"`
	Owner     OwnerDTO    `json:"owner"`
	Location  LocationDTO `json:"location"`
}
