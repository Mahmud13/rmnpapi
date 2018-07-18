package models

import (
  "time"
  "mnp_api/database"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Order struct {
  ID          uint    `json:"id"`
  CreatedAt   time.Time  `json:"created_at,omitempty"`
  UpdatedAt   time.Time `json:"updated_at,omitempty"`
  DeletedAt   *time.Time `json:"deleted_at,omitempty"`
  Service    Service  `json:"service"`
  ServiceID  uint      `valid:"Required" json:"serviceId"`
  Status      string  `valid:"Required" json:"status"`
  Customer    Customer `json:"customer"`
  CustomerID  uint      `valid:"Required" json:"customerId"`
  Retailer    Retailer  `json:"retailer"`
  RetailerID  uint      `valid:"Required" json:"retailerId"`
}

func init(){
  db := database.DB
  db.AutoMigrate(&Order{})
}
