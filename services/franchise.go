package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func CreateFranchise(franchiseInfo dto.CreateFranchiseDTO, context echo.Context) (models.Franchise, error) {

	franchiseWebsiteData, err := CreateFranchiseWebSiteData(franchiseInfo.Url, context)
	if err != nil {
		return models.Franchise{}, err
	}

	franchiseLocation, err := CreateLocation(franchiseInfo.Location)
	if err != nil {
		return models.Franchise{}, fmt.Errorf("could not create franchise location: %s", err.Error())
	}

	franchise, err := repository.CreateFranchise(franchiseWebsiteData, franchiseLocation, franchiseInfo)
	if err != nil {
		return models.Franchise{}, fmt.Errorf("could not create franchise: %s", err.Error())
	}

	return franchise, nil
}

func GetFranchiseById(franchise_id string) (models.Franchise, error) {
	numericId, err := getNumericIdFromString(franchise_id)
	if err != nil {
		return models.Franchise{}, fmt.Errorf("invalid franchise_id")
	}

	franchise, err := repository.GetFranchiseById(numericId)
	if err != nil {
		return models.Franchise{}, err
	}

	return franchise, nil
}

func GetFranchisesByFilters(name, url, company_id string) ([]models.Franchise, error) {

	numericId, err := getNumericIdFromString(company_id)
	if err != nil {
		numericId = 0
	}

	filters := dto.FranchiseSearchFilters{
		Url:           url,
		FranchiseName: name,
		CompanyId:     numericId,
	}

	return repository.GetFranchisesByFilters(filters)
}

func getNumericIdFromString(idString string) (int64, error) {
	idString = strings.TrimSpace(idString)
	numericId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return 0, err
	}
	return numericId, nil
}

func UpdateFranchiseWebSiteAutomatically(franchise_id string, context echo.Context) (models.Franchise, error) {
	numericId, err := getNumericIdFromString(franchise_id)
	if err != nil {
		return models.Franchise{}, err
	}

	currentFranchise, err := repository.GetFranchiseById(numericId)
	if err != nil {
		return models.Franchise{}, err
	}

	website, err := UpdateFranchiseWebSiteData(currentFranchise.WebsideDataId, currentFranchise.Url, context)
	currentFranchise.WebsiteData = website

	if err != nil {
		return currentFranchise, fmt.Errorf("could not update franchise website data: %s", err.Error())
	}

	return currentFranchise, nil
}

func UpdateFranchiseManually(updateFranchise dto.UpdateFranchiseDTO, franchise_id string) (models.Franchise, error) {
	numericId, err := getNumericIdFromString(franchise_id)
	if err != nil {
		return models.Franchise{}, err
	}

	//get currentFranchise
	currentFranchise, err := repository.GetFranchiseById(numericId)
	if err != nil {
		return models.Franchise{}, err
	}

	//check if location needs changes
	location, hasUpdates, err := ReviewLocationAndUpdate(currentFranchise.Location, updateFranchise.Location)
	if err != nil {
		return currentFranchise, err
	}

	if hasUpdates {
		currentFranchise.Location = location
	}

	updateModel := models.Franchise{
		Id:            numericId,
		WebsideDataId: currentFranchise.WebsideDataId,
		LocationId:    currentFranchise.LocationId,
	}

	if updateFranchise.Company_id != 0 {
		updateModel.CompanyId = updateFranchise.Company_id
	} else {
		updateModel.CompanyId = currentFranchise.CompanyId
	}

	if updateFranchise.Url != "" {
		updateModel.Url = updateFranchise.Url
	}

	if updateFranchise.Name != "" {
		updateModel.Name = updateFranchise.Name
	}

	newFranchise, err := repository.UpdateFranchiseById(updateModel, numericId)
	if err != nil {
		return newFranchise, fmt.Errorf("could not update franchise: %s", err.Error())
	}

	newFranchise.WebsiteData = currentFranchise.WebsiteData
	newFranchise.WebsideDataId = currentFranchise.WebsideDataId
	newFranchise.Location = currentFranchise.Location
	newFranchise.LocationId = currentFranchise.LocationId

	return newFranchise, nil
}
