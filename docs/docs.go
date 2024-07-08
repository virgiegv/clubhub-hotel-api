// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/virgiegv/",
            "email": "vgil_22@hotmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Check health of the service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Check if service is active",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Health"
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/company": {
            "post": {
                "description": "Create a company using its name, tax number, owner information, and location",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company Creation",
                "parameters": [
                    {
                        "description": "body",
                        "name": "companyInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCompanyDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Company"
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/company/": {
            "get": {
                "description": "Find a company by either its name or its tax number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Company search by filters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the company",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Tax number of the company",
                        "name": "tax_number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Company"
                            }
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/company/{company_id}": {
            "get": {
                "description": "Find a company by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get company by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the company",
                        "name": "company_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Company"
                        }
                    }
                }
            },
            "put": {
                "description": "Find a company by its id and update it",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Update company by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the company",
                        "name": "company_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "companyInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCompanyDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Company"
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/franchise": {
            "post": {
                "description": "Create a franchise using its name, url, associated company_id, and location",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Franchise"
                ],
                "summary": "Franchise Creation",
                "parameters": [
                    {
                        "description": "body",
                        "name": "franchiseInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFranchiseDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Franchise"
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/franchise/": {
            "get": {
                "description": "Find a franchise by name, url or associated company_id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Franchise"
                ],
                "summary": "Franchise search by filters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the franchise",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "url of the franchise",
                        "name": "url",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "id of the company that owns the franchise",
                        "name": "company_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Franchise"
                            }
                        }
                    }
                }
            }
        },
        "/clubhub/api/v1/franchise/{franchise_id}": {
            "get": {
                "description": "Find a franchise by its id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Franchise"
                ],
                "summary": "Get franchise by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the franchise",
                        "name": "franchise_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Franchise"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a given franchise's main data and location data. This will not update its website data.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Franchise"
                ],
                "summary": "Update franchise data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the franchise",
                        "name": "franchise_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "franchiseInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateFranchiseDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Franchise"
                        }
                    }
                }
            },
            "patch": {
                "description": "Given a franchise id, runs its website data analysis again to update automatically",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Franchise"
                ],
                "summary": "Update franchise website data automatically",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the franchise",
                        "name": "franchise_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Franchise"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Health": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateCompanyDTO": {
            "type": "object",
            "required": [
                "location",
                "name",
                "owner",
                "tax_number"
            ],
            "properties": {
                "location": {
                    "$ref": "#/definitions/dto.LocationDTO"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/dto.OwnerDTO"
                },
                "tax_number": {
                    "type": "string"
                }
            }
        },
        "dto.CreateFranchiseDTO": {
            "type": "object",
            "required": [
                "company_id",
                "url"
            ],
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/dto.LocationDTO"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.LocationDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "dto.OwnerDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/dto.LocationDTO"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateCompanyDTO": {
            "type": "object",
            "properties": {
                "location": {
                    "$ref": "#/definitions/dto.LocationDTO"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/dto.OwnerDTO"
                },
                "tax_number": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateFranchiseDTO": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/dto.LocationDTO"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.City": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Company": {
            "type": "object",
            "properties": {
                "franchises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Franchise"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/models.Location"
                },
                "locationId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.Owner"
                },
                "ownerId": {
                    "type": "integer"
                },
                "taxNumber": {
                    "type": "string"
                }
            }
        },
        "models.Franchise": {
            "type": "object",
            "properties": {
                "company_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/models.Location"
                },
                "location_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "website_data": {
                    "$ref": "#/definitions/models.FranchiseWebSite"
                },
                "website_data_id": {
                    "type": "integer"
                }
            }
        },
        "models.FranchiseWebEndpoint": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "ip_address": {
                    "type": "string"
                },
                "server_name": {
                    "type": "string"
                },
                "website_id": {
                    "type": "integer"
                }
            }
        },
        "models.FranchiseWebSite": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "domain_contact_email": {
                    "type": "string"
                },
                "domain_created_at": {
                    "type": "string"
                },
                "domain_expires_at": {
                    "type": "string"
                },
                "endpoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.FranchiseWebEndpoint"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "latest_error": {
                    "type": "string"
                },
                "logo_url": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                },
                "registered_to": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "$ref": "#/definitions/models.City"
                },
                "cityId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "models.Owner": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.Location"
                },
                "locationId": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Clubhub Hotel Franchise API",
	Description:      "This API manages the CRU operations of Clubhub's hotel franchises",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
