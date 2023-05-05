package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int    `json:"id" form:"id"`
	Name             string `json:"name" form:"name" binding:"required"`
	Occupation       string `json:"occupation" form:"occupation" binding:"required"`
	Email            string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
	Password         string `gorm:"not null" json:"password" form:"password" binding:"required"`
	Avatar_File_Name string `json:"avatar_file_name" form:"avatar_file_name"`
	Role             string `json:"role" form:"role"`
	Transaction      []*Transaction
	Campaign         []*Campaign
}
