package models

import (
  "time"

  "mnp_api/database"
  _ "github.com/jinzhu/gorm/dialects/mysql"

)

type Retailer struct {
  ID          uint        `json:"id"`
  CreatedAt   time.Time   `json:"created_at,omitempty"`
  UpdatedAt   time.Time   `json:"updated_at,omitempty"`
  DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
  Name        string      `json:"name"`
  Code        string      `json:code`
  Phone       string      `valid:"Required" json:"phone"`
  Imei        string      `valid:"Required" json:"imei"`
  Longitude   float32     `json:"longitude"`
  Latitude    float32     `json:"latitude"`
  Token       string      `json:"token"`
  Photo       string      `sql:"type:text" json:"photo"`
  Address     string      `json:"address"`
  AreaCode    string      `json:"areaCode"`
  Status      string      `json:"status"`
  Type        string      `json:"type"`
}

func init(){
  db := database.DB
  db.AutoMigrate(&Retailer{})
}
