package controllers

import (
	"github.com/astaxie/beego"

  "github.com/astaxie/beego/validation"
	"encoding/json"
	"mnp_api/models"
)

// Operations about object
type CustomerController struct {
	beego.Controller
}
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *CustomerController) GetAll() {
	customers := [] models.Customer{}
	db.Find(&customers)
	c.Data["json"] = customers
	c.ServeJSON()
}
// @Title Get
// @Description find customer by customer ID
// @Param	customerID		path 	string	true		"the customer id you want to get"
// @Success 200 {customer} models.Customer
// @Failure 403 :customerID is empty
// @router /:customerID [get]
func (r *CustomerController) Get() {
	customerID := r.Ctx.Input.Param(":customerID")
	if customerID != "" {
		var customer models.Customer

		if db.First(&customer, customerID).RecordNotFound() {
			r.Data["json"] = "no record found"
		}else{
			r.Data["json"] = customer
		}
	}
	r.ServeJSON()
}

// @Title Get
// @Description find customer by customer ID
// @Param	customerID		path 	string	true		"the customer id you want to get"
// @Success 200 {customer} models.Customer
// @Failure 403 :customerID is empty
// @router /by-imei/:imei [get]
func (rc *CustomerController) GetByImei() {
	imei := rc.Ctx.Input.Param(":imei")
	if imei != "" {
		var customer models.Customer
		if db.Where(&models.Customer{Imei: imei}).First(&customer).RecordNotFound() {
			rc.Data["json"] = "no record found"
		}else{
			rc.Data["json"] = customer
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
func (c *CustomerController) Post() {
	valid := validation.Validation{}
	var request models.Customer
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
	var customer models.Customer

	if db.Where(&models.Customer{Phone: request.Phone}).First(&customer).RecordNotFound() {
		if err := db.Create(&request).Error; err != nil {
			panic(err)
		}else{
			c.Data["json"] = request
		}
	}else{
		db.Model(&customer).Update(&request)
		c.Data["json"] = customer
	}

	c.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:customerId [put]
func (c *CustomerController) Put() {
	customerID := c.Ctx.Input.Param(":customerId")
	if customerID != "" {
		var request models.Customer
		var customer models.Customer
		json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if db.First(&customer, customerID).RecordNotFound() {
			c.Data["json"] = struct { Success bool `json:"success"` } {false}
		}else{
			db.Model(&customer).Update(&request)
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
// @router /:customerId [delete]
func (c *CustomerController) Delete() {
	customerId := c.Ctx.Input.Param(":customerId")
	if customerId != "" {
		var customer models.Customer
		if db.First(&customer, customerId).RecordNotFound() {
			c.Data["json"] = "no record found"
		}
		if err := db.Delete(&customer).Error; err != nil {
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
// @router /pending-orders/:customerId [get]
func (c *CustomerController) GetPendingOrders() {
  var customerId uint
	c.Ctx.Input.Bind(&customerId, ":customerId")
  var order models.Order
	if customerId != 0 {
		db.Preload("Customer").Preload("Retailer").Preload("Service").Where(&models.Order{CustomerID: customerId, Status: "pending"}).First(&order)
	}
  c.Data["json"] = order
	c.ServeJSON()
}
// @Title Get
// @Description find customer by customer ID
// @Param	customerID		path 	string	true		"the customer id you want to get"
// @Success 200 {customer} models.Customer
// @Failure 403 :customerID is empty
// @router /orders/:customerId [get]
func (c *CustomerController) GetAllOrders() {
  var customerId uint
	c.Ctx.Input.Bind(&customerId, ":customerId")
  orders := [] models.Order{}
	if customerId != 0 {
		db.Preload("Customer").Preload("Retailer").Preload("Service").Where(&models.Order{CustomerID: customerId}).Find(&orders)
	}
	c.Data["json"] = orders
	c.ServeJSON()
}
