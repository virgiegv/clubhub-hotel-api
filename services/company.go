package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
	"fmt"
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

	numericId, err := getNumericIdFromString(companyId)
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

func UpdateCompany(companyId string, company dto.UpdateCompanyDTO) (models.Company, error) {
	numericId, err := getNumericIdFromString(companyId)
	if err != nil {
		return models.Company{}, fmt.Errorf("invalid company id")
	}

	//Get current Company
	currentCompany, err := repository.GetCompanyById(numericId)

	//Check if location should be changed
	location, hadUpdate, err := reviewLocationAndUpdate(currentCompany.Location, company.Location)
	if err != nil {
		return models.Company{}, fmt.Errorf("error updating location: %s", err.Error())
	}

	//Check if owner needs updating
	owner, hadUpdateOwner, err := reviewOwnerAndUpdate(currentCompany.Owner, company.Owner)

	//Now we check the main Company body for updates
	updatedCompany, err := repository.UpdateCompanyById(currentCompany.Id, company.Name, company.TaxNumber)
	if err != nil {
		return updatedCompany, err
	}

	if hadUpdate {
		updatedCompany.Location = location
	}
	if hadUpdateOwner {
		updatedCompany.Owner = owner
	}

	return updatedCompany, nil

}

func reviewOwnerAndUpdate(oldOwner models.Owner, newOwner dto.OwnerDTO) (models.Owner, bool, error) {
	updateModel := models.Owner{Id: oldOwner.Id}
	var shouldUpdate = false

	if (oldOwner.FirstName != newOwner.FirstName) && (newOwner.FirstName != "") {
		updateModel.FirstName = newOwner.FirstName
		shouldUpdate = true
	}

	if (oldOwner.LastName != newOwner.LastName) && (newOwner.LastName != "") {
		updateModel.LastName = newOwner.LastName
		shouldUpdate = true
	}

	if (oldOwner.Email != newOwner.Email) && (newOwner.Email != "") {
		updateModel.Email = newOwner.Email
		shouldUpdate = true
	}

	if (oldOwner.Phone != newOwner.Phone) && (newOwner.Phone != "") {
		updateModel.Phone = newOwner.Phone
		shouldUpdate = true
	}

	location, hadUpdate, err := reviewLocationAndUpdate(oldOwner.Location, newOwner.Location)
	if err != nil {
		return models.Owner{}, true, fmt.Errorf("error updating owner location: %s", err.Error())
	}
	if hadUpdate {
		updateModel.Location = location
	}

	if shouldUpdate {
		owner, err := repository.UpdateOwner(oldOwner.Id, updateModel)
		if err != nil {
			return owner, true, err
		}
		return owner, true, nil
	}

	return models.Owner{}, false, nil
}

func reviewLocationAndUpdate(oldLocation models.Location, newLocation dto.LocationDTO) (models.Location, bool, error) {
	updateModel := models.Location{Id: oldLocation.Id}
	var shouldUpdate = false

	if (oldLocation.Address != newLocation.Address) && (newLocation.Address != "") {
		updateModel.Address = newLocation.Address
		shouldUpdate = true
	}

	if (oldLocation.ZipCode != newLocation.ZipCode) && (newLocation.ZipCode != "") {
		updateModel.ZipCode = newLocation.ZipCode
		shouldUpdate = true
	}

	if ((newLocation.City != "") || (newLocation.Country != "")) &&
		((oldLocation.City.Name != newLocation.City) || (oldLocation.City.Country != newLocation.Country)) {

		updateModel.City = models.City{
			Name:    newLocation.City,
			Country: newLocation.Country,
		}
		shouldUpdate = true
	}

	if shouldUpdate {
		location, err := repository.UpdateLocation(oldLocation.Id, updateModel)
		if err != nil {
			return location, true, err
		}
		return location, true, nil
	}

	return models.Location{}, false, nil

}
