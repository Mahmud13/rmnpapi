package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"mnp_api/models"
	"mnp_api/database"
  "github.com/astaxie/beego/validation"
)

var db = database.DB
// Operations about object
type RetailerController struct {
	beego.Controller
}
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *RetailerController) GetAll() {
	retailers := [] models.Retailer{}
	db.Find(&retailers)
	c.Data["json"] = retailers
	c.ServeJSON()
}
// @Title Get
// @Description find retailer by retailer ID
// @Param	retailerID		path 	string	true		"the retailer id you want to get"
// @Success 200 {retailer} models.Retailer
// @Failure 403 :retailerID is empty
// @router /:retailerID [get]
func (r *RetailerController) Get() {
	retailerID := r.Ctx.Input.Param(":retailerID")
	if retailerID != "" {
		var retailer models.Retailer

		if db.First(&retailer, retailerID).RecordNotFound() {
			r.Data["json"] = "no record found"
		}else{
			r.Data["json"] = retailer
		}
	}
	r.ServeJSON()
}

// @Title Get
// @Description find retailer by retailer ID
// @Param	retailerID		path 	string	true		"the retailer id you want to get"
// @Success 200 {retailer} models.Retailer
// @Failure 403 :retailerID is empty
// @router /by-imei/:imei [get]
func (rc *RetailerController) GetByImei() {
	imei := rc.Ctx.Input.Param(":imei")
	if imei != "" {
		var retailer models.Retailer
		if db.Where(&models.Retailer{Imei: imei}).First(&retailer).RecordNotFound() {
			rc.Data["json"] = "no record found"
		}else{
			rc.Data["json"] = retailer
		}
	}
	rc.ServeJSON()
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *RetailerController) Post() {
	valid := validation.Validation{}
	var request models.Retailer
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
	var retailer models.Retailer

	if db.Where(&models.Retailer{Phone: request.Phone}).First(&retailer).RecordNotFound() {
		if err := db.Create(&request).Error; err != nil {
			panic(err)
		}else{
			c.Data["json"] = request
		}
	}else{
		db.Model(&retailer).Update(&request)
		c.Data["json"] = retailer
	}

	c.ServeJSON()
}
// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:retailerId [put]
func (c *RetailerController) Put() {
	retailerID := c.Ctx.Input.Param(":retailerID")
	if retailerID != "" {
		var request models.Retailer
		var retailer models.Retailer
		json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if db.First(&retailer, retailerID).RecordNotFound() {
			c.Data["json"] = "no record found"
		}else{
			db.Model(&retailer).Update(&request)
			c.Data["json"] = "record updated"
		}
	}
	c.ServeJSON()
}
// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:retailerId [delete]
func (c *RetailerController) Delete() {
	retailerId := c.Ctx.Input.Param(":retailerId")
	if retailerId != "" {
		var retailer models.Retailer
		if db.First(&retailer, retailerId).RecordNotFound() {
			c.Data["json"] = "no record found"
		}
		if err := db.Delete(&retailer).Error; err != nil {
			panic(err)
		} else {
			c.Data["json"] = "delete success!"
		}
	}
	c.ServeJSON()
}
// @Title Get
// @Description find customer by customer ID
// @Param	customerID		path 	string	true		"the customer id you want to get"
// @Success 200 {customer} models.Customer
// @Failure 403 :customerID is empty
// @router /pending-orders/:retailerId [get]
func (c *RetailerController) GetPending() {
  var retailerId uint
	c.Ctx.Input.Bind(&retailerId, ":retailerId")
  var order models.Order
	if retailerId != 0{
		db.Preload("Customer").Preload("Retailer").Where(&models.Order{RetailerID: retailerId, Status: "pending"}).First(&order)
	}
  c.Data["json"] = order
	c.ServeJSON()
}

// @Title Get
// @Description find retailer by retailer ID
// @Param	retailerID		path 	string	true		"the retailer id you want to get"
// @Success 200 {retailer} models.Retailer
// @Failure 403 :retailerID is empty
// @router /orders/:retailerId [get]
func (c *RetailerController) GetAllOrders() {
  var retailerId uint
	c.Ctx.Input.Bind(&retailerId, ":retailerId")
  orders := [] models.Order{}
	if retailerId != 0 {
		db.Preload("Retailer").Preload("Retailer").Where(&models.Order{RetailerID: retailerId}).Find(&orders)
	}
	c.Data["json"] = orders
	c.ServeJSON()
}
