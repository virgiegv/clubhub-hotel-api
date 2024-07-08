package dto

type CreateFranchiseDTO struct {
	CompanyId int64       `json:"company_id" validate:"required"`
	Url       string      `json:"url" validate:"required"`
	Name      string      `json:"name"`
	Location  LocationDTO `json:"location"`
}
