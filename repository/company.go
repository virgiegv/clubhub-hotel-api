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

func GetCompanyById(companyId int64) (models.Company, error) {
	db := models.Init().DB

	var companyModel models.Company
	result := db.First(&companyModel, companyId)

	if result.Error != nil {
		return companyModel, fmt.Errorf("could not find company with id %d", companyId)
	}

	//Get company location
	location, err := GetLocationById(companyModel.LocationId)
	if err != nil {
		return companyModel, fmt.Errorf("could not find location with id %d", companyModel.LocationId)
	}

	//Get company owner
	owner, err := GetOwnerById(companyModel.OwnerId)
	if err != nil {
		return companyModel, fmt.Errorf("could not find company owner with id %d", companyModel.OwnerId)
	}

	companyModel.Location = location
	companyModel.Owner = owner

	franchises, err := GetFranchisesByCompanyId(companyId)
	if err != nil {
		franchises = []models.Franchise{}
	}
	companyModel.Franchises = franchises

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

func UpdateCompanyById(companyId int64, name, taxNumber string) (models.Company, error) {
	db := models.Init().DB

	var companyModel models.Company
	result := db.First(&companyModel, companyId)

	if result.Error != nil {
		return companyModel, fmt.Errorf("could not find company with id %d", companyId)
	}

	updateResult := db.Model(&companyModel).Updates(models.Company{
		Name:      name,
		TaxNumber: taxNumber,
	})
	if updateResult.Error != nil {
		return companyModel, fmt.Errorf("error updating company: %s", updateResult.Error.Error())
	}

	return companyModel, nil

}
