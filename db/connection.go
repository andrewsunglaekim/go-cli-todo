package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Open() error {
  var err error
  DB, err = gorm.Open("postgres", "host=localhost port=5432 user=andrewkim dbname=todo password=sandhorsebobanus sslmode=disable")
  if err != nil {
    return err
  }
  return err
}

func Close() error {
  return DB.Close()
}
