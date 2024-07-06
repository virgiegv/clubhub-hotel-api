package models

import (
	"fmt"
	"time"
)

type Location struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	CityId    int64
	City      City `gorm:"foreignKey:CityId"`
	Address   string
	ZipCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l Location) String() string {
	return fmt.Sprintf("Location<%d %s %s %s>", l.Id, l.Address, l.ZipCode, l.City)
}
