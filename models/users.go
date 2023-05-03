package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int    `json:"id" form:"id"`
	Name             string `json:"name" form:"name" binding:"required"`
	Occupation       string `json:"occupation" form:"occupation" binding:"required"`	
	Email            string `json:"email" form:"email" binding:"required, email"`
	Password         string `json:"password" form:"password" binding:"required"`
	Avatar_File_Name string `json:"avatar_file_name" form:"avatar_file_name"`
	Role             string `json:"role" form:"role"`
	Transaction []*Transaction
	Campaign []*Campaign
}
