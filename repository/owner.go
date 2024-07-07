package repository

import (
	"clubhub-hotel-api/repository/models"
	"fmt"
)

func InsertOwner(owner models.Owner) (models.Owner, error) {
	db := models.Init().DB
	result := db.Create(&owner)
	if result.Error != nil {
		return models.Owner{}, fmt.Errorf("could not insert owner: %s", result.Error)
	}
	return owner, result.Error
}

func GetOwnerById(ownerId int64) (models.Owner, error) {
	db := models.Init().DB

	var ownerModel models.Owner
	result := db.First(&ownerModel, ownerId)

	if result.Error != nil {
		return ownerModel, fmt.Errorf("could not find owner with id %d", ownerId)
	}

	//Get company location
	location, err := GetLocationById(ownerModel.LocationId)
	if err != nil {
		return ownerModel, fmt.Errorf("could not find location with id %d", ownerModel.LocationId)
	}

	ownerModel.Location = location

	return ownerModel, nil
}

func UpdateOwner(ownerId int64, updateModel models.Owner) (models.Owner, error) {

	updateModel.Id = ownerId

	db := models.Init().DB
	result := db.Save(&updateModel)

	if result.Error != nil {
		return models.Owner{}, fmt.Errorf("Error updating owner: %v\n", result.Error)
	}

	return updateModel, nil
}
