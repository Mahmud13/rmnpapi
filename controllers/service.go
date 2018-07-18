package controllers

import (
	"github.com/astaxie/beego"

  "github.com/astaxie/beego/validation"
	"encoding/json"
	"mnp_api/models"
)

// Operations about object
type ServiceController struct {
	beego.Controller
}
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *ServiceController) GetAll() {
	services := [] models.Service{}
	db.Find(&services)
	c.Data["json"] = services
	c.ServeJSON()
}
// @Title Get
// @Description find service by service ID
// @Param	serviceID		path 	string	true		"the service id you want to get"
// @Success 200 {service} models.Service
// @Failure 403 :serviceID is empty
// @router /:serviceID [get]
func (r *ServiceController) Get() {
	serviceID := r.Ctx.Input.Param(":serviceID")
	if serviceID != "" {
		var service models.Service

		if db.First(&service, serviceID).RecordNotFound() {
			r.Data["json"] = "no record found"
		}else{
			r.Data["json"] = service
		}
	}
	r.ServeJSON()
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ServiceController) Post() {
	valid := validation.Validation{}
	var request models.Service
	json.Unmarshal(c.Ctx.Input.RequestBody, &request)

	b, err := valid.Valid(&request)
	if err != nil {
		panic(err)
	}
	if !b {
    c.Ctx.Output.SetStatus(422)
		c.Data["json"] =  valid.Errors
		c.ServeJSON()
		return
	}
	var service models.Service

	if db.Where(&models.Service{Name: request.Name}).First(&service).RecordNotFound() {
		if err := db.Create(&request).Error; err != nil {
			panic(err)
		}else{
			c.Data["json"] = request
		}
	}else{
		db.Model(&service).Update(&request)
		c.Data["json"] = service
	}

	c.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:serviceId [put]
func (c *ServiceController) Put() {
	serviceID := c.Ctx.Input.Param(":serviceId")
	if serviceID != "" {
		var request models.Service
		var service models.Service
		json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if db.First(&service, serviceID).RecordNotFound() {
			c.Data["json"] = struct { Success bool `json:"success"` } {false}
		}else{
			db.Model(&service).Update(&request)
			c.Data["json"] = struct { Success bool `json:"success"` } {true}
		}
	}
	c.ServeJSON()
}
// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:serviceId [delete]
func (c *ServiceController) Delete() {
	serviceId := c.Ctx.Input.Param(":serviceId")
	if serviceId != "" {
		var service models.Service
		if db.First(&service, serviceId).RecordNotFound() {
			c.Data["json"] = struct { Success bool `json:"success"` } {false}
		}
		if err := db.Delete(&service).Error; err != nil {
			panic(err)
		} else {
			c.Data["json"] = struct { Success bool `json:"success"` } {true}
		}
	}
	c.ServeJSON()
}
