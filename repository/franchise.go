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

	var franchises []models.Franchise
	db.Where(&models.Franchise{
		CompanyId: companyId,
	}).Find(&franchises)

	if len(franchises) <= 0 {
		return franchises, fmt.Errorf("could not find franchises for company %d", companyId)
	}

	return franchises, nil
}

func UpdateFranchiseById(updateModel models.Franchise, franchiseId int64) error {
	db := models.Init().DB

	updateModel.Id = franchiseId
	result := db.Save(&updateModel)

	if result.Error != nil {
		return fmt.Errorf("Error updating franchise: %v\n", result.Error)
	}

	return nil
}
