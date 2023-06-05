package model

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ClientId     string         `json:"client_id" gorm:"primaryKey,size:20" form:"client_id"`
	ClientName   string         `json:"client_name" form:"client_name"`
	ClientSecret string         `json:"client_secret" form:"client_secret"`
	HomepageURL  string         `json:"homepage_url" form:"homepage_url"`
	RedirectURL  string         `json:"redirect_url" form:"redirect_url"`
	Description  string         `json:"description" form:"description"`
	CreatedAt    time.Time      `json:"createdAt" json:"createdAt" form:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt" json:"updatedAt" form:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index" json:"deletedAt" form:"deletedAt"`
}

func NewClient() *Client {
	return &Client{}
}
