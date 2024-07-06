package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
	"fmt"
	"github.com/labstack/echo/v4"
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
