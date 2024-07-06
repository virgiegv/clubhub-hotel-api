package repository

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository/models"
	"fmt"
)

func GetOrCreateCity(cityName string, country string) (models.City, error) {
	db := models.Init().DB
	//try to get city with same values
	var cityModel models.City
	db.Limit(1).Where(&models.City{
		Name:    cityName,
		Country: country,
	}).Find(&cityModel)

	if cityModel.Id > 0 {
		return cityModel, nil
	}

	//if it doesnt exists, create it

	cityModel = models.City{
		Name:    cityName,
		Country: country,
	}

	cityModel, err := insertCity(cityModel)
	if err != nil {
		return models.City{}, fmt.Errorf("did not find existing city, could not create it either: %s", err.Error())
	}

	return cityModel, nil

}

func CreateLocation(location dto.LocationDTO, city models.City) (models.Location, error) {

	locationModel := models.Location{
		CityId:  city.Id,
		Address: location.Address,
		ZipCode: location.ZipCode,
	}

	locationModel, err := insertLocation(locationModel)
	if err != nil {
		return models.Location{}, err
	}

	locationModel.City = city
	return locationModel, nil
}

func insertCity(city models.City) (models.City, error) {
	db := models.Init().DB
	result := db.Create(&city)
	if result.Error != nil {
		return models.City{}, fmt.Errorf("could not insert city: %s", result.Error)
	}
	return city, result.Error
}

func insertLocation(location models.Location) (models.Location, error) {
	db := models.Init().DB
	result := db.Create(&location)
	if result.Error != nil {
		return models.Location{}, fmt.Errorf("could not insert location: %s", result.Error)
	}
	return location, result.Error
}

func GetCityById(cityId int64) (models.City, error) {
	var city models.City
	db := models.Init().DB

	db.First(&city, cityId)
	if city.Id <= 0 {
		return models.City{}, fmt.Errorf("could not find city with id %d", cityId)
	}

	return city, nil
}

func GetLocationById(locationId int64) (models.Location, error) {
	var location models.Location
	db := models.Init().DB

	db.First(&location, locationId)
	if location.Id <= 0 {
		return models.Location{}, fmt.Errorf("could not find location with id %d", locationId)
	}

	city, err := GetCityById(location.CityId)
	if err != nil {
		return location, fmt.Errorf("could not find a city associated with location %d: %s", locationId, err.Error())
	}

	location.City = city
	return location, nil
}
