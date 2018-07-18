package models

import (
  "time"
  "mnp_api/database"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
  ID          uint    `json:"id"`
  CreatedAt   time.Time  `json:"created_at,omitempty"`
  UpdatedAt   time.Time `json:"updated_at,omitempty"`
  DeletedAt   *time.Time `json:"deleted_at,omitempty"`
  Name        string  `json:"name"`
  Phone       string  `valid:"Required" json:"phone"`
  Imei        string  `valid:"Required" json:"imei"`
  Longitude   float32 `json:"longitude"`
  Latitude    float32 `json:"latitude"`
  Token       string    `json:"token"`
  Photo       string    `sql:"type:text" json:"photo"`
}

func init(){
  db := database.DB
  db.AutoMigrate(&Customer{})
}
