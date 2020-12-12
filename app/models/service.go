package models

import (
	"github.com/revel/revel"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

func (service Service) Validate(v *revel.Validation) {
	v.Check(service.Name, revel.ValidRequired(), revel.ValidMinSize(3), revel.ValidMaxSize(100)).Message("Name must be between 3-100 characters long")
}

func (service Service) AddService() error {
	response := DB.Create(&service)
	return response.Error

}

func (service Service) GetService(id int64) (Service, error) {
	response := DB.First(&service, id)
	return service, response.Error
}

func (service Service) UpdateService(id int64) error {
	service.ID = uint(id)
	response := DB.Save(&service)
	return response.Error
}

func (service Service) DeleteService(id int64) error {
	service.ID = uint(id)
	response := DB.Delete(&service)
	return response.Error
}

func (service Service) ListServices() ([]Service, error) {
	services := make([]Service, 0, 0)
	response := DB.Order("id desc").Find(&services)
	return services, response.Error
}
