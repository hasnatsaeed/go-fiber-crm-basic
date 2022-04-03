package models

import (
	"github/hasnatsaeed/go-fiber-crm-basic/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func init() {
	configs.Connect()
	db = configs.GetDB()
	err := db.AutoMigrate(Lead{})
	if err != nil {
		panic(err)
	}
}
func (lead *Lead) Create() *Lead {
	db.Create(&lead)
	return lead
}

func GetLeads() []Lead {
	var leads []Lead
	db.Find(&leads)
	return leads
}

func GetLeadById(ID int) (*Lead, error) {
	var lead *Lead
	err := db.First(&lead, ID).Error
	if err != nil {
		return nil, err
	}
	return lead, nil
}

func DeleteLeadById(ID int) (*Lead, error) {
	var lead *Lead
	err := db.Delete(&lead, ID).Error
	if err != nil {
		return nil, err
	}
	return lead, nil
}
