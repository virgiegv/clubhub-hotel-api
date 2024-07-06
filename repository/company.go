package repository

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository/models"
	"fmt"
)

func InsertCompany(company models.Company) (models.Company, error) {
	db := models.Init().DB
	result := db.Create(&company)
	if result.Error != nil {
		return models.Company{}, fmt.Errorf("could not insert company: %s", result.Error)
	}
	return company, result.Error
}

func InsertOwner(owner models.Owner) (models.Owner, error) {
	db := models.Init().DB
	result := db.Create(&owner)
	if result.Error != nil {
		return models.Owner{}, fmt.Errorf("could not insert owner: %s", result.Error)
	}
	return owner, result.Error
}

func GetCompanyById(companyId int64) (models.Company, error) {
	db := models.Init().DB

	var companyModel models.Company
	result := db.First(&companyModel, companyId)

	if result.Error != nil {
		return companyModel, fmt.Errorf("could not find company with id %d", companyId)
	}

	return companyModel, nil
}

func GetCompanyByFilters(filters dto.CompanySearchFilters) ([]models.Company, error) {
	db := models.Init().DB

	var companyModels []models.Company

	db.Where(&models.Company{Name: filters.Name, TaxNumber: filters.TaxNumber}).Find(&companyModels)

	if len(companyModels) <= 0 {
		return companyModels, fmt.Errorf(
			"could not find companies with name %s or tax_number %s",
			filters.Name,
			filters.TaxNumber)
	}

	return companyModels, nil
}

func UpdateCompanyById(companyId int64, newCompany models.Company) (models.Company, error) {
	db := models.Init().DB

	var companyModel models.Company
	result := db.First(&companyModel, companyId)

	if result.Error != nil {
		return companyModel, fmt.Errorf("could not find company with id %d", companyId)
	}

	db.Model(&companyModel).Updates(models.Company{
		Name:       newCompany.Name,
		TaxNumber:  newCompany.TaxNumber,
		LocationId: newCompany.LocationId,
		OwnerId:    newCompany.OwnerId,
		//TO DO: do i need to also update structs for location and owner?
	})

	return companyModel, nil

}