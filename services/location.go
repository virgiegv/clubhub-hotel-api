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

func ReviewLocationAndUpdate(oldLocation models.Location, newLocation dto.LocationDTO) (models.Location, bool, error) {
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
		updateModel.CityId = oldLocation.CityId
		location, err := repository.UpdateLocation(oldLocation.Id, updateModel)
		if err != nil {
			return location, true, err
		}
		return location, true, nil
	}

	return models.Location{}, false, nil

}
