package controllers

import (
	"clubhub-hotel-api/dto"
	"clubhub-hotel-api/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateFranchise godoc
// @Tags Franchise
// @Summary Franchise Creation
// @Description Create a franchise using its name, url, associated company_id, and location
// @Produce json
// @Param franchiseInfo body dto.CreateFranchiseDTO true "body"
// @Success 201 {object} models.Franchise
// @Router /clubhub/api/v1/franchise [post]
func CreateFranchise(c echo.Context) error {

	franchiseInfo := dto.CreateFranchiseDTO{}
	if err := c.Bind(&franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := services.CreateFranchise(franchiseInfo, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// GetFrenchiseById godoc
// @Tags Franchise
// @Summary Get franchise by id
// @Description Find a franchise by its id
// @Produce json
// @Param franchise_id path string true "Id of the franchise"
// @Success 200 {object} models.Franchise
// @Router /clubhub/api/v1/franchise/{franchise_id} [get]
func GetFranchiseById(c echo.Context) error {
	franchise_id := c.Param("franchise_id")

	message := "Buscando por franchise_id: "
	if franchise_id != "" {
		message = message + franchise_id
	}

	franchise, err := services.GetFranchiseById(franchise_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, franchise)
}

// GetFranchiseByFilters godoc
// @Tags Franchise
// @Summary Franchise search by filters
// @Description Find a franchise by name, url or associated company_id
// @Produce json
// @Param name query string false "Name of the franchise"
// @Param url query string false "url of the franchise"
// @Param company_id query string false "id of the company that owns the franchise"
// @Success 200 {object} []models.Franchise
// @Router /clubhub/api/v1/franchise/ [get]
func GetFranchiseByFilters(c echo.Context) error {
	name := c.QueryParam("name")
	url := c.QueryParam("url")
	company_id := c.QueryParam("company_id")

	companies, err := services.GetFranchisesByFilters(name, url, company_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if len(companies) <= 0 {
		return c.JSON(http.StatusNoContent, companies)
	}

	return c.JSON(http.StatusOK, companies)
}

// UpdateFranchiseWebSiteData godoc
// @Tags Franchise
// @Summary Update franchise website data automatically
// @Description Given a franchise id, runs its website data analysis again to update automatically
// @Produce json
// @Param franchise_id path string true "Id of the franchise"
// @Success 200 {object} models.Franchise
// @Router /clubhub/api/v1/franchise/{franchise_id} [patch]
func UpdateFranchiseWebSiteDataAutomatically(c echo.Context) error {
	franchise_id := c.Param("franchise_id")

	franchise, err := services.UpdateFranchiseWebSiteAutomatically(franchise_id, c)

	if err != nil {
		if franchise.Id == 0 {
			return c.JSON(http.StatusNotFound, fmt.Errorf("could not find franchise: %s", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("could not update franchise: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, franchise)
}

// UpdateFranchise godoc
// @Tags Franchise
// @Summary Update franchise data
// @Description Update a given franchise's main data and location data. This will not update its website data.
// @Produce json
// @Param franchise_id path string true "Id of the franchise"
// @Param franchiseInfo body dto.UpdateFranchiseDTO true "body"
// @Success 200 {object} models.Franchise
// @Router /clubhub/api/v1/franchise/{franchise_id} [put]
func UpdateFranchise(c echo.Context) error {
	franchise_id := c.Param("franchise_id")

	franchiseInfo := dto.UpdateFranchiseDTO{}
	if err := c.Bind(&franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(franchiseInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	franchise, err := services.UpdateFranchiseManually(franchiseInfo, franchise_id)
	if err != nil {
		if franchise.Id == 0 {
			return c.JSON(http.StatusNotFound, fmt.Errorf("could not find franchise: %s", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("could not update franchise: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, franchise)
}
