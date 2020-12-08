package controllers

import (
	"resiliency-matrix/app/helpers"
	"resiliency-matrix/app/models"

	"github.com/revel/revel"
)

type Products struct {
	*revel.Controller
}

var (
	product       models.Product
	productErrors helpers.Error
)

func (c App) parseProductBody() error {
	err := c.Params.BindJSON(&product)
	return err
}

func (c App) AddProduct() revel.Result {
	err := c.parseProductBody()
	if err != nil {
		c.Response.Status = 400
		productErrors.Error = "Invalid JSON passed"
		return c.RenderJSON(productErrors)
	}
	product.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		productErrors.Error = productErrors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(productErrors)
	}
	response := product.AddProduct()
	if response != nil {
		c.Response.Status = 503
		productErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(productErrors)
	}
	c.Response.Status = 201
	return nil
}

func (c App) GetProduct(id int64) revel.Result {
	data, response := product.GetProduct(id)
	if response != nil {
		c.Response.Status = 400
		productErrors.Error = "Product not found"
		return c.RenderJSON(productErrors)
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) UpdateProduct(id int64) revel.Result {
	err := c.parseProductBody()
	if err != nil {
		c.Response.Status = 400
		productErrors.Error = "Invalid JSON passed"
		return c.RenderJSON(productErrors)
	}
	product.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		productErrors.Error = productErrors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(productErrors)
	}
	response := product.UpdateProduct(id)
	if response != nil {
		c.Response.Status = 503
		productErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(productErrors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) DeleteProduct(id int64) revel.Result {
	response := product.DeleteProduct(id)
	if response != nil {
		c.Response.Status = 503
		productErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(productErrors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) ListProducts(page int64, perPage int64) revel.Result {
	data, response := product.ListProducts()
	if response != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}
