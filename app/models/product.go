package models

import (
	"github.com/revel/revel"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)" json:"name"`
}

func (product Product) Validate(v *revel.Validation) {
	v.Check(product.Name, revel.ValidRequired(), revel.ValidMinSize(2), revel.ValidMaxSize(100)).Message("Name must be between 2-100 characters long")
}

func (product Product) AddProduct() error {
	response := DB.Create(&product)
	return response.Error
}

func (product Product) GetProduct(id int64) (Product, error) {
	response := DB.First(&product, id)
	return product, response.Error
}

func (product Product) UpdateProduct(id int64) error {
	product.ID = uint(id)
	response := DB.Save(&product)
	return response.Error
}

func (product Product) DeleteProduct(id int64) error {
	product.ID = uint(id)
	response := DB.Delete(&product)
	return response.Error
}

func (product Product) ListProducts() ([]Product, error) {
	products := make([]Product, 0, 0)
	response := DB.Order("id desc").Find(&products)
	return products, response.Error
}
