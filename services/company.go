package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
	"fmt"
	"strconv"
	"strings"
)

func CreateCompanyService(company dto.CreateCompanyDTO) (models.Company, error) {

	companyLocation, err := CreateLocation(company.Location)
	if err != nil {
		return models.Company{}, fmt.Errorf("error with company location: %s", err.Error())
	}

	owner, err := createOwner(company.Owner)
	if err != nil {
		return models.Company{}, fmt.Errorf("error with company owner: %s", err.Error())
	}

	companyModel, err := createCompany(company, companyLocation, owner)
	if err != nil {
		return models.Company{}, err
	}

	companyModel.Location = companyLocation
	companyModel.Owner = owner

	return companyModel, nil
}

func createCompany(company dto.CreateCompanyDTO, location models.Location, owner models.Owner) (models.Company, error) {
	companyModel := models.Company{
		Name:       company.Name,
		TaxNumber:  company.TaxNumber,
		LocationId: location.Id,
		OwnerId:    owner.Id,
	}

	companyModel, err := repository.InsertCompany(companyModel)
	if err != nil {
		return models.Company{}, err
	}

	return companyModel, nil
}

func createOwner(owner dto.OwnerDTO) (models.Owner, error) {
	//Create the owner's location
	ownerLocation, err := CreateLocation(owner.Location)
	if err != nil {
		return models.Owner{}, fmt.Errorf("error with owner location: %s", err.Error())
	}

	ownerModel := models.Owner{
		FirstName:  owner.FirstName,
		LastName:   owner.LastName,
		Email:      owner.Email,
		Phone:      owner.Phone,
		LocationId: ownerLocation.Id,
	}
	ownerModel, err = repository.InsertOwner(ownerModel)
	if err != nil {
		return models.Owner{}, err
	}

	ownerModel.Location = ownerLocation

	return ownerModel, nil
}

// TO DO: Search result models are only bringing the id of associations, not the whole struct!!!!
func GetCompanyById(companyId string) (models.Company, error) {
	companyId = strings.TrimSpace(companyId)
	numericId, err := strconv.ParseInt(companyId, 10, 64)
	if err != nil {
		return models.Company{}, fmt.Errorf("invalid company_id")
	}

	return repository.GetCompanyById(numericId)
}

func GetCompaniesByFilters(name, taxNumber string) ([]models.Company, error) {
	filters := dto.CompanySearchFilters{
		Name:      name,
		TaxNumber: taxNumber,
	}

	return repository.GetCompanyByFilters(filters)
}

func UpdateCompany(companyId string, company dto.CreateCompanyDTO) (models.Company, error) {
	companyId = strings.TrimSpace(companyId)
	_, err := strconv.ParseInt(companyId, 10, 64)
	if err != nil {
		return models.Company{}, fmt.Errorf("invalid company_id")
	}

	//TO DO: implement
	return models.Company{}, nil
}
