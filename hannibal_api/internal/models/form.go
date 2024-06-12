package models

import (
	"gorm.io/gorm"
	"mime/multipart"
)

// Form represents a model for storing form data, including JSON data and files.
type Form struct {
	gorm.Model
	Json string `gorm:"not null"`
	File *multipart.FileHeader
}
