package controllers

import (
	"github.com/astaxie/beego"
  "encoding/json"
  "github.com/astaxie/beego/validation"
	"mnp_api/models"
)
// Operations about object
type OrderController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *OrderController) GetAll() {
	orders := [] models.Order{}
	db.Find(&orders)
	c.Data["json"] = orders
	c.ServeJSON()
}
// @Title Get
// @Description find customer by customer ID
// @Param	customerID		path 	string	true		"the customer id you want to get"
// @Success 200 {customer} models.Customer
// @Failure 403 :customerID is empty
// @router /:orderID [get]
func (c *OrderController) Get() {
	orderID := c.Ctx.Input.Param(":orderID")
	if orderID != "" {
		var order models.Order

		if db.Preload("Customer").Preload("Retailer").First(&order, orderID).RecordNotFound() {
			c.Data["json"] = "no record found"
		}else{
			c.Data["json"] = order
		}
	}
	c.ServeJSON()
}
// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *OrderController) Post() {
	valid := validation.Validation{}
	var order models.Order
	json.Unmarshal(c.Ctx.Input.RequestBody, &order)
  order.Status = "pending"
	b, err := valid.Valid(&order)
	if err != nil {
		panic(err)
	}
	if !b {
    c.Ctx.Output.SetStatus(422)
		c.Data["json"] =  valid.Errors
		c.ServeJSON()
		return
	}
  if err := db.Create(&order).Error; err != nil {
    panic(err)
  }
  c.Data["json"] = order
	c.ServeJSON()
}
// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:orderId [put]
func (c *OrderController) Put() {
	orderID := c.Ctx.Input.Param(":orderId")
	if orderID != "" {
		var request models.Order
		var order models.Order
		json.Unmarshal(c.Ctx.Input.RequestBody, &request)
		if db.First(&order, orderID).RecordNotFound() {
			c.Data["json"] = struct { Success bool `json:"success"` } {false}
		}else{
			db.Model(&order).Update(&request)
			c.Data["json"] = struct { Success bool `json:"success"` } {true}
		}
	}
	c.ServeJSON()
}
