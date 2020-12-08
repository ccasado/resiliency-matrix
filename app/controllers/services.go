package controllers

import (
	"resiliency-matrix/app/helpers"
	"resiliency-matrix/app/models"

	"github.com/revel/revel"
)

type Services struct {
	*revel.Controller
}

var (
	service       models.Service
	serviceErrors helpers.Error
)

func (c App) parseServiceBody() error {
	err := c.Params.BindJSON(&service)
	return err
}

func (c App) AddService() revel.Result {
	err := c.parseServiceBody()
	if err != nil {
		c.Response.Status = 400
		serviceErrors.Error = "Invalid JSON passed"
		return c.RenderJSON(serviceErrors)
	}
	service.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		serviceErrors.Error = serviceErrors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(serviceErrors)
	}
	response := service.AddService()
	if response != nil {
		c.Response.Status = 503
		serviceErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(serviceErrors)
	}
	c.Response.Status = 201
	return nil
}

func (c App) GetService(id int64) revel.Result {
	data, response := service.GetService(id)
	if response != nil {
		c.Response.Status = 400
		serviceErrors.Error = "Service not found"
		return c.RenderJSON(serviceErrors)
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) UpdateService(id int64) revel.Result {
	err := c.parseServiceBody()
	if err != nil {
		c.Response.Status = 400
		serviceErrors.Error = "Invalid JSON passed"
		return c.RenderJSON(serviceErrors)
	}
	service.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		serviceErrors.Error = serviceErrors.FormatError(c.Validation.ErrorMap())
		return c.RenderJSON(serviceErrors)
	}
	response := service.UpdateService(id)
	if response != nil {
		c.Response.Status = 503
		serviceErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(serviceErrors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) DeleteService(id int64) revel.Result {
	response := service.DeleteService(id)
	if response != nil {
		c.Response.Status = 503
		serviceErrors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(serviceErrors)
	}
	c.Response.Status = 200
	return nil
}

func (c App) ListServices(page int64, perPage int64) revel.Result {
	data, response := service.ListServices()
	if response != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}
