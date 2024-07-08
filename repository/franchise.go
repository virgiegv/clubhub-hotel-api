package repository

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository/models"
	"fmt"
)

func insertFranchise(franchise models.Franchise) (models.Franchise, error) {
	db := models.Init().DB
	result := db.Create(&franchise)
	if result.Error != nil {
		return models.Franchise{}, fmt.Errorf("could not insert franchise: %s", result.Error)
	}
	return franchise, result.Error
}

func CreateFranchise(
	website models.FranchiseWebSite,
	location models.Location,
	franchiseInfo dto.CreateFranchiseDTO,
) (models.Franchise, error) {
	franchise, err := insertFranchise(models.Franchise{
		CompanyId:     franchiseInfo.CompanyId,
		Name:          franchiseInfo.Name,
		Url:           franchiseInfo.Url,
		WebsideDataId: website.Id,
		LocationId:    location.Id,
	})

	franchise.WebsiteData = website
	franchise.Location = location

	return franchise, err
}

func GetFranchiseById(franchiseId int64) (models.Franchise, error) {
	db := models.Init().DB

	var franchise models.Franchise
	db.First(&franchise, franchiseId)
	if franchise.Id <= 0 {
		return models.Franchise{}, fmt.Errorf("could not find franchise with id %d", franchiseId)
	}

	website, _ := GetWebSiteDataById(franchise.WebsideDataId)

	location, _ := GetLocationById(franchise.LocationId)

	franchise.WebsiteData = website
	franchise.Location = location

	return franchise, nil

}

func GetFranchisesByCompanyId(companyId int64) ([]models.Franchise, error) {
	db := models.Init().DB

	var ids []int64
	err := db.Model(&models.Franchise{}).
		Where("company_id = ?", companyId).
		Pluck("id", &ids).Error

	if err != nil {
		return []models.Franchise{}, fmt.Errorf("could not find franchises for company %d", companyId)
	}

	var result []models.Franchise
	for _, f := range ids {
		newFranchise, _ := GetFranchiseById(f)
		result = append(result, newFranchise)
	}

	return result, nil
}

func GetFranchisesByFilters(filters dto.FranchiseSearchFilters) ([]models.Franchise, error) {
	db := models.Init().DB

	query := db.Model(&models.Franchise{})

	if filters.CompanyId != 0 {
		query = query.Or("company_id = ?", filters.CompanyId)
	}

	if filters.FranchiseName != "" {
		query = query.Or("name ILIKE ?", "%"+filters.FranchiseName+"%")
	}

	if filters.Url != "" {
		query = query.Or("url ILIKE ?", "%"+filters.Url+"%")
	}

	var resultIds []int64

	err := query.Pluck("id", &resultIds).Error
	if err != nil {
		return []models.Franchise{}, fmt.Errorf("could not retrieve franchises: %s", err.Error())
	}

	var result []models.Franchise
	for _, f := range resultIds {
		newFranchise, _ := GetFranchiseById(f)
		result = append(result, newFranchise)
	}

	return result, nil
}

func UpdateFranchiseById(updateModel models.Franchise, franchiseId int64) (models.Franchise, error) {
	db := models.Init().DB

	updateModel.Id = franchiseId
	result := db.Save(&updateModel)

	if result.Error != nil {
		return models.Franchise{}, fmt.Errorf("Error updating franchise: %v\n", result.Error)
	}

	return updateModel, nil
}
