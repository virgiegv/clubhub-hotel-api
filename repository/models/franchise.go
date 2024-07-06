package models

import (
	"fmt"
	"time"
)

type Franchise struct {
	Id            int64 `json:"id" gorm:"primaryKey"`
	CompanyId     int64
	Name          string
	Url           string
	WebsideDataId int64
	WebsiteData   FranchiseWebSite `gorm:"foreignKey:WebsideDataId"`
	LocationId    int64
	Location      Location `gorm:"foreignKey:LocationId"`
}

func (f Franchise) String() string {
	return fmt.Sprintf("Franchise<%d %s %s %s>", f.Id, f.Name, f.Url) //, f.Location)
}

type FranchiseWebSite struct {
	Id                    int64                  `json:"id" gorm:"primaryKey"`
	LogoUrl               string                 `json:"logo_url"`
	WebsiteCreationDate   string                 `json:"web_created_at"`
	WebsiteExpirationDate string                 `json:"web_expires_at"`
	RegisteredTo          string                 `json:"registered_to"`
	DomainContactEmail    string                 `json:"domain_contact_email"`
	Port                  int                    `json:"port"`
	Protocol              string                 `json:"protocol"`
	Endpoints             []FranchiseWebEndpoint `json:"endpoints"`
	LatestError           string                 `json:"latest_error"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type FranchiseWebEndpoint struct {
	Id         int64  `json:"id"`
	WebsiteId  int64  `json:"website_id"`
	IpAddress  string `json:"ip_address"`
	ServerName string `json:"server_name"`
}
