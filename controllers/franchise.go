package controllers

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateCompany godoc
// @Tags Company
// @Summary Company Creation
// @Description Create a company using its name, tax number, owner information, and location
// @Produce json
// @Param companyInfo body dto.CreateFranchiseDTO true "body"
// @Success 201 {object} models.Franchise
// @Router /clubhub/api/v1/frqanchise [post]
func CreateFranchise(c echo.Context) error {

	franchiseInfo := dto.CreateFranchiseDTO{}
	if err := c.Bind(&franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := services.CreateFranchise(franchiseInfo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
