package models

import (
	"fmt"
)

type Company struct {
	Id         int64 `json:"id" gorm:"primaryKey"`
	Name       string
	TaxNumber  string
	LocationId int64 `gorm:"foreignKey:CityId"`
	Location   Location
	OwnerId    int64 `gorm:"foreignKey:CityId"`
	Owner      Owner
	Franchises []Franchise
}

func (c Company) String() string {
	franchises := "["
	for k, v := range c.Franchises {
		if k == len(c.Franchises)-1 {
			franchises = franchises + fmt.Sprintf("%s]", v)
		} else {
			franchises = franchises + fmt.Sprintf("%s,", v)
		}
	}
	return fmt.Sprintf("Company{id: %d, name: %s, tax_number: %s, location: %s, owner: %s, franchises: %s}",
		c.Id, c.Name, c.TaxNumber, c.Location, c.Owner, franchises)
}
