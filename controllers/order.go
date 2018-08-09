package controllers

import (
	"github.com/astaxie/beego"
  "encoding/json"
  "github.com/astaxie/beego/validation"
  "golang.org/x/net/context"

  firebase "firebase.google.com/go"

  "firebase.google.com/go/messaging"

  "google.golang.org/api/option"
	"mnp_api/models"
	"fmt"
	// "mnp_api/helpers"
	// "math"
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
	orders := []models.Order{}
	db.Preload("Service").Preload("Customer").Preload("Retailer").Order("created_at desc").Limit(10).Find(&orders)
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

		if db.Preload("Customer").Preload("Retailer").Preload("Service").First(&order, orderID).RecordNotFound() {
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
	// distance := 10.0;
	//
	// radius := 6371.0;
	//
	// maxlat := order.Latitude + helpers.Rad2deg(distance / radius);
	// minlat := order.Latitude - helpers.Rad2deg(distance / radius);
	//
	// maxlng := order.Longitude + helpers.Rad2deg(distance / radius / math.Cos(helpers.Deg2rad(order.Latitude)));
	// minlng := order.Longitude - helpers.Rad2deg(distance / radius / math.Cos(helpers.Deg2rad(order.Latitude)));
	var retailer models.Retailer
	// if db.Where("latitude BETWEEN ? AND ?", minlat, maxlat).Where("longitude BETWEEN ? AND ?", minlng, maxlng).First(&retailer, "type = ?", "moving").RecordNotFound() {
	// 	c.Ctx.Output.SetStatus(500)
	// 	c.Data["json"] = "No nearby retailers found"
	// 	c.ServeJSON()
	// 	return
	// }
	if db.First(&retailer, 16).RecordNotFound() {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = "No nearby retailers found"
		c.ServeJSON()
		return
	}
	order.Retailer = retailer
  if err := db.Create(&order).Error; err != nil {
    panic(err)
  }
	if db.Preload("Customer").Preload("Retailer").Preload("Service").First(&order, order.ID).RecordNotFound() {
		c.Data["json"] = `No record found`
		c.ServeJSON()
		return
	}
	_, err = sendNotification(order.Retailer.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = sendNotification(order.Customer.Token)

	if err != nil {
		fmt.Println(err.Error())
	}
  c.Data["json"] = order
	c.ServeJSON()
}
func sendNotification(registrationToken string) (string, error){
  opt := option.WithCredentialsFile("storage/robi-mnp-firebase-adminsdk-eyz6a-6a47d7b2bb.json")
  ctx := context.Background()
  app, err := firebase.NewApp(ctx, nil, opt)
  if err != nil {
		return "", err
  }
  // Access auth service from the default app
  client, err := app.Messaging(ctx)
  if err != nil {
		return "", err
  }
  // This registration token comes from the client FCM SDKs.

  // See documentation on defining a message payload.
  message := &messaging.Message{
        Data: map[string]string{
                "title": "ROBI MNP",
                "body":  "New Request found",
        },
        Token: registrationToken,
    }
    // Send a message to the device corresponding to the provided
    // registration token.
    response, err := client.Send(ctx, message)
    if err != nil {
			return "", err
		}
		return response, nil
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
