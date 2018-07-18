package models

import (
  "time"
  "mnp_api/database"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Service struct {
  ID          uint    `json:"id"`
  CreatedAt   time.Time  `json:"created_at,omitempty"`
  UpdatedAt   time.Time `json:"updated_at,omitempty"`
  DeletedAt   *time.Time `json:"deleted_at,omitempty"`
  Name        string  `valid:"Required" json:"name"`
  Price       float32  `sql:"type:decimal(5,2)" valid:"Required" json:"price"`
}

func init(){
  db := database.DB
  db.AutoMigrate(&Service{})
}
