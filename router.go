package main

import (
	"clubhub-hotel-api/controllers"
	_ "clubhub-hotel-api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server *echo.Echo
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	r.server.GET("/", controllers.HealthCheck)

	var basePath = r.server.Group("/clubhub/api/v1")
	{
		basePath.GET("/swagger/*", echoSwagger.WrapHandler)

		var companyGroup = basePath.Group("/company")
		{
			companyGroup.POST("", controllers.CreateCompany)
			companyGroup.GET("/", controllers.GetCompanyByFilters)
			companyGroup.GET("/:company_id", controllers.GetCompanyById)
			companyGroup.PUT("/:company_id", controllers.UpdateCompany)

		}
	}

}
