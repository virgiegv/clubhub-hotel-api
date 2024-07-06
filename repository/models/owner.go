package models

import (
	"fmt"
)

type Owner struct {
	Id         int64 `json:"id" gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	LocationId int64
	Location   Location `gorm:"foreignKey:LocationId"`
}

func (o Owner) String() string {
	return fmt.Sprintf("Owner<%d %s %s %s %s %s>", o.Id, o.FirstName, o.LastName, o.Email, o.Phone, o.Location)
}
