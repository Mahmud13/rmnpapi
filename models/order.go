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
  CompletedAt   *time.Time `json:"completed_at,omitempty"`
  Service    Service  `json:"service"`
  ServiceID  uint      `valid:"Required" json:"serviceId"`
  Detail      OrderDetail  `json:"detail"`
  DetailID   uint     `json:"detailId"`
  Status      string  `valid:"Required" json:"status"`
  Customer    Customer `json:"customer"`
  CustomerID  uint      `valid:"Required" json:"customerId"`
  Retailer    Retailer  `json:"retailer"`
  RetailerID  uint      `json:"retailerId"`
  Longitude   float64   `valid:"Required" json:"longitude"`
  Latitude   float64   `valid:"Required" json:"latitude"`
}
type OrderDetail struct {
  ID          uint    `json:"id"`
  Number      string  `json:"number,omitempty"`
  DateOfBirth *time.Time  `json:"dateOfBirth,omitempty"`
  PrefDeliveredAt *time.Time `json:"prefDeliveredAt,omitempty"`
}
func init(){
  db := database.DB
  db.AutoMigrate(&Order{})
  db.AutoMigrate(&OrderDetail{})
}
