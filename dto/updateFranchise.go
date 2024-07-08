package dto

type UpdateFranchiseDTO struct {
	Company_id int64       `json:"company_id"`
	Name       string      `json:"name"`
	Url        string      `json:"url"`
	Location   LocationDTO `json:"location"`
}
