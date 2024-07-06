package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
)

func CreateLocation(location dto.LocationDTO) (models.Location, error) {
	city, err := repository.GetOrCreateCity(location.City, location.Country)
	if err != nil {
		return models.Location{}, err
	}

	//I have either the existing city or new one, now lets create its location
	createdLocation, err := repository.CreateLocation(location, city)
	if err != nil {
		return models.Location{}, err
	}

	return createdLocation, nil
}
