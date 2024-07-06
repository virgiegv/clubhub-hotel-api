package controllers

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/services"
	"clubhub-hotel-api/validations"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate = validations.GetValidator()

/*func CreateCompanyTEST(c echo.Context) error {
	//TEST with whois
	//result, err := whois.Whois("marriott.com")
	//if err == nil {
	//	fmt.Println(result)
	//}

	message := ""
	//id, err := services.CreateCompanyService()
	if id > 0 && err == nil {
		message = fmt.Sprint("Created a city with id %d", id)
	} else {
		message = fmt.Sprintf("could not create city! err: %s", err.Error())
	}

	return c.JSON(http.StatusCreated, message)
}*/

// CreateCompany godoc
// @Tags Company
// @Summary Company Creation
// @Description Create a company using its name, tax number, owner information, and location
// @Produce json
// @Param companyInfo body dto.CreateCompanyDTO true "body"
// @Success 201 {object} models.Company
// @Router /clubhub/api/v1/company [post]
func CreateCompany(c echo.Context) error {

	companyInfo := dto.CreateCompanyDTO{}
	if err := c.Bind(&companyInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(companyInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := services.CreateCompanyService(companyInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// GetCompanyByFilters godoc
// @Tags Company
// @Summary Company search by filters
// @Description Find a company by either its name or its tax number
// @Produce json
// @Param name query string false "Name of the company"
// @Param tax_number query string false "Tax number of the company"
// @Success 200 {object} []models.Company
// @Router /clubhub/api/v1/company/ [get]
func GetCompanyByFilters(c echo.Context) error {
	name := c.QueryParam("name")
	taxNumber := c.QueryParam("tax_number")
	var message = "Buscando por params: "

	if name != "" {
		message = message + fmt.Sprintf("{name: %s}", name)
	}

	if taxNumber != "" {
		message = message + fmt.Sprintf("{tax_number: %s}", taxNumber)
	}

	companies, err := services.GetCompaniesByFilters(name, taxNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, companies)
}

// GetCompanyById godoc
// @Tags Company
// @Summary Get company by id
// @Description Find a company by its id
// @Produce json
// @Param company_id path string true "Id of the company"
// @Success 200 {object} models.Company
// @Router /clubhub/api/v1/company/{company_id} [get]
func GetCompanyById(c echo.Context) error {
	companyId := c.Param("company_id")

	message := "Buscando por company_id: "
	if companyId != "" {
		message = message + companyId
	}

	company, err := services.GetCompanyById(companyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, company)
}

// UpdateCompany godoc
// @Tags Company
// @Summary Update company by id
// @Description Find a company by its id and update it
// @Produce json
// @Param company_id path string true "Id of the company"
// @Success 200 {object} string
// @Router /clubhub/api/v1/company/{company_id} [put]
func UpdateCompany(c echo.Context) error {
	companyId := c.Param("company_id")

	if companyId == "" {
		message := "Not found"
		return c.JSON(http.StatusNotFound, message)
	}
	message := "Updating by id:" + companyId
	return c.JSON(http.StatusOK, message)
}
