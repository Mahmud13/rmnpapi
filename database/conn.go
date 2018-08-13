package database

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
  db, err := gorm.Open("mysql", "root:root@(172.17.0.2)/mnp_api?parseTime=true&charset=utf8&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  DB = db
}
