package repository

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository/models"
	"fmt"
)

func insertEp(ep models.FranchiseWebEndpoint) (models.FranchiseWebEndpoint, error) {
	db := models.Init().DB
	result := db.Create(&ep)
	if result.Error != nil {
		return models.FranchiseWebEndpoint{}, fmt.Errorf("could not insert ep: %s", result.Error)
	}
	return ep, result.Error
}

func GetOrCreateEndpoint(epData dto.AnalysisEndpointsDTO, websiteId int64) (models.FranchiseWebEndpoint, error, bool) {
	db := models.Init().DB
	//try to get EP with same values
	var epModel models.FranchiseWebEndpoint
	db.Limit(1).Where(&models.FranchiseWebEndpoint{
		WebsiteId:  websiteId,
		IpAddress:  epData.IpAddress,
		ServerName: epData.ServerName,
	}).Find(&epModel)

	if epModel.Id > 0 {
		return epModel, nil, false
	}

	epModel = models.FranchiseWebEndpoint{
		WebsiteId:  websiteId,
		IpAddress:  epData.IpAddress,
		ServerName: epData.ServerName,
	}

	epModel, err := insertEp(epModel)
	if err != nil {
		return models.FranchiseWebEndpoint{}, fmt.Errorf(
			"did not find existing ep, could not create it either: %s", err.Error(),
		), false
	}

	return epModel, nil, true
}

func GetEndpointsByWebsiteId(websiteId int64) ([]models.FranchiseWebEndpoint, error) {
	db := models.Init().DB

	var epModels []models.FranchiseWebEndpoint

	db.Where(&models.FranchiseWebEndpoint{
		WebsiteId: websiteId,
	}).Find(&epModels)

	if len(epModels) > 0 {
		return epModels, nil
	}

	return epModels, fmt.Errorf("could not find EPs for website id %d", websiteId)

}

func UpdateEndpointByEndpointId(endpointID int64, ipAddress, serverName string) error {
	db := models.Init().DB

	ep := models.FranchiseWebEndpoint{
		Id:         endpointID,
		IpAddress:  ipAddress,
		ServerName: serverName,
	}
	result := db.Save(&ep)
	if result.Error != nil {
		return fmt.Errorf("Error updating endpoint: %v\n", result.Error)
	}

	return nil
}

func CreateEmptyWebsiteData() (models.FranchiseWebSite, error) {

	newWebSite := models.FranchiseWebSite{}
	db := models.Init().DB
	result := db.Create(&newWebSite)
	if result.Error != nil {
		return models.FranchiseWebSite{}, fmt.Errorf("could not insert website: %s", result.Error)
	}
	return newWebSite, result.Error
}

func UpdateWebsiteData(updateModel models.FranchiseWebSite, websiteId int64) (models.FranchiseWebSite, error) {
	db := models.Init().DB

	//UpdateModel will contain initialized values only where it should change
	updateModel.Id = websiteId
	result := db.Save(&updateModel)

	if result.Error != nil {
		return models.FranchiseWebSite{}, fmt.Errorf("Error updating website data: %s", result.Error.Error())
	}

	return updateModel, nil
}

func GetWebSiteDataById(websiteId int64) (models.FranchiseWebSite, error) {
	db := models.Init().DB

	var websiteData models.FranchiseWebSite
	db.First(&websiteData, websiteId)
	if websiteData.Id <= 0 {
		return models.FranchiseWebSite{}, fmt.Errorf("could not find website data for id %d", websiteId)
	}

	//Now get the endpoints
	endpoints, err := GetEndpointsByWebsiteId(websiteId)
	if err != nil {
		endpoints = []models.FranchiseWebEndpoint{}
	}

	websiteData.Endpoints = endpoints
	return websiteData, nil
}

func UpdateWebsiteDataWithAnalysis(responseDTO dto.AnalysisResponseDTO, websiteId int64) error {
	//Get the website data
	website, err := GetWebSiteDataById(websiteId)
	if err != nil {
		return err
	}

	var endpoints []models.FranchiseWebEndpoint
	//Create endpoints
	for _, ep := range responseDTO.Endpoints {
		epModel, createError, created := GetOrCreateEndpoint(ep, websiteId)
		if !created && createError != nil {
			return fmt.Errorf("could not create endpoint: %s", err)
		}
		endpoints = append(endpoints, epModel)
	}

	website.Endpoints = endpoints
	website.Port = responseDTO.Port
	website.Protocol = responseDTO.Protocol

	_, err = UpdateWebsiteData(website, websiteId)
	if err != nil {
		return fmt.Errorf("couldnt update website with initial analysis: %s", err.Error())
	}

	return nil
}
