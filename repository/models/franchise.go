package models

import (
	"fmt"
	"time"
)

type Franchise struct {
	Id            int64            `json:"id" gorm:"primaryKey"`
	CompanyId     int64            `json:"company_id"`
	Name          string           `json:"name"`
	Url           string           `json:"url"`
	WebsideDataId int64            `json:"website_data_id"`
	WebsiteData   FranchiseWebSite `gorm:"foreignKey:WebsideDataId" json:"website_data"`
	LocationId    int64            `json:"location_id"`
	Location      Location         `gorm:"foreignKey:LocationId" json:"location"`
}

func (f Franchise) String() string {
	return fmt.Sprintf("Franchise<%d %s %s %s>", f.Id, f.Name, f.Url) //, f.Location)
}

type FranchiseWebSite struct {
	Id                    int64                  `json:"id" gorm:"primaryKey"`
	LogoUrl               string                 `json:"logo_url"`
	WebsiteCreationDate   string                 `json:"domain_created_at"`
	WebsiteExpirationDate string                 `json:"domain_expires_at"`
	RegisteredTo          string                 `json:"registered_to"`
	DomainContactEmail    string                 `json:"domain_contact_email"`
	Port                  int                    `json:"port"`
	Protocol              string                 `json:"protocol"`
	Endpoints             []FranchiseWebEndpoint `json:"endpoints" gorm:"foreignKey:WebsiteId"`
	LatestError           string                 `json:"latest_error"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

type FranchiseWebEndpoint struct {
	Id         int64  `json:"id"`
	WebsiteId  int64  `json:"website_id"`
	IpAddress  string `json:"ip_address"`
	ServerName string `json:"server_name"`
}
