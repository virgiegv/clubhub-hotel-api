package models

import (
	"fmt"
)

type City struct {
	Id      int64  `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"uniqueIndex:idx_name_country"`
	Country string `json:"country" gorm:"uniqueIndex:idx_name_country"`
}

func (c City) String() string {
	return fmt.Sprintf("City<%d %s %s>", c.Id, c.Name, c.Country)
}
