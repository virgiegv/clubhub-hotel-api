package services

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/repository"
	"clubhub-hotel-api/repository/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"net/http"
	urlPackage "net/url"
	"time"
)

func CreateFranchiseWebSiteData(url string, context echo.Context) (models.FranchiseWebSite, error) {
	var errors []error

	franchiseWeb, err := repository.CreateEmptyWebsiteData()
	if err != nil {
		return models.FranchiseWebSite{}, fmt.Errorf("could not create website data for franchise: %s", err.Error())
	}

	go ProduceWebsiteAnalysis(franchiseWeb.Id, url, context)

	whoIsModel, err := getFromWhoIs(url)
	if err != nil {
		errors = append(errors, err)
	}

	whoIsModel.LogoUrl = getLogoUrl(url)

	franchiseWeb, err = repository.UpdateWebsiteData(whoIsModel, franchiseWeb.Id)
	if err != nil {
		return models.FranchiseWebSite{}, fmt.Errorf(
			"could not update franchise website data with whoIs and logo information: %s", err.Error(),
		)
	}

	return franchiseWeb, nil
}

func getLogoUrl(url string) string {
	return fmt.Sprintf("https://logo.clearbit.com/%s", url)
}

func ProduceWebsiteAnalysis(websiteId int64, url string, context echo.Context) {
	//This function will be called inside a goroutine

	//Execute Analysis
	result, analysisErr := executeAnalysisWithRetries(url, 10, 5*time.Second)
	if analysisErr != nil {
		context.Logger().Error("could not perform analysis of franchise website")
		context.Logger().Errorf("error received: %s", analysisErr.Error())
	}

	err := repository.UpdateWebsiteDataWithAnalysis(result, websiteId)
	if err != nil {
		context.Logger().Error("could not perform initial website data update with analysis results")
		context.Logger().Errorf("error received: %s", err.Error())
	}

	context.Logger().Infof("Successfully performed analysis of website for id %s", websiteId)

}

func getURLAnalysis(url string) (dto.AnalysisResponseDTO, error) {

	analysisBody := dto.AnalysisResponseDTO{}
	resp, err := http.Get(url)

	if err != nil {
		//No se pudo ejecutar
		return dto.AnalysisResponseDTO{},
			fmt.Errorf("request received error <%s>", err.Error())
	} //Si retorno error, voy a intentar de nuevo llamar esta función

	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			return dto.AnalysisResponseDTO{},
				fmt.Errorf("request received code <%d>", resp.StatusCode)
		}
	}

	//Ahora, si hay respuesta legible:
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&analysisBody)
	if err != nil {
		return dto.AnalysisResponseDTO{},
			fmt.Errorf("could not decode response: %s", err.Error())
	}

	//Check status of the query process
	if analysisBody.Status == "READY" {
		return analysisBody, nil
	}

	return analysisBody, fmt.Errorf("Analysis not yet ready")
}

func executeAnalysisWithRetries(url string,
	retries int,
	delay time.Duration) (dto.AnalysisResponseDTO, error) {

	var latestError error

	baseURL := "https://api.ssllabs.com/api/v3/analyze"

	analysisURL, err := urlPackage.Parse(baseURL)
	if err != nil {
		return dto.AnalysisResponseDTO{}, err
	}

	q := analysisURL.Query()
	q.Set("host", url)
	analysisURL.RawQuery = q.Encode()

	for i := 0; i < retries; i++ {
		analysisResult, err := getURLAnalysis(analysisURL.String())
		if err == nil && analysisResult.Status == "READY" {
			return analysisResult, nil
		}

		if err != nil {
			latestError = err
		}

		time.Sleep(delay)
		delay = delay * 2
	}

	//We timed out on the analysis, so the WebsiteData object Latest Error will have to be filled
	errorMessage := fmt.Errorf("Exceeded retry attempts. Latest error: %s", latestError.Error())
	return dto.AnalysisResponseDTO{
		ErrorMessage: errorMessage.Error(),
	}, errorMessage

}

func getFromWhoIs(url string) (models.FranchiseWebSite, error) {

	raw, err := whois.Whois(url)
	if err != nil {
		return models.FranchiseWebSite{}, fmt.Errorf("could not get domain information")
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		return models.FranchiseWebSite{}, fmt.Errorf("could not parse domain information")
	}

	var registrant string

	if result.Registrant.Name == "" {
		registrant = result.Registrant.Organization
	} else {
		registrant = result.Registrant.Name
	}

	return models.FranchiseWebSite{
		WebsiteCreationDate:   result.Domain.CreatedDate,
		WebsiteExpirationDate: result.Domain.ExpirationDate,
		RegisteredTo:          registrant,
		DomainContactEmail:    result.Registrant.Email,
	}, nil

}

func UpdateFranchiseWebSiteData(websiteId int64, url string, context echo.Context) (models.FranchiseWebSite, error) {
	var errors []error

	go ProduceWebsiteAnalysis(websiteId, url, context)

	whoIsModel, err := getFromWhoIs(url)
	if err != nil {
		errors = append(errors, err)
	}

	whoIsModel.LogoUrl = getLogoUrl(url)

	franchiseWeb, err := repository.UpdateWebsiteData(whoIsModel, websiteId)
	if err != nil {
		return models.FranchiseWebSite{}, fmt.Errorf(
			"could not update franchise website data with whoIs and logo information: %s", err.Error(),
		)
	}

	return franchiseWeb, nil
}
