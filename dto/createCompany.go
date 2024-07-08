package dto

type CreateCompanyDTO struct {
	Name      string      `json:"name" validate:"required"`
	TaxNumber string      `json:"tax_number" validate:"required"`
	Owner     OwnerDTO    `json:"owner" validate:"required"`
	Location  LocationDTO `json:"location" validate:"required"`
}

type LocationDTO struct {
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type OwnerDTO struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Location  LocationDTO `json:"location"`
}
